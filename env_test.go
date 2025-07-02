package goconf

import (
	"testing"
)

func TestLoadEnv(t *testing.T) {
	type TestConf struct {
		String         string `env:"SIMPLE_STRING"`
		Integer        int    `env:"INT_VALUE"`
		Boolean        bool   `env:"BOOLEAN_VALUE"`
		DefaultString  string `env:"DEFAULT_STRING" default:"default string"`
		UnqottedString string `env:"QUOTTED_STRING"`
	}

	var tc TestConf

	t.Setenv("SIMPLE_STRING", "simple")
	t.Setenv("INT_VALUE", "42")
	t.Setenv("BOOLEAN_VALUE", "true")
	t.Setenv("QUOTTED_STRING", "\"noquotes\"")

	if err := LoadEnv(&tc); err != nil {
		t.Error("Error while loading environment: ", err)
	}

	if tc.String != "simple" {
		t.Error("String vaue is not set")
	}

	if tc.Integer != 42 {
		t.Errorf("Integer vaue is not set, want:%d got:%d", 42, tc.Integer)
	}

	if tc.Boolean != true {
		t.Errorf("Boolean vaue is not set, want: %t got:%t", true, tc.Boolean)
	}

	if tc.DefaultString != "default string" {
		t.Errorf("Fafault string not set to declared default value")
	}

	if tc.UnqottedString != "noquotes" {
		t.Errorf("Quotes not removed, want: noquotes got:%s", tc.UnqottedString)
	}
}
