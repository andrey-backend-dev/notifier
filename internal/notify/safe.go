package notify

import "fmt"

func SafeHandle(notification Notification, handler SendFunc) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Возникла паника в обработчике: %v", r)
		}
	}()
	return handler(notification)
}
