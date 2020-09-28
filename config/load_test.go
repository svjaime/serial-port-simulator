package config

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

//how do I test this?
func TestLoad(t *testing.T) {
	viper.AddConfigPath("testdata")
	x, err := Load()

	if err == nil && x == nil {
		t.Error("something went wrong")
	}

	if err != nil && x != nil {
		t.Error("something went wrong")
	}
}

func BenchmarkLoad(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Load()
	}
}

func TestSave(t *testing.T) {

	type test struct {
		data   *Config
		answer error
	}

	tests := []test{
		{
			data:   &Config{Name: "test1", Path: "abc", Protocol: inv},
			answer: nil,
		}, {
			data:   &Config{Name: "test1", Path: "xyz", Protocol: inv},
			answer: fmt.Errorf("error"),
		},
	}

	for _, v := range tests {
		x := save(v.data)
		if (x == nil && v.answer != nil) || (x != nil && v.answer == nil) {
			t.Error("got", x, "want", v.answer)
		}
	}
}

func BenchmarkSave(b *testing.B) {

	c := &Config{
		Name:     "name123",
		Path:     "path",
		Protocol: inv,
		BaudRate: 28800,
	}

	for i := 0; i < b.N; i++ {
		save(c)
	}
}
