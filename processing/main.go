package main

import (
	"fmt"

	cryptixdPackage "github.com/cryptix-network/cryptix-graph-inspector/processing/cryptixd"
	databasePackage "github.com/cryptix-network/cryptix-graph-inspector/processing/database"
	configPackage "github.com/cryptix-network/cryptix-graph-inspector/processing/infrastructure/config"
	"github.com/cryptix-network/cryptix-graph-inspector/processing/infrastructure/logging"
	processingPackage "github.com/cryptix-network/cryptix-graph-inspector/processing/processing"
	versionPackage "github.com/cryptix-network/cryptix-graph-inspector/processing/version"
	"github.com/cryptix-network/cryptixd/version"
)

func main() {
	fmt.Println("=================================================")
	fmt.Println("Cryptix Graph Inspector (KGI)   -   Processing Tier")
	fmt.Println("=================================================")

	config, err := configPackage.LoadConfig()
	if err != nil {
		logging.LogErrorAndExit("Could not parse command line arguments.\n%s", err)
	}

	logging.Logger().Infof("Application version %s", versionPackage.Version())
	logging.Logger().Infof("Embedded cryptixd version %s", version.Version())
	logging.Logger().Infof("Network %s", config.ActiveNetParams.Name)

	database, err := databasePackage.Connect(config.DatabaseConnectionString)
	if err != nil {
		logging.LogErrorAndExit("Could not connect to database %s: %s", config.DatabaseConnectionString, err)
	}
	defer database.Close()

	cryptixd, err := cryptixdPackage.New(config)
	if err != nil {
		logging.LogErrorAndExit("Could not create cryptixd: %s", err)
	}
	processing, err := processingPackage.NewProcessing(config, database, cryptixd)
	if err != nil {
		logging.LogErrorAndExit("Could not initialize processing: %s", err)
	}

	// This is no longer useful since cryptixd v0.12.2
	// that introduce a consensus event channel.
	// See processing.initConsensusEventsHandler.

	// cryptixd.SetOnBlockAddedListener(func(block *externalapi.DomainBlock) {
	// 	blockHash := consensushashing.BlockHash(block)
	// 	blockInfo, err := cryptixd.Domain().Consensus().GetBlockInfo(blockHash)
	// 	if err != nil {
	// 		logging.LogErrorAndExit("Consensus ValidateAndInsertBlock listener could not get block info for block %s: %s", blockHash, err)
	// 	}
	// 	logging.Logger().Debugf("Consensus ValidateAndInsertBlock listener gets block %s with status %s", blockHash, blockInfo.BlockStatus.String())
	// })

	cryptixd.SetOnVirtualResolvedListener(func() {
		err := processing.ResyncVirtualSelectedParentChain()
		if err != nil {
			logging.LogErrorAndExit("Could not resync the virtual selected parent chain: %s", err)
		}
	})
	cryptixd.SetOnConsensusResetListener(func() {
		err := processing.ResyncDatabase()
		if err != nil {
			logging.LogErrorAndExit("Could not resync database: %s", err)
		}
	})
	err = cryptixd.Start()
	if err != nil {
		logging.LogErrorAndExit("Could not start cryptixd: %s", err)
	}

	<-make(chan struct{})
}
