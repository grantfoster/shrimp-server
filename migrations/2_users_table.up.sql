create table users (
  id uuid primary key default uuid_generate_v4(),
  email text,
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
);
