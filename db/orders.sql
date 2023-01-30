CREATE table orders (
    id int primary key ,
    customer_id int,
    amount decimal(10,2),
    address varchar(500),
    cart text,
    preferred_date date
)