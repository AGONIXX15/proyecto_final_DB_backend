package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/AGONIXX15/db_proyecto_final/internal/admin"
	"github.com/AGONIXX15/db_proyecto_final/internal/cliente"
	"github.com/AGONIXX15/db_proyecto_final/internal/colegio"
	"github.com/AGONIXX15/db_proyecto_final/internal/database"
	"github.com/AGONIXX15/db_proyecto_final/internal/detalle_pedido"
	"github.com/AGONIXX15/db_proyecto_final/internal/factura"
	"github.com/AGONIXX15/db_proyecto_final/internal/materia_prima"
	"github.com/AGONIXX15/db_proyecto_final/internal/pedido"
	"github.com/AGONIXX15/db_proyecto_final/internal/producto_terminado"
	"github.com/AGONIXX15/db_proyecto_final/internal/proveedor"
	"github.com/AGONIXX15/db_proyecto_final/internal/reportes"
	"github.com/AGONIXX15/db_proyecto_final/internal/uniforme"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	// err := godotenv.Load("../.env")
	// if err != nil {
	// 	log.Fatalln("error loading .env")
	// }

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatalln("error no DATABASE_URL in .env")
	}
	fmt.Println(dsn)

	database.InitDB(dsn)

	router := gin.Default()

	 router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))

	adminRepo := admin.NewAdminRepository(database.DB)
	adminService := admin.NewAdminService(adminRepo)
	adminHandler := admin.NewAdminHandler(adminService)
	admin.RegisterAdminRoutes(router, adminHandler)
	clienteRepo := cliente.NewClienteRepository(database.DB)
	clienteService := cliente.NewClienteService(clienteRepo)
	clienteHandler := cliente.NewClienteHandler(clienteService)
	cliente.RegisterClienteRoutes(router, clienteHandler)

	colegioRepo := colegio.NewColegioRepository(database.DB)
	colegioService := colegio.NewColegioService(colegioRepo)
	colegioHandler := colegio.NewColegioHandler(colegioService)
	colegio.RegisterColegioRoutes(router, colegioHandler)

	uniformeRepo := uniforme.NewUniformeRepository(database.DB)
	uniformeService := uniforme.NewUniformeService(uniformeRepo)
	uniformeHandler := uniforme.NewUniformeHandler(uniformeService)
	uniforme.RegisterUniformeRoutes(router, uniformeHandler)

	proveedorRepo := proveedor.NewProveedorRepository(database.DB)
	proveedorService := proveedor.NewProveedorService(proveedorRepo)
	proveedorHandler := proveedor.NewProveedorHandler(proveedorService)
	proveedor.RegisterProveedorRoutes(router, proveedorHandler)

	materiaRepo := materia_prima.NewMateriaPrimaRepository(database.DB)
	materiaService := materia_prima.NewMateriaPrimaService(materiaRepo)
	materiaHandler := materia_prima.NewMateriaPrimaHandler(materiaService)
	materia_prima.RegisterMateriaPrimaRoutes(router, materiaHandler)

	detalleRepo := detalle_pedido.NewDetallePedidoRepository(database.DB)
	detalleService := detalle_pedido.NewDetallePedidoService(detalleRepo)
	detalleHandler := detalle_pedido.NewDetallePedidoHandler(detalleService)
	detalle_pedido.RegisterDetallePedidoRoutes(router, detalleHandler)

	pedidoRepo := pedido.NewPedidoRepository(database.DB)
	pedidoService := pedido.NewPedidoService(pedidoRepo)
	pedidoHandler := pedido.NewPedidoHandler(pedidoService)
	pedido.RegisterPedidoRoutes(router, pedidoHandler)

	facturaRepo := factura.NewFacturaRepository(database.DB)
	facturaService := factura.NewFacturaService(facturaRepo)
	facturaHandler := factura.NewFacturaHandler(facturaService)
	factura.RegisterFacturaRoutes(router, facturaHandler)

	productoRepo := producto_terminado.NewProductoTerminadoRepository(database.DB)
	productoService := producto_terminado.NewProductoTerminadoService(productoRepo)
	productoHandler := producto_terminado.NewProductoTerminadoHandler(productoService)
	producto_terminado.RegisterProductoTerminadoRoutes(router,productoHandler)


	reportesRepo := reportes.NewReporteRepository(database.DB)
	reportesService := reportes.NewReporteService(reportesRepo)
	reportesHandler := reportes.NewReporteHandler(reportesService)
	reportes.RegisterRoutesReportes(router, reportesHandler)

	router.Use(cors.Default())
	router.Run(":8080")
}
