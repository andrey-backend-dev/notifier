package notify

type NotificationQueue struct {
	slice []*Notification
}

func (queue *NotificationQueue) Push(n *Notification) {
	queue.slice = append(queue.slice, n)
}

func (queue *NotificationQueue) Pop() (*Notification, bool) {
	if len(queue.slice) == 0 {
		return nil, false
	}

	firstNotification := queue.slice[0]

	clear(queue.slice[0:1])
	queue.slice = queue.slice[1:]

	return firstNotification, true
}
