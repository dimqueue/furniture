-- +migrate Up

create table managers
(
    uuid          text not null primary key,
    login         text not null,
    manager_email text not null
);

create table statuses
(
    id     bigserial not null primary key,
    text_code text      not null
);

create table materials
(
    id    bigserial not null primary key,
    title text      not null
);

create table products
(
    id          bigserial not null primary key,
    title       text      not null,
    price       bigint    not null,
    status_id   bigint default 2 null references statuses (id),
    material_id bigint    not null references materials (id)
);

create table orders
(
    uuid        text   not null primary key,
    description text   not null,
    first_name  text   not null,
    last_name   text   not null,
    delivery    text not null,
    order_email       text   not null,
    product_id  bigint not null references products (id),
    manager_id  text default null references managers (uuid)
);
-- +migrate Down

drop table orders;
drop table products;
drop table materials;
drop table statuses;
drop table managers;