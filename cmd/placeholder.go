package cmd

import (
	"fmt"
	"strings"
)

type kvPairs map[string]string

func (k *kvPairs) String() string {
	var result []string
	for key, val := range *k {
		result = append(result, fmt.Sprintf("%s=%s", key, val))
	}
	return strings.Join(result, ", ")
}

func (k *kvPairs) Set(value string) error {
	parts := strings.SplitN(value, "=", 2)
	if len(parts) != 2 {
		return fmt.Errorf("invalid format: expected key=value, got %q", value)
	}
	if *k == nil {
		*k = make(map[string]string)
	}
	(*k)[parts[0]] = parts[1]
	return nil
}

func (k *kvPairs) Type() string {
	return "key=value"
}
