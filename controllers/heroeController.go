package controllers

import (
	"TheNexusBattlesII_Card_API_Reborn/models"
	"TheNexusBattlesII_Card_API_Reborn/utils"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// HeroeGet godoc
// @Router /heroes/{id} [get]
// @Tags Heroes
// @Summary trae un documento tipo Heroe
// @Description este método devuelve un documento tipo Heroe por búsqueda de id
// @Param id path string true "ID del heroe"
// @Produce json
// @Success 200 {object} models.HeroeModel "Documento tipo Heroe"
// @Failure 400 {object} string "Id de búsqueda inválido"
// @Failure 404 {object} string "Documento no existente en la base de datos"
func GetHeroe(c echo.Context) error {
	//[1. Validar el id]
	id := c.Param("id")
	id_bson, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message:": "Bad Request: Invalid id"})
	}
	//[2. Traer el documento]
	var heroe models.HeroeModel
	if err := heroe.GetObject(id_bson); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message:": "404: Document not found"})
	}
	return c.JSON(http.StatusOK, heroe)
}

// HeroeListGet godoc
// @Router /heroes/ [get]
// @Tags Heroes
// @Summary trae una lista de documentos tipo Heroe
// @Description este método devuelve una lista de documentos tipo Heroe
// @Param size query int true "Tamaño de la paginación"
// @Param page query int true "Página de los documentos"
// @Param statusFilter query bool true "Si se activa el filtro, solo trae las cartas activas"
// @Produce json
// @Success 200 {object} []models.HeroeModel "Lista de documentos tipo Heroe"
// @Failure 400 {object} string "Parámetros de paginación inválidos"
// @Failure 500 {object} string "Error interno en el servidor"
func GetHeroeList(c echo.Context) error {
	//[1. Validar paginación]
	var pagination models.PaginationModel
	if err := c.Bind(&pagination); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message:": "Bad Request: Invalid query parameters"})
	}
	if err := pagination.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message:": "Bad Request: Invalid query parameters"})
	}
	//[2. Traer los documentos]
	var heroeList []models.HeroeModel
	if err := models.HeroeModel.GetObjectList(models.HeroeModel{}, pagination, &heroeList); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message:": "Internal Server Error: Documents not found"})
	}
	return c.JSON(http.StatusOK, heroeList)
}

// HeroeListByNameGet godoc
// @Router /heroes/name/ [get]
// @Tags Heroes
// @Summary trae una lista de documentos tipo Heroe filtrado por nombre
// @Description este método devuelve una lista de documentos tipo Heroe filtrado por nombre
// @Param size query int true "Tamaño de la paginación"
// @Param page query int true "Página de los documentos"
// @Param keyword query string true "palabra clave para filtrar"
// @Accept json
// @Produce json
// @Success 200 {object} []models.HeroeModel "Lista de documentos tipo Heroe"
// @Failure 400 {object} string "Parámetros de paginación inválidos"
// @Failure 500 {object} string "Error interno en el servidor"
func GetHeroeListByName(c echo.Context) error {
	//[2. Validar paginación y keyword]
	var pagination models.PaginationModel
	if err := c.Bind(&pagination); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message:": "Bad Request: Invalid parameters"})
	}
	keyword := c.QueryParam("keyword")
	if err := pagination.Validate(); err != nil && keyword == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message:": "Bad Request: Invalid parameters"})
	}
	//[2. Traer los documentos]
	var heroeList []models.HeroeModel
	if err := models.HeroeModel.GetObjectListByName(models.HeroeModel{}, pagination, keyword, &heroeList); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message:": "Internal Server Error: Documents not found"})
	}
	return c.JSON(http.StatusOK, heroeList)
}

// HeroePost godoc
// @Router /heroes/ [post]
// @Tags Heroes
// @Summary crea un documento tipo Heroe
// @Description este método recibe los parámetros de tipo Heroe y crea la carta
// @Param imagen formData file true "Imagen de la carta"
// @Param icono formData string true "Icono de la carta"
// @Param nombre formData string true "Nombre de la carta"
// @Param clase formData string true "Clase de la carta"
// @Param tipo formData string true "Tipo de la carta"
// @Param poder formData int true "Poder de la carta"
// @Param vida formData int true "Vida de la carta"
// @Param defensa formData int true "defensa de la carta"
// @Param ataqueBase formData int true "Ataque base de la carta"
// @Param ataqueRnd formData int true "Ataque aleatorio de la carta"
// @Param daño formData int true "Daño de la carta"
// @Param descripcion formData string true "Descripción de la carta"
// @Param precio formData int true "[Ecommerce] Precio de la carta"
// @Param descuento formData int true "[Ecommerce] Descuento de la carta"
// @Param stock formData int true "[Ecommerce] Stock de la carta"
// @Accept mpfd
// @Produce json
// @Success 200 {object} map[string]string "Documento creado exitosamente"
// @Failure 400 {object} map[string]string "Parámetros de paginación inválidos"
// @Failure 500 {object} map[string]string "Error interno en el servidor"
func PostHeroe(c echo.Context) error {
	//[1. Validar datos del objeto - asignar valores]
	var heroe models.HeroeModel
	heroe.Imagen = &[]string{"tempImageName"}[0]
	heroe.Coleccion = &[]string{"Heroes"}[0]
	heroe.Estado = &[]bool{true}[0]
	if err := c.Bind(&heroe); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message:": "Bad Request: Invalid parameters"})
	}
	if err := heroe.Validate(false); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message:": "Bad Request: Invalid parameters"})
	}
	//[2. Validar carga de imagen]
	var imagenName string
	if imagen, err := c.FormFile("imagen"); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Bad Request: Error uploading image"})
	} else {
		if imagenName, err = utils.UploadImageHandler(imagen, []string{".jpg", ".jpeg", ".png", ".gif"}); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error: Error processing image"})
		}
		heroe.Imagen = &imagenName
	}
	//[4. Subir objeto a la base de datos]
	if err := heroe.PostObject(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error: Error uploading document"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Document created successfully"})
}

// HeroePost godoc
// @Router /heroes/ [patch]
// @Tags Heroes
// @Summary actualiza un documento tipo Heroe
// @Description este método recibe los parámetros de tipo Heroe y actualiza la carta
// @Param id formData string true "Id de la carta"
// @Param imagen formData file false "Imagen de la carta"
// @Param icono formData string false "Icono de la carta"
// @Param nombre formData string false "Nombre de la carta"
// @Param poder formData int false "Poder de la carta"
// @Param vida formData int false "Vida de la carta"
// @Param defensa formData int false "defensa de la carta"
// @Param ataqueBase formData int false "Ataque base de la carta"
// @Param ataqueRnd formData int false "Ataque aleatorio de la carta"
// @Param daño formData int false "Daño de la carta"
// @Param descripcion formData string false "Descripción de la carta"
// @Param precio formData int false "[Ecommerce] Precio de la carta"
// @Param descuento formData int false "[Ecommerce] Descuento de la carta"
// @Param stock formData int false "[Ecommerce] Stock de la carta"
// @Param estado formData bool false "Activar o desactivar la carta"
// @Accept mpfd
// @Produce json
// @Success 200 {object} map[string]string "Documento actualizado exitosamente"
// @Failure 400 {object} map[string]string "Parámetros de paginación inválidos"
// @Failure 500 {object} map[string]string "Error interno en el servidor"
func PatchHeroe(c echo.Context) error {
	//[1. Validar el id]
	id := c.FormValue("id")
	id_bson, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message:": "Bad Request: Invalid id"})
	}
	//[2. Validar datos - asignar valores]
	var heroe models.HeroeModel
	if err := c.Bind(&heroe); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message:": "Bad Request: Invalid parameters"})
	}
	if err := heroe.Validate(true); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message:": "Bad Request: Invalid parameters"})
	}
	heroe.Id = id_bson
	heroe.Coleccion = &[]string{"Heroes"}[0]
	//[3. Validar carga de imagen (si la hay)]
	var imagenName string
	if imagen, err := c.FormFile("imagen"); err != nil && imagen != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Bad Request: Error uploading image"})
	} else if imagen != nil {
		fmt.Println("entre aqui")
		if imagenName, err = utils.UploadImageHandler(imagen, []string{".jpg", ".jpeg", ".png", ".gif"}); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error: Error processing image"})
		}
		heroe.Imagen = &imagenName
	}
	//[4. Actualizar objeto a la base de datos]
	if err := heroe.UpdateObject(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error: Error updating document"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Document updated successfully"})
}
