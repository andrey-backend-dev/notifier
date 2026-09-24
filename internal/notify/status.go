package notify

import "fmt"

type Status byte

const (
	StatusUnidentified Status = iota
	StatusFailed
	StatusCreated
	StatusDelivering
	StatusDelivered

	lastStatus
)

func (status Status) Valid() bool {
	return status >= StatusFailed && status < lastStatus
}

func ParseStatus(status string) (Status, error) {
	switch status {
	case "failed":
		return StatusFailed, nil
	case "created":
		return StatusCreated, nil
	case "delivering":
		return StatusDelivering, nil
	case "delivered":
		return StatusDelivered, nil
	default:
		return StatusUnidentified, fmt.Errorf("unknown status: %q", status)
	}
}

func (s Status) IsFinal() bool {
	switch s {
	case StatusFailed, StatusDelivered:
		return true
	default:
		return false
	}
}

func (s Status) String() string {
	switch s {
	case StatusUnidentified:
		return "unidentified"
	case StatusFailed:
		return "failed"
	case StatusCreated:
		return "created"
	case StatusDelivering:
		return "delivering"
	case StatusDelivered:
		return "delivered"
	default:
		return fmt.Sprintf("Status(%d)", s)
	}
}
