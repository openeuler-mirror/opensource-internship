package handler

import (
	"encoding/json"
	"time"
)

type GetCurrentStatusOfHistoricalIssueQueryResult struct {
	Shards struct {
		Failed     int `json:"failed"`
		Skipped    int `json:"skipped"`
		Successful int `json:"successful"`
		Total      int `json:"total"`
	} `json:"_shards"`
	Aggregations struct {
		Num2 struct {
			Buckets []struct {
				Num3 struct {
					Buckets []struct {
						DocCount int    `json:"doc_count"`
						Key      string `json:"key"`
					} `json:"buckets"`
					DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
					SumOtherDocCount        int `json:"sum_other_doc_count"`
				} `json:"3"`
				DocCount    int       `json:"doc_count"`
				Key         int64     `json:"key"`
				KeyAsString time.Time `json:"key_as_string"`
			} `json:"buckets"`
		} `json:"2"`
	} `json:"aggregations"`
	Hits struct {
		Hits     []interface{} `json:"hits"`
		MaxScore int           `json:"max_score"`
		Total    int           `json:"total"`
	} `json:"hits"`
	TimedOut bool `json:"timed_out"`
	Took     int  `json:"took"`
}

type GetIssueStateDistributionQueryResult struct {
	Shards struct {
		Failed     int `json:"failed"`
		Skipped    int `json:"skipped"`
		Successful int `json:"successful"`
		Total      int `json:"total"`
	} `json:"_shards"`
	Aggregations struct {
		State struct {
			Buckets []struct {
				DocCount int    `json:"doc_count"`
				Key      string `json:"key"`
			} `json:"buckets"`
			DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
			SumOtherDocCount        int `json:"sum_other_doc_count"`
		} `json:"state"`
	} `json:"aggregations"`
	ClusterName string `json:"cluster_name"`
	ClusterUUID string `json:"cluster_uuid"`
	Hits        struct {
		Hits     []interface{} `json:"hits"`
		MaxScore int           `json:"max_score"`
		Total    int           `json:"total"`
	} `json:"hits"`
	Name     string `json:"name"`
	Tagline  string `json:"tagline"`
	TimedOut bool   `json:"timed_out"`
	Took     int    `json:"took"`
	Version  struct {
		BuildDate                        time.Time `json:"build_date"`
		BuildFlavor                      string    `json:"build_flavor"`
		BuildHash                        string    `json:"build_hash"`
		BuildSnapshot                    bool      `json:"build_snapshot"`
		BuildType                        string    `json:"build_type"`
		LuceneVersion                    string    `json:"lucene_version"`
		MinimumIndexCompatibilityVersion string    `json:"minimum_index_compatibility_version"`
		MinimumWireCompatibilityVersion  string    `json:"minimum_wire_compatibility_version"`
		Number                           string    `json:"number"`
	} `json:"version"`
}

type VersionResult struct {
	Server string `json:"server"`
	Client string `json:"client"`
}

type GetIssueStateDistributionData struct {
	Legend []string `json:"legend"`
	Series []struct {
		Value int    `json:"value"`
		Name  string `json:"name"`
	} `json:"series"`
}

type GetIssueOpenDaysDistributionData struct {
	Zero_0_1_0 struct {
		DocCount int64 `json:"doc_count"`
		From     int64 `json:"from"`
		To       int64 `json:"to"`
	} `json:"0.0-1.0"`
	One_0_3_0 struct {
		DocCount int64 `json:"doc_count"`
		From     int64 `json:"from"`
		To       int64 `json:"to"`
	} `json:"1.0-3.0"`
	One4_0_30_0 struct {
		DocCount int64 `json:"doc_count"`
		From     int64 `json:"from"`
		To       int64 `json:"to"`
	} `json:"14.0-30.0"`
	One80_0_360_0 struct {
		DocCount int64 `json:"doc_count"`
		From     int64 `json:"from"`
		To       int64 `json:"to"`
	} `json:"180.0-360.0"`
	Three_0_7_0 struct {
		DocCount int64 `json:"doc_count"`
		From     int64 `json:"from"`
		To       int64 `json:"to"`
	} `json:"3.0-7.0"`
	Three0_0_180_0 struct {
		DocCount int64 `json:"doc_count"`
		From     int64 `json:"from"`
		To       int64 `json:"to"`
	} `json:"30.0-180.0"`
	Three60_0 struct {
		DocCount int64 `json:"doc_count"`
		From     int64 `json:"from"`
	} `json:"360.0-*"`
	Seven_0_14_0 struct {
		DocCount int64 `json:"doc_count"`
		From     int64 `json:"from"`
		To       int64 `json:"to"`
	} `json:"7.0-14.0"`
}

type GetCurrentStatusOfHistoricalIssueData struct {
	Open        []int    `json:"open"`
	Closed      []int    `json:"closed"`
	Rejected    []int    `json:"rejected"`
	Progressing []int    `json:"progressing"`
	Date        []string `json:"date"`
}

type GetDeveloperIssueBehaviorRecordData struct {
	Comment []int    `json:"comment"`
	Creater []int    `json:"creater"`
	Date    []string `json:"date"`
}

type GetIssueDaysDistributionData struct {
	Open        []int `json:"open"`
	Closed      []int `json:"closed"`
	Rejected    []int `json:"rejected"`
	Progressing []int `json:"progressing"`
	Day         []int `json:"day"`
}

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    getIssueOpenDaysDistributionQueryResult, err := UnmarshalGetIssueOpenDaysDistributionQueryResult(bytes)
//    bytes, err = getIssueOpenDaysDistributionQueryResult.Marshal()

func UnmarshalGetIssueOpenDaysDistributionQueryResult(data []byte) (GetIssueOpenDaysDistributionQueryResult, error) {
	var r GetIssueOpenDaysDistributionQueryResult
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GetIssueOpenDaysDistributionQueryResult) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type GetIssueOpenDaysDistributionQueryResult struct {
	Shards       Shards       `json:"_shards"`
	Aggregations Aggregations `json:"aggregations"`
	Hits         Hits         `json:"hits"`
	TimedOut     bool         `json:"timed_out"`
	Took         int64        `json:"took"`
}

type Aggregations struct {
	The2 The2 `json:"2"`
}

type The2 struct {
	Buckets Buckets `json:"buckets"`
}

type Buckets struct {
	The0010     The0010 `json:"0.0-1.0"`
	The1030     The0010 `json:"1.0-3.0"`
	The140300   The0010 `json:"14.0-30.0"`
	The18003600 The0010 `json:"180.0-360.0"`
	The3070     The0010 `json:"3.0-7.0"`
	The3001800  The0010 `json:"30.0-180.0"`
	The3600     The3600 `json:"360.0-*"`
	The70140    The0010 `json:"7.0-14.0"`
}

type The0010 struct {
	DocCount int64 `json:"doc_count"`
	From     int64 `json:"from"`
	To       int64 `json:"to"`
}

type The3600 struct {
	DocCount int64 `json:"doc_count"`
	From     int64 `json:"from"`
}

type Hits struct {
	Hits     []interface{} `json:"hits"`
	MaxScore int64         `json:"max_score"`
	Total    int64         `json:"total"`
}

type Shards struct {
	Failed     int64 `json:"failed"`
	Skipped    int64 `json:"skipped"`
	Successful int64 `json:"successful"`
	Total      int64 `json:"total"`
}

type GetIssueDaysDistributionQueryResult struct {
	Shards struct {
		Failed     int `json:"failed"`
		Skipped    int `json:"skipped"`
		Successful int `json:"successful"`
		Total      int `json:"total"`
	} `json:"_shards"`
	Aggregations struct {
		Num2 struct {
			Buckets []struct {
				Num3 struct {
					Buckets []struct {
						DocCount int    `json:"doc_count"`
						Key      string `json:"key"`
					} `json:"buckets"`
					DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
					SumOtherDocCount        int `json:"sum_other_doc_count"`
				} `json:"3"`
				DocCount int `json:"doc_count"`
				Key      int `json:"key"`
			} `json:"buckets"`
		} `json:"2"`
	} `json:"aggregations"`
	Hits struct {
		Hits     []interface{} `json:"hits"`
		MaxScore int           `json:"max_score"`
		Total    int           `json:"total"`
	} `json:"hits"`
	TimedOut bool `json:"timed_out"`
	Took     int  `json:"took"`
}

type GetIssueFirstAttentionTimeDistributionQueryResult struct {
	Shards struct {
		Failed     int `json:"failed"`
		Skipped    int `json:"skipped"`
		Successful int `json:"successful"`
		Total      int `json:"total"`
	} `json:"_shards"`
	Aggregations struct {
		Num2 struct {
			Buckets struct {
				Zero010 struct {
					DocCount int `json:"doc_count"`
					From     int `json:"from"`
					To       int `json:"to"`
				} `json:"0.0-1.0"`
				One030 struct {
					DocCount int `json:"doc_count"`
					From     int `json:"from"`
					To       int `json:"to"`
				} `json:"1.0-3.0"`
				One40300 struct {
					DocCount int `json:"doc_count"`
					From     int `json:"from"`
					To       int `json:"to"`
				} `json:"14.0-30.0"`
				One8003600 struct {
					DocCount int `json:"doc_count"`
					From     int `json:"from"`
					To       int `json:"to"`
				} `json:"180.0-360.0"`
				Three070 struct {
					DocCount int `json:"doc_count"`
					From     int `json:"from"`
					To       int `json:"to"`
				} `json:"3.0-7.0"`
				Three001800 struct {
					DocCount int `json:"doc_count"`
					From     int `json:"from"`
					To       int `json:"to"`
				} `json:"30.0-180.0"`
				Three600 struct {
					DocCount int `json:"doc_count"`
					From     int `json:"from"`
				} `json:"360.0-*"`
				Seven0140 struct {
					DocCount int `json:"doc_count"`
					From     int `json:"from"`
					To       int `json:"to"`
				} `json:"7.0-14.0"`
			} `json:"buckets"`
		} `json:"2"`
	} `json:"aggregations"`
	Hits struct {
		Hits     []interface{} `json:"hits"`
		MaxScore int           `json:"max_score"`
		Total    int           `json:"total"`
	} `json:"hits"`
	TimedOut bool `json:"timed_out"`
	Took     int  `json:"took"`
}

type GetDevelopersCreateIssueCategoriesQueryResult struct {
	Shards struct {
		Failed     int `json:"failed"`
		Skipped    int `json:"skipped"`
		Successful int `json:"successful"`
		Total      int `json:"total"`
	} `json:"_shards"`
	Aggregations struct {
		Num2 struct {
			Buckets []struct {
				Num3 struct {
					Buckets []struct {
						DocCount int    `json:"doc_count"`
						Key      string `json:"key"`
					} `json:"buckets"`
					DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
					SumOtherDocCount        int `json:"sum_other_doc_count"`
				} `json:"3"`
				DocCount    int       `json:"doc_count"`
				Key         int64     `json:"key"`
				KeyAsString time.Time `json:"key_as_string"`
			} `json:"buckets"`
		} `json:"2"`
	} `json:"aggregations"`
	Hits struct {
		Hits     []interface{} `json:"hits"`
		MaxScore int           `json:"max_score"`
		Total    int           `json:"total"`
	} `json:"hits"`
	TimedOut bool `json:"timed_out"`
	Took     int  `json:"took"`
}

type GetDeveloperIssueBehaviorRecordQueryResult struct {
	Shards struct {
		Failed     int `json:"failed"`
		Skipped    int `json:"skipped"`
		Successful int `json:"successful"`
		Total      int `json:"total"`
	} `json:"_shards"`
	Aggregations struct {
		Num2 struct {
			Buckets []struct {
				Num3 struct {
					Buckets []struct {
						DocCount int    `json:"doc_count"`
						Key      string `json:"key"`
					} `json:"buckets"`
					DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
					SumOtherDocCount        int `json:"sum_other_doc_count"`
				} `json:"3"`
				DocCount    int       `json:"doc_count"`
				Key         int64     `json:"key"`
				KeyAsString time.Time `json:"key_as_string"`
			} `json:"buckets"`
		} `json:"2"`
	} `json:"aggregations"`
	Hits struct {
		Hits     []interface{} `json:"hits"`
		MaxScore int           `json:"max_score"`
		Total    int           `json:"total"`
	} `json:"hits"`
	TimedOut bool `json:"timed_out"`
	Took     int  `json:"took"`
}

func UnmarshalGetIssueOpenDaysDistributionData(data []byte) (GetIssueOpenDaysDistributionData, error) {
	var r GetIssueOpenDaysDistributionData
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GetIssueOpenDaysDistributionData) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
