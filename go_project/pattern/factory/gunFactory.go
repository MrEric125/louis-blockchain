package main

import "fmt"

func getGun(gunType string) (IGun,error)  {

	if gunType=="ak47" {
		return newAk47(),nil
	}
	if gunType=="musket" {
		return newMusket(),nil
	}

	return nil,fmt.Errorf("Wrong gun type passed")

}

func main() {
	ak47,_:=getGun("ak47")
	musket,_:=getGun("musket")

	printDetail(ak47)

	printDetail(musket)


}

func printDetail(g IGun)  {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()

}