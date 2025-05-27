package actions

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/av-ugolkov/yask/internal/config"
	keywords "github.com/av-ugolkov/yask/internal/key-words"
	"github.com/av-ugolkov/yask/internal/regex"
)

func CreateFile(rootPath, file string, value string) error {
	file = config.AsPlaceholder(file)
	file = keywords.RemoveInsulator(file)
	value = config.AsPlaceholder(value)

	if !regex.IsValidate(file) {
		return fmt.Errorf("%v: [%s]", regex.ErrInvalidFileName, file)
	}

	f, err := os.Create(path.Join(rootPath, file))
	if err != nil {
		return fmt.Errorf("couldn't create a file [%s] in the folder [%s]: %v", file, rootPath, err)
	}

	switch {
	case strings.HasPrefix(value, string(keywords.FileLink)):
		content, err := os.ReadFile(value[1:])
		if err != nil {
			return err
		}
		return writeBytes(f, content)
	case strings.HasPrefix(value, string(keywords.ContentLink)):
		content := config.GetContent(value)
		return writeString(f, content)
	default:
		return writeString(f, value)
	}
}

func CreateFiles(rootPath string, files map[string]any) error {
	for file, value := range files {
		err := CreateFile(rootPath, file, fmt.Sprintf("%v", value))
		if err != nil {
			return err
		}
	}

	return nil
}

func writeString(file *os.File, content string) error {
	_, err := file.WriteString(content)
	if err != nil {
		return fmt.Errorf("couldn't write the data in the file [%s]: %v", file.Name(), err)
	}

	return nil
}

func writeBytes(file *os.File, content []byte) error {
	_, err := file.Write(content)
	if err != nil {
		return fmt.Errorf("couldn't write the data in the file [%s]: %v", file.Name(), err)
	}

	return nil
}
