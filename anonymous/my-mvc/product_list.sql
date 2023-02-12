create table product_list
(
    Id          int unsigned auto_increment
        primary key,
    Name        varchar(20)  not null,
    Price       int          not null,
    Quantity    varchar(50)  null,
    Description varchar(200) null,
    Action      varchar(50)  null
);

