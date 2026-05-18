package config_file

import (
	"testing"

	. "github.com/infrago/base"
	"github.com/infrago/config"
)

func TestFileDriverNoConfigFile(t *testing.T) {
	cfg, err := (&FileConfigDriver{}).Load(Map{})
	if err != nil {
		t.Fatal(err)
	}
	if cfg != nil {
		t.Fatalf("cfg=%v, want nil", cfg)
	}
}

func TestDecodeYamlConfig(t *testing.T) {
	cfg, err := config.Decode([]byte("name: demo\n"), "yaml")
	if err != nil {
		t.Fatal(err)
	}
	if cfg["name"] != "demo" {
		t.Fatalf("name=%v, want demo", cfg["name"])
	}
}
