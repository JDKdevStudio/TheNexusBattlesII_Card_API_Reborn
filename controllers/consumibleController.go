package controllers

import (
	"TheNexusBattlesII_Card_API_Reborn/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ConsumibleListGet godoc
// @Router /consumible/ [get]
// @Tags Consumible
// @Summary trae una lista de documentos tipo Consumible
// @Description este método devuelve una lista de documentos tipo Consumible
// @Param size query int true "Tamaño de la paginación"
// @Param page query int true "Página de los documentos"
// @Param coleccion query string true "Coleccion del documento"
// @Param statusFilter query bool true "Si se activa el filtro, solo trae las cartas activas"
// @Produce json
// @Success 200 {object} []models.ConsumibleModel "Lista de documentos tipo Consumible"
// @Failure 400 {object} string "Parámetros de paginación inválidos"
// @Failure 500 {object} string "Error interno en el servidor"
func GetConsumibleList(c echo.Context) error {
	//[1. Validar paginación]
	var pagination models.PaginationModel
	if err := c.Bind(&pagination); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message:": "Bad Request: Invalid query parameters"})
	}
	if err := pagination.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message:": "Bad Request: Invalid query parameters"})
	}
	//[2. Traer los documentos]
	coleccion := c.QueryParam("coleccion")
	var consumibleList []models.ConsumibleModel
	if err := models.ConsumibleModel.GetObjectList(models.ConsumibleModel{}, pagination, coleccion, &consumibleList); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message:": "Internal Server Error: Documents not found"})
	}
	return c.JSON(http.StatusOK, consumibleList)
}

// ConsumibleListByNameGet godoc
// @Router /consumible/name/ [get]
// @Tags Consumible
// @Summary trae una lista de documentos tipo Consumible filtrado por nombre
// @Description este método devuelve una lista de documentos tipo Consumible filtrado por nombre
// @Param size query int true "Tamaño de la paginación"
// @Param page query int true "Página de los documentos"
// @Param keyword query string true "palabra clave para filtrar"
// @Param coleccion query string true "Coleccion del documento"
// @Accept json
// @Produce json
// @Success 200 {object} []models.ConsumibleModel "Lista de documentos tipo Consumible"
// @Failure 400 {object} string "Parámetros de paginación inválidos"
// @Failure 500 {object} string "Error interno en el servidor"
func GetConsumibleListByName(c echo.Context) error {
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
	var consumibleList []models.ConsumibleModel
	coleccion := c.QueryParam("coleccion")
	if err := models.ConsumibleModel.GetObjectListByName(models.ConsumibleModel{}, pagination, keyword, coleccion, &consumibleList); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message:": "Internal Server Error: Documents not found"})
	}
	return c.JSON(http.StatusOK, consumibleList)
}
