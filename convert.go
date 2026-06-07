package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/sunshineplan/imgconv"
)

func convertImg(inputFile, outputFile string) error {
	//src, err := imgconv.Open("/home/ian/workspace/tempPic/IanChild517.tif")
	src, err := imgconv.Open(inputFile)
	if err != nil {
		return fmt.Errorf("failed to open image: %w", err)
	}
	fileBaseName := filepath.Base(inputFile)
	//Slice off old extension to make way for new extension
	newExt := strings.TrimSuffix(fileBaseName, filepath.Ext(fileBaseName))
	pngExt := ".png"
	newExt += pngExt
	outputFile += newExt
	fmt.Printf("converting %s to %s\n", inputFile, outputFile)

	//err = imgconv.Save("/home/ian/workspace/tempPic/IanChild517.png", src, &imgconv.FormatOption{Format: imgconv.PNG})
	err = imgconv.Save(outputFile, src, &imgconv.FormatOption{Format: imgconv.PNG})
	if err != nil {
		return fmt.Errorf("failed to open image: %w", err)
	}
	return nil
}
