package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type LineInfo struct {
	lineNo int
	line string
}

type FindInfo struct {
	filename string
	lines []LineInfo
}

func main() {
	
	if len(os.Args) < 3  {
		fmt.Println("at list 2 more arguments is needed. ex) ex26.1 word filepath")
		return
	}

	word := os.Args[1]
	paths := os.Args[2:]
	findInfos := []FindInfo{}

	for _, path := range paths {
		findInfos = append(findInfos, FindWordInAllFiles(word,path)...)
	}
	for _, findInfo := range findInfos {
		fmt.Println(findInfo.filename)
		fmt.Println("--------------------------------------------------")
		for _, lineInfo := range findInfo.lines {
			fmt.Println("\t", lineInfo.lineNo, "\t", lineInfo.line)
		}
		fmt.Println("--------------------------------------------------")
		fmt.Println()

	}

	fmt.Printf("a word to find : %s , paths : %s \n" , word , paths)
	PrintAllFiles(paths)
	
}


func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

func PrintAllFiles (paths []string) {
	for _ , path := range paths {
		filelist, err := GetFileList(path)
		fmt.Println("get file list result :",filelist)
		if err != nil {
			fmt.Println("incorrect file path. err:", err ,"path:",path)
			return
		}
		fmt.Println("filelist to find")
		for _, name := range filelist {
			fmt.Println(name)
		}
	}
}

func FindWordInAllFiles( word,path string) []FindInfo {
	findInfos := []FindInfo{}

	filelist, err := GetFileList(path)
	if err != nil {
		fmt.Println("invalid file path, err :", err, "path:",path)
		return findInfos
	}
	for _, filename := range filelist {
		findInfos = append(findInfos, FindWordInFile(word,filename))
	}
	return findInfos
}

func FindWordInFile(word,filename string) FindInfo {
	findInfo := FindInfo{filename,[]LineInfo{}}
	file,err := os.Open(filename)
	if err != nil {
		fmt.Println("Can not find a file", filename)
		return findInfo
	}
	defer file.Close()

	lineNo := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		if strings.Contains(line, word){
			findInfo.lines = append(findInfo.lines, LineInfo{lineNo, line})
		}
		lineNo++
	}
	return findInfo


}



func PrintFile(filename string) {
	file,err := os.Open(filename)
	if err != nil {
		fmt.Println("Can't find a file", filename)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	
}