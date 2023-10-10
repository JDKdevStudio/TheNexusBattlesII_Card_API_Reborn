package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func UploadImageHandler(file *multipart.FileHeader, allowedExtensions []string) (string, error) {
	// Obtén la extensión del archivo
	ext := filepath.Ext(file.Filename)

	// Comprueba si la extensión está en la lista de extensiones permitidas
	var isAllowedExtension bool
	for _, allowedExt := range allowedExtensions {
		if ext == allowedExt {
			isAllowedExtension = true
			break
		}
	}

	if !isAllowedExtension {
		return "", fmt.Errorf("la extensión del archivo no es válida")
	}

	// Carpeta de destino donde se almacenará la imagen
	destinationFolder := os.Getenv("ASSETSFOLDER")
	// Nombre de archivo deseado
	destinationFilename := GenerateUniqueName() + ext
	// Ruta completa del archivo de destino
	destinationPath := filepath.Join(destinationFolder, destinationFilename)

	// Crea la carpeta de destino si no existe
	if err := os.MkdirAll(destinationFolder, 0755); err != nil {
		return "", err
	}

	// Abre el archivo de origen
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Crea el archivo de destino
	dst, err := os.Create(destinationPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copia el contenido desde el archivo de origen al archivo de destino
	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	// Retorna la ruta completa del archivo de destino para su posterior uso si es necesario
	return destinationFilename, nil
}
