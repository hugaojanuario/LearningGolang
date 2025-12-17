package main

import (
	"fmt"
)

func main() {

	var noteOneInput int;
	var noteTwoInput int;
	var noteThreeInput int;

	fmt.Print("Informe a primeira nota: ")
	fmt.Scan(&noteOneInput)

	fmt.Print("Informe a segunda nota: ")
	fmt.Scan(&noteTwoInput)

	fmt.Print("Informe a terceira nota: ")
	fmt.Scan(&noteThreeInput)

	mediaNotas(noteOneInput, noteTwoInput, noteThreeInput)
	
}