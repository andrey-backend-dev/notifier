package notify

import "fmt"

var (
	sequence uint64
)

type MemStore struct {
	items map[uint64]Notification
}

func NewMemStore() MemStore {
	return MemStore{items: make(map[uint64]Notification)}
}

func (m MemStore) Save(notification Notification) (uint64, error) {
	if notification.ID == 0 {
		sequence++
		notification.ID = sequence
	}

	if old, ok := m.items[notification.ID]; ok && old.Status == StatusDelivered {
		return notification.ID, fmt.Errorf("%w: id=%d", ErrAlreadySent, notification.ID)
	}

	m.items[notification.ID] = notification

	return notification.ID, nil
}

func (m MemStore) Load(id uint64) (Notification, error) {
	if v, ok := m.items[id]; ok {
		return v, nil
	}
	return Notification{}, fmt.Errorf("%w: id=%d", ErrNotFound, id)
}

func (m MemStore) Delete(id uint64) error {
	if _, ok := m.items[id]; !ok {
		return fmt.Errorf("%w: id=%d", ErrNotFound, id)
	}

	delete(m.items, id)

	return nil
}
