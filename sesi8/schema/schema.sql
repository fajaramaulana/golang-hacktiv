CREATE TABLE orders (
    order_id serial primary key,
    customer_name varchar(50),
    ordered_at timestamp
)

CREATE TABLE items (
    item_id serial primary key,
    item_code varchar(10) not null,
    description varchar(100),
    quantity int,
    order_id int
)