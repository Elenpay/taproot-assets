// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package universerpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UniverseClient is the client API for Universe service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UniverseClient interface {
	// tapcli: `universe multiverse`
	// MultiverseRoot returns the root of the multiverse tree. This is useful to
	// determine the equality of two multiverse trees, since the root can directly
	// be compared to another multiverse root to find out if a sync is required.
	MultiverseRoot(ctx context.Context, in *MultiverseRootRequest, opts ...grpc.CallOption) (*MultiverseRootResponse, error)
	// tapcli: `universe roots`
	// AssetRoots queries for the known Universe roots associated with each known
	// asset. These roots represent the supply/audit state for each known asset.
	AssetRoots(ctx context.Context, in *AssetRootRequest, opts ...grpc.CallOption) (*AssetRootResponse, error)
	// tapcli: `universe roots`
	// QueryAssetRoots attempts to locate the current Universe root for a specific
	// asset. This asset can be identified by its asset ID or group key.
	QueryAssetRoots(ctx context.Context, in *AssetRootQuery, opts ...grpc.CallOption) (*QueryRootResponse, error)
	// tapcli: `universe delete`
	// DeleteAssetRoot deletes the Universe root for a specific asset, including
	// all asoociated universe keys, leaves, and events.
	DeleteAssetRoot(ctx context.Context, in *DeleteRootQuery, opts ...grpc.CallOption) (*DeleteRootResponse, error)
	// tapcli: `universe keys`
	// AssetLeafKeys queries for the set of Universe keys associated with a given
	// asset_id or group_key. Each key takes the form: (outpoint, script_key),
	// where outpoint is an outpoint in the Bitcoin blockchain that anchors a
	// valid Taproot Asset commitment, and script_key is the script_key of
	// the asset within the Taproot Asset commitment for the given asset_id or
	// group_key.
	AssetLeafKeys(ctx context.Context, in *AssetLeafKeysRequest, opts ...grpc.CallOption) (*AssetLeafKeyResponse, error)
	// tapcli: `universe leaves`
	// AssetLeaves queries for the set of asset leaves (the values in the Universe
	// MS-SMT tree) for a given asset_id or group_key. These represents either
	// asset issuance events (they have a genesis witness) or asset transfers that
	// took place on chain. The leaves contain a normal Taproot Asset proof, as
	// well as details for the asset.
	AssetLeaves(ctx context.Context, in *ID, opts ...grpc.CallOption) (*AssetLeafResponse, error)
	// tapcli: `universe proofs query`
	// QueryProof attempts to query for an issuance or transfer proof for a given
	// asset based on its UniverseKey. A UniverseKey is composed of the Universe
	// ID (asset_id/group_key) and also a leaf key (outpoint || script_key). If
	// found, then the issuance proof is returned that includes an inclusion proof
	// to the known Universe root, as well as a Taproot Asset state transition or
	// issuance proof for the said asset.
	QueryProof(ctx context.Context, in *UniverseKey, opts ...grpc.CallOption) (*AssetProofResponse, error)
	// tapcli: `universe proofs insert`
	// InsertProof attempts to insert a new issuance or transfer proof into the
	// Universe tree specified by the UniverseKey. If valid, then the proof is
	// inserted into the database, with a new Universe root returned for the
	// updated asset_id/group_key.
	InsertProof(ctx context.Context, in *AssetProof, opts ...grpc.CallOption) (*AssetProofResponse, error)
	// tapcli: `universe info`
	// Info returns a set of information about the current state of the Universe.
	Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error)
	// tapcli: `universe sync`
	// SyncUniverse takes host information for a remote Universe server, then
	// attempts to synchronize either only the set of specified asset_ids, or all
	// assets if none are specified. The sync process will attempt to query for
	// the latest known root for each asset, performing tree based reconciliation
	// to arrive at a new shared root.
	SyncUniverse(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncResponse, error)
	// tapcli: `universe federation list`
	// ListFederationServers lists the set of servers that make up the federation
	// of the local Universe server. This servers are used to push out new proofs,
	// and also periodically call sync new proofs from the remote server.
	ListFederationServers(ctx context.Context, in *ListFederationServersRequest, opts ...grpc.CallOption) (*ListFederationServersResponse, error)
	// tapcli: `universe federation add`
	// AddFederationServer adds a new server to the federation of the local
	// Universe server. Once a server is added, this call can also optionally be
	// used to trigger a sync of the remote server.
	AddFederationServer(ctx context.Context, in *AddFederationServerRequest, opts ...grpc.CallOption) (*AddFederationServerResponse, error)
	// tapcli: `universe federation del`
	// DeleteFederationServer removes a server from the federation of the local
	// Universe server.
	DeleteFederationServer(ctx context.Context, in *DeleteFederationServerRequest, opts ...grpc.CallOption) (*DeleteFederationServerResponse, error)
	// tapcli: `universe stats`
	// UniverseStats returns a set of aggregate statistics for the current state
	// of the Universe. Stats returned include: total number of syncs, total
	// number of proofs, and total number of known assets.
	UniverseStats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (*StatsResponse, error)
	// tapcli `universe stats assets`
	// QueryAssetStats returns a set of statistics for a given set of assets.
	// Stats can be queried for all assets, or based on the: asset ID, name, or
	// asset type. Pagination is supported via the offset and limit params.
	// Results can also be sorted based on any of the main query params.
	QueryAssetStats(ctx context.Context, in *AssetStatsQuery, opts ...grpc.CallOption) (*UniverseAssetStats, error)
	// tapcli `universe stats events`
	// QueryEvents returns the number of sync and proof events for a given time
	// period, grouped by day.
	QueryEvents(ctx context.Context, in *QueryEventsRequest, opts ...grpc.CallOption) (*QueryEventsResponse, error)
	// SetFederationSyncConfig sets the configuration of the universe federation
	// sync.
	SetFederationSyncConfig(ctx context.Context, in *SetFederationSyncConfigRequest, opts ...grpc.CallOption) (*SetFederationSyncConfigResponse, error)
	// tapcli: `universe federation config info`
	// QueryFederationSyncConfig queries the universe federation sync configuration
	// settings.
	QueryFederationSyncConfig(ctx context.Context, in *QueryFederationSyncConfigRequest, opts ...grpc.CallOption) (*QueryFederationSyncConfigResponse, error)
}

type universeClient struct {
	cc grpc.ClientConnInterface
}

func NewUniverseClient(cc grpc.ClientConnInterface) UniverseClient {
	return &universeClient{cc}
}

func (c *universeClient) MultiverseRoot(ctx context.Context, in *MultiverseRootRequest, opts ...grpc.CallOption) (*MultiverseRootResponse, error) {
	out := new(MultiverseRootResponse)
	err := c.cc.Invoke(ctx, "/universerpc.Universe/MultiverseRoot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universeClient) AssetRoots(ctx context.Context, in *AssetRootRequest, opts ...grpc.CallOption) (*AssetRootResponse, error) {
	out := new(AssetRootResponse)
	err := c.cc.Invoke(ctx, "/universerpc.Universe/AssetRoots", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universeClient) QueryAssetRoots(ctx context.Context, in *AssetRootQuery, opts ...grpc.CallOption) (*QueryRootResponse, error) {
	out := new(QueryRootResponse)
	err := c.cc.Invoke(ctx, "/universerpc.Universe/QueryAssetRoots", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universeClient) DeleteAssetRoot(ctx context.Context, in *DeleteRootQuery, opts ...grpc.CallOption) (*DeleteRootResponse, error) {
	out := new(DeleteRootResponse)
	err := c.cc.Invoke(ctx, "/universerpc.Universe/DeleteAssetRoot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universeClient) AssetLeafKeys(ctx context.Context, in *AssetLeafKeysRequest, opts ...grpc.CallOption) (*AssetLeafKeyResponse, error) {
	out := new(AssetLeafKeyResponse)
	err := c.cc.Invoke(ctx, "/universerpc.Universe/AssetLeafKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universeClient) AssetLeaves(ctx context.Context, in *ID, opts ...grpc.CallOption) (*AssetLeafResponse, error) {
	out := new(AssetLeafResponse)
	err := c.cc.Invoke(ctx, "/universerpc.Universe/AssetLeaves", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universeClient) QueryProof(ctx context.Context, in *UniverseKey, opts ...grpc.CallOption) (*AssetProofResponse, error) {
	out := new(AssetProofResponse)
	err := c.cc.Invoke(ctx, "/universerpc.Universe/QueryProof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universeClient) InsertProof(ctx context.Context, in *AssetProof, opts ...grpc.CallOption) (*AssetProofResponse, error) {
	out := new(AssetProofResponse)
	err := c.cc.Invoke(ctx, "/universerpc.Universe/InsertProof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universeClient) Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := c.cc.Invoke(ctx, "/universerpc.Universe/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universeClient) SyncUniverse(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncResponse, error) {
	out := new(SyncResponse)
	err := c.cc.Invoke(ctx, "/universerpc.Universe/SyncUniverse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universeClient) ListFederationServers(ctx context.Context, in *ListFederationServersRequest, opts ...grpc.CallOption) (*ListFederationServersResponse, error) {
	out := new(ListFederationServersResponse)
	err := c.cc.Invoke(ctx, "/universerpc.Universe/ListFederationServers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universeClient) AddFederationServer(ctx context.Context, in *AddFederationServerRequest, opts ...grpc.CallOption) (*AddFederationServerResponse, error) {
	out := new(AddFederationServerResponse)
	err := c.cc.Invoke(ctx, "/universerpc.Universe/AddFederationServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universeClient) DeleteFederationServer(ctx context.Context, in *DeleteFederationServerRequest, opts ...grpc.CallOption) (*DeleteFederationServerResponse, error) {
	out := new(DeleteFederationServerResponse)
	err := c.cc.Invoke(ctx, "/universerpc.Universe/DeleteFederationServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universeClient) UniverseStats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (*StatsResponse, error) {
	out := new(StatsResponse)
	err := c.cc.Invoke(ctx, "/universerpc.Universe/UniverseStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universeClient) QueryAssetStats(ctx context.Context, in *AssetStatsQuery, opts ...grpc.CallOption) (*UniverseAssetStats, error) {
	out := new(UniverseAssetStats)
	err := c.cc.Invoke(ctx, "/universerpc.Universe/QueryAssetStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universeClient) QueryEvents(ctx context.Context, in *QueryEventsRequest, opts ...grpc.CallOption) (*QueryEventsResponse, error) {
	out := new(QueryEventsResponse)
	err := c.cc.Invoke(ctx, "/universerpc.Universe/QueryEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universeClient) SetFederationSyncConfig(ctx context.Context, in *SetFederationSyncConfigRequest, opts ...grpc.CallOption) (*SetFederationSyncConfigResponse, error) {
	out := new(SetFederationSyncConfigResponse)
	err := c.cc.Invoke(ctx, "/universerpc.Universe/SetFederationSyncConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universeClient) QueryFederationSyncConfig(ctx context.Context, in *QueryFederationSyncConfigRequest, opts ...grpc.CallOption) (*QueryFederationSyncConfigResponse, error) {
	out := new(QueryFederationSyncConfigResponse)
	err := c.cc.Invoke(ctx, "/universerpc.Universe/QueryFederationSyncConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UniverseServer is the server API for Universe service.
// All implementations must embed UnimplementedUniverseServer
// for forward compatibility
type UniverseServer interface {
	// tapcli: `universe multiverse`
	// MultiverseRoot returns the root of the multiverse tree. This is useful to
	// determine the equality of two multiverse trees, since the root can directly
	// be compared to another multiverse root to find out if a sync is required.
	MultiverseRoot(context.Context, *MultiverseRootRequest) (*MultiverseRootResponse, error)
	// tapcli: `universe roots`
	// AssetRoots queries for the known Universe roots associated with each known
	// asset. These roots represent the supply/audit state for each known asset.
	AssetRoots(context.Context, *AssetRootRequest) (*AssetRootResponse, error)
	// tapcli: `universe roots`
	// QueryAssetRoots attempts to locate the current Universe root for a specific
	// asset. This asset can be identified by its asset ID or group key.
	QueryAssetRoots(context.Context, *AssetRootQuery) (*QueryRootResponse, error)
	// tapcli: `universe delete`
	// DeleteAssetRoot deletes the Universe root for a specific asset, including
	// all asoociated universe keys, leaves, and events.
	DeleteAssetRoot(context.Context, *DeleteRootQuery) (*DeleteRootResponse, error)
	// tapcli: `universe keys`
	// AssetLeafKeys queries for the set of Universe keys associated with a given
	// asset_id or group_key. Each key takes the form: (outpoint, script_key),
	// where outpoint is an outpoint in the Bitcoin blockchain that anchors a
	// valid Taproot Asset commitment, and script_key is the script_key of
	// the asset within the Taproot Asset commitment for the given asset_id or
	// group_key.
	AssetLeafKeys(context.Context, *AssetLeafKeysRequest) (*AssetLeafKeyResponse, error)
	// tapcli: `universe leaves`
	// AssetLeaves queries for the set of asset leaves (the values in the Universe
	// MS-SMT tree) for a given asset_id or group_key. These represents either
	// asset issuance events (they have a genesis witness) or asset transfers that
	// took place on chain. The leaves contain a normal Taproot Asset proof, as
	// well as details for the asset.
	AssetLeaves(context.Context, *ID) (*AssetLeafResponse, error)
	// tapcli: `universe proofs query`
	// QueryProof attempts to query for an issuance or transfer proof for a given
	// asset based on its UniverseKey. A UniverseKey is composed of the Universe
	// ID (asset_id/group_key) and also a leaf key (outpoint || script_key). If
	// found, then the issuance proof is returned that includes an inclusion proof
	// to the known Universe root, as well as a Taproot Asset state transition or
	// issuance proof for the said asset.
	QueryProof(context.Context, *UniverseKey) (*AssetProofResponse, error)
	// tapcli: `universe proofs insert`
	// InsertProof attempts to insert a new issuance or transfer proof into the
	// Universe tree specified by the UniverseKey. If valid, then the proof is
	// inserted into the database, with a new Universe root returned for the
	// updated asset_id/group_key.
	InsertProof(context.Context, *AssetProof) (*AssetProofResponse, error)
	// tapcli: `universe info`
	// Info returns a set of information about the current state of the Universe.
	Info(context.Context, *InfoRequest) (*InfoResponse, error)
	// tapcli: `universe sync`
	// SyncUniverse takes host information for a remote Universe server, then
	// attempts to synchronize either only the set of specified asset_ids, or all
	// assets if none are specified. The sync process will attempt to query for
	// the latest known root for each asset, performing tree based reconciliation
	// to arrive at a new shared root.
	SyncUniverse(context.Context, *SyncRequest) (*SyncResponse, error)
	// tapcli: `universe federation list`
	// ListFederationServers lists the set of servers that make up the federation
	// of the local Universe server. This servers are used to push out new proofs,
	// and also periodically call sync new proofs from the remote server.
	ListFederationServers(context.Context, *ListFederationServersRequest) (*ListFederationServersResponse, error)
	// tapcli: `universe federation add`
	// AddFederationServer adds a new server to the federation of the local
	// Universe server. Once a server is added, this call can also optionally be
	// used to trigger a sync of the remote server.
	AddFederationServer(context.Context, *AddFederationServerRequest) (*AddFederationServerResponse, error)
	// tapcli: `universe federation del`
	// DeleteFederationServer removes a server from the federation of the local
	// Universe server.
	DeleteFederationServer(context.Context, *DeleteFederationServerRequest) (*DeleteFederationServerResponse, error)
	// tapcli: `universe stats`
	// UniverseStats returns a set of aggregate statistics for the current state
	// of the Universe. Stats returned include: total number of syncs, total
	// number of proofs, and total number of known assets.
	UniverseStats(context.Context, *StatsRequest) (*StatsResponse, error)
	// tapcli `universe stats assets`
	// QueryAssetStats returns a set of statistics for a given set of assets.
	// Stats can be queried for all assets, or based on the: asset ID, name, or
	// asset type. Pagination is supported via the offset and limit params.
	// Results can also be sorted based on any of the main query params.
	QueryAssetStats(context.Context, *AssetStatsQuery) (*UniverseAssetStats, error)
	// tapcli `universe stats events`
	// QueryEvents returns the number of sync and proof events for a given time
	// period, grouped by day.
	QueryEvents(context.Context, *QueryEventsRequest) (*QueryEventsResponse, error)
	// SetFederationSyncConfig sets the configuration of the universe federation
	// sync.
	SetFederationSyncConfig(context.Context, *SetFederationSyncConfigRequest) (*SetFederationSyncConfigResponse, error)
	// tapcli: `universe federation config info`
	// QueryFederationSyncConfig queries the universe federation sync configuration
	// settings.
	QueryFederationSyncConfig(context.Context, *QueryFederationSyncConfigRequest) (*QueryFederationSyncConfigResponse, error)
	mustEmbedUnimplementedUniverseServer()
}

// UnimplementedUniverseServer must be embedded to have forward compatible implementations.
type UnimplementedUniverseServer struct {
}

func (UnimplementedUniverseServer) MultiverseRoot(context.Context, *MultiverseRootRequest) (*MultiverseRootResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MultiverseRoot not implemented")
}
func (UnimplementedUniverseServer) AssetRoots(context.Context, *AssetRootRequest) (*AssetRootResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssetRoots not implemented")
}
func (UnimplementedUniverseServer) QueryAssetRoots(context.Context, *AssetRootQuery) (*QueryRootResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryAssetRoots not implemented")
}
func (UnimplementedUniverseServer) DeleteAssetRoot(context.Context, *DeleteRootQuery) (*DeleteRootResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAssetRoot not implemented")
}
func (UnimplementedUniverseServer) AssetLeafKeys(context.Context, *AssetLeafKeysRequest) (*AssetLeafKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssetLeafKeys not implemented")
}
func (UnimplementedUniverseServer) AssetLeaves(context.Context, *ID) (*AssetLeafResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssetLeaves not implemented")
}
func (UnimplementedUniverseServer) QueryProof(context.Context, *UniverseKey) (*AssetProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryProof not implemented")
}
func (UnimplementedUniverseServer) InsertProof(context.Context, *AssetProof) (*AssetProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertProof not implemented")
}
func (UnimplementedUniverseServer) Info(context.Context, *InfoRequest) (*InfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedUniverseServer) SyncUniverse(context.Context, *SyncRequest) (*SyncResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncUniverse not implemented")
}
func (UnimplementedUniverseServer) ListFederationServers(context.Context, *ListFederationServersRequest) (*ListFederationServersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFederationServers not implemented")
}
func (UnimplementedUniverseServer) AddFederationServer(context.Context, *AddFederationServerRequest) (*AddFederationServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFederationServer not implemented")
}
func (UnimplementedUniverseServer) DeleteFederationServer(context.Context, *DeleteFederationServerRequest) (*DeleteFederationServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFederationServer not implemented")
}
func (UnimplementedUniverseServer) UniverseStats(context.Context, *StatsRequest) (*StatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UniverseStats not implemented")
}
func (UnimplementedUniverseServer) QueryAssetStats(context.Context, *AssetStatsQuery) (*UniverseAssetStats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryAssetStats not implemented")
}
func (UnimplementedUniverseServer) QueryEvents(context.Context, *QueryEventsRequest) (*QueryEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryEvents not implemented")
}
func (UnimplementedUniverseServer) SetFederationSyncConfig(context.Context, *SetFederationSyncConfigRequest) (*SetFederationSyncConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetFederationSyncConfig not implemented")
}
func (UnimplementedUniverseServer) QueryFederationSyncConfig(context.Context, *QueryFederationSyncConfigRequest) (*QueryFederationSyncConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryFederationSyncConfig not implemented")
}
func (UnimplementedUniverseServer) mustEmbedUnimplementedUniverseServer() {}

// UnsafeUniverseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UniverseServer will
// result in compilation errors.
type UnsafeUniverseServer interface {
	mustEmbedUnimplementedUniverseServer()
}

func RegisterUniverseServer(s grpc.ServiceRegistrar, srv UniverseServer) {
	s.RegisterService(&Universe_ServiceDesc, srv)
}

func _Universe_MultiverseRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiverseRootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniverseServer).MultiverseRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/universerpc.Universe/MultiverseRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniverseServer).MultiverseRoot(ctx, req.(*MultiverseRootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Universe_AssetRoots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssetRootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniverseServer).AssetRoots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/universerpc.Universe/AssetRoots",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniverseServer).AssetRoots(ctx, req.(*AssetRootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Universe_QueryAssetRoots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssetRootQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniverseServer).QueryAssetRoots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/universerpc.Universe/QueryAssetRoots",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniverseServer).QueryAssetRoots(ctx, req.(*AssetRootQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Universe_DeleteAssetRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRootQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniverseServer).DeleteAssetRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/universerpc.Universe/DeleteAssetRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniverseServer).DeleteAssetRoot(ctx, req.(*DeleteRootQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Universe_AssetLeafKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssetLeafKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniverseServer).AssetLeafKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/universerpc.Universe/AssetLeafKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniverseServer).AssetLeafKeys(ctx, req.(*AssetLeafKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Universe_AssetLeaves_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniverseServer).AssetLeaves(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/universerpc.Universe/AssetLeaves",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniverseServer).AssetLeaves(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Universe_QueryProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UniverseKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniverseServer).QueryProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/universerpc.Universe/QueryProof",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniverseServer).QueryProof(ctx, req.(*UniverseKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _Universe_InsertProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssetProof)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniverseServer).InsertProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/universerpc.Universe/InsertProof",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniverseServer).InsertProof(ctx, req.(*AssetProof))
	}
	return interceptor(ctx, in, info, handler)
}

func _Universe_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniverseServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/universerpc.Universe/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniverseServer).Info(ctx, req.(*InfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Universe_SyncUniverse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniverseServer).SyncUniverse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/universerpc.Universe/SyncUniverse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniverseServer).SyncUniverse(ctx, req.(*SyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Universe_ListFederationServers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFederationServersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniverseServer).ListFederationServers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/universerpc.Universe/ListFederationServers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniverseServer).ListFederationServers(ctx, req.(*ListFederationServersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Universe_AddFederationServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFederationServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniverseServer).AddFederationServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/universerpc.Universe/AddFederationServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniverseServer).AddFederationServer(ctx, req.(*AddFederationServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Universe_DeleteFederationServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFederationServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniverseServer).DeleteFederationServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/universerpc.Universe/DeleteFederationServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniverseServer).DeleteFederationServer(ctx, req.(*DeleteFederationServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Universe_UniverseStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniverseServer).UniverseStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/universerpc.Universe/UniverseStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniverseServer).UniverseStats(ctx, req.(*StatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Universe_QueryAssetStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssetStatsQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniverseServer).QueryAssetStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/universerpc.Universe/QueryAssetStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniverseServer).QueryAssetStats(ctx, req.(*AssetStatsQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Universe_QueryEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniverseServer).QueryEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/universerpc.Universe/QueryEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniverseServer).QueryEvents(ctx, req.(*QueryEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Universe_SetFederationSyncConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetFederationSyncConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniverseServer).SetFederationSyncConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/universerpc.Universe/SetFederationSyncConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniverseServer).SetFederationSyncConfig(ctx, req.(*SetFederationSyncConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Universe_QueryFederationSyncConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFederationSyncConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniverseServer).QueryFederationSyncConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/universerpc.Universe/QueryFederationSyncConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniverseServer).QueryFederationSyncConfig(ctx, req.(*QueryFederationSyncConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Universe_ServiceDesc is the grpc.ServiceDesc for Universe service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Universe_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "universerpc.Universe",
	HandlerType: (*UniverseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MultiverseRoot",
			Handler:    _Universe_MultiverseRoot_Handler,
		},
		{
			MethodName: "AssetRoots",
			Handler:    _Universe_AssetRoots_Handler,
		},
		{
			MethodName: "QueryAssetRoots",
			Handler:    _Universe_QueryAssetRoots_Handler,
		},
		{
			MethodName: "DeleteAssetRoot",
			Handler:    _Universe_DeleteAssetRoot_Handler,
		},
		{
			MethodName: "AssetLeafKeys",
			Handler:    _Universe_AssetLeafKeys_Handler,
		},
		{
			MethodName: "AssetLeaves",
			Handler:    _Universe_AssetLeaves_Handler,
		},
		{
			MethodName: "QueryProof",
			Handler:    _Universe_QueryProof_Handler,
		},
		{
			MethodName: "InsertProof",
			Handler:    _Universe_InsertProof_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _Universe_Info_Handler,
		},
		{
			MethodName: "SyncUniverse",
			Handler:    _Universe_SyncUniverse_Handler,
		},
		{
			MethodName: "ListFederationServers",
			Handler:    _Universe_ListFederationServers_Handler,
		},
		{
			MethodName: "AddFederationServer",
			Handler:    _Universe_AddFederationServer_Handler,
		},
		{
			MethodName: "DeleteFederationServer",
			Handler:    _Universe_DeleteFederationServer_Handler,
		},
		{
			MethodName: "UniverseStats",
			Handler:    _Universe_UniverseStats_Handler,
		},
		{
			MethodName: "QueryAssetStats",
			Handler:    _Universe_QueryAssetStats_Handler,
		},
		{
			MethodName: "QueryEvents",
			Handler:    _Universe_QueryEvents_Handler,
		},
		{
			MethodName: "SetFederationSyncConfig",
			Handler:    _Universe_SetFederationSyncConfig_Handler,
		},
		{
			MethodName: "QueryFederationSyncConfig",
			Handler:    _Universe_QueryFederationSyncConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "universerpc/universe.proto",
}
