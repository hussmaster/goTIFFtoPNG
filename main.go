package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//TODO
//Take in cmdline arg for root of folder you want to convert - done
//second cmdline arg for where you want to save to - done
//Create output folder ie. /home/ian/PhotoScansPNG - done
//Double check that that folder doens't already exist, if it does quit program - done
//After creating output folder as you walk the directory, create folder for each directory - done
//Converts and saves image to png with same filename in the appropriate folder - done
//so PhotoScans/1992/IanChild1.tif saves to PhotoScansPNG/1992/IanChild1.png - done
//Setup program to wait until go routines finish
//Limit go routines to 10 at a time? To not overload cpu/memory
//put in option for different output file formats: prompt for png, jpeg, etc

//implement log file that tracks what files could not be converted

func main() {
	if len(os.Args) != 3 {
		fmt.Println("incorrect number of arguments\nUsage: gotifftopng '<input directory>' '<output directory>'")
		os.Exit(1)
	}

	inputDir := os.Args[1]
	//Checks if desired output directory exists
	err := pathCheck(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	//if _, err := os.Stat(os.Args[2]); err != nil {
	//	if os.IsNotExist(err) {
	//		fmt.Printf("Target output folder does not exist, creating directory %s\n", os.Args[2])
	//		if err := os.Mkdir(os.Args[2], 0755); err != nil {
	//			log.Fatal(err)
	//		}
	//	} else {
	//		log.Fatal(err)
	//	}
	//}
	filepath.WalkDir(inputDir, walk)
}

func walk(input string, d fs.DirEntry, err error) error {
	inputDir := os.Args[1]
	outputDir := os.Args[2]
	//Check for and append trailing slash of output directory
	oDirSuf := strings.HasSuffix(outputDir, "/")
	if oDirSuf == false {
		outputDir += "/"
	}
	fmt.Println(outputDir)
	if err != nil {
		return err
	}
	if !d.IsDir() {
		//Get nested folder of original path
		nestedFolder := strings.Split(input, inputDir)
		outputDir += nestedFolder[1]
		fmt.Printf("Converting %s to %s...\n", input, outputDir)
		go convertImg(input, outputDir)
		//if err != nil {
		//	fmt.Println(err)
		//}
		//fmt.Print("Press 'Enter' to continue...")
		//bufio.NewReader(os.Stdin).ReadBytes('\n')
		//count += 1
	} else {
		println(input)
		//slashIndex := strings.LastIndex(input, "/")
		//name := input[slashIndex+1:]
		//fmt.Println(name)
		//newOutputDir := outputDir + name + "/"
		newOutputDir := oDirAppend(input, outputDir)
		fmt.Println(newOutputDir)
		//Path check for nested folder
		err = pathCheck(newOutputDir)
		if err != nil {
			log.Fatal(err)
		}
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
	return nil
}
