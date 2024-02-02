CREATE TABLE IF NOT EXISTS project(
    id serial primary key,
    name varchar not null,
    -- in mCPU
    cpu bigint not null,
    -- in MiB
    memory bigint not null,
    -- in MiB
    storage bigint not null
)