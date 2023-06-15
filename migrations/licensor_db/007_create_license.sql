-- Write your migrate up statements here

CREATE TABLE application.license (
    license_id uuid DEFAULT uuid_generate_v4() not null primary key,
    issuer uuid not null,
    org_uuid uuid not null,
    verifier uuid not null,
    valid_from timestamp not null,
    valid_until timestamp not null,
    validity_period interval not null,
    active boolean not null default true,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);

---- create above / drop below ----

DROP TABLE application.license;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
