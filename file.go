package config_file

import (
	"fmt"
	"os"

	. "github.com/infrago/base"
	"github.com/infrago/config"
	"github.com/infrago/infra"
)

type FileConfigDriver struct{}

func init() {
	infra.Register("file", &FileConfigDriver{})
}

func (d *FileConfigDriver) Load(params Map) (Map, error) {
	file := ""
	if vv, ok := params["file"].(string); ok {
		file = vv
	}
	if vv, ok := params["path"].(string); ok {
		file = vv
	}
	if vv, ok := params["config"].(string); ok {
		file = vv
	}
	if file == "" {
		return nil, nil
	}

	data, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("read config file %q: %w", file, err)
	}

	format, _ := params["format"].(string)
	if format == "" {
		format = config.FormatFromPath(file)
	}
	if format == "" {
		format = config.DetectFormat(data)
	}

	return config.Decode(data, format)
}
