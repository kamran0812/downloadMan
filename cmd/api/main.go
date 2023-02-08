package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	//flags
	urlPtr := flag.String("url", "", "valid download url")
	targetPtr := flag.String("t", "", "target path of file including file name")
	connPtr := flag.Int("con", 3, "number of partitions")
	flag.Parse()

	flagValidator(*urlPtr, *targetPtr)

	startTime := time.Now()

	d := NewDownload(*urlPtr, *targetPtr, *connPtr)

	size, err := request(d)
	if err != nil {
		log.Fatalf("error making connection:%s", err)
	}
	err = d.Start(int(size))
	if err != nil {
		log.Fatalf("error occured while downloading..:%s\n", err)
	}

	fmt.Printf("Download Completed In: %v seconds\n", time.Since(startTime).Seconds())

}

func flagValidator(url string, target string) {
	if url == "" || target == "" {
		fmt.Println("url or target cannot be empty.. use: downman -h for help")
		os.Exit(1)
	}
}
