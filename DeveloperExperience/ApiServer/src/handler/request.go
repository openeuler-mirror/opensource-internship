package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/elastic/go-elasticsearch/v6"
)

func InitElasticSearch() *elasticsearch.Client {
	cfg := elasticsearch.Config{
		Addresses: []string{ConfigContent.ElasticSearch.Url},
		Username:  ConfigContent.ElasticSearch.Username,
		Password:  ConfigContent.ElasticSearch.Password,
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	return es
}

func ElasticSearchQuery(es *elasticsearch.Client, index string, body *bytes.Buffer) map[string]interface{} {
	var (
		r map[string]interface{}
	)
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(index),
		es.Search.WithBody(body),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
		es.Search.WithIgnoreUnavailable(true),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	return r
}

func GetElasticSearchVersion() VersionResult {
	var (
		r map[string]interface{}
	)
	res, err := InitElasticSearch().Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	var VersionResult VersionResult
	VersionResult.Client = elasticsearch.Version
	VersionResult.Server = r["version"].(map[string]interface{})["number"].(string)
	return VersionResult
}

func GetIssueStateDistribution() GetIssueStateDistributionData {
	body := &bytes.Buffer{}
	body.WriteString(`{"size":0,"aggs":{"state":{"terms":{"field":"state"}}}}`)
	index := "gitee_issues-enriched"
	r := ElasticSearchQuery(InitElasticSearch(), index, body)
	ResultJsonByte, _ := json.Marshal(r)
	var QueryResult GetIssueStateDistributionQueryResult
	json.Unmarshal(ResultJsonByte, &QueryResult)
	Buckets := QueryResult.Aggregations.State.Buckets
	var Data GetIssueStateDistributionData
	for i := 0; i < len(Buckets); i++ {
		Data.Legend = append(Data.Legend, Buckets[i].Key)
		Data.Series = append(Data.Series, struct {
			Value int    "json:\"value\""
			Name  string "json:\"name\""
		}{Buckets[i].DocCount, Buckets[i].Key})
	}
	return Data
}

func GetCurrentStatusOfHistoricalIssue() GetCurrentStatusOfHistoricalIssueData {
	body := &bytes.Buffer{}
	body.WriteString(`{"aggs":{"2":{"date_histogram":{"field":"created_at","interval":"1w","time_zone":"Asia/Shanghai","min_doc_count":1},"aggs":{"3":{"terms":{"field":"state","size":1000,"order":{"_count":"desc"}}}}}},"size":0,"script_fields":{},"docvalue_fields":[{"field":"closed_at","format":"date_time"},{"field":"created_at","format":"date_time"},{"field":"grimoire_creation_date","format":"date_time"},{"field":"metadata__enriched_on","format":"date_time"},{"field":"metadata__timestamp","format":"date_time"},{"field":"metadata__updated_on","format":"date_time"},{"field":"updated_at","format":"date_time"}],"query":{"bool":{"must":[{"query_string":{"query":"pull_request:false","analyze_wildcard":true,"default_field":"*"}},{"query_string":{"query":"pull_request:false","analyze_wildcard":true,"default_field":"*"}},{"range":{"metadata__timestamp":{"gte":1635145582100,"lte":1642921582100,"format":"epoch_millis"}}}],"filter":[],"should":[],"must_not":[]}},"timeout":"30000ms"}`)
	index := "gitee_issues-enriched"
	r := ElasticSearchQuery(InitElasticSearch(), index, body)
	ResultJsonByte, _ := json.Marshal(r)
	var QueryResult GetCurrentStatusOfHistoricalIssueQueryResult
	json.Unmarshal(ResultJsonByte, &QueryResult)
	var Data GetCurrentStatusOfHistoricalIssueData
	Buckets := QueryResult.Aggregations.Num2.Buckets
	for i := 0; i < len(Buckets); i++ {
		Data.Date = append(Data.Date, time.Time(Buckets[i].KeyAsString).Format("2006-01-02"))
		Closed := 0
		Open := 0
		Rejected := 0
		Progressing := 0
		for j := 0; j < len(Buckets[i].Num3.Buckets); j++ {
			if Buckets[i].Num3.Buckets[j].Key == "closed" {
				Closed = Buckets[i].Num3.Buckets[j].DocCount
			} else if Buckets[i].Num3.Buckets[j].Key == "open" {
				Open = Buckets[i].Num3.Buckets[j].DocCount
			} else if Buckets[i].Num3.Buckets[j].Key == "rejected" {
				Rejected = Buckets[i].Num3.Buckets[j].DocCount
			} else if Buckets[i].Num3.Buckets[j].Key == "progressing" {
				Progressing = Buckets[i].Num3.Buckets[j].DocCount
			}
		}
		Data.Closed = append(Data.Closed, Closed)
		Data.Open = append(Data.Open, Open)
		Data.Rejected = append(Data.Rejected, Rejected)
		Data.Progressing = append(Data.Progressing, Progressing)
	}
	return Data
}

func GetIssueOpenDaysDistribution() map[string]interface{} {
	body := &bytes.Buffer{}
	body.WriteString(`{"aggs":{"2":{"range":{"field":"time_open_days","ranges":[{"from":0,"to":1},{"from":1,"to":3},{"from":3,"to":7},{"from":7,"to":14},{"from":14,"to":30},{"from":30,"to":180},{"from":180,"to":360},{"from":360}],"keyed":true}}},"size":0,"_source":{"excludes":[]},"stored_fields":["*"],"script_fields":{},"docvalue_fields":[{"field":"closed_at","format":"date_time"},{"field":"created_at","format":"date_time"},{"field":"grimoire_creation_date","format":"date_time"},{"field":"metadata__enriched_on","format":"date_time"},{"field":"metadata__timestamp","format":"date_time"},{"field":"metadata__updated_on","format":"date_time"},{"field":"updated_at","format":"date_time"}],"query":{"bool":{"must":[{"match_all":{}},{"match_all":{}},{"range":{"metadata__timestamp":{"gte":1635162849530,"lte":1642938849530,"format":"epoch_millis"}}}],"filter":[],"should":[],"must_not":[]}},"timeout":"30000ms"}`)
	index := "gitee_issues-enriched"
	r := ElasticSearchQuery(InitElasticSearch(), index, body)
	return r
}

func GetIssueDaysDistribution() GetIssueDaysDistributionData {
	body := &bytes.Buffer{}
	body.WriteString(`{"aggs":{"2":{"histogram":{"field":"time_open_days","interval":10,"min_doc_count":1},"aggs":{"3":{"terms":{"field":"state","size":10,"order":{"_count":"desc"}}}}}},"size":0,"_source":{"excludes":[]},"stored_fields":["*"],"script_fields":{},"docvalue_fields":[{"field":"closed_at","format":"date_time"},{"field":"created_at","format":"date_time"},{"field":"grimoire_creation_date","format":"date_time"},{"field":"metadata__enriched_on","format":"date_time"},{"field":"metadata__timestamp","format":"date_time"},{"field":"metadata__updated_on","format":"date_time"},{"field":"updated_at","format":"date_time"}],"query":{"bool":{"must":[{"match_all":{}},{"match_all":{}},{"range":{"metadata__timestamp":{"gte":1635171505034,"lte":1642947505034,"format":"epoch_millis"}}}],"filter":[],"should":[],"must_not":[]}},"timeout":"30000ms"}`)
	index := "gitee_issues-enriched"
	r := ElasticSearchQuery(InitElasticSearch(), index, body)
	ResultJsonByte, _ := json.Marshal(r)
	var QueryResult GetIssueDaysDistributionQueryResult
	json.Unmarshal(ResultJsonByte, &QueryResult)
	Buckets := QueryResult.Aggregations.Num2.Buckets
	var Data GetIssueDaysDistributionData
	for i := 0; i < len(Buckets); i++ {
		Data.Day = append(Data.Day, Buckets[i].Key)
		Closed := 0
		Open := 0
		Rejected := 0
		Progressing := 0
		for j := 0; j < len(Buckets[i].Num3.Buckets); j++ {
			if Buckets[i].Num3.Buckets[j].Key == "closed" {
				Closed = Buckets[i].Num3.Buckets[j].DocCount
			} else if Buckets[i].Num3.Buckets[j].Key == "open" {
				Open = Buckets[i].Num3.Buckets[j].DocCount
			} else if Buckets[i].Num3.Buckets[j].Key == "rejected" {
				Rejected = Buckets[i].Num3.Buckets[j].DocCount
			} else if Buckets[i].Num3.Buckets[j].Key == "progressing" {
				Progressing = Buckets[i].Num3.Buckets[j].DocCount
			}
		}
		Data.Closed = append(Data.Closed, Closed)
		Data.Open = append(Data.Open, Open)
		Data.Rejected = append(Data.Rejected, Rejected)
		Data.Progressing = append(Data.Progressing, Progressing)
	}
	return Data
}

func GetIssueFirstAttentionTimeDistribution() GetIssueFirstAttentionTimeDistributionQueryResult {
	body := &bytes.Buffer{}
	body.WriteString(`{"aggs":{"2":{"range":{"field":"time_to_first_attention_without_bot","ranges":[{"from":0,"to":1},{"from":1,"to":3},{"from":3,"to":7},{"from":7,"to":14},{"from":14,"to":30},{"from":30,"to":180},{"from":180,"to":360},{"from":360}],"keyed":true}}},"size":0,"_source":{"excludes":[]},"stored_fields":["*"],"script_fields":{},"docvalue_fields":[{"field":"closed_at","format":"date_time"},{"field":"created_at","format":"date_time"},{"field":"grimoire_creation_date","format":"date_time"},{"field":"metadata__enriched_on","format":"date_time"},{"field":"metadata__timestamp","format":"date_time"},{"field":"metadata__updated_on","format":"date_time"},{"field":"updated_at","format":"date_time"}],"query":{"bool":{"must":[{"match_all":{}},{"match_all":{}},{"range":{"metadata__timestamp":{"gte":1635182749430,"lte":1642958749430,"format":"epoch_millis"}}}],"filter":[],"should":[],"must_not":[]}},"timeout":"30000ms"}`)
	index := "gitee_issues-enriched"
	r := ElasticSearchQuery(InitElasticSearch(), index, body)
	ResultJsonByte, _ := json.Marshal(r)
	var QueryResult GetIssueFirstAttentionTimeDistributionQueryResult
	json.Unmarshal(ResultJsonByte, &QueryResult)
	return QueryResult
}

func GetDeveloperIssueBehaviorRecord() GetDeveloperIssueBehaviorRecordData {
	body := &bytes.Buffer{}
	body.WriteString(`{"aggs":{"2":{"date_histogram":{"field":"update_at","interval":"1d","time_zone":"Asia/Shanghai","min_doc_count":1},"aggs":{"3":{"terms":{"field":"action.keyword","size":5,"order":{"_count":"desc"}}}}}},"size":0,"_source":{"excludes":[]},"stored_fields":["*"],"script_fields":{},"docvalue_fields":[{"field":"issue_labels.created_at","format":"date_time"},{"field":"issue_labels.updated_at","format":"date_time"},{"field":"update_at","format":"date_time"}],"query":{"bool":{"must":[{"match_all":{}},{"match_all":{}},{"range":{"update_at":{"gte":1635183667428,"lte":1642959667428,"format":"epoch_millis"}}}],"filter":[],"should":[],"must_not":[]}},"timeout":"30000ms"}`)
	index := "issues_comment-raw"
	r := ElasticSearchQuery(InitElasticSearch(), index, body)
	ResultJsonByte, _ := json.Marshal(r)
	var QueryResult GetDeveloperIssueBehaviorRecordQueryResult
	json.Unmarshal(ResultJsonByte, &QueryResult)
	Buckets := QueryResult.Aggregations.Num2.Buckets
	var Data GetDeveloperIssueBehaviorRecordData
	for i := 0; i < len(Buckets); i++ {
		Data.Date = append(Data.Date, time.Time(Buckets[i].KeyAsString).Format("2006-01-02"))
		Comment := 0
		Creater := 0
		for j := 0; j < len(Buckets[i].Num3.Buckets); j++ {
			if Buckets[i].Num3.Buckets[j].Key == "comment" {
				Comment = Buckets[i].Num3.Buckets[j].DocCount
			} else if Buckets[i].Num3.Buckets[j].Key == "creater" {
				Creater = Buckets[i].Num3.Buckets[j].DocCount
			}
		}
		Data.Creater = append(Data.Creater, Creater)
		Data.Comment = append(Data.Comment, Comment)
	}
	return Data
}

func GetDevelopersCreateIssueCategories() GetDevelopersCreateIssueCategoriesQueryResult {
	body := &bytes.Buffer{}
	body.WriteString(`{"aggs":{"2":{"date_histogram":{"field":"created_at","interval":"1d","time_zone":"Asia/Shanghai","min_doc_count":1},"aggs":{"3":{"terms":{"field":"labels","size":10,"order":{"_count":"desc"}}}}}},"size":0,"_source":{"excludes":[]},"stored_fields":["*"],"script_fields":{},"docvalue_fields":[{"field":"closed_at","format":"date_time"},{"field":"created_at","format":"date_time"},{"field":"grimoire_creation_date","format":"date_time"},{"field":"metadata__enriched_on","format":"date_time"},{"field":"metadata__timestamp","format":"date_time"},{"field":"metadata__updated_on","format":"date_time"},{"field":"updated_at","format":"date_time"}],"query":{"bool":{"must":[{"match_all":{}},{"match_all":{}},{"range":{"metadata__timestamp":{"gte":1635182838911,"lte":1642958838911,"format":"epoch_millis"}}}],"filter":[],"should":[],"must_not":[]}},"timeout":"30000ms"}`)
	index := "gitee_issues-enriched"
	r := ElasticSearchQuery(InitElasticSearch(), index, body)
	ResultJsonByte, _ := json.Marshal(r)
	var QueryResult GetDevelopersCreateIssueCategoriesQueryResult
	json.Unmarshal(ResultJsonByte, &QueryResult)
	return QueryResult
}
