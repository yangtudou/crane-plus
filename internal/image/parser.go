package image

import (
	"fmt"
	"strings"
)

func Parse(reference string) (*Image, error) {
	if reference == "" {
		return nil, fmt.Errorf("empty image reference")
	}

	original := reference

	tag := "latest"

	parts := strings.Split(reference, "/")

	last := parts[len(parts)-1]

	if index := strings.LastIndex(last, ":"); index > -1 {
		tag = last[index+1:]
		last = last[:index]
		parts[len(parts)-1] = last
	}

	var registry string
	var namespace string
	var name string

	switch len(parts) {
	case 1:
		registry = "docker.io"
		namespace = "library"
		name = parts[0]

	case 2:
		if isRegistry(parts[0]) {
			registry = parts[0]
			name = parts[1]
		} else {
			registry = "docker.io"
			namespace = parts[0]
			name = parts[1]
		}

	default:
		registry = parts[0]
		namespace = strings.Join(parts[1:len(parts)-1], "/")
		name = parts[len(parts)-1]
	}

	fullName := name

	if namespace != "" {
		fullName = namespace + "/" + name
	}

	return &Image{
		Reference: original + addDefaultTag(original, tag),
		Registry:  NormalizeRegistry(registry),
		Namespace: namespace,
		Name:      fullName,
		Tag:       tag,
	}, nil
}

func addDefaultTag(reference string, tag string) string {
	if strings.Contains(reference, ":") {
		return ""
	}

	if tag == "latest" {
		return ":latest"
	}

	return ""
}

func isRegistry(value string) bool {
	return strings.Contains(value, ".") ||
		strings.Contains(value, ":") ||
		value == "localhost"
}
