package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"
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
	//Channel to send file names to
	var wg sync.WaitGroup
	fileChan := make(chan string, 300)
	inputDir := os.Args[1]
	//Checks if desired output directory exists
	err := pathCheck(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	// Start worker pool, get number of parallel options from user
	fmt.Println("Enter number of parallel image conversions:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	numWorkers, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Print("must be valid integer\n")
		log.Fatal(err)
	}
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go workerConvert(fileChan, &wg)
	}
	// Walk directory and send paths to the channel
	filepath.WalkDir(inputDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			fileChan <- path
		}
		return nil
	})

	if err != nil {
		fmt.Printf("error walking path %v: %v\n", inputDir, err)
	}
	// Close channel once paths are exhausted
	close(fileChan)
	// Wait to end program until all go routines/conversions have finished
	wg.Wait()
	fmt.Println("Image converting complete!")
}

/*
	func walk(input string, d fs.DirEntry, err error) error {
		//var wg sync.WaitGroup
		inputDir := os.Args[1]
		outputDir := os.Args[2]
		//Check for and append trailing slash of output directory
		oDirSuf := strings.HasSuffix(outputDir, "/")
		if oDirSuf == false {
			outputDir += "/"
		}
		//fmt.Println(outputDir)
		if err != nil {
			return err
		}
		if !d.IsDir() {
			//Get nested folder of original path
			//fmt.Printf("input: %s inputDir: %s\n", input, inputDir)
			nestedFolder := strings.Split(input, inputDir)
			//fmt.Printf("nestedFolder: %v\n", nestedFolder[1])
			//fmt.Printf("outputDir: %s\n", outputDir)
			outputDir += nestedFolder[1]
			fmt.Printf("outputDir: %s\n", outputDir)
			//fmt.Printf("Converting %s to %s...\n", input, outputDir)
			//bufio.NewReader(os.Stdin).ReadBytes('\n')
			convertImg(input, outputDir)

		}
		return nil
	}
*/
// Function to process paths sent to the channel
func workerConvert(files <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	inputDir := os.Args[1]
	outputDir := os.Args[2]
	for path := range files {
		//bufio.NewReader(os.Stdin).ReadBytes('\n')
		convertImg(path, inputDir, outputDir)
	}
}
