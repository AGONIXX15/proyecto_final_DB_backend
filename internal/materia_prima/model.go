package materia_prima


type MateriaPrima struct {
	CodMateria    int    `gorm:"primaryKey" json:"cod_materia"`
	TipoMateria   string `gorm:"size:50;not null" json:"tipo_materia"`
	Descripcion   string `gorm:"size:200" json:"descripcion"`
	CantidadExistencia int `json:"cantidad_existencia"`
	UnidadMedida  string `gorm:"size:30" json:"unidad_medida"`
	NITProveedor  int    `json:"nit_proveedor"`
}
