package cliente


type Cliente struct {
	Documento       uint    `gorm:"primaryKey" json:"documento" binding:"required"`       // clave primaria
	NombreCompleto  string `gorm:"size:150" json:"nombre" binding:"required"`         // VARCHAR(150)
	Telefono        string `gorm:"size:20" json:"telefono" binding:"required"`          // VARCHAR(20)
}
