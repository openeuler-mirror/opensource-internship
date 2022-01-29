package handler

import (
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
	Series []struct {
		Value int    `json:"value"`
		Name  string `json:"name"`
	} `json:"series"`
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

type GetIssueOpenDaysDistributionQueryResult struct {
	Data struct {
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
	} `json:"data"`
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
