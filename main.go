package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Please provide the path to find the dir and delete")
		fmt.Println("usage : [path/to/find] [dir_name_to_delete]")
		return
	}

	givenPath := args[0]
	dirName := args[1]
	fmt.Printf("Finding the dir : '%s' in the givenPath : '%s' recursively to proceed with deletion\n", dirName, givenPath)
	// curntWorkingDir, _ := os.Getwd()
	// fmt.Printf("curntWorkingDir: %v\n", curntWorkingDir)
	givenPathDirList, givenpathDirNameList, err := getDirFromGivenPath(givenPath)
	if err != nil {
		fmt.Printf("error while getting Dir from the given Path\n%v\n", err)
		return
	}
	findAndDeleteDirInGivenPath(givenPath, givenPathDirList, givenpathDirNameList, dirName)
}

func getDirFromGivenPath(findDirInPath string) ([]fs.FileInfo, []string, error) {
	files, err := ioutil.ReadDir(findDirInPath)
	if err != nil {
		return nil, nil, err
	}
	var givenPathDirList []fs.FileInfo
	var givenPathDirNameList []string
	for _, file := range files {
		if file.IsDir() {
			givenPathDirList = append(givenPathDirList, file)
			givenPathDirNameList = append(givenPathDirNameList, file.Name())
		}
	}
	return givenPathDirList, givenPathDirNameList, nil
}

func findAndDeleteDirInGivenPath(searchPath string, dirListInGivenPath []fs.FileInfo, dirNameListInGivenPath []string, delDirName string) {
	fmt.Printf("Directories in the given path(%q) : %q\n", searchPath, dirNameListInGivenPath)

	for _, cDir := range dirListInGivenPath {
		if delDirName == cDir.Name() {
			deleteDir(cDir, searchPath)
		} else {
			updatedSerachPath := path.Join([]string{searchPath, cDir.Name()}...)
			dirList, dirNameList, err := getDirFromGivenPath(updatedSerachPath)
			if err != nil {
				fmt.Printf("error while getting dir list from the path\n%v\n", err)
			}
			findAndDeleteDirInGivenPath(updatedSerachPath, dirList, dirNameList, delDirName)
		}

	}
}

func deleteDir(dirObject fs.FileInfo, givenDirPath string) {
	deletionPath := path.Join([]string{givenDirPath, dirObject.Name()}...)
	fmt.Printf("\nDeleting Dir: %q and in the path : %q\n", dirObject.Name(), deletionPath)

	err := os.RemoveAll(deletionPath)
	if err != nil {
		fmt.Printf("errored while deleting the Dir\n %v\n", err)
	}
	fmt.Printf("Deleted Dir: %q and in the path : %q\n\n", dirObject.Name(), deletionPath)

}
