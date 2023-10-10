package models

type EfectoModel struct {
	Estadistica   string `json:"estadistica"   validate:"required,oneof=Poder Vida Defensa Ataque Daño"`
	ValorAfectado int    `json:"valorAfectado" validate:"required"`
	TurnosValidos int    `json:"turnosValidos" validate:"required,gte=-1"`
	Id_Estrategia int    `json:"id_Estrategia" validate:"required,gt=0"`
}
