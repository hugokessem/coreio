package valueobject

type FetchBy string

const (
	FetchByCustomerNumber FetchBy = "Customer Numner"
	FetchByAccountNumber  FetchBy = "Account Number"
)

func (f FetchBy) IsValid() bool {
	switch f {
	case FetchByCustomerNumber,
		FetchByAccountNumber:
		return true
	default:
		return false
	}
}

func (f FetchBy) String() string {
	return string(f)
}

func (f FetchBy) Equal(value string) bool {
	return f.String() == value
}
