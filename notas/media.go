package main

import (
	"fmt"
)

func mediaNotas(noteOne int, noteTwo int, noteThree int) int {

	mediaNotes := (noteOne + noteTwo + noteThree) / 3

	fmt.Printf("A Media das notas é de:%d ", mediaNotes)

	return mediaNotes

}
