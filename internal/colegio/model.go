package colegio



type Colegio struct {
	ID       int    `gorm:"primaryKey;autoIncrement" json:"id_colegio"`
	Nombre   string `gorm:"size:100;not null" json:"nombre"`
	Direccion string `gorm:"size:200" json:"direccion"`
	Telefono  string `gorm:"size:20" json:"telefono"`
}

