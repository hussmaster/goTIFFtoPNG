package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/sunshineplan/imgconv"
)

func convertImg(inputFile, inputDir, outputFile string) error {
	//Strip out filename from outputFile to check for path
	outputDir := filepath.Dir(outputFile)
	// Check for trailing slash
	oDirSuf := strings.HasSuffix(outputDir, "/")
	if oDirSuf == false {
		outputDir += "/"
	}
	// Get nested folder from input directory
	nestedFolder := strings.Split(inputFile, inputDir)
	// Open the source file
	src, err := imgconv.Open(inputFile)
	if err != nil {
		return fmt.Errorf("failed to open image: %w", err)
	}
	fileBaseName := filepath.Base(inputFile)
	// Split the nested folder from the filename
	nestedFolder = strings.Split(nestedFolder[1], fileBaseName)
	outputDir += nestedFolder[0]
	// Check and create filepath if needed
	pathCheck(outputDir)
	// Again check for trailing slash before re-adding new filename
	oDirSuf = strings.HasSuffix(outputDir, "/")
	if oDirSuf == false {
		outputDir += "/"
	}
	//Slice off old extension to make way for new extension
	newExt := strings.TrimSuffix(fileBaseName, filepath.Ext(fileBaseName))
	// Make this configurable in the future
	pngExt := ".png"
	newExt += pngExt
	outputDir += newExt
	fmt.Printf("Converting %s to %s...\n", inputFile, outputDir)
	// Start the image conversion
	err = imgconv.Save(outputDir, src, &imgconv.FormatOption{Format: imgconv.PNG})
	if err != nil {
		return fmt.Errorf("failed to open image: %w", err)
	}
	return nil
}
