package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateUniqueName() string {
	// Crear un generador de números aleatorios local
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// Obtener un número aleatorio para evitar colisiones
	randomNumber := random.Intn(1000)

	// Obtener la hora actual en nanosegundos
	timestamp := time.Now().UnixNano()

	// Construir la cadena única
	uniqueName := fmt.Sprintf("%d_%d", timestamp, randomNumber)

	return uniqueName
}
