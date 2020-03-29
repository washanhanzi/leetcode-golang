package openTheLock

import "testing"

func TestOpenLock(t *testing.T) {
	got := openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202")
	want := 6

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
