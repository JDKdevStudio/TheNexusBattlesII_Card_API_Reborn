package utils

import (
	"os"
	"path/filepath"
)

func DeleteImageHandler(filename *string) error {
	// Carpeta donde se encuentra la imagen
	imageFolder := "assets" // Reemplaza con la ruta correcta

	// Ruta completa de la imagen a borrar
	imagePath := filepath.Join(imageFolder, *filename)

	// Verifica si el archivo existe antes de intentar borrarlo
	_, err := os.Stat(imagePath)
	if os.IsNotExist(err) {
		return err
	}

	// Intenta borrar la imagen
	if err := os.Remove(imagePath); err != nil {
		return err
	}

	return nil
}
