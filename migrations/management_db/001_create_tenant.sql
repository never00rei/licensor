-- This is a sample migration.

create table tenant(
  org_id bigserial not null primary key,
  org_name text not null,
  org_uuid uuid DEFAULT uuid_generate_v4(),
  schema_name text not null,
  created_at timestamp not null default now(),
  updated_at timestamp not null default now()
);

---- create above / drop below ----

drop table tenant;
