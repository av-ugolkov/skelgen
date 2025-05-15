package config

import (
	"os"
	"regexp"
	"strings"

	"github.com/av-ugolkov/gopkg/logger"
	kw "github.com/av-ugolkov/yask/internal/key-words"

	"gopkg.in/yaml.v3"
)

type config struct {
	skel    map[string]any
	content map[string]any
	dynamic map[string]string
}

var inst config

func Load(path string, dynamic map[string]string) error {
	f, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var mapConfig map[string]any
	err = yaml.Unmarshal(f, &mapConfig)
	if err != nil {
		return err
	}

	dynamicM := make(map[string]string, len(dynamic))
	for k, v := range dynamic {
		dynamicM[k] = v
	}

	inst = config{
		skel:    mapConfig[string(kw.Skel)].(map[string]any),
		dynamic: dynamicM,
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

func AsDynamic(s string) string {
	var value string
	if kw.IsDynamic(s) {
		value = inst.dynamic[extractOnePlaceholder(s)]
		if value != "" {
			return replacePlaceholders(s, value)
		}
	}
	return s
}

func replacePlaceholders(input, replaceValue string) string {
	re := regexp.MustCompile(`\$\{\{.*?\}\}|\$\{.*?\}`)
	return re.ReplaceAllString(input, replaceValue)
}

func extractOnePlaceholder(input string) string {
	re := regexp.MustCompile(`\$\{([^{}]+)\}`)
	match := re.FindStringSubmatch(input)
	if len(match) > 1 {
		return match[1] // значение внутри ${...}
	}
	return ""
}
