package main

import (
	"os"
	"path"
	"path/filepath"

	"github.com/av-ugolkov/gopkg/logger"

	"gopkg.in/yaml.v3"
)

func generateStructure(pathFile string, instance any) error {
	f, err := os.ReadFile(pathFile)
	if err != nil {
		return err
	}

	var projectMap map[string]any
	err = yaml.Unmarshal(f, &projectMap)
	if err != nil {
		return err
	}

	// Create the folder structure
	for key, value := range projectMap {
		// Create the top-level folder
		err = os.Mkdir(key, 0755)
		if err != nil && !os.IsExist(err) {
			return err
		}

		// Recursively create subfolders
		createSubfolders(key, value)
	}

	return nil
}

func createSubfolders(parent string, value any) {
	if valueMap, ok := value.(map[string]any); ok {
		for key, subValue := range valueMap {
			if filepath.Ext(key) != "" {
				f, err := os.Create(path.Join(parent, key))
				if err != nil {
					logger.Errorf("%v", err)
				}
				_, err = f.WriteString(subValue.(string))
				if err != nil {
					logger.Errorf("%v", err)
				}
			} else {
				err := os.Mkdir(path.Join(parent, key), 0755)
				if err != nil && !os.IsExist(err) {
					logger.Errorf("%v", err)
				}

				createSubfolders(path.Join(parent, key), subValue)
			}
		}
	}
}
