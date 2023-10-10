package controllers

import (
	"TheNexusBattlesII_Card_API_Reborn/models"
	"TheNexusBattlesII_Card_API_Reborn/utils"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ConsumiblePost godoc
// @Router /consumible/ [post]
// @Tags Consumible
// @Summary crea un documento tipo Consumible
// @Description este método recibe los parámetros de tipo Consumible y crea la carta
// @Param imagen formData file true "Imagen de la carta"
// @Param icono formData string true "Icono de la carta"
// @Param nombre formData string true "Nombre de la carta"
// @Param clase formData string true "Clase de la carta"
// @Param tipo formData string true "Tipo de la carta"
// @Param coleccion formData string true "Colección de la carta"
// @Param efecto formData string true "Efecto de la carta"
// @Param efectoHeroe formData string false "Efecto tipo héroe de la carta"
// @Param descripcion formData string true "Descripción de la carta"
// @Param precio formData int true "[Ecommerce] Precio de la carta"
// @Param descuento formData int true "[Ecommerce] Descuento de la carta"
// @Param stock formData int true "[Ecommerce] Stock de la carta"
// @Param estado formData bool true "Activar o desactivar la carta"
// @Accept mpfd
// @Produce json
// @Success 200 {object} map[string]string "Documento creado exitosamente"
// @Failure 400 {object} map[string]string "Parámetros de paginación inválidos"
// @Failure 500 {object} map[string]string "Error interno en el servidor"
func PostConsumible(c echo.Context) error {
	//[1- Validar datos del objeto - asignar valores]
	var consumible models.ConsumibleModel
	consumible_imagen := "tempImageName"
	consumible.Imagen = &consumible_imagen
	//[1.1 validar efectos - asignar valores]
	efecto := c.FormValue("efecto")
	if err := json.Unmarshal([]byte(efecto), &consumible.Efecto); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if efecto_heroe := c.FormValue("efectoHeroe"); len(efecto_heroe) != 0 {
		if err := json.Unmarshal([]byte(efecto_heroe), &consumible.EfectoHeroe); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
	}
	if err := c.Bind(&consumible); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := consumible.Validate(true); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message:": "Bad Request: Invalid parameters"})
	}
	//[2. Validar carga de imagen]
	var imagenName string
	if imagen, err := c.FormFile("imagen"); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Bad Request: Error uploading image"})
	} else {
		if imagenName, err = utils.UploadImageHandler(imagen, []string{".webp"}); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error: Error processing image"})
		}
		consumible.Imagen = &imagenName
	}
	//[3. Subir objeto a la base de datos]
	if err := consumible.PostObject(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error: Error uploading document"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Document created successfully"})
}

// ConsumiblePatch godoc
// @Router /consumible/ [patch]
// @Tags Consumible
// @Summary actualiza un documento tipo Consumible
// @Description este método recibe los parámetros de tipo Consumible y actualiza la carta
// @Param id formData string true "Id de la carta"
// @Param imagen formData file false "Imagen de la carta"
// @Param icono formData string false "Icono de la carta"
// @Param nombre formData string false "Nombre de la carta"
// @Param efecto formData string false "Efecto de la carta"
// @Param efectoHeroe formData string false "Efecto tipo héroe de la carta"
// @Param descripcion formData string false "Descripción de la carta"
// @Param precio formData int false "[Ecommerce] Precio de la carta"
// @Param descuento formData int false "[Ecommerce] Descuento de la carta"
// @Param stock formData int false "[Ecommerce] Stock de la carta"
// @Param estado formData bool false "Activar o desactivar la carta"
// @Accept mpfd
// @Produce json
// @Success 200 {object} map[string]string "Documento creado exitosamente"
// @Failure 400 {object} map[string]string "Parámetros de paginación inválidos"
// @Failure 500 {object} map[string]string "Error interno en el servidor"
func PatchConsumible(c echo.Context) error {
	//[1. Validar el id]
	id := c.FormValue("id")
	id_bson, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message:": "Bad Request: Invalid id"})
	}
	//[2. validar datos -asignar valores]
	var consumible models.ConsumibleModel
	if efecto := c.FormValue("efecto"); len(efecto) != 0 {
		if err := json.Unmarshal([]byte(efecto), &consumible.Efecto); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
	}
	if efecto_heroe := c.FormValue("efectoHeroe"); len(efecto_heroe) != 0 {
		if err := json.Unmarshal([]byte(efecto_heroe), &consumible.EfectoHeroe); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
	}
	if err := c.Bind(&consumible); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message:": "Bad Request: Invalid parameters"})
	}
	if err := consumible.Validate(true); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message:": "Bad Request: Invalid parameters"})
	}
	consumible.Id = id_bson
	//[3. Validar carga de imagen (si la hay)]
	if imagen, err := c.FormFile("imagen"); err != nil && imagen != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Bad Request: Error uploading image"})
	} else if imagen != nil {
		var imagenName string
		if imagenName, err = utils.UploadImageHandler(imagen, []string{".webp"}); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error: Error processing image"})
		}
		consumible.Imagen = &imagenName
	}
	//[4. Actualizar objeto a la base de datos]
	if err := consumible.UpdateObject(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error: Error updating document"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Document updated successfully"})
}
