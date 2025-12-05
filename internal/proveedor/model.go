package proveedor


type Proveedor struct {
	NIT           int    `gorm:"primaryKey" json:"nit"`
	Nombre        string `gorm:"size:100;not null" json:"nombre"`
	Direccion     string `gorm:"size:200" json:"direccion"`
	Telefono      string `gorm:"size:20" json:"telefono"`
	NombreContacto string `gorm:"size:100" json:"nombre_contacto"`
}

