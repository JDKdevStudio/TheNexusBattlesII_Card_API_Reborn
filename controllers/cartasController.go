package controllers

import (
	"TheNexusBattlesII_Card_API_Reborn/functions"
	"TheNexusBattlesII_Card_API_Reborn/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CartasGet godoc
// @Router /cartas/{id} [get]
// @Tags Cartas
// @Summary trae un documento tipo Carta
// @Description este método devuelve un documento tipo Carta por búsqueda de id
// @Param id path string true "ID de la Carta"
// @Produce json
// @Success 200 {object} interface{} "Documento tipo Carta"
// @Failure 400 {object} string "Id de búsqueda inválido"
// @Failure 404 {object} string "Documento no existente en la base de datos"
func GetObject(c echo.Context) error {
	//[1. Validar el id]
	id := c.Param("id")
	id_bson, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message:": "Bad Request: Invalid id"})
	}
	//[2. Traer el documento]
	document := make(map[string]interface{})
	if err := functions.GetObject(id_bson, document); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message:": "404: Document not found"})
	}
	return c.JSON(http.StatusOK, document)
}

// CartasListGet godoc
// @Router /cartas/ [get]
// @Tags Cartas
// @Summary trae una lista de documentos tipo Carta
// @Description este método devuelve una lista de documentos tipo Carta
// @Param size query int true "Tamaño de la paginación"
// @Param page query int true "Página de los documentos"
// @Param keyword query string false "palabra clave para filtrar por nombre"
// @Param coleccion query string true "Coleccion del documento"
// @Param onlyActives query bool true "Si está activo trae SOLO los documentos activos"
// @Accept json
// @Produce json
// @Success 200 {object} []map[string]interface{} "Lista de documentos tipo Carta"
// @Failure 400 {object} string "Parámetros de paginación inválidos"
// @Failure 500 {object} string "Error interno en el servidor"
func GetObjectList(c echo.Context) error {
	//[2. Validar paginación]
	var pagination models.PaginationModel
	if err := c.Bind(&pagination); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message:": "Bad Request: Invalid parameters"})
	}
	if err := pagination.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message:": "Bad Request: Invalid parameters"})
	}
	//[2. Traer los documentos]
	list := make([]map[string]interface{}, 0)
	if err := functions.GetObjectList(pagination, &list); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message:": "Internal Server Error: Documents not found"})
	}
	return c.JSON(http.StatusOK, list)
}
