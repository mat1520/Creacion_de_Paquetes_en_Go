package Conversor_Divisas

import "fmt"

// Package Conversor_Divisas proporciona una función sencilla para convertir
// una cantidad ingresada en dólares (USD) a varias monedas de ejemplo.
// Está pensado como ejemplo para la creación de paquetes en Go.

// Repositorio: https://github.com/mat1520/Creacion_de_Paquetes_en_Go

// Comandos recomendados (desde el directorio del proyecto):
// 1. Inicializar el módulo (si aún no existe):
//    go mod init github.com/mat1520/Creacion_de_Paquetes_en_Go
// 2. Crear un archivo `main.go` de ejemplo (ver ejemplo comentado al final)
// 3. Ejecutar el ejemplo:
//    go run main.go

// Ejemplo de uso (crear un `main.go` y ejecutar `go run main.go`):
// package main

// import "github.com/mat1520/Creacion_de_Paquetes_en_Go"

//	func main() {
//	    Conversor_Divisas.ConversorDivisas()
//	}

func ConversorDivisas() {
	var dolares float64
	var moneda string
	var resultado float64

	fmt.Println("Ingrese la cantidad en dólares:")
	fmt.Scanln(&dolares)

	fmt.Println("Ingrese la moneda a la que desea convertir (Euros, LB, Won, BTC):")
	fmt.Scanln(&moneda)

	switch moneda {
	case "Euros":
		resultado = dolares * 0.85
		fmt.Printf("%.2f dólares son %.2f Euros\n", dolares, resultado)
	case "LB":
		resultado = dolares * 0.75
		fmt.Printf("%.2f dólares son %.2f Libras Esterlinas\n", dolares, resultado)
	case "Won":
		resultado = dolares * 1100.00
		fmt.Printf("%.2f dólares son %.2f Won\n", dolares, resultado)
	case "BTC":
		resultado = dolares * 0.000022
		fmt.Printf("%.2f dólares son %.6f BTC\n", dolares, resultado)
	default:
		fmt.Println("Moneda no reconocida. Por favor ingrese una moneda válida.")
	}
}
