package lineio

import (
	"bufio"
	"os"
	"strings"
)

func ForEach(name string, cb func(line string) error) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.TrimRight(s.Text(), "\r\n")
		if err := cb(line); err != nil {
			return err
		}
	}
	if err := s.Err(); err != nil {
		return err
	}
	return f.Close()
}
