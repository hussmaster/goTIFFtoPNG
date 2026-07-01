# goTIFFtoPNG
Go program that utilizes the Go library [imgconv](https://github.com/sunshineplan/imgconv).\
I created this to solve a small problem for myself which is to convert large TIFF images from a large repository of scanned childhood photos to PNG for smaller, relatively, sharing.\
I also utilized this as a small way to dip my toes into go routines so rather than iterating through a directory structure of images it can be done simultaneously.\
This also is not exclusively limited to just TIFF files as the input, right now it just takes the input file and outputs to PNG.

## Requirements
Go version >= 1.25

## Setup
To install this project you can run 
```
go install github.com/hussmaster/goTIFFtoPNG
```
Alternatively you can clone the project and run a go build from the root of the project directory

### Usage

It's a fairly simple program that takes in an input directory of your images, and then an output directory of where you want the converted images to be saved.\
It will also recreate the directory structure of the input directory
<b>Example usage</b>
go run . /home/username/scannedPhotos/ /home/username/output/\
You will also be prompted for how many you would like to consecutively convert at once to either strain or lessen the load on the machine.

### Improvements
Small improvement would be to also prompt the user what output extension rather than a hardcoded PNG file format.\
