package core

import "gitlab.com/bersufekadgetachew/cbe-super-app-shared/shared/survey"

func (c *CBECoreAPI) mapSurveyResult(surveyRules []survey.SurveyRule) []survey.SurveyRule {
	rules := make([]survey.SurveyRule, 0, len(surveyRules))
	for i := 0; i < len(surveyRules); i++ {
		rule := surveyRules[i]
		if rule.SuccessThreshold == nil &&
			rule.EnabledBranch == nil &&
			rule.FirstTransaction == nil &&
			rule.TimebaseSurvey == nil &&
			rule.SuperappRole == nil &&
			rule.SingleTransactionAmount == nil {
			continue
		}

		rules = append(rules, rule)
	}
	return rules
}
