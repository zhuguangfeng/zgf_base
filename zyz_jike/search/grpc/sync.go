package grpc

import (
	"context"
	"google.golang.org/grpc"
	searchv1 "zyz_jike/api/proto/gen/search/v1"
	"zyz_jike/search/domain"
	"zyz_jike/search/service"
)

type SyncServiceServer struct {
	searchv1.UnimplementedSyncServiceServer
	syncSvc service.SyncService
}

func NewSyncServiceServer(syncSvc service.SyncService) *SyncServiceServer {
	return &SyncServiceServer{
		syncSvc: syncSvc,
	}
}

func (s *SyncServiceServer) Register(server grpc.ServiceRegistrar) {
	searchv1.RegisterSyncServiceServer(server, s)
}

func (s *SyncServiceServer) InputUser(ctx context.Context, request *searchv1.InputUserRequest) (*searchv1.InputUserResponse, error) {
	err := s.syncSvc.InputUser(ctx, s.toDomainUser(request.GetUser()))
	return &searchv1.InputUserResponse{}, err
}

func (s *SyncServiceServer) InputArticle(ctx context.Context, request *searchv1.InputArticleRequest) (*searchv1.InputArticleResponse, error) {
	err := s.syncSvc.InputArticle(ctx, s.toDomainArticle(request.GetArticle()))
	return &searchv1.InputArticleResponse{}, err
}

func (s *SyncServiceServer) toDomainUser(vuser *searchv1.User) domain.User {
	return domain.User{
		Id:       vuser.Id,
		Nickname: vuser.Nickname,
		Phone:    vuser.Phone,
	}
}
func (s *SyncServiceServer) toDomainArticle(art *searchv1.Article) domain.Article {
	return domain.Article{
		Id:              art.Id,
		Uid:             art.Uid,
		Category:        int8(art.Category),
		ArticleCategory: int8(art.ArticleCategory),
		Title:           art.Title,
		Content:         art.Content,
		RichText:        art.RichText,
		Pic:             art.Pic,
		Pics:            art.Pics,
		Status:          int8(art.Status),
	}
}
