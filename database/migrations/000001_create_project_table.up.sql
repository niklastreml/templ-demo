CREATE TABLE IF NOT EXISTS project(
    id serial primary key,
    name varchar not null,
    -- in mCPU
    cpu integer not null,
    -- in bytes
    memory integer not null,
    -- in bytes
    storage integer not null
)