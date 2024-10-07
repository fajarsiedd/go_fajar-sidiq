-- INSERT
INSERT INTO operators (name) VALUES 
('Operator A'), ('Operator B'), ('Operator C'), ('Operator D'), ('Operator E');

INSERT INTO product_types (name) VALUES
('Elektronik'), ('Pakaian'), ('Makanan');

INSERT INTO products (product_type_id, operator_id, name, code, price, status) VALUES
(1, 3, 'Samsung A55', 'SM001', 5000000.0, 1),
(1, 3, 'Laptop ASUS', 'LP001', 8000000.0, 1),
(2, 1, 'Baju Oversized', 'KO001', 350000.0, 1),
(2, 1, 'Celana Denim', 'CD001', 150000.0, 1),
(2, 1, 'Sweater Rajut', 'SR001', 125000.0, 1),
(3, 4, 'Seblak Uhuy', 'SBLK01', 15000.0, 1),
(3, 4, 'Martabak Bangkar', 'MTBK01', 25000.0, 1),
(3, 4, 'Odeng', 'OD001', 20000.0, 1);

INSERT INTO product_descriptions (product_id, description) VALUES
(1, 'Smartphone mid-range yang canggih'),
(2, 'Laptop canggih, ringan dan murah'),
(3, 'Kaos kekinian agar kamu lebih skena'),
(4, 'Celana denim bahan stretch, nyaman dipakai'),
(5, 'Sweater rajut yang adem, bagus seperti oppa oppa korea'),
(6, 'Seblak pedas maknyos!'),
(7, 'Martabak asoy geboy gurih wauw'),
(8, 'Odeng rebus dengan kuah gurih mantap');

INSERT INTO payment_methods (name, status) VALUES
('Mandiri Virtual Account', 1),
('Mandiri Transfer Bank', 1),
('Gopay', 1);

INSERT INTO users (status, dob, gender) VALUES
(1, '1999-11-24', 'L'),
(1, '2003-12-23', 'P'),
(0, '1971-03-03', 'P'),
(1, '2005-09-01', 'L'),
(0, '1991-10-14', 'L');

INSERT INTO transactions (user_id, payment_method_id, total_qty, total) VALUES
(1, 1, 0, 0),
(2, 1, 0, 0),
(3, 1, 0, 0);

-- trigger untuk mengubah quantity di tabel transactions
DELIMITER $$
CREATE TRIGGER update_total_qty
AFTER INSERT ON transaction_details
FOR EACH ROW
BEGIN
    UPDATE transactions SET
        total_qty = (SELECT SUM(qty) FROM transaction_details WHERE transaction_id = NEW.transaction_id),
        total = (SELECT SUM(price) FROM transaction_details WHERE transaction_id = NEW.transaction_id)
    WHERE id = NEW.transaction_id;
END;
$$

INSERT INTO transaction_details (transaction_id, product_id, status, qty, price) VALUES
(4, 1, 1, 2, 5000000.0),
(4, 2, 1, 1, 8000000.0),
(4, 3, 1, 1, 350000.0),
(5, 3, 1, 2, 350000.0),
(5, 4, 1, 2, 150000.0),
(5, 5, 1, 1, 125000.0),
(6, 6, 1, 4, 15000.0),
(6, 7, 1, 2, 25000.0),
(6, 8, 1, 3, 20000.0);

-- SELECT
SELECT name FROM users WHERE gender = 'L';

SELECT * FROM products WHERE id = 3;

SELECT * FROM users WHERE created_at >= DATE_SUB(NOW(), INTERVAL 7 DAY) AND name LIKE '%a%';

SELECT COUNT(*) AS jumlah FROM users WHERE gender = 'P';

SELECT * FROM users ORDER BY name ASC;

SELECT * FROM products LIMIT 5;

-- UPDATE
UPDATE products SET name = 'product dummy' WHERE id = 1;

UPDATE transaction_details SET qty = 3 WHERE product_id = 1;

-- DELETE
DELETE FROM products WHERE id = 1;