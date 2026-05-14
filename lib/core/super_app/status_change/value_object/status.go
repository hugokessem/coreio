package valueobject

import "errors"

type Status string

const (
	Active   Status = "ACTIVE"
	InActive Status = "INACTIVE"
	Disabled Status = "DESABLED"
	Band     Status = "BARD"
	Expired  Status = "EXPIRED"
	Locked   Status = "LOCKED"
	Blocked  Status = "BLOCKED"
)

var ErrInvalidSuperappStatus = errors.New("invalid superapp status")

func (s Status) IsValid() error {
	switch s {
	case Active, InActive, Disabled, Band, Expired, Locked, Blocked:
		return nil
	default:
		return ErrInvalidSuperappStatus
	}
}

func (s Status) String() string {
	return string(s)
}
