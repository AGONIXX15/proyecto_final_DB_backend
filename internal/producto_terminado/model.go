package producto_terminado

type ProductoTerminado struct {
	CodProducto        uint    `gorm:"primaryKey;column:cod_producto" json:"cod_producto"`
	Nombre string `gorm:"column:nombre" json:"nombre"`
	Categoria string `gorm:"column:categoria" json:"categoria"` //accesorios, camisas, pantalones y otros
	Descripcion        string  `gorm:"column:descripcion" json:"descripcion"`
	Talla              string  `gorm:"column:talla" json:"talla"`
	Sexo               string  `gorm:"column:sexo;check: sexo IN ('M','F', 'U')" json:"sexo"`
	Precio             float64 `gorm:"column:precio" json:"precio"`
	CantidadExistencia int     `gorm:"column:cantidad_existencia" json:"cantidad_existencia"`
}



