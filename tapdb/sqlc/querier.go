// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package sqlc

import (
	"context"
	"database/sql"
	"time"
)

type Querier interface {
	AllAssets(ctx context.Context) ([]Asset, error)
	AllInternalKeys(ctx context.Context) ([]InternalKey, error)
	AllMintingBatches(ctx context.Context) ([]AllMintingBatchesRow, error)
	AnchorGenesisPoint(ctx context.Context, arg AnchorGenesisPointParams) error
	AnchorPendingAssets(ctx context.Context, arg AnchorPendingAssetsParams) error
	ApplyPendingOutput(ctx context.Context, arg ApplyPendingOutputParams) (int64, error)
	AssetsByGenesisPoint(ctx context.Context, prevOut []byte) ([]AssetsByGenesisPointRow, error)
	AssetsDBSizePostgres(ctx context.Context) (int64, error)
	AssetsDBSizeSqlite(ctx context.Context) (int32, error)
	AssetsInBatch(ctx context.Context, rawKey []byte) ([]AssetsInBatchRow, error)
	BindMintingBatchWithTapSibling(ctx context.Context, arg BindMintingBatchWithTapSiblingParams) error
	BindMintingBatchWithTx(ctx context.Context, arg BindMintingBatchWithTxParams) error
	ConfirmChainAnchorTx(ctx context.Context, arg ConfirmChainAnchorTxParams) error
	ConfirmChainTx(ctx context.Context, arg ConfirmChainTxParams) error
	DeleteAllNodes(ctx context.Context, namespace string) (int64, error)
	DeleteAssetWitnesses(ctx context.Context, assetID int64) error
	DeleteExpiredUTXOLeases(ctx context.Context, now sql.NullTime) error
	DeleteFederationProofSyncLog(ctx context.Context, arg DeleteFederationProofSyncLogParams) error
	DeleteManagedUTXO(ctx context.Context, outpoint []byte) error
	DeleteMultiverseLeaf(ctx context.Context, arg DeleteMultiverseLeafParams) error
	DeleteNode(ctx context.Context, arg DeleteNodeParams) (int64, error)
	DeleteRoot(ctx context.Context, namespace string) (int64, error)
	DeleteTapscriptTreeEdges(ctx context.Context, rootHash []byte) error
	DeleteTapscriptTreeNodes(ctx context.Context) error
	DeleteTapscriptTreeRoot(ctx context.Context, rootHash []byte) error
	DeleteUTXOLease(ctx context.Context, outpoint []byte) error
	DeleteUniverseEvents(ctx context.Context, namespaceRoot string) error
	DeleteUniverseLeaves(ctx context.Context, namespace string) error
	DeleteUniverseRoot(ctx context.Context, namespaceRoot string) error
	DeleteUniverseServer(ctx context.Context, arg DeleteUniverseServerParams) error
	FetchAddrByTaprootOutputKey(ctx context.Context, taprootOutputKey []byte) (FetchAddrByTaprootOutputKeyRow, error)
	FetchAddrEvent(ctx context.Context, id int64) (FetchAddrEventRow, error)
	FetchAddrEventByAddrKeyAndOutpoint(ctx context.Context, arg FetchAddrEventByAddrKeyAndOutpointParams) (FetchAddrEventByAddrKeyAndOutpointRow, error)
	FetchAddrs(ctx context.Context, arg FetchAddrsParams) ([]FetchAddrsRow, error)
	FetchAllNodes(ctx context.Context) ([]MssmtNode, error)
	FetchAssetID(ctx context.Context, arg FetchAssetIDParams) ([]int64, error)
	FetchAssetMeta(ctx context.Context, metaID int64) (FetchAssetMetaRow, error)
	FetchAssetMetaByHash(ctx context.Context, metaDataHash []byte) (FetchAssetMetaByHashRow, error)
	FetchAssetMetaForAsset(ctx context.Context, assetID []byte) (FetchAssetMetaForAssetRow, error)
	FetchAssetProof(ctx context.Context, arg FetchAssetProofParams) ([]FetchAssetProofRow, error)
	FetchAssetProofs(ctx context.Context) ([]FetchAssetProofsRow, error)
	FetchAssetProofsByAssetID(ctx context.Context, assetID []byte) ([]FetchAssetProofsByAssetIDRow, error)
	FetchAssetProofsSizes(ctx context.Context) ([]FetchAssetProofsSizesRow, error)
	FetchAssetWitnesses(ctx context.Context, assetID sql.NullInt64) ([]FetchAssetWitnessesRow, error)
	FetchAssetsByAnchorTx(ctx context.Context, anchorUtxoID sql.NullInt64) ([]Asset, error)
	// We use a LEFT JOIN here as not every asset has a meta data entry.
	// We use a LEFT JOIN here as not every asset has a group key, so this'll
	// generate rows that have NULL values for the group key fields if an asset
	// doesn't have a group key. See the comment in fetchAssetSprouts for a work
	// around that needs to be used with this query until a sqlc bug is fixed.
	FetchAssetsForBatch(ctx context.Context, rawKey []byte) ([]FetchAssetsForBatchRow, error)
	FetchChainTx(ctx context.Context, txid []byte) (ChainTxn, error)
	FetchChildren(ctx context.Context, arg FetchChildrenParams) ([]FetchChildrenRow, error)
	FetchChildrenSelfJoin(ctx context.Context, arg FetchChildrenSelfJoinParams) ([]FetchChildrenSelfJoinRow, error)
	FetchGenesisByAssetID(ctx context.Context, assetID []byte) (GenesisInfoView, error)
	FetchGenesisByID(ctx context.Context, genAssetID int64) (FetchGenesisByIDRow, error)
	FetchGenesisID(ctx context.Context, arg FetchGenesisIDParams) (int64, error)
	FetchGenesisPointByAnchorTx(ctx context.Context, anchorTxID sql.NullInt64) (GenesisPoint, error)
	FetchGroupByGenesis(ctx context.Context, genesisID int64) (FetchGroupByGenesisRow, error)
	// Sort and limit to return the genesis ID for initial genesis of the group.
	FetchGroupByGroupKey(ctx context.Context, groupKey []byte) (FetchGroupByGroupKeyRow, error)
	FetchGroupedAssets(ctx context.Context) ([]FetchGroupedAssetsRow, error)
	FetchInternalKeyLocator(ctx context.Context, rawKey []byte) (FetchInternalKeyLocatorRow, error)
	FetchManagedUTXO(ctx context.Context, arg FetchManagedUTXOParams) (FetchManagedUTXORow, error)
	FetchManagedUTXOs(ctx context.Context) ([]FetchManagedUTXOsRow, error)
	FetchMintingBatch(ctx context.Context, rawKey []byte) (FetchMintingBatchRow, error)
	FetchMintingBatchesByInverseState(ctx context.Context, batchState int16) ([]FetchMintingBatchesByInverseStateRow, error)
	FetchMultiverseRoot(ctx context.Context, namespaceRoot string) (FetchMultiverseRootRow, error)
	FetchRootNode(ctx context.Context, namespace string) (MssmtNode, error)
	FetchScriptKeyByTweakedKey(ctx context.Context, tweakedScriptKey []byte) (FetchScriptKeyByTweakedKeyRow, error)
	FetchScriptKeyIDByTweakedKey(ctx context.Context, tweakedScriptKey []byte) (int64, error)
	FetchSeedlingByID(ctx context.Context, seedlingID int64) (AssetSeedling, error)
	FetchSeedlingID(ctx context.Context, arg FetchSeedlingIDParams) (int64, error)
	FetchSeedlingsForBatch(ctx context.Context, rawKey []byte) ([]FetchSeedlingsForBatchRow, error)
	// Sort the nodes by node_index here instead of returning the indices.
	FetchTapscriptTree(ctx context.Context, rootHash []byte) ([]FetchTapscriptTreeRow, error)
	FetchTransferInputs(ctx context.Context, transferID int64) ([]FetchTransferInputsRow, error)
	FetchTransferOutputs(ctx context.Context, transferID int64) ([]FetchTransferOutputsRow, error)
	FetchUniverseKeys(ctx context.Context, arg FetchUniverseKeysParams) ([]FetchUniverseKeysRow, error)
	FetchUniverseRoot(ctx context.Context, namespace string) (FetchUniverseRootRow, error)
	GenesisAssets(ctx context.Context) ([]GenesisAsset, error)
	GenesisPoints(ctx context.Context) ([]GenesisPoint, error)
	GetRootKey(ctx context.Context, id []byte) (Macaroon, error)
	HasAssetProof(ctx context.Context, tweakedScriptKey []byte) (bool, error)
	InsertAddr(ctx context.Context, arg InsertAddrParams) (int64, error)
	InsertAssetSeedling(ctx context.Context, arg InsertAssetSeedlingParams) error
	InsertAssetSeedlingIntoBatch(ctx context.Context, arg InsertAssetSeedlingIntoBatchParams) error
	InsertAssetTransfer(ctx context.Context, arg InsertAssetTransferParams) (int64, error)
	InsertAssetTransferInput(ctx context.Context, arg InsertAssetTransferInputParams) error
	InsertAssetTransferOutput(ctx context.Context, arg InsertAssetTransferOutputParams) error
	InsertBranch(ctx context.Context, arg InsertBranchParams) error
	InsertBurn(ctx context.Context, arg InsertBurnParams) (int64, error)
	InsertCompactedLeaf(ctx context.Context, arg InsertCompactedLeafParams) error
	InsertLeaf(ctx context.Context, arg InsertLeafParams) error
	InsertNewProofEvent(ctx context.Context, arg InsertNewProofEventParams) error
	InsertNewSyncEvent(ctx context.Context, arg InsertNewSyncEventParams) error
	InsertPassiveAsset(ctx context.Context, arg InsertPassiveAssetParams) error
	InsertRootKey(ctx context.Context, arg InsertRootKeyParams) error
	InsertUniverseServer(ctx context.Context, arg InsertUniverseServerParams) error
	LogProofTransferAttempt(ctx context.Context, arg LogProofTransferAttemptParams) error
	LogServerSync(ctx context.Context, arg LogServerSyncParams) error
	NewMintingBatch(ctx context.Context, arg NewMintingBatchParams) error
	// We use a LEFT JOIN here as not every asset has a group key, so this'll
	// generate rows that have NULL values for the group key fields if an asset
	// doesn't have a group key. See the comment in fetchAssetSprouts for a work
	// around that needs to be used with this query until a sqlc bug is fixed.
	QueryAssetBalancesByAsset(ctx context.Context, arg QueryAssetBalancesByAssetParams) ([]QueryAssetBalancesByAssetRow, error)
	QueryAssetBalancesByGroup(ctx context.Context, arg QueryAssetBalancesByGroupParams) ([]QueryAssetBalancesByGroupRow, error)
	// BETWEEN is inclusive for both start and end values.
	QueryAssetStatsPerDayPostgres(ctx context.Context, arg QueryAssetStatsPerDayPostgresParams) ([]QueryAssetStatsPerDayPostgresRow, error)
	QueryAssetStatsPerDaySqlite(ctx context.Context, arg QueryAssetStatsPerDaySqliteParams) ([]QueryAssetStatsPerDaySqliteRow, error)
	QueryAssetTransfers(ctx context.Context, arg QueryAssetTransfersParams) ([]QueryAssetTransfersRow, error)
	// We use a LEFT JOIN here as not every asset has a group key, so this'll
	// generate rows that have NULL values for the group key fields if an asset
	// doesn't have a group key. See the comment in fetchAssetSprouts for a work
	// around that needs to be used with this query until a sqlc bug is fixed.
	// This clause is used to select specific assets for a asset ID, general
	// channel balances, and also coin selection. We use the sqlc.narg feature to
	// make the entire statement evaluate to true, if none of these extra args are
	// specified.
	QueryAssets(ctx context.Context, arg QueryAssetsParams) ([]QueryAssetsRow, error)
	QueryBurns(ctx context.Context, arg QueryBurnsParams) ([]QueryBurnsRow, error)
	QueryEventIDs(ctx context.Context, arg QueryEventIDsParams) ([]QueryEventIDsRow, error)
	QueryFederationGlobalSyncConfigs(ctx context.Context) ([]FederationGlobalSyncConfig, error)
	// Join on mssmt_nodes to get leaf related fields.
	// Join on genesis_info_view to get leaf related fields.
	QueryFederationProofSyncLog(ctx context.Context, arg QueryFederationProofSyncLogParams) ([]QueryFederationProofSyncLogRow, error)
	QueryFederationUniSyncConfigs(ctx context.Context) ([]FederationUniSyncConfig, error)
	QueryMultiverseLeaves(ctx context.Context, arg QueryMultiverseLeavesParams) ([]QueryMultiverseLeavesRow, error)
	QueryPassiveAssets(ctx context.Context, transferID int64) ([]QueryPassiveAssetsRow, error)
	QueryProofTransferAttempts(ctx context.Context, arg QueryProofTransferAttemptsParams) ([]time.Time, error)
	// TODO(roasbeef): use the universe id instead for the grouping? so namespace
	// root, simplifies queries
	QueryUniverseAssetStats(ctx context.Context, arg QueryUniverseAssetStatsParams) ([]QueryUniverseAssetStatsRow, error)
	QueryUniverseLeaves(ctx context.Context, arg QueryUniverseLeavesParams) ([]QueryUniverseLeavesRow, error)
	QueryUniverseServers(ctx context.Context, arg QueryUniverseServersParams) ([]UniverseServer, error)
	QueryUniverseStats(ctx context.Context) (QueryUniverseStatsRow, error)
	ReAnchorPassiveAssets(ctx context.Context, arg ReAnchorPassiveAssetsParams) error
	SetAddrManaged(ctx context.Context, arg SetAddrManagedParams) error
	SetAssetSpent(ctx context.Context, arg SetAssetSpentParams) (int64, error)
	SetTransferOutputProofDeliveryStatus(ctx context.Context, arg SetTransferOutputProofDeliveryStatusParams) error
	UniverseLeaves(ctx context.Context) ([]UniverseLeafe, error)
	UniverseRoots(ctx context.Context, arg UniverseRootsParams) ([]UniverseRootsRow, error)
	UpdateBatchGenesisTx(ctx context.Context, arg UpdateBatchGenesisTxParams) error
	UpdateMintingBatchState(ctx context.Context, arg UpdateMintingBatchStateParams) error
	UpdateUTXOLease(ctx context.Context, arg UpdateUTXOLeaseParams) error
	UpsertAddrEvent(ctx context.Context, arg UpsertAddrEventParams) (int64, error)
	UpsertAsset(ctx context.Context, arg UpsertAssetParams) (int64, error)
	UpsertAssetGroupKey(ctx context.Context, arg UpsertAssetGroupKeyParams) (int64, error)
	UpsertAssetGroupWitness(ctx context.Context, arg UpsertAssetGroupWitnessParams) (int64, error)
	UpsertAssetMeta(ctx context.Context, arg UpsertAssetMetaParams) (int64, error)
	UpsertAssetProofByID(ctx context.Context, arg UpsertAssetProofByIDParams) error
	UpsertAssetWitness(ctx context.Context, arg UpsertAssetWitnessParams) error
	UpsertChainTx(ctx context.Context, arg UpsertChainTxParams) (int64, error)
	UpsertFederationGlobalSyncConfig(ctx context.Context, arg UpsertFederationGlobalSyncConfigParams) error
	UpsertFederationProofSyncLog(ctx context.Context, arg UpsertFederationProofSyncLogParams) (int64, error)
	UpsertFederationUniSyncConfig(ctx context.Context, arg UpsertFederationUniSyncConfigParams) error
	UpsertGenesisAsset(ctx context.Context, arg UpsertGenesisAssetParams) (int64, error)
	UpsertGenesisPoint(ctx context.Context, prevOut []byte) (int64, error)
	UpsertInternalKey(ctx context.Context, arg UpsertInternalKeyParams) (int64, error)
	UpsertManagedUTXO(ctx context.Context, arg UpsertManagedUTXOParams) (int64, error)
	UpsertMultiverseLeaf(ctx context.Context, arg UpsertMultiverseLeafParams) (int64, error)
	UpsertMultiverseRoot(ctx context.Context, arg UpsertMultiverseRootParams) (int64, error)
	UpsertRootNode(ctx context.Context, arg UpsertRootNodeParams) error
	UpsertScriptKey(ctx context.Context, arg UpsertScriptKeyParams) (int64, error)
	UpsertTapscriptTreeEdge(ctx context.Context, arg UpsertTapscriptTreeEdgeParams) (int64, error)
	UpsertTapscriptTreeNode(ctx context.Context, rawNode []byte) (int64, error)
	UpsertTapscriptTreeRootHash(ctx context.Context, arg UpsertTapscriptTreeRootHashParams) (int64, error)
	UpsertUniverseLeaf(ctx context.Context, arg UpsertUniverseLeafParams) error
	UpsertUniverseRoot(ctx context.Context, arg UpsertUniverseRootParams) (int64, error)
}

var _ Querier = (*Queries)(nil)
