package main

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

func generateSkeleton(pathFile string, instance any) error {
	f, err := os.ReadFile(pathFile)
	if err != nil {
		return err
	}

	var mapConfig map[string]any
	err = yaml.Unmarshal(f, &mapConfig)
	if err != nil {
		return err
	}

	switch mapConfig["ver"].(int) {
	case 1:
		projectMap := make(map[string]any, len(mapConfig)-1)
		copyMap(mapConfig, projectMap)
		return genFolders(projectMap)
	default:
		return errors.New("unknown version")
	}
}

func copyMap(map1 map[string]any, map2 map[string]any) {
	for key, value := range map1 {
		if key == "ver" {
			continue
		}
		map2[key] = value
	}
}
