package config

import (
	"fmt"
	"testing"
)

func TestToString(t *testing.T) {

	c := &Config{
		Name:     "name123",
		Path:     "path",
		Protocol: inv,
		BaudRate: 28800,
	}

	answer := fmt.Sprintf("Loaded configuration:\n\tName: %v\n\tPath: %v\n\tDataFilePath: %v\n\tProtocol: %v\n\tBaudrate: %v\n",
		c.Name, c.Path, c.DataFilePath, c.Protocol, c.BaudRate)

	x := c.toString()
	if x != answer {
		t.Error("got", x, "want", answer)
	}
}

func BenchmarkToString(b *testing.B) {

	c := &Config{
		Name:     "name123",
		Path:     "path",
		Protocol: inv,
		BaudRate: 28800,
	}

	for i := 0; i < b.N; i++ {
		c.toString()
	}
}
