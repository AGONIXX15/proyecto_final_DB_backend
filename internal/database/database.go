package database

import (

	"github.com/AGONIXX15/db_proyecto_final/internal/admin"
	"github.com/AGONIXX15/db_proyecto_final/internal/cliente"
	"github.com/AGONIXX15/db_proyecto_final/internal/colegio"
	"github.com/AGONIXX15/db_proyecto_final/internal/detalle_pedido"
	"github.com/AGONIXX15/db_proyecto_final/internal/factura"
	"github.com/AGONIXX15/db_proyecto_final/internal/materia_prima"
	"github.com/AGONIXX15/db_proyecto_final/internal/pedido"
	"github.com/AGONIXX15/db_proyecto_final/internal/producto_terminado"
	"github.com/AGONIXX15/db_proyecto_final/internal/proveedor"
	"github.com/AGONIXX15/db_proyecto_final/internal/uniforme"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dsn string) {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}


	DB = db
	db.AutoMigrate(&cliente.Cliente{})
	db.AutoMigrate(&admin.Admin{})
	db.AutoMigrate(&colegio.Colegio{})
	db.AutoMigrate(&detalle_pedido.DetallePedido{})
	db.AutoMigrate(&factura.Factura{})
	db.AutoMigrate(&materia_prima.MateriaPrima{})
	db.AutoMigrate(&pedido.Pedido{})
	db.AutoMigrate(&proveedor.Proveedor{})
	db.AutoMigrate(&uniforme.Uniforme{})
	db.AutoMigrate(&producto_terminado.ProductoTerminado{})
}
