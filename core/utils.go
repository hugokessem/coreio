package core

import (
	"context"
	"strconv"
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
		survey.SuccessThresholdRule |
		survey.SingleTransactionAmountRule
}

// Rule holds exactly one typed survey rule.
type Rule[T surveyRuleConstraint] struct {
	SurveyType valueobject.SamplingMethod
	Rule       T
}

// SurveyParam carries request context plus one typed rule.
type SurveyParam[T surveyRuleConstraint] struct {
	RedisKey     string
	BranchCode   string
	SuperappRole string
	Amount       uint64
	Currency     string
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

func findRule[T surveyRuleConstraint](param SurveyParam[T], method valueobject.SamplingMethod) (Rule[T], bool) {
	if !method.IsValid() {
		return Rule[T]{}, false
	}

	if param.Rule.SurveyType == method {
		return Rule[T]{
			SurveyType: method,
			Rule:       param.Rule.Rule,
		}, true
	}

	return Rule[T]{}, false
}

func (c *CBECoreAPI) initSurvey(ctx context.Context, redisKey, branchCode, superappRole, amount, currency string, surveyRules []survey.SurveyRule) survey.SurveyResult {
	if c.config.RedisClient == nil {
		return survey.SurveyResult{
			SurveyType: nil,
			Result:     false,
			Url:        nil,
		}
	}

	parsedAmount := parseSurveyAmount(amount)

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
					Amount:       parsedAmount,
					Currency:     currency,
					Rule: Rule[survey.SuccessThresholdRule]{
						SurveyType: valueobject.SamplingHighValue,
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
					Amount:       parsedAmount,
					Currency:     currency,
					Rule: Rule[survey.EnabledBranchRule]{
						SurveyType: valueobject.SamplingBranchBased,
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
					Amount:       parsedAmount,
					Currency:     currency,
					Rule: Rule[survey.FirstTransactionRule]{
						SurveyType: valueobject.SamplingFirstTransaction,
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
					Amount:       parsedAmount,
					Currency:     currency,
					Rule: Rule[survey.TimebaseSurveyRule]{
						SurveyType: valueobject.SamplingTimeBased,
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
					Amount:       parsedAmount,
					Currency:     currency,
					Rule: Rule[survey.SuperappRoleRule]{
						SurveyType: valueobject.SamplingCustomerSegment,
						Rule:       *rule.SuperappRole,
					},
				}), order: rule.SuperappRole.Order,
			})
		}

		if rule.SingleTransactionAmount != nil {
			results = append(results, SurveyResult{
				result: c.triggerSingleTransactionAmountRule(SurveyParam[survey.SingleTransactionAmountRule]{
					RedisKey:     redisKey,
					BranchCode:   branchCode,
					SuperappRole: superappRole,
					Amount:       parsedAmount,
					Currency:     currency,
					Rule: Rule[survey.SingleTransactionAmountRule]{
						SurveyType: valueobject.SamplingSingleTransaction,
						Rule:       *rule.SingleTransactionAmount,
					},
				}), order: rule.SingleTransactionAmount.Order,
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
	rule, ok := findRule(param, valueobject.SamplingBranchBased)
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
	rule, ok := findRule(param, valueobject.SamplingFirstTransaction)
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
	rule, ok := findRule(param, valueobject.SamplingTimeBased)
	if !ok || len(rule.Rule.TimeBasedConfig) == 0 {
		return survey.SurveyResult{
			SurveyType: nil,
			Result:     false,
			Url:        nil,
		}
	}

	for i := 0; i < len(rule.Rule.TimeBasedConfig); i++ {
		cfg := rule.Rule.TimeBasedConfig[i]
		if !cfg.StartDate.IsValid() || !cfg.EndDate.IsValid() {
			continue
		}
		if cfg.Weekdays.IsValid() && !cfg.Weekdays.IsToday() {
			continue
		}
		if valueobject.IsNowBetween(cfg.StartDate, cfg.EndDate) {
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

func (c *CBECoreAPI) triggerSuperappRoleSurvey(param SurveyParam[survey.SuperappRoleRule]) survey.SurveyResult {
	rule, ok := findRule(param, valueobject.SamplingCustomerSegment)
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
	rule, ok := findRule(param, valueobject.SamplingHighValue)
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
		case valueobject.Percentage:
			total := successCount + failedCount
			if total > 0 {
				rate := (successCount * 100) / total
				if rate >= int(threshold.Value) {
					met = true
				}
			}
		case valueobject.Frequency:
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

func (c *CBECoreAPI) triggerSingleTransactionAmountRule(param SurveyParam[survey.SingleTransactionAmountRule]) survey.SurveyResult {
	rule, ok := findRule(param, valueobject.SamplingSingleTransaction)
	if !ok || len(rule.Rule.Amount) == 0 {
		return survey.SurveyResult{
			SurveyType: nil,
			Result:     false,
			Url:        nil,
		}
	}

	for i := 0; i < len(rule.Rule.Amount); i++ {
		condition := rule.Rule.Amount[i]
		if condition.Currency != "" && param.Currency != "" &&
			!strings.EqualFold(strings.TrimSpace(condition.Currency), strings.TrimSpace(param.Currency)) {
			continue
		}
		if !condition.Operand.IsValid() {
			continue
		}

		if condition.Operand.Match(int(param.Amount), int(condition.Amount)) {
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

func parseSurveyAmount(amount string) uint64 {
	amount = strings.TrimSpace(amount)
	if amount == "" {
		return 0
	}
	if v, err := strconv.ParseUint(amount, 10, 64); err == nil {
		return v
	}
	f, err := strconv.ParseFloat(amount, 64)
	if err != nil || f < 0 {
		return 0
	}
	return uint64(f)
}
