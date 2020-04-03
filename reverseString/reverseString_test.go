package reverseString

import "testing"

func TestReverseString(t *testing.T) {
	got := []byte{'a', 'b', 'c'}
	reverseString(got)
	want := []byte{'c', 'b', 'a'}

	for i := range want {
		if want[i] != got[i] {
			t.Errorf("got %q want %q", got, want)
		}
	}
}
