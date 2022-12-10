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

func getFiles(dir string, fileExtension string) (map[int][]file, error) {
	var files = make(map[int][]file)

	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}

		if fileExtension == "." || filepath.Ext(path) == fileExtension {
			intSize := int(info.Size())
			files[intSize] = append(files[intSize], file{Size: intSize, Path: path})
		}
		return err
	})

	return files, err
}

func printSortedFiles(files map[int][]file, sortingOption int) {
	var sortedKeys []int
	for size := range files {
		sortedKeys = append(sortedKeys, size)
	}

	if sortingOption == 1 {
		sort.Sort(sort.Reverse(sort.IntSlice(sortedKeys)))
	} else {
		sort.Ints(sortedKeys)
	}

	for _, size := range sortedKeys {
		sort.Slice(files[size], func(i, j int) bool {
			return files[size][i].Path > files[size][i].Path
		})
	}

	for _, size := range sortedKeys {
		fmt.Printf("%d bytes\n", size)
		//fmt.Println(strings.Join(files[size], "\n"))
		for i := range files[size] {
			fmt.Println(files[size][i].Path)
		}
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

func getDuplicateFiles(files map[int][]string) {
	//duplicateFiles := make(map[int][]string)
	md5Hash := md5.New()

	for size, paths := range files {
		fmt.Println(size)
		for i := range paths {
			file, _ := os.Open(paths[i])
			io.Copy(md5Hash, file)
			fmt.Println(md5Hash.Sum(nil))
			md5Hash.Reset()
		}
	}
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

	printSortedFiles(files, sortingOption)

	//shouldPrintDuplicateFiles, err := getShouldPrintDuplicateFiles(reader)
	//if shouldPrintDuplicateFiles {
		//getDuplicateFiles(files)
	//}
}
