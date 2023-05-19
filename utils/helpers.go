package utils

import (
	"log"
	"os/user"
	"path/filepath"
	"strings"
)

func ParseTilde(path string) string {
	if !strings.HasPrefix(path, "~") {
		return path
	}

	user, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(user.HomeDir, path[2:])
}
