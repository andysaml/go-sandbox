package main

import (
	"fmt"
	"os"
)

func openFile(filename string) {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		//return
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		//return
	}
	fmt.Println("File info:", fileinfo)

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	bytesread, err := file.Read(buffer)
	if err != nil {
		fmt.Println(err)
		//return
	}

	fmt.Println("bytes read: ", bytesread)
	//fmt.Println("bytestream to string: ", string(buffer))

}
