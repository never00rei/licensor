-- Write your migrate up statements here

CREATE TABLE application.customer(
    customer_id bigserial not null primary key,
    org_uuid uuid DEFAULT uuid_generate_v4(),
    name text not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    active boolean not null default true,
    deleted boolean not null default false,
    deleted_at timestamp,
    UNIQUE(org_uuid, name),
    UNIQUE(name)
);

---- create above / drop below ----

DROP TABLE application.customer;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
