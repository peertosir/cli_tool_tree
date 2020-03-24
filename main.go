package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	dir := flag.String("dir", ".", "Directory to show. Default = current dir")
	showFiles := flag.Bool("f", true, "Toggle file view. Default = true")
	flag.Parse()
	result := readDirsAndFilesTree(*dir, "", *showFiles)
	fmt.Println(result)
}

func readDirsAndFilesTree(name, prefix string, showFiles bool) string {
	result := fmt.Sprintf("%v%v/\n", prefix, name)
	file, errorOpen := os.Open(name)
	if errorOpen != nil {
		err := fmt.Errorf("Error with accessing %v occures: ", name)
		fmt.Println(err)
	}
	files, errorReadDir := file.Readdir(0)
	if errorReadDir != nil {
		err := fmt.Errorf("Error with opening dir occures: ")
		fmt.Println(err)
	}
	if len(files) == 0 {
		return result
	}
	for _, value := range files {
		if value.IsDir() == true {
			result += readDirsAndFilesTree(value.Name(), prefix+"\t", showFiles)
		} else {
			result += fmt.Sprintf("%v\t%v\n", prefix, value.Name())
		}
	}
	return result
}
