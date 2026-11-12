package notify

import "testing"

func TestPush(t *testing.T) {
	queue := NotificationQueue{}
	notification1 := Notification{ID: 1}
	notification2 := Notification{ID: 67}

	queue.Push(&notification1)
	queue.Push(&notification2)

	if len(queue.slice) != 2 {
		t.Fatalf("Length of the queue should be 2, actual: %d", len(queue.slice))
	}

	if queue.slice[0] != &notification1 {
		t.Errorf("Expected: %v, actual: %v", notification1, queue.slice[0])
	}

	if queue.slice[1] != &notification2 {
		t.Errorf("Expected: %v, actual: %v", notification2, queue.slice[1])
	}
}

func TestPop(t *testing.T) {
	queue := NotificationQueue{slice: make([]*Notification, 0, 5)}
	notification1 := Notification{ID: 1}
	notification2 := Notification{ID: 2}

	queue.Push(&notification1)
	queue.Push(&notification2)

	if result, ok := queue.Pop(); ok != true || result != &notification1 {
		if ok {
			t.Fatalf("Expected: %v, actual: %v", notification1, result)
		} else {
			t.Fatalf("Expected to be ok")
		}
	}

	if result, ok := queue.Pop(); ok != true || result != &notification2 {
		if ok {
			t.Fatalf("Expected: %v, actual: %v", notification2, result)
		} else {
			t.Fatalf("Expected to be ok")
		}
	}

	if _, ok := queue.Pop(); ok == true {
		t.Fatalf("Expected to be not ok")
	}
}
