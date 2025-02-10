-- +goose Up

create table if not exists update_list (
        item_id varchar(5) primary key
    );

-- +goose Down
drop table if exists update_list;