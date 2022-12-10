package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type file struct {
	Size int
	Path string
	Hash string
}

func (f *file) hash() {
	file, err := os.Open(f.Path)
	if err != nil {
		log.Fatal(err)
	}
	md5Hash := md5.New()
	_, err = io.Copy(md5Hash, file)
	if err != nil {
		log.Fatal(err)
	}
	f.Hash = fmt.Sprintf("%x", md5Hash.Sum(nil))
}

func getInput(reader *bufio.Reader) (string, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(input, "\n"), nil
}

func getFileExtension(reader *bufio.Reader) (string, error) {
	fmt.Println("Enter file format:")
	fileFormat, err := getInput(reader)
	return "." + fileFormat, err
}

func getSortingOption(reader *bufio.Reader) (int, error) {
	fmt.Println("Enter a sorting option:")
	input, err := getInput(reader)
	sortingOption, err := strconv.Atoi(input)

	for sortingOption != 1 && sortingOption != 2 {
		fmt.Println("Wrong option")
		return getSortingOption(reader)
	}
	return sortingOption, err
}

func getFiles(dir string, fileExtension string) ([]file, error) {
	var files []file

	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}

		if fileExtension == "." || filepath.Ext(path) == fileExtension {
			intSize := int(info.Size())
			files = append(files, file{Size: intSize, Path: path})
		}
		return err
	})

	return files, err
}

func sortFilesBySize(files []file, sortingOption int) []file {
	sort.Slice(files, func(i, j int) bool {
		if sortingOption == 1 {
			return files[i].Size > files[j].Size
		}
		return files[i].Size < files[j].Size
	})
	return files
}

func printSortedFiles(files []file) {
	lastSize := 0
	for i := range files {
		if lastSize != files[i].Size {
			lastSize = files[i].Size
			fmt.Printf("%d bytes\n", lastSize)
		}
		fmt.Println(files[i].Path)
	}
}

func getShouldPrintDuplicateFiles(reader *bufio.Reader) (bool, error) {
	fmt.Println("Check for duplicates?")
	input, err := getInput(reader)
	if input != "yes" && input != "no" {
		return getShouldPrintDuplicateFiles(reader)
	}
	return input == "yes", err
}

func getDuplicateFiles(files []file) []file {
	var duplicateFiles []file

	var hashes = make(map[string]int)
	for i := range files {
		files[i].hash()
		hashes[files[i].Hash] += 1
	}

	for hash, amount := range hashes {
		if amount <= 1 {
			delete(hashes, hash)
		}
	}

	for i := range files {
		if _, ok := hashes[files[i].Hash]; ok {
			duplicateFiles = append(duplicateFiles, files[i])
		}
	}

	return duplicateFiles
}

func printDuplicateFiles(files []file) {
	lastSize := 0
	lastHash := ""
	for i := range files {
		if lastSize != files[i].Size {
			lastSize = files[i].Size
			fmt.Printf("%d bytes\n", lastSize)
		}
		if lastHash != files[i].Hash {
			lastHash = files[i].Hash
			fmt.Printf("Hash: %s\n", files[i].Hash)
		}
		fmt.Printf("%d. %s\n", i+1, files[i].Path)
	}
}

func getShouldDeleteFiles(reader *bufio.Reader) (string, error) {
	fmt.Println("Delete files?")
	input, err := getInput(reader)
	if input != "yes" && input != "no" {
		fmt.Println("Wrong option")
		return getShouldDeleteFiles(reader)
	}
	return input, err
}

func deleteFiles(reader *bufio.Reader, files []file) []int {
	fmt.Println("Enter file numbers to delete:")
	input, err := getInput(reader)

	rawSlice := strings.Split(input, " ")
	filesToDelete := make([]int, len(rawSlice))
	for i, s := range rawSlice {
		filesToDelete[i], err = strconv.Atoi(s)
		filesToDelete[i] -= 1
	}

	for i := range filesToDelete {
		if filesToDelete[i] > len(files) {
			err = fmt.Errorf("file %d not found", filesToDelete[i])
		}
	}

	if err != nil {
		fmt.Println("Wrong format")
		return deleteFiles(reader, files)
	}

	var freedSpace int
	for i := range filesToDelete {
		freedSpace += files[filesToDelete[i]].Size
		os.Remove(files[filesToDelete[i]].Path)
	}

	fmt.Printf("Total freed up space: %d bytes", freedSpace)

	return filesToDelete
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Directory is not specified")
		os.Exit(1)
	}

	dir := os.Args[1]

	reader := bufio.NewReader(os.Stdin)

	fileExtension, err := getFileExtension(reader)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Size sorting options:")
	fmt.Println("1. Descending")
	fmt.Println("2. Ascending")

	sortingOption, err := getSortingOption(reader)
	if err != nil {
		log.Fatal(err)
	}

	files, err := getFiles(dir, fileExtension)
	if err != nil {
		log.Fatal(err)
	}

	sortFilesBySize(files, sortingOption)
	printSortedFiles(files)

	shouldPrintDuplicateFiles, err := getShouldPrintDuplicateFiles(reader)
	if shouldPrintDuplicateFiles {
		duplicateFiles := getDuplicateFiles(files)
		printDuplicateFiles(duplicateFiles)

		if shouldDelete, _ := getShouldDeleteFiles(reader); shouldDelete == "yes" {
			deleteFiles(reader, duplicateFiles)
		}
	}
}
