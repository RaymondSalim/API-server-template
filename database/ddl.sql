create table if not exists "counter-table" (
    id              bigserial primary key,
    "count"           bigint
    created_at      timestamp with time zone,
    updated_at      timestamp with time zone,
    deleted_at      timestamp with time zone,
);

create index if not exists "idx_counter-table_deleted_at"
    on "counter-table" (deleted_at);

create table if not exists "foo-table" (
    id          bigserial primary key,
    name        text
);
