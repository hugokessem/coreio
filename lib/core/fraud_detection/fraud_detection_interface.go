package frauddetection

type FraudAPIInterface interface {
	Call(param FraudAPIPayload) (*FraudAPIResponse, error)
}
