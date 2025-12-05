package uniforme



type Uniforme struct {
	ID          int    `gorm:"primaryKey;autoIncrement" json:"id_uniforme"`
	TipoUniforme string `gorm:"size:50;not null" json:"tipo_uniforme"`
	Color        string `gorm:"size:50" json:"color"`
	TipoTela     string `gorm:"size:50" json:"tipo_tela"`
	Bordado      string `gorm:"size:50" json:"bordado"`
	Estampado    string `gorm:"size:50" json:"estampado"`
	Detalles     string `gorm:"size:200" json:"detalles"`
	IDColegio    int    `json:"id_colegio"`
}

