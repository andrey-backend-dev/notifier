package notify

import "fmt"

type Notification struct {
	ID        uint64
	Recipient string
	Body      string
	Status    Status
	Attempts  uint32
}

func (n *Notification) Retry() {
	if n.Status == StatusDelivered {
		return
	}

	n.Status = StatusDelivering
	n.Attempts++
}

func (n Notification) String() string {
	return fmt.Sprintf("notification#%d to %s [%s, attempt %d]", n.ID, n.Recipient, n.Status, n.Attempts)
}
