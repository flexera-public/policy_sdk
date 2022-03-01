//go:build !windows
// +build !windows

package config

import (
	"os/user"
	"path/filepath"
)

func DefaultFile(name string) string {
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	return filepath.Join(currentUser.HomeDir, name)
}
