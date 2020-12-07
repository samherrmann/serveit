package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
	"syscall"
)

func main() {
	err := syscall.Mkfifo("../my-pipe", 0666)
	if err != nil && !os.IsExist(err) {
		log.Fatalln(err)
	}

	//to open pipe to read
	file, err := os.OpenFile("../my-pipe", syscall.O_RDONLY, os.ModeNamedPipe)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		data, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(strings.TrimSuffix(data, "\n"))
	}
}
