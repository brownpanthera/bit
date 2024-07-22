package main

import (
	"fmt"
	"os"

	"github.com/jackpal/bencode-go"
)

var url string = "puppy.torrent"

func main() {
	iodata, err := os.Open(url)
	if err != nil {
		return
	}

	content, err := os.ReadFile(url)
	if err != nil {
		return
	}
	fmt.Println("Raw data:", string(content))

	decoderData, err := bencode.Decode(iodata)
	if err != nil {
		return
	}
	fmt.Println("bencode:", decoderData)

	err = os.WriteFile("torrent.txt", content, 0644)
	if err != nil {
		return
	}

}
