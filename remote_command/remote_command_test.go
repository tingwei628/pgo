package remote_command_test

// go test -v remote_command_test.go
import (
	"pgo/remote_command"
	"testing"
)

func TestCommnd(t *testing.T) {
	rc := remote_command.RC{}
	total := rc.Command()
	if total != 100 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 100)
	}
}
