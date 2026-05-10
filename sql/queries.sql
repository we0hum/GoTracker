INSERT INTO orders (id, customer, address)
VALUES (1,
        'Ivan',
        'ул. Ростовская 5'
        );

SELECT * FROM orders;

SELECT * FROM orders
WHERE customer = "Ivan" AND is_delivered = true;

UPDATE orders
SET is_delivered = true
WHERE id = 1;

DELETE FROM orders
WHERE id = 1;

SELECT is_delivered, COUNT(is_delivered) FROM orders
GROUP BY is_delivered