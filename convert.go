package main

import (
	"fmt"

	"github.com/sunshineplan/imgconv"
)

func convertImg(inputFile, outputFile string) error {
	//src, err := imgconv.Open("/home/ian/workspace/tempPic/IanChild517.tif")
	src, err := imgconv.Open(inputFile)
	if err != nil {
		return fmt.Errorf("failed to open image: %w", err)
	}
	//err = imgconv.Save("/home/ian/workspace/tempPic/IanChild517.png", src, &imgconv.FormatOption{Format: imgconv.PNG})
	err = imgconv.Save("/mnt/ntfs/wd_hdd/PhotoScans/1992png/test01.png", src, &imgconv.FormatOption{Format: imgconv.PNG})
	if err != nil {
		return fmt.Errorf("failed to open image: %w", err)
	}
	return nil
}
