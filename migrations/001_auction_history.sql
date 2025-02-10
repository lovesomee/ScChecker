-- +goose Up

create table if not exists auction_history (
    id serial primary key,
    item_id varchar(5),
    amount int,
    price int,
    time timestamp,
    additional jsonb,
    created_time timestamp default now(),
    unique (price, time)
);

create index auction_history_item_id on auction_history (item_id);

-- +goose Down
drop table if exists auction_history;
drop index auction_history_item_id;
