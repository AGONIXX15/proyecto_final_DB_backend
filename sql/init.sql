CREATE OR REPLACE FUNCTION crear_pedido(
    p_cliente_id INT,
    p_fecha_encargo DATE,
    p_fecha_entrega DATE,
    p_abono NUMERIC,
    p_detalles JSON)
RETURNS INT AS $$
DECLARE
    nuevo_pedido_id INT;
    d JSON;
    tipo TEXT;
    cod INT;
    cant INT;
    stock_actual INT;
BEGIN
    -- 1. Crear pedido
    INSERT INTO pedido (cliente_id, fecha_encargo, fecha_entrega, abono, estado)
    VALUES (p_cliente_id, p_fecha_encargo, p_fecha_entrega, p_abono, 'pendiente')
    RETURNING num_pedido INTO nuevo_pedido_id;

    -- 2. Recorrer los detalles
    FOR d IN SELECT * FROM json_array_elements(p_detalles)
    LOOP
        tipo := d->>'type_item';
        cod  := (d->>'cod_item')::int;
        cant := (d->>'cantidad')::int;

        INSERT INTO detalle_pedido(num_pedido, type_item, cod_item, cantidad, medidas, observaciones)
        VALUES (nuevo_pedido_id, tipo, cod, cant, d->>'medidas', d->>'observaciones');

        -- 3. Si es producto → descontar inventario
        IF tipo = 'producto' THEN
            SELECT cantidad_existencia INTO stock_actual
            FROM producto_terminado WHERE cod_producto = cod FOR UPDATE;

            IF stock_actual < cant THEN
                RAISE EXCEPTION 'Stock insuficiente para producto %', cod;
            END IF;

            UPDATE producto_terminado
            SET cantidad_existencia = cantidad_existencia - cant,
                estado = 'encargado'
            WHERE cod_producto = cod;
        END IF;

    END LOOP;

    RETURN nuevo_pedido_id;
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
    d_old RECORD;
    d_new JSON;
    key_old TEXT;
    tipo_old TEXT;
    cod_old INT;
    cant_old INT;

    tipo_new TEXT;
    cod_new INT;
    cant_new INT;

    diff INT;
    stock_actual INT;
BEGIN
    -- 1. Actualizar datos del pedido
    UPDATE pedido
    SET fecha_entrega = p_fecha_entrega,
        abono = p_abono
    WHERE num_pedido = p_num_pedido;

    -- 2. Procesar detalles antiguos
    FOR d_old IN
        SELECT * FROM detalle_pedido WHERE num_pedido = p_num_pedido
    LOOP
        tipo_old := d_old.type_item;
        cod_old  := d_old.cod_item;
        cant_old := d_old.cantidad;

        -- Buscar en los nuevos detalles
        SELECT elem INTO d_new
        FROM json_array_elements(p_detalles) elem
        WHERE (elem->>'type_item') = tipo_old
          AND (elem->>'cod_item')::int = cod_old;

        IF d_new IS NULL THEN
            -- detalle eliminado → devolver stock
            IF tipo_old = 'producto' THEN
                UPDATE producto_terminado
                SET cantidad_existencia = cantidad_existencia + cant_old,
                    estado = 'disponible'
                WHERE cod_producto = cod_old;
            END IF;

            DELETE FROM detalle_pedido
            WHERE num_pedido = p_num_pedido
              AND cod_item = cod_old;

        ELSE
            tipo_new := d_new->>'type_item';
            cod_new  := (d_new->>'cod_item')::int;
            cant_new := (d_new->>'cantidad')::int;

            diff := cant_new - cant_old;

            IF tipo_new = 'producto' THEN
                IF diff > 0 THEN
                    -- reservar más
                    SELECT cantidad_existencia INTO stock_actual
                    FROM producto_terminado WHERE cod_producto = cod_new FOR UPDATE;

                    IF stock_actual < diff THEN
                        RAISE EXCEPTION 'Stock insuficiente para producto %', cod_new;
                    END IF;

                    UPDATE producto_terminado
                    SET cantidad_existencia = cantidad_existencia - diff
                    WHERE cod_producto = cod_new;

                ELSIF diff < 0 THEN
                    -- devolver stock
                    UPDATE producto_terminado
                    SET cantidad_existencia = cantidad_existencia + ABS(diff)
                    WHERE cod_producto = cod_new;
                END IF;
            END IF;

            -- actualizar detalle
            UPDATE detalle_pedido
            SET cantidad = cant_new,
                medidas = d_new->>'medidas',
                observaciones = d_new->>'observaciones'
            WHERE num_pedido = p_num_pedido
              AND cod_item = cod_old;
        END IF;
    END LOOP;

    -- 3. Insertar detalles nuevos que no existían
    FOR d_new IN SELECT * FROM json_array_elements(p_detalles)
    LOOP
        tipo_new := d_new->>'type_item';
        cod_new  := (d_new->>'cod_item')::int;

        PERFORM 1 FROM detalle_pedido
        WHERE num_pedido = p_num_pedido
          AND type_item = tipo_new
          AND cod_item = cod_new;

        IF NOT FOUND THEN
            cant_new := (d_new->>'cantidad')::int;

            INSERT INTO detalle_pedido(num_pedido, type_item, cod_item, cantidad, medidas, observaciones)
            VALUES (p_num_pedido, tipo_new, cod_new, cant_new, d_new->>'medidas', d_new->>'observaciones');

            IF tipo_new = 'producto' THEN
                SELECT cantidad_existencia INTO stock_actual
                FROM producto_terminado WHERE cod_producto = cod_new FOR UPDATE;

                IF stock_actual < cant_new THEN
                    RAISE EXCEPTION 'Stock insuficiente para producto %', cod_new;
                END IF;

                UPDATE producto_terminado
                SET cantidad_existencia = cantidad_existencia - cant_new
                WHERE cod_producto = cod_new;
            END IF;
        END IF;
    END LOOP;
END;
$$ LANGUAGE plpgsql;


CREATE OR REPLACE FUNCTION cancelar_pedido(p_num_pedido INT)
RETURNS VOID AS $$
DECLARE
    d RECORD;
BEGIN
    UPDATE pedido
    SET estado = 'cancelado'
    WHERE num_pedido = p_num_pedido;

    -- devolver stock
    FOR d IN
        SELECT * FROM detalle_pedido WHERE num_pedido = p_num_pedido
    LOOP
        IF d.type_item = 'producto' THEN
            UPDATE producto_terminado
            SET cantidad_existencia = cantidad_existencia + d.cantidad,
                estado = 'disponible'
            WHERE cod_producto = d.cod_item;
        END IF;
    END LOOP;
END;
$$ LANGUAGE plpgsql;


