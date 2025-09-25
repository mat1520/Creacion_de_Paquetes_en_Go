package conversor

import "fmt"

// ConversorDivisas convierte dólares a varias monedas.
// Uso: import "github.com/mat1520/Creacion_de_Paquetes_en_Go/conversor"
// luego llamar conversor.ConversorDivisas()
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
		resultado = dolares * 0.85 // Ejemplo de tasa de conversión
		fmt.Printf("%.2f dólares son %.2f Euros\n", dolares, resultado)
	case "LB":
		resultado = dolares * 0.75 // Ejemplo de tasa de conversión
		fmt.Printf("%.2f dólares son %.2f Libras Esterlinas\n", dolares, resultado)
	case "Won":
		resultado = dolares * 1100.00 // Ejemplo de tasa de conversión
		fmt.Printf("%.2f dólares son %.2f Won\n", dolares, resultado)
	case "BTC":
		resultado = dolares * 0.000022 // Ejemplo de tasa de conversión
		fmt.Printf("%.2f dólares son %.6f BTC\n", dolares, resultado)
	default:
		fmt.Println("Moneda no reconocida. Por favor ingrese una moneda válida.")
	}
}
