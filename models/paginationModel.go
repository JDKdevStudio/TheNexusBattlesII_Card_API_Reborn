package models

import "github.com/go-playground/validator/v10"

type PaginationModel struct {
	Size        int    `query:"size"         validate:"required,gt=0"`
	Page        int    `query:"page"         validate:"required,gt=0"`
	Coleccion   string `query:"coleccion"    validate:"required,oneof=All Heroes Armas Armaduras Items Epicas"`
	Keyword     string `query:"keyword"`
	OnlyActives *bool  `query:"onlyActives" validate:"required"`
}

func (p PaginationModel) Validate() error {
	validate := validator.New()
	err := validate.Struct(p)
	if err != nil {
		return err
	}
	return nil
}
