-- This is a sample migration.

create table management.tenant(
  org_id bigserial not null primary key,
  org_name text not null,
  org_uuid uuid DEFAULT uuid_generate_v4(),
  created_at timestamp not null default now(),
  updated_at timestamp not null default now(),
  UNIQUE(org_name)
);

---- create above / drop below ----

drop table management.tenant;
