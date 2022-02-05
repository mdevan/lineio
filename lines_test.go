package lineio

import (
	"strings"
	"testing"
)

func TestForEach(t *testing.T) {
	err := ForEach("/etc/hosts", func(line string) error {
		if strings.HasPrefix(line, "127.0.0.1") {
			t.Log(line)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
