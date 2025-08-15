package config

import (
	"errors"
	"os"
	"regexp"
	"strings"

	"github.com/av-ugolkov/gopkg/logger"
	kw "github.com/av-ugolkov/yask/internal/key-words"

	"gopkg.in/yaml.v3"
)

type config struct {
	skel         map[string]any
	content      map[string]any
	placeholders map[string]string
}

var inst config

func Load(path string, ph map[string]string) error {
	f, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var mapConfig map[string]any
	err = yaml.Unmarshal(f, &mapConfig)
	if err != nil {
		return err
	}

	if mapConfig[string(kw.Skel)] == nil {
		return errors.New("skeleton is empty")
	}

	placeholders := make(map[string]string, len(ph))
	for k, v := range ph {
		placeholders[k] = v
	}

	inst = config{
		skel:         mapConfig[string(kw.Skel)].(map[string]any),
		placeholders: placeholders,
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

func AsPlaceholder(s string) string {
	var value string
	if hasPlaceholders(s) {
		value = inst.placeholders[extractOnePlaceholder(s)]
		if value != "" {
			return replacePlaceholders(s, value)
		}
	}
	return s
}

var (
	regexReplacePlaceholder = regexp.MustCompile(`\$\{\{.*?\}\}|\$\{.*?\}`)
	regexFindPlaceholder    = regexp.MustCompile(`\$\{([^{}]+)\}`)
)

func replacePlaceholders(input, replaceValue string) string {
	return regexReplacePlaceholder.ReplaceAllString(input, replaceValue)
}

func hasPlaceholders(s string) bool {
	return regexFindPlaceholder.MatchString(s)
}

func extractOnePlaceholder(input string) string {
	match := regexFindPlaceholder.FindStringSubmatch(input)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}
