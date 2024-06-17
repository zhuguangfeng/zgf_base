package dao

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"strconv"
	"strings"
)

const ArticleIndexName = "article_index"
const TagIndexName = "tags_index"

type Article struct {
	Id              int64    `json:"id"`
	Uid             int64    `json:"uid"`
	Category        int8     `json:"category"`
	ArticleCategory int8     `json:"article_category"`
	Title           string   `json:"title"`
	Content         string   `json:"content"`
	RichText        string   `json:"rich_text"`
	Pic             string   `json:"pic"`
	Pics            []string `json:"pics"`
	Status          int8     `json:"status"`
}

type ArticleElasticDao struct {
	client *elastic.Client
}

func NewArticleElasticDao(client *elastic.Client) ArticleDao {
	return &ArticleElasticDao{
		client: client,
	}
}

func (dao *ArticleElasticDao) InputArticle(ctx context.Context, article Article) error {
	_, err := dao.client.Index().Index(ArticleIndexName).Id(strconv.FormatInt(article.Id, 10)).BodyJson(article).Do(ctx)
	return err
}

func (dao *ArticleElasticDao) Search(ctx context.Context, tagArtIds []int64, keywords []string) ([]Article, error) {
	queryString := strings.Join(keywords, " ")

	status := elastic.NewTermsQuery("status", 2)

	title := elastic.NewMatchQuery("title", queryString)
	content := elastic.NewMatchQuery("content", queryString)

	//或查询
	or := elastic.NewBoolQuery().Should(title, content)
	//and查询
	and := elastic.NewBoolQuery().Must(status, or)

	resp, err := dao.client.Search(ArticleIndexName).Query(and).Do(ctx)
	if err != nil {
		return nil, err
	}

	var res []Article

	for _, hit := range resp.Hits.Hits {
		var art Article
		err := json.Unmarshal(hit.Source, &art)
		if err != nil {
			return nil, err
		}
		res = append(res, art)
	}
	return res, nil
}
