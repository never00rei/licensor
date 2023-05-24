-- Write your migrate up statements here

create table management_user(
    user_id bigserial not null primary key,
    username text not null,
    api_key text not null,
    email text not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    active boolean not null default true,
    deleted boolean not null default false,
    deleted_at timestamp,
    is_admin boolean not null default false
)

---- create above / drop below ----

drop table management_user;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
