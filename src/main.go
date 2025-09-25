package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mat1520/Creacion_de_Paquetes_en_Go/Contador_de_Vocales"
	"github.com/mat1520/Creacion_de_Paquetes_en_Go/Conversor_Divisas"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("--- Menú de ejemplo ---")
		fmt.Println("1) Conversor de divisas")
		fmt.Println("2) Contador de vocales")
		fmt.Println("0) Salir")
		fmt.Print("Seleccione una opción: ")

		var opcion int
		_, err := fmt.Fscanln(reader, &opcion)
		if err != nil {
			// limpiar buffer y continuar
			reader.ReadString('\n')
			fmt.Println("Entrada inválida")
			continue
		}

		switch opcion {
		case 1:
			Conversor_Divisas.ConversorDivisas()
		case 2:
			Contador_de_Vocales.ContadorDeVocales()
		case 0:
			fmt.Println("Adiós")
			return
		default:
			fmt.Println("Opción no válida")
		}
	}
}
