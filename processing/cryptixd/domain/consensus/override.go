package consensus

import "github.com/cryptix-network/cryptixd/domain/consensus/model/externalapi"

func (c *Consensus) BuildBlock(coinbaseData *externalapi.DomainCoinbaseData, transactions []*externalapi.DomainTransaction) (*externalapi.DomainBlock, error) {
	return c.cryptixdConsensus.BuildBlock(coinbaseData, transactions)
}

func (c *Consensus) BuildBlockTemplate(coinbaseData *externalapi.DomainCoinbaseData, transactions []*externalapi.DomainTransaction) (*externalapi.DomainBlockTemplate, error) {
	return c.cryptixdConsensus.BuildBlockTemplate(coinbaseData, transactions)
}

func (c *Consensus) ValidateTransactionAndPopulateWithConsensusData(transaction *externalapi.DomainTransaction) error {
	return c.cryptixdConsensus.ValidateTransactionAndPopulateWithConsensusData(transaction)
}

func (c *Consensus) GetBlock(blockHash *externalapi.DomainHash) (*externalapi.DomainBlock, error) {
	return c.cryptixdConsensus.GetBlock(blockHash)
}

func (c *Consensus) GetBlockHeader(blockHash *externalapi.DomainHash) (externalapi.BlockHeader, error) {
	return c.cryptixdConsensus.GetBlockHeader(blockHash)
}

func (c *Consensus) GetBlockInfo(blockHash *externalapi.DomainHash) (*externalapi.BlockInfo, error) {
	return c.cryptixdConsensus.GetBlockInfo(blockHash)
}

func (c *Consensus) GetHashesBetween(lowHash, highHash *externalapi.DomainHash, maxBlueScoreDifference uint64) (
	[]*externalapi.DomainHash, *externalapi.DomainHash, error) {

	return c.cryptixdConsensus.GetHashesBetween(lowHash, highHash, maxBlueScoreDifference)
}

func (c *Consensus) GetAnticone(blockHash, contextHash *externalapi.DomainHash, maxBlocks uint64) (hashes []*externalapi.DomainHash, err error) {
	return c.cryptixdConsensus.GetAnticone(blockHash, contextHash, maxBlocks)
}

func (c *Consensus) GetPruningPointUTXOs(expectedPruningPointHash *externalapi.DomainHash, fromOutpoint *externalapi.DomainOutpoint, limit int) ([]*externalapi.OutpointAndUTXOEntryPair, error) {
	return c.cryptixdConsensus.GetPruningPointUTXOs(expectedPruningPointHash, fromOutpoint, limit)
}

func (c *Consensus) GetVirtualUTXOs(expectedVirtualParents []*externalapi.DomainHash, fromOutpoint *externalapi.DomainOutpoint, limit int) ([]*externalapi.OutpointAndUTXOEntryPair, error) {
	return c.cryptixdConsensus.GetVirtualUTXOs(expectedVirtualParents, fromOutpoint, limit)
}

func (c *Consensus) PruningPoint() (*externalapi.DomainHash, error) {
	return c.cryptixdConsensus.PruningPoint()
}

func (c *Consensus) ClearImportedPruningPointData() error {
	return c.cryptixdConsensus.ClearImportedPruningPointData()
}

func (c *Consensus) AppendImportedPruningPointUTXOs(outpointAndUTXOEntryPairs []*externalapi.OutpointAndUTXOEntryPair) error {
	return c.cryptixdConsensus.AppendImportedPruningPointUTXOs(outpointAndUTXOEntryPairs)
}

func (c *Consensus) GetVirtualSelectedParent() (*externalapi.DomainHash, error) {
	return c.cryptixdConsensus.GetVirtualSelectedParent()
}

func (c *Consensus) CreateHeadersSelectedChainBlockLocator(lowHash, highHash *externalapi.DomainHash) (externalapi.BlockLocator, error) {
	return c.cryptixdConsensus.CreateHeadersSelectedChainBlockLocator(lowHash, highHash)
}

func (c *Consensus) CreateFullHeadersSelectedChainBlockLocator() (externalapi.BlockLocator, error) {
	return c.cryptixdConsensus.CreateFullHeadersSelectedChainBlockLocator()
}

func (c *Consensus) GetSyncInfo() (*externalapi.SyncInfo, error) {
	return c.cryptixdConsensus.GetSyncInfo()
}

func (c *Consensus) Tips() ([]*externalapi.DomainHash, error) {
	return c.cryptixdConsensus.Tips()
}

func (c *Consensus) GetVirtualInfo() (*externalapi.VirtualInfo, error) {
	return c.cryptixdConsensus.GetVirtualInfo()
}

func (c *Consensus) IsValidPruningPoint(blockHash *externalapi.DomainHash) (bool, error) {
	return c.cryptixdConsensus.IsValidPruningPoint(blockHash)
}

func (c *Consensus) GetVirtualSelectedParentChainFromBlock(blockHash *externalapi.DomainHash) (*externalapi.SelectedChainPath, error) {
	return c.cryptixdConsensus.GetVirtualSelectedParentChainFromBlock(blockHash)
}

func (c *Consensus) IsInSelectedParentChainOf(blockHashA *externalapi.DomainHash, blockHashB *externalapi.DomainHash) (bool, error) {
	return c.cryptixdConsensus.IsInSelectedParentChainOf(blockHashA, blockHashB)
}

func (c *Consensus) GetHeadersSelectedTip() (*externalapi.DomainHash, error) {
	return c.cryptixdConsensus.GetHeadersSelectedTip()
}

func (c *Consensus) Anticone(blockHash *externalapi.DomainHash) ([]*externalapi.DomainHash, error) {
	return c.cryptixdConsensus.Anticone(blockHash)
}

func (c *Consensus) GetBlockRelations(blockHash *externalapi.DomainHash) (
	parents []*externalapi.DomainHash, children []*externalapi.DomainHash, err error) {

	return c.cryptixdConsensus.GetBlockRelations(blockHash)
}

func (s *Consensus) GetBlockAcceptanceData(blockHash *externalapi.DomainHash) (externalapi.AcceptanceData, error) {
	return s.cryptixdConsensus.GetBlockAcceptanceData(blockHash)
}

func (s *Consensus) GetBlocksAcceptanceData(blockHashes []*externalapi.DomainHash) ([]externalapi.AcceptanceData, error) {
	return s.cryptixdConsensus.GetBlocksAcceptanceData(blockHashes)
}

func (c *Consensus) GetBlockEvenIfHeaderOnly(blockHash *externalapi.DomainHash) (*externalapi.DomainBlock, error) {
	return c.cryptixdConsensus.GetBlockEvenIfHeaderOnly(blockHash)
}

func (c *Consensus) EstimateNetworkHashesPerSecond(startHash *externalapi.DomainHash, windowSize int) (uint64, error) {
	return c.cryptixdConsensus.EstimateNetworkHashesPerSecond(startHash, windowSize)
}

func (c *Consensus) GetVirtualDAAScore() (uint64, error) {
	return c.cryptixdConsensus.GetVirtualDAAScore()
}

func (c *Consensus) Init(shouldNotAddGenesis bool) error {
	return c.cryptixdConsensus.Init(shouldNotAddGenesis)
}

func (c *Consensus) PruningPointAndItsAnticone() ([]*externalapi.DomainHash, error) {
	return c.cryptixdConsensus.PruningPointAndItsAnticone()
}

func (c *Consensus) ValidateAndInsertImportedPruningPoint(newPruningPoint *externalapi.DomainHash) error {
	return c.cryptixdConsensus.ValidateAndInsertImportedPruningPoint(newPruningPoint)
}

func (c *Consensus) CreateBlockLocatorFromPruningPoint(highHash *externalapi.DomainHash, limit uint32) (externalapi.BlockLocator, error) {
	return c.cryptixdConsensus.CreateBlockLocatorFromPruningPoint(highHash, limit)
}

func (c *Consensus) PopulateMass(transaction *externalapi.DomainTransaction) {
	c.cryptixdConsensus.PopulateMass(transaction)
}

func (c *Consensus) ValidateAndInsertBlockWithTrustedData(block *externalapi.BlockWithTrustedData, validateUTXO bool) error {
	return c.cryptixdConsensus.ValidateAndInsertBlockWithTrustedData(block, validateUTXO)
}

func (c *Consensus) GetMissingBlockBodyHashes(highHash *externalapi.DomainHash) ([]*externalapi.DomainHash, error) {
	return c.cryptixdConsensus.GetMissingBlockBodyHashes(highHash)
}

func (c *Consensus) ImportPruningPoints(pruningPoints []externalapi.BlockHeader) error {
	return c.cryptixdConsensus.ImportPruningPoints(pruningPoints)
}

func (c *Consensus) BuildPruningPointProof() (*externalapi.PruningPointProof, error) {
	return c.cryptixdConsensus.BuildPruningPointProof()
}

func (c *Consensus) ValidatePruningPointProof(pruningPointProof *externalapi.PruningPointProof) error {
	return c.cryptixdConsensus.ValidatePruningPointProof(pruningPointProof)
}

func (c *Consensus) ApplyPruningPointProof(pruningPointProof *externalapi.PruningPointProof) error {
	return c.cryptixdConsensus.ApplyPruningPointProof(pruningPointProof)
}

func (c *Consensus) PruningPointHeaders() ([]externalapi.BlockHeader, error) {
	return c.cryptixdConsensus.PruningPointHeaders()
}

func (c *Consensus) ArePruningPointsViolatingFinality(pruningPoints []externalapi.BlockHeader) (bool, error) {
	return c.cryptixdConsensus.ArePruningPointsViolatingFinality(pruningPoints)
}

func (c *Consensus) BlockDAAWindowHashes(blockHash *externalapi.DomainHash) ([]*externalapi.DomainHash, error) {
	return c.cryptixdConsensus.BlockDAAWindowHashes(blockHash)
}

func (c *Consensus) TrustedDataDataDAAHeader(trustedBlockHash, daaBlockHash *externalapi.DomainHash, daaBlockWindowIndex uint64) (*externalapi.TrustedDataDataDAAHeader, error) {
	return c.cryptixdConsensus.TrustedDataDataDAAHeader(trustedBlockHash, daaBlockHash, daaBlockWindowIndex)
}

func (c *Consensus) TrustedBlockAssociatedGHOSTDAGDataBlockHashes(blockHash *externalapi.DomainHash) ([]*externalapi.DomainHash, error) {
	return c.cryptixdConsensus.TrustedBlockAssociatedGHOSTDAGDataBlockHashes(blockHash)
}

func (c *Consensus) TrustedGHOSTDAGData(blockHash *externalapi.DomainHash) (*externalapi.BlockGHOSTDAGData, error) {
	return c.cryptixdConsensus.TrustedGHOSTDAGData(blockHash)
}

func (c *Consensus) IsChainBlock(blockHash *externalapi.DomainHash) (bool, error) {
	return c.cryptixdConsensus.IsChainBlock(blockHash)
}

func (c *Consensus) VirtualMergeDepthRoot() (*externalapi.DomainHash, error) {
	return c.cryptixdConsensus.VirtualMergeDepthRoot()
}

func (c *Consensus) IsNearlySynced() (bool, error) {
	return c.cryptixdConsensus.IsNearlySynced()
}
