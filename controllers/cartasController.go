package controllers

import (
	"TheNexusBattlesII_Card_API_Reborn/functions"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CartasGet godoc
// @Router /cartas/getid/{id} [get]
// @Tags Cartas
// @Summary trae un documento tipo Carta
// @Description este método devuelve un documento tipo Carta por búsqueda de id
// @Param id path string true "ID de la Carta"
// @Produce json
// @Success 200 {object} interface{} "Documento tipo Carta"
// @Failure 400 {object} string "Id de búsqueda inválido"
// @Failure 404 {object} string "Documento no existente en la base de datos"
func GetObjectById(c echo.Context) error {
	//[1. Validar el id]
	id := c.Param("id")
	id_bson, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message:": "Bad Request: Invalid id"})
	}
	//[2. Traer el documento]
	document := make(map[string]interface{})
	if err := functions.GetObjectById(id_bson, document); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message:": "404: Document not found"})
	}
	return c.JSON(http.StatusOK, document)
}
