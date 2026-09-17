package notify

import "fmt"

type Status int8

const (
	Unidentified Status = iota
	Failed
	Created
	Delivering
	Delivered
	lastStatus
)

func (status Status) Valid() bool {
	return status >= Failed && status < lastStatus
}

func ParseStatus(status string) (Status, error) {
	switch status {
	case "failed":
		return Failed, nil
	case "created":
		return Created, nil
	case "delivering":
		return Delivering, nil
	case "delivered":
		return Delivered, nil
	default:
		return Unidentified, fmt.Errorf("unknown status: %q", status)
	}
}
