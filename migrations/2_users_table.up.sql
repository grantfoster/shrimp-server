create table users (
  id uuid primary key unique not null default uuid_generate_v4(),
  username text unique, 
  email text unique,
  password text,
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
);
