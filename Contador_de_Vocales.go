package Contador_de_Vocales

import (
	"fmt"
	"strings"
)

// Package Contador_de_Vocales: cuenta vocales en una frase.
// Uso muy simple: importar el paquete y llamar ContadorDeVocales().

// ContadorDeVocales lee una frase desde stdin y muestra el conteo
// de cada vocal encontrada (a, e, i, o, u).
// Ejemplo de uso:
//
//	// en main.go
//	// import "github.com/mat1520/Creacion_de_Paquetes_en_Go"
//	// func main() { Contador_de_Vocales.ContadorDeVocales() }
func ContadorDeVocales() {
	var frase string
	fmt.Println("Ingrese una frase:")
	fmt.Scanln(&frase)

	frase = strings.ToLower(frase)
	vocales := "aeiou"
	contador := make(map[rune]int)

	for _, letra := range frase {
		if strings.ContainsRune(vocales, letra) {
			contador[letra]++
		}
	}

	fmt.Println("Conteo de vocales:")
	for vocal, cantidad := range contador {
		fmt.Printf("%c: %d\n", vocal, cantidad)
	}
}
