package contador

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ContadorDeVocales lee una frase y muestra el conteo de vocales (a,e,i,o,u).
// Uso: import "github.com/mat1520/Creacion_de_Paquetes_en_Go/contador"
// luego llamar contador.ContadorDeVocales()
func ContadorDeVocales() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingrese una frase:")
	frase, _ := reader.ReadString('\n')

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
