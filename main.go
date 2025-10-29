// ===========================================
// Minecraft Mod Pack Auto-Updater
// Created by: MacelroyDev
// Date: October 2025
// ===========================================
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("=====================================")
	fmt.Println("   Minecraft Mod Pack Auto-Updater   ")
	fmt.Println("=====================================")

	// Locate the mods folder
	modsPath, err := getModsFolder()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		fmt.Println("Update failed. Press Enter to exit...")
		fmt.Scanln() // Wait for user input to keep the console open
		os.Exit(1)
	}

	// Get the latest URL
	latestURL, err := getLatestModPackURL()
	if err != nil {
		fmt.Printf("ERROR fetching update link: %v\n", err)
		fmt.Println("Update failed. Press Enter to exit...")
		fmt.Scanln()
		os.Exit(1)
	}

	// Download and extract the new files
	if err := downloadAndExtract(modsPath, latestURL); err != nil {
		fmt.Printf("ERROR during download/extraction: %v\n", err)
		fmt.Println("Update failed. Press Enter to exit...")
		fmt.Scanln()
		os.Exit(1)
	}

	fmt.Println("ALL DONE! Your mod pack is now up-to-date!")
	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
}
