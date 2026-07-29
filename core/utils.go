package core

import (
	"context"
	"strings"
	"time"

	"gitlab.com/bersufekadgetachew/cbe-super-app-shared/shared/survey"
)

type SurveyResult struct {
	SurveyType survey.SamplingMethod
	Result     bool
	Url        *string
}

type surveyRuleConstraint interface {
	survey.EnabledBranchRule |
		survey.FirstTransactionRule |
		survey.TimebaseSurveyRule |
		survey.SuperappRoleRule |
		survey.SuccessThresholdRule
}

// Rule holds exactly one typed survey rule.
type Rule[T surveyRuleConstraint] struct {
	SurveyType survey.SamplingMethod
	Rule       T
}

// SurveyParam carries request context plus one typed rule.
type SurveyParam[T surveyRuleConstraint] struct {
	RedisKey     string
	BranchCode   string
	SuperappRole string
	Rule         Rule[T]
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

func findRule[T surveyRuleConstraint](param SurveyParam[T], method survey.SamplingMethod) (Rule[T], bool) {
	if param.Rule.SurveyType == method {
		return param.Rule, true
	}
	return Rule[T]{SurveyType: method}, false
}

func (c *CBECoreAPI) initSurvey(ctx context.Context, redisKey, branchCode, superappRole string, surveyRules []survey.SurveyRule) []SurveyResult {
	if c.config.RedisClient == nil {
		return nil
	}

	results := make([]SurveyResult, 0, len(surveyRules))
	for i := 0; i < len(surveyRules); i++ {
		rule := surveyRules[i]

		if rule.SuccessThreshold != nil {
			results = append(results, c.triggerTresholdSurvey(ctx, SurveyParam[survey.SuccessThresholdRule]{
				RedisKey:     redisKey,
				BranchCode:   branchCode,
				SuperappRole: superappRole,
				Rule: Rule[survey.SuccessThresholdRule]{
					SurveyType: survey.SamplingHighValue,
					Rule:       *rule.SuccessThreshold,
				},
			}))
		}
		if rule.EnabledBranch != nil {
			results = append(results, c.triggerBranchSurvey(SurveyParam[survey.EnabledBranchRule]{
				RedisKey:     redisKey,
				BranchCode:   branchCode,
				SuperappRole: superappRole,
				Rule: Rule[survey.EnabledBranchRule]{
					SurveyType: survey.SamplingBranchBased,
					Rule:       *rule.EnabledBranch,
				},
			}))
		}
		if rule.FirstTransaction != nil {
			results = append(results, c.triggerFirstTransactionSurvey(SurveyParam[survey.FirstTransactionRule]{
				RedisKey:     redisKey,
				BranchCode:   branchCode,
				SuperappRole: superappRole,
				Rule: Rule[survey.FirstTransactionRule]{
					SurveyType: survey.SamplingFirstTransaction,
					Rule:       *rule.FirstTransaction,
				},
			}))
		}
		if rule.TimebaseSurvey != nil {
			results = append(results, c.triggerTimebaseSurvey(SurveyParam[survey.TimebaseSurveyRule]{
				RedisKey:     redisKey,
				BranchCode:   branchCode,
				SuperappRole: superappRole,
				Rule: Rule[survey.TimebaseSurveyRule]{
					SurveyType: survey.SamplingTimeBased,
					Rule:       *rule.TimebaseSurvey,
				},
			}))
		}
		if rule.SuperappRole != nil {
			results = append(results, c.triggerSuperappRoleSurvey(SurveyParam[survey.SuperappRoleRule]{
				RedisKey:     redisKey,
				BranchCode:   branchCode,
				SuperappRole: superappRole,
				Rule: Rule[survey.SuperappRoleRule]{
					SurveyType: survey.SamplingCustomerSegment,
					Rule:       *rule.SuperappRole,
				},
			}))
		}
	}

	return results
}

func (c *CBECoreAPI) triggerBranchSurvey(param SurveyParam[survey.EnabledBranchRule]) SurveyResult {
	rule, ok := findRule(param, survey.SamplingBranchBased)
	if !ok {
		return SurveyResult{
			SurveyType: survey.SamplingBranchBased,
			Result:     false,
			Url:        nil,
		}
	}

	for i := 0; i < len(rule.Rule.Branches); i++ {
		branch := rule.Rule.Branches[i]
		if strings.TrimSpace(branch) == strings.TrimSpace(param.BranchCode) {
			return SurveyResult{
				SurveyType: survey.SamplingBranchBased,
				Result:     true,
				Url:        rule.Rule.Url,
			}
		}
	}

	return SurveyResult{
		SurveyType: survey.SamplingBranchBased,
		Result:     false,
		Url:        nil,
	}
}

func (c *CBECoreAPI) triggerFirstTransactionSurvey(param SurveyParam[survey.FirstTransactionRule]) SurveyResult {
	rule, ok := findRule(param, survey.SamplingFirstTransaction)
	if !ok {
		return SurveyResult{
			SurveyType: survey.SamplingFirstTransaction,
			Result:     false,
			Url:        nil,
		}
	}

	return SurveyResult{
		SurveyType: survey.SamplingFirstTransaction,
		Result:     rule.Rule.Enabled,
		Url:        rule.Rule.Url,
	}
}

func (c *CBECoreAPI) triggerTimebaseSurvey(param SurveyParam[survey.TimebaseSurveyRule]) SurveyResult {
	rule, ok := findRule(param, survey.SamplingTimeBased)
	if !ok {
		return SurveyResult{
			SurveyType: survey.SamplingTimeBased,
			Result:     false,
			Url:        nil,
		}
	}

	startTimestamp := rule.Rule.StartTimestamp
	endTimestamp := rule.Rule.EndTimestamp

	if startTimestamp == "" || endTimestamp == "" {
		return SurveyResult{
			SurveyType: survey.SamplingTimeBased,
			Result:     false,
			Url:        nil,
		}
	}

	startTimestampTime, err := time.Parse(time.RFC3339, startTimestamp)
	if err != nil {
		return SurveyResult{
			SurveyType: survey.SamplingTimeBased,
			Result:     false,
			Url:        nil,
		}
	}

	endTimestampTime, err := time.Parse(time.RFC3339, endTimestamp)
	if err != nil {
		return SurveyResult{
			SurveyType: survey.SamplingTimeBased,
			Result:     false,
			Url:        nil,
		}
	}

	if time.Now().Before(startTimestampTime) || time.Now().After(endTimestampTime) {
		return SurveyResult{
			SurveyType: survey.SamplingTimeBased,
			Result:     false,
			Url:        nil,
		}
	}

	return SurveyResult{
		SurveyType: survey.SamplingTimeBased,
		Result:     true,
		Url:        rule.Rule.Url,
	}
}

func (c *CBECoreAPI) triggerSuperappRoleSurvey(param SurveyParam[survey.SuperappRoleRule]) SurveyResult {
	rule, ok := findRule(param, survey.SamplingCustomerSegment)
	if !ok || len(rule.Rule.Roles) == 0 {
		return SurveyResult{
			SurveyType: survey.SamplingCustomerSegment,
			Result:     false,
			Url:        nil,
		}
	}

	for i := 0; i < len(rule.Rule.Roles); i++ {
		role := rule.Rule.Roles[i]
		if role == param.SuperappRole {
			return SurveyResult{
				SurveyType: survey.SamplingCustomerSegment,
				Result:     true,
				Url:        rule.Rule.Url,
			}
		}
	}

	return SurveyResult{
		SurveyType: survey.SamplingCustomerSegment,
		Result:     false,
		Url:        nil,
	}
}

func (c *CBECoreAPI) triggerTresholdSurvey(ctx context.Context, param SurveyParam[survey.SuccessThresholdRule]) SurveyResult {
	rule, ok := findRule(param, survey.SamplingHighValue)
	if !ok || rule.Rule.SuccessThreshold.Value <= 0 {
		return SurveyResult{
			SurveyType: survey.SamplingHighValue,
			Result:     false,
			Url:        nil,
		}
	}

	redisKey := c.redisKey("success_ft_count", param.RedisKey)
	successCount, err := c.config.RedisClient.Get(ctx, redisKey).Int()
	if err != nil {
		return SurveyResult{
			SurveyType: survey.SamplingHighValue,
			Result:     false,
			Url:        nil,
		}
	}

	threshold := rule.Rule.SuccessThreshold
	switch threshold.ThresholdType {
	case survey.Percentage:
		successCount = (successCount * 100) / threshold.Value
		if successCount >= threshold.Value {
			return SurveyResult{
				SurveyType: survey.SamplingHighValue,
				Result:     true,
				Url:        rule.Rule.Url,
			}
		}
	case survey.Absolute:
		if successCount > threshold.Value {
			return SurveyResult{
				SurveyType: survey.SamplingHighValue,
				Result:     true,
				Url:        rule.Rule.Url,
			}
		}
	}

	return SurveyResult{
		SurveyType: survey.SamplingHighValue,
		Result:     false,
		Url:        nil,
	}
}
