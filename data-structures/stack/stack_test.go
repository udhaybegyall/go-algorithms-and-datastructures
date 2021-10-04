package stack

import "testing"

func TestStack_stack(t *testing.T) {
	var stack Stack

	stack.push(4)
	stack.push(5)
	stack.push(1)

	empty := stack.isEmpty()
	want_empty := false

	pop := stack.pop()
	want_pop := 1

	if pop != want_pop {
		t.Errorf("got %q, wanted %q", pop, want_pop)
	}

	if empty != want_empty {
		t.Errorf("got %t, wanted %t", empty, want_empty)
	}
}
