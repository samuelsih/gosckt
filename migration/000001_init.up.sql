CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table users(
  id uuid default uuid_generate_v4() primary key,
  email varchar(100) not null,
  name varchar(50) not null,
  password bytea not null,
  UNIQUE(email)
);