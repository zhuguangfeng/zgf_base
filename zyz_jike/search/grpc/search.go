package grpc

import (
	"context"
	"github.com/ecodeclub/ekit/slice"
	"google.golang.org/grpc"
	searchv1 "zyz_jike/api/proto/gen/search/v1"
	"zyz_jike/search/domain"
	"zyz_jike/search/service"
)

type SearchServiceServer struct {
	searchv1.UnimplementedSearchServiceServer
	svc service.SearchService
}

func NewSearchService(svc service.SearchService) *SearchServiceServer {
	return &SearchServiceServer{
		svc: svc,
	}
}

func (s *SearchServiceServer) Register(server grpc.ServiceRegistrar) {
	searchv1.RegisterSearchServiceServer(server, s)
}

func (s *SearchServiceServer) Search(ctx context.Context, req *searchv1.SearchRequest) (*searchv1.SearchResponse, error) {
	resp, err := s.svc.Search(ctx, req.GetUid(), req.Expression)
	if err != nil {
		return nil, err
	}
	return &searchv1.SearchResponse{
		User: &searchv1.UserResult{
			Users: slice.Map(resp.Users, func(idx int, src domain.User) *searchv1.User {
				return &searchv1.User{
					Id:       src.Id,
					Nickname: src.Nickname,
					Phone:    src.Phone,
				}
			}),
		},
		Article: &searchv1.ArticleResult{
			Articles: slice.Map(resp.Articles, func(idx int, src domain.Article) *searchv1.Article {
				return &searchv1.Article{
					Id: src.Id,

					Title:           src.Title,
					Content:         src.Content,
					Status:          int32(src.Status),
					Pic:             src.Pic,
					Pics:            src.Pics,
					Category:        int32(src.Category),
					ArticleCategory: int32(src.ArticleCategory),
					RichText:        src.RichText,
					CreatedAt:       src.CreatedAt,
					Uid:             src.Uid,
				}
			}),
		},
	}, nil
}
