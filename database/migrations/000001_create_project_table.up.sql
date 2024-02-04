CREATE TABLE IF NOT EXISTS project(
    id serial primary key,
    name varchar not null,
    -- in mCPU
    cpu bigint not null,
    -- in MiB
    memory bigint not null,
    -- in MiB
    storage bigint not null,
    cluster varchar not null
);

CREATE TYPE OperationalModel AS ENUM ('UNKNOWN', 'DE3', 'DE4', 'DE4_VSNFD', 'DE6');
CREATE TYPE Platform AS ENUM ('UNKNOWN', 'AWS', 'GCP');
CREATE TYPE WorkerType AS ENUM ('UNKNOWN', 'SMALL', 'SMALL_CPU_HEAVY', 'SMALL_MEMORY_HEAVY', 'MEDIUM', 'MEDIUM_CPU_HEAVY', 'MEDIUM_MEMORY_HEAVY', 'LARGE', 'LARGE_CPU_HEAVY', 'LARGE_MEMORY_HEAVY');

CREATE TABLE IF NOT EXISTS clusters (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    psp_element VARCHAR(255) NOT NULL,
    operational_model OperationalModel NOT NULL,
    platform Platform NOT NULL,
    worker_type WorkerType NOT NULL,
    is_gpu_worker BOOLEAN NOT NULL,
    max_worker_count INT NOT NULL,
    worker_count_number BIGINT,
    longhorn_storage BIGINT NOT NULL,
    nfs_storage BIGINT,
    reason TEXT NOT NULL,
    notes_text TEXT
);
