package cryptixd

import (
	configPackage "github.com/cryptix-network/cryptix-graph-inspector/processing/infrastructure/config"
	"github.com/cryptix-network/cryptix-graph-inspector/processing/infrastructure/database"

	//"github.com/cryptix-network/cryptix-graph-inspector/processing/infrastructure/logging"
	"net"

	domainPackage "github.com/cryptix-network/cryptix-graph-inspector/processing/cryptixd/domain"
	consensusPackage "github.com/cryptix-network/cryptix-graph-inspector/processing/cryptixd/domain/consensus"
	"github.com/cryptix-network/cryptixd/app/protocol"
	"github.com/cryptix-network/cryptixd/domain/consensus/model/externalapi"
	cryptixdConfigPackage "github.com/cryptix-network/cryptixd/infrastructure/config"
	"github.com/cryptix-network/cryptixd/infrastructure/network/addressmanager"
	"github.com/cryptix-network/cryptixd/infrastructure/network/connmanager"
	"github.com/cryptix-network/cryptixd/infrastructure/network/netadapter"
	"github.com/cryptix-network/cryptixd/infrastructure/network/netadapter/router"
)

type Cryptixd struct {
	config            *configPackage.Config
	domain            *domainPackage.Domain
	netAdapter        *netadapter.NetAdapter
	addressManager    *addressmanager.AddressManager
	connectionManager *connmanager.ConnectionManager
	protocolManager   *protocol.Manager
}

func New(config *configPackage.Config) (*Cryptixd, error) {
	cryptixdConfig := cryptixdConfigPackage.DefaultConfig()
	cryptixdConfig.ConnectPeers = config.ConnectPeers
	cryptixdConfig.DNSSeed = config.DNSSeed
	cryptixdConfig.GRPCSeed = config.GRPCSeed
	cryptixdConfig.NetworkFlags = config.NetworkFlags
	cryptixdConfig.Lookup = net.LookupIP

	//logging.UpdateLogLevels()

	databaseContext, err := database.Open(config)
	if err != nil {
		return nil, err
	}
	domain, err := domainPackage.New(config.NetworkFlags.ActiveNetParams, databaseContext)
	if err != nil {
		return nil, err
	}
	netAdapter, err := netadapter.NewNetAdapter(cryptixdConfig)
	if err != nil {
		return nil, err
	}
	netAdapter.SetRPCRouterInitializer(func(router *router.Router, connection *netadapter.NetConnection) {})
	addressManager, err := addressmanager.New(addressmanager.NewConfig(cryptixdConfig), databaseContext)
	if err != nil {
		return nil, err
	}
	connectionManager, err := connmanager.New(cryptixdConfig, netAdapter, addressManager)
	if err != nil {
		return nil, err
	}
	protocolManager, err := protocol.NewManager(cryptixdConfig, domain, netAdapter, addressManager, connectionManager)
	if err != nil {
		return nil, err
	}
	return &Cryptixd{
		config:            config,
		domain:            domain,
		netAdapter:        netAdapter,
		addressManager:    addressManager,
		connectionManager: connectionManager,
		protocolManager:   protocolManager,
	}, nil
}

func (k *Cryptixd) SetOnBlockAddedListener(listener consensusPackage.OnBlockAddedListener) {
	k.domain.SetOnBlockAddedListener(listener)
}

func (k *Cryptixd) SetOnVirtualResolvedListener(listener consensusPackage.OnVirtualResolvedListener) {
	k.domain.SetOnVirtualResolvedListener(listener)
}

func (k *Cryptixd) SetOnConsensusResetListener(listener domainPackage.OnConsensusResetListener) {
	k.domain.SetOnConsensusResetListener(listener)
}

func (k *Cryptixd) BlockGHOSTDAGData(blockHash *externalapi.DomainHash) (*externalapi.BlockGHOSTDAGData, error) {
	return k.domain.BlockGHOSTDAGData(blockHash)
}

func (k *Cryptixd) Start() error {
	err := k.netAdapter.Start()
	if err != nil {
		return err
	}
	k.connectionManager.Start()
	return nil
}

func (k *Cryptixd) Domain() *domainPackage.Domain {
	return k.domain
}
