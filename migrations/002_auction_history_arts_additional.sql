-- +goose Up

create table if not exists auction_history_arts_additional (
    auction_history_id serial references auction_history (id),
    quality int,
    potential int,
    qlt_percent float
    );

create index auction_history_arts_additional_quality on auction_history_arts_additional (quality);
create index auction_history_arts_additional_potential on auction_history_arts_additional (potential);
create index auction_history_arts_additional_qlt_percent on auction_history_arts_additional (qlt_percent);

-- +goose Down
drop table if exists auction_history_arts_additional;
drop index auction_history_arts_additional_quality;
drop index auction_history_arts_additional_potential;
drop index auction_history_arts_additional_qlt_percent;

