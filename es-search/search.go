package es_search

import (
	"encoding/json"
	"github.com/beego/beego/v2/client/httplib"
)

var ESUrl string

func init() {
	ESUrl = "127.0.0.1:9200/"
}

// EsSearch 搜索
// indexName 库的名称  fields,要搜索的字段名称 ，keyword 搜索关键字 limit 分页字段  page展示几条
func GetEsSearch(indexName string, limit, page int, query map[string]interface{}) HitsData {
	from := limit
	size := page
	sort := []interface{}{}
	aggs := map[string]interface{}{}
	// 传递 搜索条件主题
	body := map[string]interface{}{
		"query": query,
		"from":  from,
		"size":  size,
		"sort":  sort,
		"aggs":  aggs,
	}
	req := httplib.Post(ESUrl + indexName + "/_search/")
	// 解析body
	req.JSONBody(body)
	s, _ := req.String()
	// 实例化一个结构体
	var str ReqSearchData
	// 将结果集转换成 结构体
	json.Unmarshal([]byte(s), &str)
	return str.Hits
	// 获得数据后按照以下方式输出
	//var data ResData
	//var datalist []ResData
	//
	//for _, v := range hits.Hits {
	//	data.Highlight = v.Highlight
	//	json.Unmarshal(v.Source, &data)
	//	datalist = append(datalist, data)
	//}
	//c.Data["json"] = map[string]interface{}{"code": 0, "data": datalist}
	//c.ServeJSON()
	//return
}

// 封装一个es搜索的方法
func EsSearch(indexName string, limit, page int, query map[string]interface{}) HitsData {
	return GetEsSearch(indexName, limit, page, query)
}
