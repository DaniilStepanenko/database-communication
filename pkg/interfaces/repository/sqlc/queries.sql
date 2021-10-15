-- name: CreateCustomer :one
INSERT INTO dvds.customer (store_id,
                           first_name,
                           last_name,
                           email,
                           address_id,
                           activebool,
                           create_date,
                           last_update,
                           active)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING customer.customer_id;

-- name: ReadCustomer :one
SELECT customer.store_id,
       customer.first_name,
       customer.last_name,
       customer.email,
       customer.address_id,
       customer.activebool,
       customer.create_date,
       customer.last_update,
       customer.active
FROM dvds.customer
WHERE customer.customer_id = $1;

-- name: UpdateCustomer :exec
UPDATE dvds.customer
SET (store_id, first_name, last_name, email, address_id, activebool, create_date, last_update,
     active) = ($1, $2, $3, $4, $5, $6, $7, $8, $9)
WHERE customer.customer_id = $10;

-- name: DeleteCustomer :exec
DELETE FROM dvds.customer
WHERE customer.customer_id = $1;

-- name: FindActiveCustomers :many
SELECT customer.customer_id,
       customer.store_id,
       customer.first_name,
       customer.last_name,
       customer.email,
       customer.address_id,
       customer.activebool,
       customer.create_date,
       customer.last_update,
       customer.active
FROM dvds.customer
         INNER JOIN (
    SELECT payment.customer_id,
           MAX(payment.payment_date) AS "max_payment_date"
    FROM dvds.payment
    GROUP BY payment.customer_id
) AS payment ON customer.customer_id = payment.customer_id
WHERE max_payment_date >= $1::TIMESTAMP WITHOUT TIME ZONE;
