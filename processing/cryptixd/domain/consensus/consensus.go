package consensus

import (
	cryptixdConsensus "github.com/cryptix-network/cryptixd/domain/consensus"
	consensusDatabase "github.com/cryptix-network/cryptixd/domain/consensus/database"
	"github.com/cryptix-network/cryptixd/domain/consensus/datastructures/ghostdagdatastore"
	"github.com/cryptix-network/cryptixd/domain/consensus/model"
	"github.com/cryptix-network/cryptixd/domain/consensus/model/externalapi"
	"github.com/cryptix-network/cryptixd/domain/prefixmanager/prefix"
	"github.com/cryptix-network/cryptixd/infrastructure/db/database"
)

func New(consensusConfig *cryptixdConsensus.Config, databaseContext database.Database, dbPrefix *prefix.Prefix, consensusEventsChan chan externalapi.ConsensusEvent) (*Consensus, bool, error) {
	cryptixdConsensusFactory := cryptixdConsensus.NewFactory()
	cryptixdConsensusInstance, shouldMigrate, err := cryptixdConsensusFactory.NewConsensus(consensusConfig, databaseContext, dbPrefix, consensusEventsChan)
	if err != nil {
		return nil, false, err
	}

	dbManager := consensusDatabase.New(databaseContext)
	pruningWindowSizeForCaches := int(consensusConfig.Params.PruningDepth())
	prefixBucket := consensusDatabase.MakeBucket(dbPrefix.Serialize())
	ghostdagDataStore := ghostdagdatastore.New(prefixBucket.Bucket([]byte{byte(0)}), pruningWindowSizeForCaches, true)

	return &Consensus{
		dbManager:         dbManager,
		cryptixdConsensus: cryptixdConsensusInstance,
		ghostdagDataStore: ghostdagDataStore,
	}, shouldMigrate, nil
}

type Consensus struct {
	dbManager         model.DBManager
	cryptixdConsensus externalapi.Consensus
	ghostdagDataStore model.GHOSTDAGDataStore

	onBlockAddedListener      OnBlockAddedListener
	onVirtualResolvedListener OnVirtualResolvedListener
}

func (c *Consensus) ValidateAndInsertBlock(block *externalapi.DomainBlock, shouldValidateAgainstUTXO bool) error {
	err := c.cryptixdConsensus.ValidateAndInsertBlock(block, shouldValidateAgainstUTXO)
	if err != nil {
		return err
	}
	if c.onBlockAddedListener != nil {
		c.onBlockAddedListener(block)
	}
	return nil
}

func (c *Consensus) ResolveVirtual(progressReportCallback func(uint64, uint64)) error {
	err := c.cryptixdConsensus.ResolveVirtual(progressReportCallback)
	if err != nil {
		return err
	}
	if c.onVirtualResolvedListener != nil {
		c.onVirtualResolvedListener()
	}
	return nil
}

type OnBlockAddedListener func(*externalapi.DomainBlock)
type OnVirtualResolvedListener func()

func (c *Consensus) SetOnBlockAddedListener(listener OnBlockAddedListener) {
	c.onBlockAddedListener = listener
}

func (c *Consensus) SetOnVirtualResolvedListener(listener OnVirtualResolvedListener) {
	c.onVirtualResolvedListener = listener
}

func (c *Consensus) BlockGHOSTDAGData(blockHash *externalapi.DomainHash) (*externalapi.BlockGHOSTDAGData, error) {
	return c.ghostdagDataStore.Get(c.dbManager, model.NewStagingArea(), blockHash, false)
}
