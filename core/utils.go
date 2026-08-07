package core

import (
	"context"
	"strings"
	"time"

	"gitlab.com/bersufekadgetachew/cbe-super-app-shared/shared/survey"
	valueobject "gitlab.com/bersufekadgetachew/cbe-super-app-shared/shared/survey/value_object"
)

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

func (c *CBECoreAPI) initSurvey(ctx context.Context, redisKey, branchCode, superappRole string, surveyRules []survey.SurveyRule) survey.SurveyResult {
	if c.config.RedisClient == nil {
		return survey.SurveyResult{
			SurveyType: nil,
			Result:     false,
			Url:        nil,
		}
	}

	type SurveyResult struct {
		result survey.SurveyResult
		order  uint8
	}
	results := make([]SurveyResult, 0, len(surveyRules))
	for i := 0; i < len(surveyRules); i++ {
		rule := surveyRules[i]

		if rule.SuccessThreshold != nil {
			results = append(results, SurveyResult{
				result: c.triggerTresholdSurvey(ctx, SurveyParam[survey.SuccessThresholdRule]{
					RedisKey:     redisKey,
					BranchCode:   branchCode,
					SuperappRole: superappRole,
					Rule: Rule[survey.SuccessThresholdRule]{
						SurveyType: survey.SamplingHighValue,
						Rule:       *rule.SuccessThreshold,
					},
				}), order: rule.SuccessThreshold.Order,
			})
		}

		if rule.EnabledBranch != nil {
			results = append(results, SurveyResult{
				result: c.triggerBranchSurvey(SurveyParam[survey.EnabledBranchRule]{
					RedisKey:     redisKey,
					BranchCode:   branchCode,
					SuperappRole: superappRole,
					Rule: Rule[survey.EnabledBranchRule]{
						SurveyType: survey.SamplingBranchBased,
						Rule:       *rule.EnabledBranch,
					},
				}), order: rule.EnabledBranch.Order,
			})
		}
		if rule.FirstTransaction != nil {
			results = append(results, SurveyResult{
				result: c.triggerFirstTransactionSurvey(SurveyParam[survey.FirstTransactionRule]{
					RedisKey:     redisKey,
					BranchCode:   branchCode,
					SuperappRole: superappRole,
					Rule: Rule[survey.FirstTransactionRule]{
						SurveyType: survey.SamplingFirstTransaction,
						Rule:       *rule.FirstTransaction,
					},
				}), order: rule.FirstTransaction.Order,
			})
		}
		if rule.TimebaseSurvey != nil {
			results = append(results, SurveyResult{
				result: c.triggerTimebaseSurvey(SurveyParam[survey.TimebaseSurveyRule]{
					RedisKey:     redisKey,
					BranchCode:   branchCode,
					SuperappRole: superappRole,
					Rule: Rule[survey.TimebaseSurveyRule]{
						SurveyType: survey.SamplingTimeBased,
						Rule:       *rule.TimebaseSurvey,
					},
				}), order: rule.TimebaseSurvey.Order,
			})
		}
		if rule.SuperappRole != nil {
			results = append(results, SurveyResult{
				result: c.triggerSuperappRoleSurvey(SurveyParam[survey.SuperappRoleRule]{
					RedisKey:     redisKey,
					BranchCode:   branchCode,
					SuperappRole: superappRole,
					Rule: Rule[survey.SuperappRoleRule]{
						SurveyType: survey.SamplingCustomerSegment,
						Rule:       *rule.SuperappRole,
					},
				}), order: rule.SuperappRole.Order,
			})
		}
	}

	if len(results) == 0 {
		return survey.SurveyResult{
			SurveyType: nil,
			Result:     false,
			Url:        nil,
		}
	}

	temp := results[0]
	for i := 0; i < len(results); i++ {
		if results[i].order > temp.order {
			temp = results[i]
		}
	}

	return survey.SurveyResult{
		SurveyType: temp.result.SurveyType,
		Result:     temp.result.Result,
		Url:        temp.result.Url,
	}
}

func (c *CBECoreAPI) triggerBranchSurvey(param SurveyParam[survey.EnabledBranchRule]) survey.SurveyResult {
	rule, ok := findRule(param, survey.SamplingBranchBased)
	if !ok {
		return survey.SurveyResult{
			SurveyType: nil,
			Result:     false,
			Url:        nil,
		}
	}

	for i := 0; i < len(rule.Rule.Branches); i++ {
		branch := rule.Rule.Branches[i]
		if strings.TrimSpace(branch) == strings.TrimSpace(param.BranchCode) {
			return survey.SurveyResult{
				SurveyType: &rule.SurveyType,
				Result:     true,
				Url:        rule.Rule.Url,
			}
		}
	}

	return survey.SurveyResult{
		SurveyType: nil,
		Result:     false,
		Url:        nil,
	}
}

func (c *CBECoreAPI) triggerFirstTransactionSurvey(param SurveyParam[survey.FirstTransactionRule]) survey.SurveyResult {
	rule, ok := findRule(param, survey.SamplingFirstTransaction)
	if !ok {
		return survey.SurveyResult{
			SurveyType: nil,
			Result:     false,
			Url:        nil,
		}
	}

	return survey.SurveyResult{
		SurveyType: &rule.SurveyType,
		Result:     rule.Rule.Enabled,
		Url:        rule.Rule.Url,
	}
}

func (c *CBECoreAPI) triggerTimebaseSurvey(param SurveyParam[survey.TimebaseSurveyRule]) survey.SurveyResult {
	rule, ok := findRule(param, survey.SamplingTimeBased)
	if !ok {
		return survey.SurveyResult{
			SurveyType: nil,
			Result:     false,
			Url:        nil,
		}
	}

	startTimestamp := rule.Rule.StartTimestamp
	endTimestamp := rule.Rule.EndTimestamp

	if !startTimestamp.IsValid() || !endTimestamp.IsValid() {
		return survey.SurveyResult{
			SurveyType: nil,
			Result:     false,
			Url:        nil,
		}
	}

	if valueobject.IsNowBetween(startTimestamp, endTimestamp) {
		return survey.SurveyResult{
			SurveyType: &rule.SurveyType,
			Result:     true,
			Url:        rule.Rule.Url,
		}
	}

	return survey.SurveyResult{
		SurveyType: nil,
		Result:     false,
		Url:        nil,
	}
}

func (c *CBECoreAPI) triggerSuperappRoleSurvey(param SurveyParam[survey.SuperappRoleRule]) survey.SurveyResult {
	rule, ok := findRule(param, survey.SamplingCustomerSegment)
	if !ok || len(rule.Rule.Roles) == 0 {
		return survey.SurveyResult{
			SurveyType: nil,
			Result:     false,
			Url:        nil,
		}
	}

	for i := 0; i < len(rule.Rule.Roles); i++ {
		role := rule.Rule.Roles[i]
		if role == param.SuperappRole {
			return survey.SurveyResult{
				SurveyType: &rule.SurveyType,
				Result:     true,
				Url:        rule.Rule.Url,
			}
		}
	}

	return survey.SurveyResult{
		SurveyType: nil,
		Result:     false,
		Url:        nil,
	}
}

func (c *CBECoreAPI) triggerTresholdSurvey(ctx context.Context, param SurveyParam[survey.SuccessThresholdRule]) survey.SurveyResult {
	rule, ok := findRule(param, survey.SamplingHighValue)
	if !ok || len(rule.Rule.SuccessThreshold) == 0 {
		return survey.SurveyResult{
			SurveyType: nil,
			Result:     false,
			Url:        nil,
		}
	}

	successKey := param.RedisKey
	if successKey == "" || !strings.HasPrefix(successKey, "success_ft_count") {
		successKey = c.redisKey("success_ft_count", param.RedisKey)
	}

	successCount, err := c.config.RedisClient.Get(ctx, successKey).Int()
	if err != nil {
		successCount = 0
	}

	failedKey := strings.Replace(successKey, "success_ft_count", "failed_ft_count", 1)
	failedCount, err := c.config.RedisClient.Get(ctx, failedKey).Int()
	if err != nil {
		failedCount = 0
	}

	met := false
	for i := 0; i < len(rule.Rule.SuccessThreshold); i++ {
		threshold := rule.Rule.SuccessThreshold[i]
		if threshold.Value == 0 {
			continue
		}

		switch threshold.ThresholdType {
		case survey.Percentage:
			total := successCount + failedCount
			if total > 0 {
				rate := (successCount * 100) / total
				if rate >= int(threshold.Value) {
					met = true
				}
			}
		case survey.Frequency:
			if successCount >= int(threshold.Value) {
				met = true
			}
		}

		if met {
			break
		}
	}

	if met {
		return survey.SurveyResult{
			SurveyType: &rule.SurveyType,
			Result:     true,
			Url:        rule.Rule.Url,
		}
	}

	return survey.SurveyResult{
		SurveyType: nil,
		Result:     false,
		Url:        nil,
	}
}
