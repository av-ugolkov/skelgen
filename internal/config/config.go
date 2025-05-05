package config

import (
	"os"
	"strings"

	"github.com/av-ugolkov/gopkg/logger"
	kw "github.com/av-ugolkov/yask/internal/key-words"

	"gopkg.in/yaml.v3"
)

type config struct {
	skel    map[string]any
	content map[string]any
}

var inst config

func Load(path string) error {
	f, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var mapConfig map[string]any
	err = yaml.Unmarshal(f, &mapConfig)
	if err != nil {
		return err
	}

	inst = config{
		skel: mapConfig[string(kw.Skel)].(map[string]any),
	}

	mapContent := mapConfig[string(kw.Content)]
	if mapContent != nil {
		inst.content = mapContent.(map[string]any)
	}

	return nil
}

func Skel() map[string]any {
	return inst.skel
}

func GetContent(path string) string {
	pathToContent := strings.Split(path, "/")
	content := inst.content
	for _, p := range pathToContent[2:] {
		if c, ok := content[p]; ok {
			switch c.(type) {
			case map[string]any:
				content = c.(map[string]any)
			case string:
				return c.(string)
			default:
				logger.Warnf("unknown type")
				return ""
			}
		}
	}

	logger.Warnf("not found content by path [%s]", path)
	return ""
}
