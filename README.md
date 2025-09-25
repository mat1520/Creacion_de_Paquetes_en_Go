# Conversor de Divisas (Ejemplo de paquete en Go)

Este repositorio contiene un ejemplo sencillo de una función para convertir
una cantidad en dólares (USD) a distintas monedas. Está pensado como un
pequeño paquete didáctico para aprender a crear paquetes en Go.

Repositorio: https://github.com/mat1520/Creacion_de_Paquetes_en_Go

Contenido relevante:
- `Conversor_Divisas.go` - Implementación y documentación de ejemplo.

Cómo probarlo (pasos mínimos)

1) Desde el directorio del proyecto, inicializa el módulo de Go (si no existe):

```bash
go mod init github.com/mat1520/Creacion_de_Paquetes_en_Go
```

2) Crea un `main.go` de ejemplo (archivo nuevo) con este contenido:

```go
package main

import "github.com/mat1520/Creacion_de_Paquetes_en_Go"

func main() {
    Conversor_Divisas.ConversorDivisas()
}
```

3) Ejecuta el ejemplo:

```bash
go run main.go
```

Notas
- Las tasas de conversión incluidas son ejemplos estáticos. Para uso real
  deberías obtener tasas actualizadas desde una API externa.
- Si prefieres, puedes mover `Conversor_Divisas.go` a un subdirectorio
  (por ejemplo `conversor`) y ajustar las rutas de importación en el `main.go`.

Contribuciones y licencias
- Este repositorio es un ejemplo; añade mejoras como tests, manejo de errores
  y llamadas a APIs para hacerlo más robusto.
