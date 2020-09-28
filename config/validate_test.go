package config

import (
	"fmt"
	"testing"
)

func TestValidate(t *testing.T) {

	type test struct {
		data   *Config
		answer error
	}

	tests := []test{
		{
			data:   &Config{},
			answer: fmt.Errorf("error"),
		}, {
			data:   &Config{Name: "test1"},
			answer: fmt.Errorf("error"),
		}, {
			data:   &Config{Name: "test1", Path: "abc"},
			answer: fmt.Errorf("error"),
		}, {
			data:   &Config{Name: "test1", Path: "abc", Protocol: inv, BaudRate: 123},
			answer: fmt.Errorf("error"),
		}, {
			data:   &Config{Name: "test1", Path: "abc", DataFilePath: "testdata", Protocol: inv},
			answer: fmt.Errorf("error"),
		}, {
			data:   &Config{Name: "test1", Path: "abc", DataFilePath: "testdata/datafile", Protocol: inv},
			answer: nil,
		}, {
			data:   &Config{Name: "test1", Path: "abc", DataFilePath: "testdata/datafile", Protocol: inv, BaudRate: 1200},
			answer: nil,
		}, {
			data:   &Config{Name: "test1", Path: "abc", DataFilePath: "zzzzzzzzz", Protocol: inv, BaudRate: 1200},
			answer: fmt.Errorf("error"),
		}, {
			data:   &Config{Name: "test1", Path: "abc", DataFilePath: "", Protocol: inv, BaudRate: 1200},
			answer: fmt.Errorf("error"),
		},
	}

	for _, v := range tests {
		x := v.data.Validate()
		if (x == nil && v.answer != nil) || (x != nil && v.answer == nil) {
			t.Error("got", x, "want", v.answer)
		}
	}
}

func BenchmarkValidate(b *testing.B) {

	c := &Config{
		Name:     "name123",
		Path:     "path",
		Protocol: inv,
		BaudRate: 28800,
	}

	for i := 0; i < b.N; i++ {
		c.Validate()
	}
}

func TestValidateName(t *testing.T) {

	type test struct {
		data   *Config
		answer error
	}

	tests := []test{
		{
			data:   &Config{Name: ""},
			answer: fmt.Errorf("error"),
		}, {
			data:   &Config{Name: "aa"},
			answer: fmt.Errorf("error"),
		}, {
			data:   &Config{Name: "aaaaaaaaaaaaaaaaaaaaa"},
			answer: fmt.Errorf("error"),
		}, {
			data:   &Config{Name: "asd.123"},
			answer: fmt.Errorf("error"),
		},
		{
			data:   &Config{Name: "asd123"},
			answer: nil,
		},
	}

	for _, v := range tests {
		x := validateName(v.data)
		if (x == nil && v.answer != nil) || (x != nil && v.answer == nil) {
			t.Error("got", x, "want", v.answer)
		}
	}
}

func BenchmarkValidateName(b *testing.B) {

	c := &Config{
		Name:     "name123",
		Path:     "path",
		Protocol: inv,
		BaudRate: 28800,
	}

	for i := 0; i < b.N; i++ {
		validateName(c)
	}
}

func TestValidatePath(t *testing.T) {

	type test struct {
		data   *Config
		answer error
	}

	tests := []test{
		{
			data:   &Config{Path: ""},
			answer: fmt.Errorf("error"),
		}, {
			data:   &Config{Path: "aa"},
			answer: nil,
		},
	}

	for _, v := range tests {
		x := validatePath(v.data)
		if (x == nil && v.answer != nil) || (x != nil && v.answer == nil) {
			t.Error("got", x, "want", v.answer)
		}
	}
}

func BenchmarkValidatePath(b *testing.B) {

	c := &Config{
		Name:     "name123",
		Path:     "path",
		Protocol: inv,
		BaudRate: 28800,
	}

	for i := 0; i < b.N; i++ {
		validatePath(c)
	}
}

func TestValidateDataFilePath(t *testing.T) {

	type test struct {
		data   *Config
		answer error
	}

	tests := []test{
		{
			data:   &Config{DataFilePath: ""},
			answer: fmt.Errorf("error"),
		}, {
			data:   &Config{DataFilePath: "testdata/datafile"},
			answer: nil,
		},
	}

	for _, v := range tests {
		x := validateDataFilePath(v.data)
		if (x == nil && v.answer != nil) || (x != nil && v.answer == nil) {
			t.Error("got", x, "want", v.answer)
		}
	}
}

func BenchmarkValidateDataFilePath(b *testing.B) {

	c := &Config{
		Name:         "name123",
		Path:         "path",
		DataFilePath: "datafile",
		Protocol:     inv,
		BaudRate:     28800,
	}

	for i := 0; i < b.N; i++ {
		validateDataFilePath(c)
	}
}

func TestValidateProtocol(t *testing.T) {

	type test struct {
		data   *Config
		answer error
	}

	tests := []test{
		{
			data:   &Config{Protocol: "inv"},
			answer: fmt.Errorf("error"),
		}, {
			data:   &Config{Protocol: inv},
			answer: nil,
		},
	}

	for _, v := range tests {
		x := validateProtocol(v.data)
		if (x == nil && v.answer != nil) || (x != nil && v.answer == nil) {
			t.Error("got", x, "want", v.answer)
		}
	}
}

func BenchmarkValidateProtocol(b *testing.B) {

	c := &Config{
		Name:     "name123",
		Path:     "path",
		Protocol: inv,
		BaudRate: 28800,
	}

	for i := 0; i < b.N; i++ {
		validateProtocol(c)
	}
}

func TestValidateBaudRate(t *testing.T) {

	type test struct {
		data   *Config
		answer error
	}

	tests := []test{
		{
			data:   &Config{BaudRate: 123},
			answer: fmt.Errorf("error"),
		}, {
			data:   &Config{},
			answer: nil,
		}, {
			data:   &Config{BaudRate: 9600},
			answer: nil,
		},
	}

	for _, v := range tests {
		x := validateBaudRate(v.data)
		if (x == nil && v.answer != nil) || (x != nil && v.answer == nil) {
			t.Error("got", x, "want", v.answer)
		}
	}
}

func BenchmarkValidateBaudRate(b *testing.B) {

	c := &Config{
		Name:     "name123",
		Path:     "path",
		Protocol: inv,
		BaudRate: 28800,
	}

	for i := 0; i < b.N; i++ {
		validateBaudRate(c)
	}
}
