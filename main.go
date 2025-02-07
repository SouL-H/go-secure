package main

import (
	"log"
	"os"
)

func main() {
	if err:=InSecureCode(); err!=nil{
		log.Println("Error: ", err)
		return
	}
}

func InSecureCode() error{
	err := os.Mkdir("./example", 0777)
	if err != nil {
		return err
	}
	return nil
}
