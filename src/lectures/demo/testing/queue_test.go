package queue

import "testing"

func TestAppend(t *testing.T) {
	queue := New(5)
	for i := 0; i < 5; i++ {
		if len(queue.items) != i {
			t.Errorf("Incorrect queue element count %v, want %v", len(queue.items), i)
		}
		if !queue.Append(i) {
			t.Errorf("failed to append item %v", i)
		}
	}
	if queue.Append(6) {
		t.Errorf("Queue was full, should not have appended")
	}
}

func TestNext(t *testing.T) {
	queue := New(5)
	for i := 0; i <= 5; i++ {
		queue.Append(i)
	}

	for i := 0; i < 5; i++ {
		value, valid := queue.Next()

		if !valid {
			t.Errorf("should have returned a value %v", i)
		}
		if value != i {
			t.Errorf("should have returned a value %v instead of %v", i, value)
		}
	}

	value, valid := queue.Next()
	if valid {
		t.Errorf("Should not have returned a value %v", value)
	}
}
