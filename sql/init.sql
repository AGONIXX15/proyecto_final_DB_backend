CREATE OR REPLACE FUNCTION crear_pedido(
    p_doc_cliente BIGINT,
    p_fecha_encargo DATE,
    p_fecha_entrega DATE,
    p_abono NUMERIC,
    p_detalles JSON
)
RETURNS INT AS $$
DECLARE
    v_num_pedido INT;
    d JSON;
    tipo_item TEXT;
    cod_item INT;
    cant INT;
    stock_actual INT;
    precio_unidad NUMERIC;
BEGIN
    -- 1. Insertar pedido y devolver num_pedido
    INSERT INTO pedidos(doc_cliente, fecha_encargo, fecha_entrega, abono, estado)
    VALUES (p_doc_cliente, p_fecha_encargo, p_fecha_entrega, p_abono, 'Pendiente')
    RETURNING num_pedido INTO v_num_pedido;

    -- 2. Recorrer los detalles del JSON
    FOR d IN SELECT * FROM json_array_elements(p_detalles) AS detalle_json
    LOOP
        tipo_item := d->>'tipo_producto';
        cod_item  := (d->>'cod_producto')::INT;
        cant      := (d->>'cantidad')::INT;
        precio_unidad := (d->>'precio_unidad')::NUMERIC;

        -- Insertar detalle
        INSERT INTO detalle_pedidos(
            num_pedido,
            type_item,
            cod_item,
            precio_unitario,
            cantidad,
            medidas,
            observaciones,
            subtotal
        )
        VALUES (
            v_num_pedido,
            tipo_item,
            cod_item,
            precio_unidad,
            cant,
            d->>'medidas',         
            d->>'observaciones',
            precio_unidad * cant 
        );

        -- Restar stock si es producto
        IF tipo_item = 'producto' THEN
            SELECT cantidad_existencia INTO stock_actual
            FROM producto_terminados
            WHERE cod_producto = cod_item
            FOR UPDATE;

            IF stock_actual < cant THEN
                RAISE EXCEPTION 'Stock insuficiente para producto %', cod_item;
            END IF;

            UPDATE producto_terminados
            SET cantidad_existencia = cantidad_existencia - cant
            WHERE cod_producto = cod_item;
        END IF;
    END LOOP;

    RETURN v_num_pedido;
END;
$$ LANGUAGE plpgsql;


CREATE OR REPLACE FUNCTION actualizar_pedido(
    p_num_pedido INT,
    p_fecha_entrega DATE,
    p_abono NUMERIC,
    p_detalles JSON
)
RETURNS VOID AS $$
DECLARE
    d RECORD;
    tipo_item TEXT;
    cod_item INT;
    cant INT;
    stock_actual INT;
    precio_unidad NUMERIC;
    cant_anterior INT;
BEGIN
    -- 1️⃣ Actualizar datos generales del pedido
    UPDATE pedidos
    SET fecha_entrega = p_fecha_entrega,
        abono = p_abono
    WHERE num_pedido = p_num_pedido;

    -- 2️⃣ Recuperar stock de los productos existentes y devolver stock
    FOR d IN
        SELECT dp.cod_item AS cod_item, dp.type_item AS type_item, dp.cantidad AS cantidad
        FROM detalle_pedidos dp
        WHERE dp.num_pedido = p_num_pedido
    LOOP
        cant_anterior := d.cantidad;

        IF d.type_item = 'producto' THEN
            -- Devolver stock antiguo
            UPDATE producto_terminados
            SET cantidad_existencia = cantidad_existencia + cant_anterior
            WHERE cod_producto = d.cod_item;
        END IF;
    END LOOP;

    -- 3️⃣ Borrar detalles antiguos
    DELETE FROM detalle_pedidos
    WHERE num_pedido = p_num_pedido;

    -- 4️⃣ Insertar nuevos detalles y restar stock
    FOR d IN
        SELECT json_array_elements(p_detalles) AS detalle
    LOOP
        tipo_item := d.detalle->>'tipo_producto';
        cod_item  := (d.detalle->>'cod_producto')::INT;
        cant      := (d.detalle->>'cantidad')::INT;
        precio_unidad := (d.detalle->>'precio_unidad')::NUMERIC;

        -- Insertar detalle
        INSERT INTO detalle_pedidos(
            num_pedido,
            type_item,
            cod_item,
            precio_unitario,
            cantidad,
            medidas,
            observaciones,
            subtotal
        )
        VALUES (
            p_num_pedido,
            tipo_item,
            cod_item,
            precio_unidad,
            cant,
            d.detalle->>'medidas',
            d.detalle->>'observaciones',
            precio_unidad * cant
        );

        IF tipo_item = 'producto' THEN
            -- Verificar stock
            SELECT cantidad_existencia INTO stock_actual
            FROM producto_terminados
            WHERE cod_producto = cod_item
            FOR UPDATE;

            IF stock_actual < cant THEN
                RAISE EXCEPTION 'Stock insuficiente para producto %', cod_item;
            END IF;

            -- Restar stock
            UPDATE producto_terminados
            SET cantidad_existencia = cantidad_existencia - cant
            WHERE cod_producto = cod_item;
        END IF;
    END LOOP;
END;
$$ LANGUAGE plpgsql;



CREATE OR REPLACE FUNCTION cancelar_pedido(p_num_pedido INT)
RETURNS void AS $$
BEGIN
    -- Verificar que el pedido existe y no esté cancelado
    IF NOT EXISTS (
        SELECT 1 
        FROM pedidos 
        WHERE num_pedido = p_num_pedido 
          AND estado <> 'Cancelado'
    ) THEN
        RAISE NOTICE 'Pedido no existe o ya está cancelado';
        RETURN;
    END IF;

    -- Devolver el stock de los productos
    UPDATE producto_terminados pt
    SET cantidad_existencia = cantidad_existencia + dp.cantidad
    FROM detalle_pedidos dp
    WHERE dp.num_pedido = p_num_pedido
      AND pt.cod_producto = dp.cod_item;

    -- Borrar los detalles del pedido
    DELETE FROM detalle_pedidos
    WHERE num_pedido = p_num_pedido;

    -- Marcar el pedido como cancelado
    UPDATE pedidos
    SET estado = 'Cancelado'
    WHERE num_pedido = p_num_pedido;
END;
$$ LANGUAGE plpgsql;


CREATE OR REPLACE FUNCTION entregar_pedido(p_num_pedido INT)
RETURNS void AS $$
DECLARE
    total NUMERIC;
BEGIN
    -- 1️⃣ Verificar que el pedido existe y no esté cancelado
    IF NOT EXISTS (
        SELECT 1 
        FROM pedidos 
        WHERE num_pedido = p_num_pedido 
          AND estado <> 'Cancelado' or estado <> 'Entregado'
    ) THEN
        RAISE EXCEPTION 'Pedido no existe o está cancelado o ya ha sido entregado';
    END IF;

    -- 2️⃣ Actualizar el estado del pedido a Entregado
    UPDATE pedidos
    SET estado = 'Entregado'
    WHERE num_pedido = p_num_pedido;

    -- 3️⃣ Calcular el total del pedido
    SELECT SUM(subtotal) INTO total
    FROM detalle_pedidos
    WHERE num_pedido = p_num_pedido;

    -- 4️⃣ Crear la factura
    INSERT INTO facturas(fecha, total, num_pedido)
    VALUES (NOW(), total, p_num_pedido);
END;
$$ LANGUAGE plpgsql;

