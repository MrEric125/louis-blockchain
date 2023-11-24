package main

import "fmt"

type Mac struct {

}

func (*Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}