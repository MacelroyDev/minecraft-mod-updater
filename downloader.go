// ===========================================
// Minecraft Mod Pack Auto-Updater
// Created by: MacelroyDev
// Date: October 2025
// ===========================================
package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/schollz/progressbar/v3"
)

// !!!! CHANGE THIS TO YOUR LINK !!!!
const ManifestURL = "https://www.vanguard-extraction-solutions.com/latest_url.txt"

// Fetch the permanent ManifestURL and returns the direct ZIP download link
func getLatestModPackURL() (string, error) {
	fmt.Printf("Checking for latest mod pack URL from: %s\n", ManifestURL)

	resp, err := http.Get(ManifestURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch manifest URL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status received from manifest: %s", resp.Status)
	}

	// Read the body (should be a single URL string)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read manifest body: %w", err)
	}

	// Trim any whitespace or newlines and convert to string
	return strings.TrimSpace(string(body)), nil
}

func downloadAndExtract(targetPath string, ModPackURL string) error {
	fmt.Println("Starting mod pack download...")

	// Create a temp file for the modpack .zip
	tempFile, err := os.CreateTemp("", "modpack_download_*.zip")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}

	// Clean up temp file on func exit
	// Defer delays until func returns
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	// Download .zip
	resp, err := http.Get(ModPackURL)
	if err != nil {
		return fmt.Errorf("failed to download mod pack: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status received: %s", resp.Status)
	}

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"Downloading Mod Pack:",
	)

	// Copy download to temp file while updating progress bar
	_, err = io.Copy(io.MultiWriter(tempFile, bar), resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save download: %w", err)
	}

	// Extract and overwrite .zip contents
	fmt.Println("\nExtracting mod pack...")

	archive, err := zip.OpenReader(tempFile.Name())
	if err != nil {
		return fmt.Errorf("failed to open downloaded zip: %w", err)
	}
	defer archive.Close()

	for _, f := range archive.File {
		// Only extract jar files
		if !strings.HasSuffix(f.Name, ".jar") {
			continue
		}

		// Only use filename to avoid creating subfolders
		fpath := filepath.Join(targetPath, filepath.Base(f.Name))

		// Skip directories if they for some ungodly reason have a .jar suffix (Looking at you Austin)
		if f.FileInfo().IsDir() {
			continue
		}

		// O_TRUNC will overwrite files and O_CREATE will create them
		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return fmt.Errorf("failed to create file %s: %w", fpath, err)
		}

		rc, err := f.Open()
		if err != nil {
			outFile.Close()
			return fmt.Errorf("failed to open file in zip %s: %w", f.Name, err)
		}

		// Copy data from the .zip to the new one
		_, err = io.Copy(outFile, rc)

		// Close files to release resources
		outFile.Close()
		rc.Close()

		if err != nil {
			return fmt.Errorf("failed to copy data for file %s: %w", f.Name, err)
		}
	}

	fmt.Println("Mod pack updated and extracted successfully!")
	return nil
}
