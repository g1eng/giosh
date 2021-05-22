package shell

import "testing"

func TestNew(t *testing.T) {
	cmd := New()
	if cmd.lineno != 1 {
		t.Fail()
	}
}
