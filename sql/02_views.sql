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
JOIN pedidos p ON c.documento = p.doc_cliente
JOIN detalle_pedidos dp ON p.num_pedido = dp.num_pedido
WHERE p.fecha_entrega >= CURRENT_DATE OR p.fecha_entrega IS NULL;


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
CREATE OR REPLACE VIEW v_reportes_total_ventas AS
SELECT SUM(f.total) AS total_ventas
FROM facturas f;

CREATE OR REPLACE VIEW v_reportes_productos_pendientes AS
SELECT
    p.num_pedido,
    p.fecha_encargo,
    p.fecha_entrega,
    p.estado,
    c.nombre_completo AS cliente_nombre,
    c.documento AS cliente_documento,
    dp.type_item AS tipo_producto,
    dp.cod_item AS cod_producto,
    pt.nombre AS nombre_producto,
    pt.categoria AS categoria,
    dp.cantidad,
    dp.medidas,
    dp.observaciones,
    dp.precio_unitario
FROM pedidos p
JOIN clientes c ON c.documento = p.doc_cliente
JOIN detalle_pedidos dp ON dp.num_pedido = p.num_pedido
LEFT JOIN producto_terminados pt 
       ON dp.cod_item = pt.cod_producto 
      AND dp.type_item = 'producto'
WHERE p.estado = 'Pendiente';



CREATE OR REPLACE VIEW v_reportes_productos_cliente AS
SELECT
    p.num_pedido,
    p.fecha_encargo,
    p.fecha_entrega,
    p.estado,
    c.documento AS cliente_documento,
    dp.type_item AS tipo_producto,
    dp.cod_item AS cod_producto,
    pt.nombre AS nombre_producto,
    dp.cantidad,
    dp.medidas,
    dp.precio_unitario,
    dp.subtotal
FROM pedidos p
JOIN clientes c ON p.doc_cliente = c.documento
JOIN detalle_pedidos dp ON dp.num_pedido = p.num_pedido
LEFT JOIN producto_terminados pt 
       ON dp.cod_item = pt.cod_producto 
      AND dp.type_item = 'producto';

CREATE OR REPLACE VIEW v_reportes_ventas_colegio AS
WITH pedidos_uniformes AS (
    SELECT
        u.id_colegio,
        c.nombre AS nombre_colegio,
        p.num_pedido,
        p.fecha_encargo,
        p.fecha_entrega,
        p.estado,
        SUM(dp.subtotal) AS total_pedido
    FROM detalle_pedidos dp
    JOIN uniformes u ON u.id = dp.cod_item
    JOIN pedidos p ON p.num_pedido = dp.num_pedido
    JOIN colegios c ON c.id = u.id_colegio
    WHERE dp.type_item = 'uniforme'
    GROUP BY u.id_colegio, c.nombre, p.num_pedido, p.fecha_encargo, p.fecha_entrega, p.estado
)
SELECT
    pu.id_colegio,
    pu.nombre_colegio,
    pu.num_pedido,
    pu.fecha_encargo,
    pu.fecha_entrega,
    pu.estado,
    pu.total_pedido,
    SUM(pu.total_pedido) OVER (PARTITION BY pu.id_colegio) AS total_por_colegio
FROM pedidos_uniformes pu
ORDER BY pu.id_colegio, pu.num_pedido;



CREATE OR REPLACE VIEW v_reportes_colegios_uniformes AS
SELECT
    u.id_colegio,
    c.nombre AS nombre_colegio,
    c.direccion,
    c.telefono,
    dp.cod_item AS id_uniforme,
    pt.nombre AS tipo_uniforme,
    pt.categoria,
    pt.descripcion AS detalles
FROM detalle_pedidos dp
JOIN pedidos p ON p.num_pedido = dp.num_pedido
JOIN clientes cl ON cl.documento = p.doc_cliente
JOIN uniformes u ON u.id = dp.cod_item
JOIN producto_terminados pt ON pt.cod_producto = dp.cod_item
JOIN colegios c ON c.id = u.id_colegio
WHERE dp.type_item = 'uniforme';
