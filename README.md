flowchart TD
  A[FT SOAP parsed OK] --> B[Incr success_ft_count]
  B --> C{IsSurveyEnabled OR SurveyRules?}
  C -->|No| Z[Return result]
  C -->|Yes| D[redisKey success_ft_count]
  D --> E[mapSurveyResult filter empty]
  E --> F{RedisClient set?}
  F -->|No| G[surveyResults = nil]
  F -->|Yes| H[For each rule]
  H --> I{Which fields set?}
  I --> J[Threshold]
  I --> K[Branch]
  I --> L[FirstTxn]
  I --> M[Timebase]
  I --> N[SuperappRole]
  J --> O[Collect SurveyResults]
  K --> O
  L --> O
  M --> O
  N --> O
  O --> P[result.SurveyResults = surveyResults]
  G --> P
  P --> Z
