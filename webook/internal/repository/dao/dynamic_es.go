package dao

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"strconv"
	"webook/internal/model"
)

type OlivereDynamicEsDao struct {
	esCli *elastic.Client
}

func NewOlivereDynamicEsDao(esCli *elastic.Client) DynamicEsDao {
	return &OlivereDynamicEsDao{
		esCli: esCli,
	}
}

func (dao *OlivereDynamicEsDao) InputDynamic(ctx context.Context, dynamic model.DynamicEs) error {
	_, err := dao.esCli.Index().Index(model.DynamicIndexName).Id(strconv.Itoa(int(dynamic.Id))).BodyJson(dynamic).Do(ctx)
	return err
}

func (dao *OlivereDynamicEsDao) SearchDynamic(ctx context.Context, keyword string, category int8, page, size int) ([]model.DynamicEs, error) {
	titleQuery := elastic.NewMatchQuery("title", keyword).Boost(1)
	contentQuery := elastic.NewMatchQuery("content", keyword).Boost(0.5)
	//or查询
	or := elastic.NewBoolQuery().Should(titleQuery, contentQuery)
	form := (page - 1) * size
	resp, err := dao.esCli.Search(model.DynamicIndexName).Query(or).From(form).Size(size).Do(ctx)
	if err != nil {
		return nil, err
	}
	var res = make([]model.DynamicEs, 0, resp.Hits.TotalHits.Value)
	for _, hit := range resp.Hits.Hits {
		var dnc model.DynamicEs
		err := json.Unmarshal(hit.Source, &dnc)
		if err != nil {
			return nil, err
		}
		res = append(res, dnc)
	}
	return res, nil
}
