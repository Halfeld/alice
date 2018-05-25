package main

import (
	"fmt"

	"github.com/alice/googleApis"
)

func main() {
	// text := "Olá, meu nome é Igor! Tudo bem contigo?"
	// text := "Eu odeio java, é realmente muito ruim"
	// googleApis.AnalizeSentiment(text)

	faces, _ := googleApis.DetectFaces("./metadata/igor.jpg")

	fmt.Printf("Faces: %v\n", faces)
}
