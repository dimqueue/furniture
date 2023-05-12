package types

import "gitlab.com/distributed_lab/logan/v3/errors"

type Status struct {
	ID       StatusID       `db:"id" json:"id"`
	TextCode StatusTextCode `db:"text_code" json:"text_code"`
}

type StatusID int64
type StatusTextCode string

const (
	Locked   StatusTextCode = "LOCKED"
	UnLocked StatusTextCode = "UNLOCKED"
	Sold     StatusTextCode = "SOLD"
)

const (
	LockedInt StatusID = iota + 1
	UnLockedInt
	SoldInt
)

var (
	errInvalidStatus = errors.New("invalid status")
)

func (s StatusID) ToStatus() (status Status, err error) {
	status.ID = s

	switch s {
	case LockedInt:
		status.TextCode = Locked
	case UnLockedInt:
		status.TextCode = UnLocked
	case SoldInt:
		status.TextCode = Sold
	default:
		err = errInvalidStatus
	}

	return
}

func (s StatusTextCode) ToStatus() (status Status, err error) {
	status.TextCode = s

	switch s {
	case Locked:
		status.ID = LockedInt
	case UnLocked:
		status.ID = UnLockedInt
	case Sold:
		status.ID = SoldInt
	default:
		err = errInvalidStatus
	}

	return
}
