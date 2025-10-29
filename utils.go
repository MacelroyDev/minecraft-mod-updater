// ===========================================
// Minecraft Mod Pack Auto-Updater
// Created by: MacelroyDev
// Date: October 2025
// ===========================================
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func getModsFolder() (string, error) {
	// := is a short-declare operator, no need to set type
	home, err := os.UserHomeDir()
	// check for error
	if err != nil {
		return "", err
	}

	var path string
	switch runtime.GOOS {
	case "windows":
		appdata := os.Getenv("APPDATA")
		path = filepath.Join(appdata, ".minecraft", "mods")
	case "darwin": // macOS
		// File path for macOS was given by gemini as I don't use mac lol, may be incorrect
		path = filepath.Join(home, "Library", "Application Support", "minecraft", "mods")
	case "linux":
		path = filepath.Join(home, ".minecraft", "mods")
	default:
		return "", fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	// Double check path exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", fmt.Errorf("minecraft mods folder not found at %s", path)
	}

	return path, nil
}
