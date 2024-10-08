-- 1
SELECT * FROM transactions WHERE user_id = 1
UNION ALL
SELECT * FROM transactions WHERE user_id = 2;

-- 2
SELECT SUM(total) AS 'jumlah_harga' FROM transactions WHERE user_id = 1;

-- 3
SELECT COUNT(*) 'total_transaksi' FROM transaction_details td
    INNER JOIN products p
        ON p.id = td.product_id
    WHERE p.product_type_id = 2;

-- 4
SELECT p.*, pt.name AS 'type_name' FROM products p
    INNER JOIN product_types pt
        ON pt.id = p.product_type_id;

-- 5
SELECT t.*, u.name AS 'user_name', p.name AS 'product_name' FROM transaction_details td
    INNER JOIN products p
        ON p.id = td.product_id
    INNER JOIN transactions t
        ON t.id = td.transaction_id
    INNER JOIN  users u 
        ON u.id = t.user_id;

-- 6
DELIMITER $$
CREATE TRIGGER on_delete_trx
BEFORE DELETE ON transactions
FOR EACH ROW
BEGIN
    DELETE FROM transaction_details WHERE transaction_id = OLD.id;
END;
$$

-- 7
DELIMITER $$
CREATE TRIGGER on_delete_trx_detail
AFTER DELETE ON transaction_details
FOR EACH ROW
BEGIN
    UPDATE transactions
    SET total_qty = (SELECT SUM(qty) FROM transaction_details WHERE transaction_id = OLD.transaction_id),
        total = (SELECT SUM(price) FROM transaction_details WHERE transaction_id = OLD.transaction_id);
END;
$$

-- 8
SELECT * FROM products 
WHERE id NOT IN (
    SELECT product_id
    FROM transaction_details
);