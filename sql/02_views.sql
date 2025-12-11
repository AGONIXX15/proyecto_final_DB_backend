-- ===========================================
-- VISTAS DE REPORTES DEL SISTEMA
-- ===========================================

-- 1️⃣ Listado de productos encargados pendientes por entregar (ordenados por fecha)
CREATE OR REPLACE VIEW productos_pendientes AS
SELECT p.num_pedido,
       dp.cod_item,
       dp.type_item,
       dp.cantidad,
       p.fecha_entrega
FROM pedidos p
JOIN detalle_pedidos dp ON p.num_pedido = dp.num_pedido
WHERE p.fecha_entrega IS NULL
ORDER BY p.fecha_encargo;

-- 2️⃣ Productos encargados por cada cliente que no han sido entregados
CREATE OR REPLACE VIEW productos_cliente AS
SELECT c.documento,
       c.nombre_completo,
       dp.cod_item,
       dp.type_item,
       dp.cantidad,
       p.num_pedido
FROM clientes c
JOIN pedidos p ON c.documento = p.cliente_id
JOIN detalle_pedidos dp ON p.num_pedido = dp.num_pedido
WHERE p.fecha_entrega IS NULL;


CREATE OR REPLACE VIEW colegios_uniformes AS
SELECT DISTINCT c.id,
       c.nombre
FROM colegios c
JOIN uniformes u ON u.id_colegio = c.id;

-- 5️⃣ Características del uniforme de cada colegio
CREATE OR REPLACE VIEW uniforme_colegio AS
SELECT u.id_colegio,
       u.tipo_uniforme,
       u.color,
       u.tipo_tela,
       u.bordado,
       u.estampado,
       u.detalles
FROM uniformes u;

-- 6️⃣ Total de uniformes encargados por colegio
CREATE OR REPLACE VIEW ventas_por_colegio AS
SELECT u.id_colegio,
       c.nombre AS colegio_nombre,
       SUM(dp.cantidad) AS total_vendido
FROM detalle_pedidos dp
JOIN uniformes u ON dp.cod_item = u.id AND dp.type_item = 'uniforme'
JOIN colegios c ON u.id_colegio = c.id
GROUP BY u.id_colegio, c.nombre;

-- 7️⃣ Total de ventas
CREATE OR REPLACE VIEW total_ventas AS
SELECT SUM(f.total) AS total_ventas
FROM facturas f;

