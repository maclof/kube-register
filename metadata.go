package main

import (
	"errors"
	"strings"

	"github.com/coreos/fleet/machine"
)

func hasMetadata(ms machine.MachineState, metadata map[string][]string) bool {
	for k, v := range metadata {
		for _, s := range v {
			if ms.Metadata[k] == s {
				return true
			}
		}
	}
	return false
}

func parseMetadata(rawMetadata string) (map[string][]string, error) {
	metadataList := strings.Split(rawMetadata, ",")
	metadata := make(map[string][]string)
	for _, kv := range metadataList {
		i := strings.Index(kv, "=")
		if i > 0 {
			if _, ok := metadata[kv[:i]]; !ok {
				metadata[kv[:i]] = []string{}
			}
			metadata[kv[:i]] = append(metadata[kv[:i]], kv[i+1:])
		} else {
			return nil, errors.New("invalid key/value pair " + kv)
		}
	}
	return metadata, nil
}
