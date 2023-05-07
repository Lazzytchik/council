package main

import (
	"crypto/md5"
	"encoding/hex"
	"log"
)

func main() {
	log.Println(hex.EncodeToString(md5.New().Sum([]byte("dude"))))

}
