package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start(addr string) {
	r := gin.Default()
	es := r.Group("/api/es")
	{
		es.GET("/Ping", Ping)
		es.GET("/ElasticSearchVersion", ElasticSearchVersion)
		// Issue 状态分布 --> Issue Status Distribution
		es.GET("/IssueStateDistribution", IssueStateDistribution)
		// 历史 Issue 当前状态 --> Current Status Of Historical Issue
		es.GET("/CurrentStatusOfHistoricalIssue", CurrentStatusOfHistoricalIssue)
		// Issue OpenDays 分布 --> Issue OpenDays Distribution
		es.GET("/IssueOpenDaysDistribution", IssueOpenDaysDistribution)
		// Issue 天数分布 --> Issue Days Distribution
		es.GET("/IssueDaysDistribution", IssueDaysDistribution)
		// Issue FirstAttentionTime 分布 --> Issue First Attention Time Distribution
		es.GET("/IssueFirstAttentionTimeDistribution", IssueFirstAttentionTimeDistribution)
		// 开发者 Issue 行为记录 --> Developer Issue Behavior Record
		es.GET("/DeveloperIssueBehaviorRecord", DeveloperIssueBehaviorRecord)
	}
	err := r.Run(addr)
	CheckError(err)
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Pong",
	})
}

func ElasticSearchVersion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": GetElasticSearchVersion(),
	})
}

func IssueStateDistribution(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": GetIssueStateDistribution(),
	})
}

func CurrentStatusOfHistoricalIssue(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": GetCurrentStatusOfHistoricalIssue(),
	})
}

func IssueOpenDaysDistribution(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": GetIssueOpenDaysDistribution(),
	})
}

func IssueDaysDistribution(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": GetIssueDaysDistribution(),
	})
}

func IssueFirstAttentionTimeDistribution(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": GetIssueFirstAttentionTimeDistribution(),
	})
}

func DeveloperIssueBehaviorRecord(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": GetDeveloperIssueBehaviorRecord(),
	})
}
