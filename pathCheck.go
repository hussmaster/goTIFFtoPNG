package main

import (
	"fmt"
	"os"
)

// Checks if directory path exists, if not creates path
func pathCheck(path string) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Output folder doesn't exist, creating directory %s\n", path)
			if err := os.Mkdir(path, 0755); err != nil {
				return fmt.Errorf("issue creating directory: %w\n", err)
			}
		} else {
			return fmt.Errorf("issue creating directory: %w\n", err)
		}
	}
	return nil
}
