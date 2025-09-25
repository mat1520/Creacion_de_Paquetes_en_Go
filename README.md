# Conversor de Divisas y Contador de Vocales (Ejemplo de paquetes en Go)

Este repositorio contiene ejemplos sencillos de funciones implementadas como paquetes en Go. Está pensado como un pequeño proyecto didáctico para aprender a crear y usar paquetes en Go.

Repositorio: [https://github.com/mat1520/Creacion_de_Paquetes_en_Go](https://github.com/mat1520/Creacion_de_Paquetes_en_Go)

## Contenido relevante

- `conversor/conversor.go`: Implementación de un conversor de divisas.
- `contador/contador.go`: Implementación de un contador de vocales.
- `src/main.go`: Ejemplo de uso interactivo de los paquetes.

## Cómo clonar y probar el proyecto

1. Clona este repositorio:

```bash
git clone https://github.com/mat1520/Creacion_de_Paquetes_en_Go.git
cd Creacion_de_Paquetes_en_Go
```

1. Inicializa el módulo de Go (si no lo has hecho):

```bash
go mod init github.com/mat1520/Creacion_de_Paquetes_en_Go
```

1. Ejecuta el archivo `main.go` para probar las funcionalidades:

```bash
go run src/main.go
```

## Uso de los paquetes

### Conversor de Divisas

El paquete `conversor` permite convertir dólares a varias monedas. Para usarlo:

```go
import "github.com/mat1520/Creacion_de_Paquetes_en_Go/conversor"

func main() {
    conversor.ConversorDivisas()
}
```

### Contador de Vocales

El paquete `contador` permite contar las vocales en una frase ingresada por el usuario. Para usarlo:

```go
import "github.com/mat1520/Creacion_de_Paquetes_en_Go/contador"

func main() {
    contador.ContadorDeVocales()
}
```