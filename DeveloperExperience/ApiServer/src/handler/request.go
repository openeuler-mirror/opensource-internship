package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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
	var VersionResult VersionResult

	Value := RedisGetKey("GetElasticSearchVersion")
	if len(Value) != 0 {
		if err := json.Unmarshal([]byte(Value), &VersionResult); err == nil {
			fmt.Println("GetElasticSearchVersion Use Cache.")
			return VersionResult
		}
	}
	fmt.Println("GetElasticSearchVersion Request ES")

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
	VersionResult.Client = elasticsearch.Version
	VersionResult.Server = r["version"].(map[string]interface{})["number"].(string)

	if b, err := json.Marshal(VersionResult); err == nil {
		RedisSetKey("GetElasticSearchVersion", string(b))
	}

	return VersionResult
}

func GetIssueStateDistribution() GetIssueStateDistributionData {
	var Data GetIssueStateDistributionData

	Value := RedisGetKey("GetIssueStateDistribution")
	if len(Value) != 0 {
		if err := json.Unmarshal([]byte(Value), &Data); err == nil {
			fmt.Println("GetIssueStateDistribution Use Cache.")
			return Data
		}
	}
	fmt.Println("GetIssueStateDistribution Request ES")

	body := &bytes.Buffer{}
	body.WriteString(`{"size":0,"aggs":{"state":{"terms":{"field":"state"}}}}`)
	index := "gitee_issues-enriched"
	r := ElasticSearchQuery(InitElasticSearch(), index, body)
	ResultJsonByte, _ := json.Marshal(r)
	var QueryResult GetIssueStateDistributionQueryResult
	json.Unmarshal(ResultJsonByte, &QueryResult)
	Buckets := QueryResult.Aggregations.State.Buckets

	for i := 0; i < len(Buckets); i++ {
		Data.Legend = append(Data.Legend, Buckets[i].Key)
		Data.Series = append(Data.Series, struct {
			Value int    "json:\"value\""
			Name  string "json:\"name\""
		}{Buckets[i].DocCount, Buckets[i].Key})
	}

	if b, err := json.Marshal(Data); err == nil {
		RedisSetKey("GetIssueStateDistribution", string(b))
	}

	return Data
}

func GetCurrentStatusOfHistoricalIssue() GetCurrentStatusOfHistoricalIssueData {
	var Data GetCurrentStatusOfHistoricalIssueData

	Value := RedisGetKey("GetCurrentStatusOfHistoricalIssue")
	if len(Value) != 0 {
		if err := json.Unmarshal([]byte(Value), &Data); err == nil {
			fmt.Println("GetCurrentStatusOfHistoricalIssue Use Cache.")
			return Data
		}
	}
	fmt.Println("GetCurrentStatusOfHistoricalIssue Request ES")

	body := &bytes.Buffer{}
	body.WriteString(`{"aggs":{"2":{"date_histogram":{"field":"created_at","interval":"1w","time_zone":"Asia/Shanghai","min_doc_count":1},"aggs":{"3":{"terms":{"field":"state","size":1000,"order":{"_count":"desc"}}}}}},"size":0,"script_fields":{},"docvalue_fields":[{"field":"closed_at","format":"date_time"},{"field":"created_at","format":"date_time"},{"field":"grimoire_creation_date","format":"date_time"},{"field":"metadata__enriched_on","format":"date_time"},{"field":"metadata__timestamp","format":"date_time"},{"field":"metadata__updated_on","format":"date_time"},{"field":"updated_at","format":"date_time"}],"query":{"bool":{"must":[{"query_string":{"query":"pull_request:false","analyze_wildcard":true,"default_field":"*"}},{"query_string":{"query":"pull_request:false","analyze_wildcard":true,"default_field":"*"}},{"range":{"metadata__timestamp":{"gte":1635145582100,"lte":1642921582100,"format":"epoch_millis"}}}],"filter":[],"should":[],"must_not":[]}},"timeout":"30000ms"}`)
	index := "gitee_issues-enriched"
	r := ElasticSearchQuery(InitElasticSearch(), index, body)
	ResultJsonByte, _ := json.Marshal(r)
	var QueryResult GetCurrentStatusOfHistoricalIssueQueryResult
	json.Unmarshal(ResultJsonByte, &QueryResult)
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

	if b, err := json.Marshal(Data); err == nil {
		RedisSetKey("GetCurrentStatusOfHistoricalIssue", string(b))
	}

	return Data
}

func GetIssueOpenDaysDistribution() Buckets {
	var Data Buckets

	Value := RedisGetKey("GetIssueOpenDaysDistribution")
	if len(Value) != 0 {
		if err := json.Unmarshal([]byte(Value), &Data); err == nil {
			fmt.Println("GetIssueOpenDaysDistribution Use Cache.")
			return Data
		}
	}
	fmt.Println("GetIssueOpenDaysDistribution Request ES")

	body := &bytes.Buffer{}
	body.WriteString(`{"aggs":{"2":{"range":{"field":"time_open_days","ranges":[{"from":0,"to":1},{"from":1,"to":3},{"from":3,"to":7},{"from":7,"to":14},{"from":14,"to":30},{"from":30,"to":180},{"from":180,"to":360},{"from":360}],"keyed":true}}},"size":0,"_source":{"excludes":[]},"stored_fields":["*"],"script_fields":{},"docvalue_fields":[{"field":"closed_at","format":"date_time"},{"field":"created_at","format":"date_time"},{"field":"grimoire_creation_date","format":"date_time"},{"field":"metadata__enriched_on","format":"date_time"},{"field":"metadata__timestamp","format":"date_time"},{"field":"metadata__updated_on","format":"date_time"},{"field":"updated_at","format":"date_time"}],"query":{"bool":{"must":[{"match_all":{}},{"match_all":{}},{"range":{"metadata__timestamp":{"gte":1635162849530,"lte":1642938849530,"format":"epoch_millis"}}}],"filter":[],"should":[],"must_not":[]}},"timeout":"30000ms"}`)
	index := "gitee_issues-enriched"
	r := ElasticSearchQuery(InitElasticSearch(), index, body)
	ResultJsonByte, _ := json.Marshal(r)
	var QueryResult GetIssueOpenDaysDistributionQueryResult
	json.Unmarshal(ResultJsonByte, &QueryResult)
	Data = QueryResult.Aggregations.The2.Buckets

	if b, err := json.Marshal(Data); err == nil {
		RedisSetKey("GetIssueOpenDaysDistribution", string(b))
	}

	return Data
}

func GetIssueDaysDistribution() GetIssueDaysDistributionData {
	var Data GetIssueDaysDistributionData

	Value := RedisGetKey("GetIssueDaysDistribution")
	if len(Value) != 0 {
		if err := json.Unmarshal([]byte(Value), &Data); err == nil {
			fmt.Println("GetIssueDaysDistribution Use Cache.")
			return Data
		}
	}
	fmt.Println("GetIssueDaysDistribution Request ES")

	body := &bytes.Buffer{}
	body.WriteString(`{"aggs":{"2":{"histogram":{"field":"time_open_days","interval":10,"min_doc_count":1},"aggs":{"3":{"terms":{"field":"state","size":10,"order":{"_count":"desc"}}}}}},"size":0,"_source":{"excludes":[]},"stored_fields":["*"],"script_fields":{},"docvalue_fields":[{"field":"closed_at","format":"date_time"},{"field":"created_at","format":"date_time"},{"field":"grimoire_creation_date","format":"date_time"},{"field":"metadata__enriched_on","format":"date_time"},{"field":"metadata__timestamp","format":"date_time"},{"field":"metadata__updated_on","format":"date_time"},{"field":"updated_at","format":"date_time"}],"query":{"bool":{"must":[{"match_all":{}},{"match_all":{}},{"range":{"metadata__timestamp":{"gte":1635171505034,"lte":1642947505034,"format":"epoch_millis"}}}],"filter":[],"should":[],"must_not":[]}},"timeout":"30000ms"}`)
	index := "gitee_issues-enriched"
	r := ElasticSearchQuery(InitElasticSearch(), index, body)
	ResultJsonByte, _ := json.Marshal(r)
	var QueryResult GetIssueDaysDistributionQueryResult
	json.Unmarshal(ResultJsonByte, &QueryResult)
	Buckets := QueryResult.Aggregations.Num2.Buckets
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

	if b, err := json.Marshal(Data); err == nil {
		RedisSetKey("GetIssueDaysDistribution", string(b))
	}

	return Data
}

func GetIssueFirstAttentionTimeDistribution() GetIssueFirstAttentionTimeDistributionQueryResult {
	var QueryResult GetIssueFirstAttentionTimeDistributionQueryResult

	Value := RedisGetKey("GetIssueFirstAttentionTimeDistribution")
	if len(Value) != 0 {
		if err := json.Unmarshal([]byte(Value), &QueryResult); err == nil {
			fmt.Println("GetIssueFirstAttentionTimeDistribution Use Cache.")
			return QueryResult
		}
	}
	fmt.Println("GetIssueFirstAttentionTimeDistribution Request ES")

	body := &bytes.Buffer{}
	body.WriteString(`{"aggs":{"2":{"range":{"field":"time_to_first_attention_without_bot","ranges":[{"from":0,"to":1},{"from":1,"to":3},{"from":3,"to":7},{"from":7,"to":14},{"from":14,"to":30},{"from":30,"to":180},{"from":180,"to":360},{"from":360}],"keyed":true}}},"size":0,"_source":{"excludes":[]},"stored_fields":["*"],"script_fields":{},"docvalue_fields":[{"field":"closed_at","format":"date_time"},{"field":"created_at","format":"date_time"},{"field":"grimoire_creation_date","format":"date_time"},{"field":"metadata__enriched_on","format":"date_time"},{"field":"metadata__timestamp","format":"date_time"},{"field":"metadata__updated_on","format":"date_time"},{"field":"updated_at","format":"date_time"}],"query":{"bool":{"must":[{"match_all":{}},{"match_all":{}},{"range":{"metadata__timestamp":{"gte":1635182749430,"lte":1642958749430,"format":"epoch_millis"}}}],"filter":[],"should":[],"must_not":[]}},"timeout":"30000ms"}`)
	index := "gitee_issues-enriched"
	r := ElasticSearchQuery(InitElasticSearch(), index, body)
	ResultJsonByte, _ := json.Marshal(r)
	json.Unmarshal(ResultJsonByte, &QueryResult)

	if b, err := json.Marshal(QueryResult); err == nil {
		RedisSetKey("GetIssueFirstAttentionTimeDistribution", string(b))
	}

	return QueryResult
}

func GetDeveloperIssueBehaviorRecord() GetDeveloperIssueBehaviorRecordData {
	var Data GetDeveloperIssueBehaviorRecordData

	Value := RedisGetKey("GetDeveloperIssueBehaviorRecord")
	if len(Value) != 0 {
		if err := json.Unmarshal([]byte(Value), &Data); err == nil {
			fmt.Println("GetDeveloperIssueBehaviorRecord Use Cache.")
			return Data
		}
	}
	fmt.Println("GetDeveloperIssueBehaviorRecord Request ES")

	body := &bytes.Buffer{}
	body.WriteString(`{"aggs":{"2":{"date_histogram":{"field":"update_at","interval":"1d","time_zone":"Asia/Shanghai","min_doc_count":1},"aggs":{"3":{"terms":{"field":"action.keyword","size":5,"order":{"_count":"desc"}}}}}},"size":0,"_source":{"excludes":[]},"stored_fields":["*"],"script_fields":{},"docvalue_fields":[{"field":"issue_labels.created_at","format":"date_time"},{"field":"issue_labels.updated_at","format":"date_time"},{"field":"update_at","format":"date_time"}],"query":{"bool":{"must":[{"match_all":{}},{"match_all":{}},{"range":{"update_at":{"gte":1635183667428,"lte":1642959667428,"format":"epoch_millis"}}}],"filter":[],"should":[],"must_not":[]}},"timeout":"30000ms"}`)
	index := "issues_comment-raw"
	r := ElasticSearchQuery(InitElasticSearch(), index, body)
	ResultJsonByte, _ := json.Marshal(r)
	var QueryResult GetDeveloperIssueBehaviorRecordQueryResult
	json.Unmarshal(ResultJsonByte, &QueryResult)
	Buckets := QueryResult.Aggregations.Num2.Buckets
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

	if b, err := json.Marshal(Data); err == nil {
		RedisSetKey("GetDeveloperIssueBehaviorRecord", string(b))
	}

	return Data
}

func GetDevelopersCreateIssueCategories() GetDevelopersCreateIssueCategoriesQueryResult {
	var QueryResult GetDevelopersCreateIssueCategoriesQueryResult

	Value := RedisGetKey("GetDevelopersCreateIssueCategories")
	if len(Value) != 0 {
		if err := json.Unmarshal([]byte(Value), &QueryResult); err == nil {
			fmt.Println("GetDevelopersCreateIssueCategories Use Cache.")
			return QueryResult
		}
	}
	fmt.Println("GetDevelopersCreateIssueCategories Request ES")

	body := &bytes.Buffer{}
	body.WriteString(`{"aggs":{"2":{"date_histogram":{"field":"created_at","interval":"1d","time_zone":"Asia/Shanghai","min_doc_count":1},"aggs":{"3":{"terms":{"field":"labels","size":10,"order":{"_count":"desc"}}}}}},"size":0,"_source":{"excludes":[]},"stored_fields":["*"],"script_fields":{},"docvalue_fields":[{"field":"closed_at","format":"date_time"},{"field":"created_at","format":"date_time"},{"field":"grimoire_creation_date","format":"date_time"},{"field":"metadata__enriched_on","format":"date_time"},{"field":"metadata__timestamp","format":"date_time"},{"field":"metadata__updated_on","format":"date_time"},{"field":"updated_at","format":"date_time"}],"query":{"bool":{"must":[{"match_all":{}},{"match_all":{}},{"range":{"metadata__timestamp":{"gte":1635182838911,"lte":1642958838911,"format":"epoch_millis"}}}],"filter":[],"should":[],"must_not":[]}},"timeout":"30000ms"}`)
	index := "gitee_issues-enriched"
	r := ElasticSearchQuery(InitElasticSearch(), index, body)
	ResultJsonByte, _ := json.Marshal(r)
	json.Unmarshal(ResultJsonByte, &QueryResult)

	if b, err := json.Marshal(QueryResult); err == nil {
		RedisSetKey("GetDevelopersCreateIssueCategories", string(b))
	}

	return QueryResult
}
