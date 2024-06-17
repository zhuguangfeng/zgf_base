package grpc

import searchv1 "zyz_jike/api/proto/gen/search/v1"

type SyncServiceServer struct {
	searchv1.UnimplementedSyncServiceServer
	syncSvc service.SyncService
}
