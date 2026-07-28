package core

import (
	"context"
	"strings"
	"time"
)

type SurveyRule struct {
	SuccessThreshold SuccessThreshold
	EnabledBranch    []string
}

type ThresholdType string

const (
	Percentage ThresholdType = "percentage"
	Absolute   ThresholdType = "absolute"
)

type SuccessThreshold struct {
	ThresholdType ThresholdType
	Value         int
}

func (c *CBECoreAPI) redisKey(prefix, key string) string {

	if key == "" {
		now := time.Now().Format("20060102")
		return prefix + "_" + now
	}

	now := time.Now().Format("20060102")
	temp_key := prefix + "_" + now
	if temp_key != key {
		return temp_key
	}

	return key
}

func (c *CBECoreAPI) triggerBranchSurvey(branchCode string, rule SurveyRule) bool {
	for i := 0; i < len(rule.EnabledBranch); i++ {
		branch := rule.EnabledBranch[i]
		if strings.TrimSpace(branch) == strings.TrimSpace(branchCode) {
			return true
		}
	}

	return false
}

func (c *CBECoreAPI) triggerTresholdSurvey(ctx context.Context, key string, count int, rule SurveyRule) bool {

	if c.config.RedisClient == nil || rule.SuccessThreshold.Value <= 0 {
		return false
	}

	redisKey := c.redisKey("success_ft_count", key)
	successCount, err := c.config.RedisClient.Get(ctx, redisKey).Int()
	if err != nil {
		return false
	}

	if rule.SuccessThreshold.Value <= 0 {
		return false
	}

	switch rule.SuccessThreshold.ThresholdType {
	case Percentage:
		successCount = (count * 100) / rule.SuccessThreshold.Value
		if successCount >= rule.SuccessThreshold.Value {
			return true
		}
	case Absolute:
		if count > rule.SuccessThreshold.Value {
			return true
		}
	}

	return false
}
