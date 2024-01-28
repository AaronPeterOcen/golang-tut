package queue

import "testing"

func TestAddQueue(test *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		if len(q.items) != i {
			test.Errorf("Incorrect queue element count: %v, want %v", len(q.items), i)
		}
		if !q.Append(i) {
			test.Errorf("failure %v", i)
		}
	}
	if q.Append(4) {
		test.Errorf("not able to add")
	}
}

func TestNext(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		q.Append(i)
	}

	for i := 0; i < 3; i++ {
		item, ok := q.Next()
		if !ok {
			t.Errorf("should be able o queue")
		}
		if item != i {
			t.Errorf("lol")
		}
	}
	item, ok := q.Next()
	if ok {
		t.Errorf("should not: %v", item)
	}
}