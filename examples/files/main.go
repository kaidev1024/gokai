package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main1() {
	f, err := os.Create("./testFile.txt")
	if err != nil {
		log.Println("Error creating file")
		return
	}
	defer f.Close()

	log.Println("Successfully created a file")
}

func main2() {
	f, err := os.OpenFile("./testFile.txt", os.O_RDWR, 6)
	if err != nil {
		log.Println("error opening file")
		return
	}
	defer f.Close()

	n, err := f.WriteString("Hello world\n")
	if err != nil {
		log.Println("error writing file")
		return
	}
	log.Printf("Write %d chars\n", n)
}

func main3() {
	f, err := os.OpenFile("./testFile.txt", os.O_RDWR, 6)
	if err != nil {
		log.Println("error opening file: ", err)
		return
	}
	defer f.Close()
	log.Println("successfully open file")

	n, err := f.Seek(0, io.SeekCurrent) //io.SeekCurrent, io.SeekStart
	fmt.Println("offset: ", n)
	f.WriteString("****")
	n, err = f.Seek(0, io.SeekCurrent) //io.SeekCurrent, io.SeekStart
	fmt.Println("offset: ", n)
	f.WriteAt([]byte("11111"), 5)
}

func main4() {
	f, err := os.OpenFile("./testFile.txt", os.O_RDWR, 6)
	if err != nil {
		log.Println("error opening file: ", err)
		return
	}
	defer f.Close()
	log.Println("successfully open file")

	reader := bufio.NewReader(f)
	for {
		buf, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				log.Println("read complete: ", err)
				return
			}
			log.Println("read error: ", err)
		}
		log.Println(string(buf))
	}
}

func main() {
	var path string
	fmt.Scan(&path)
	fmt.Println(path)
	// open directory
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("OpenFile: ", err)
		return
	}
	defer f.Close()

	infos, err := f.Readdir(-1) // read all
	for _, info := range infos {
		fmt.Println(info.Name(), " is directiory: ", info.IsDir())
	}
}
