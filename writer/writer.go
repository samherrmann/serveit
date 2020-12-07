package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"time"
)

func main() {
	err := syscall.Mkfifo("../my-pipe", 0666)
	if err != nil && !os.IsExist(err) {
		log.Fatalln(err)
	}

	// to open pipe to write
	file, err := os.OpenFile("../my-pipe", syscall.O_WRONLY|syscall.O_CREAT, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	for i := 0; i >= 0; i++ {
		_, err = file.WriteString(fmt.Sprint(i) + "\n")
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(i)
		time.Sleep(1 * time.Second / 60)
	}
}
