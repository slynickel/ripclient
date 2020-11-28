package config

import (
	"fmt"
	"os"
	"path"
)

func Get() (string, error) {
	p := os.Getenv("%APPDATA%")
	if p == "" {
		return "", fmt.Errorf("APPDATA path could not be parsed")
	}
	return path.Join(p, ".slynickel", "rip", "config.json"), nil
}
