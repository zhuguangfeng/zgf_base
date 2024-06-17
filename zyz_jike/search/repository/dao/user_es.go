package dao

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"strconv"
	"strings"
)

const UserIndexName = "user_index"

type User struct {
	Id       int64  `json:"id"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
}

type UserElasticDao struct {
	client *elastic.Client
}

func NewUserElasticDao(client *elastic.Client) UserDao {
	return &UserElasticDao{
		client: client,
	}
}

func (dao *UserElasticDao) InputUser(ctx context.Context, user User) error {
	_, err := dao.client.Index().Index(UserIndexName).Id(strconv.FormatInt(user.Id, 10)).BodyJson(user).Do(ctx)
	return err
}

func (dao *UserElasticDao) Search(ctx context.Context, keywords []string) ([]User, error) {
	queryString := strings.Join(keywords, " ")
	resp, err := dao.client.Search(UserIndexName).Query(elastic.NewMatchQuery("nickname", queryString)).Do(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]User, 0, resp.Hits.TotalHits.Value)
	for _, hit := range resp.Hits.Hits {
		var user User
		if err = json.Unmarshal(hit.Source, &user); err != nil {
			return nil, err
		}
		res = append(res, user)
	}
	return res, nil
}
