CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table users (
    email text PRIMARY KEY UNIQUE,
    username text NOT NULL
);

create table groups (
  id uuid PRIMARY KEY DEFAULT public.uuid_generate_v4(),
  name text NOT NULL,
  host_email text references users(email) NOT NULL,
  player02_email text references users(email),
  player03_email text references users(email),
  player04_email text references users(email)
);

create table saves (
    hash text primary key unique NOT NULL,
    group_id uuid references groups(id) NOT NULL,
    save_owner text references users(email) NOT NULL,
    folder_name text,
    full_local_path text,
    cdn_path text
);

ALTER TABLE groups
ADD COLUMN last_save text references saves(hash);


-- Ten random Users
WITH numbers AS (
  SELECT generate_series(1,10) AS id
)

INSERT INTO users (username, email)
SELECT
  'user' || id AS username,
  'user' || id || '@example.com' AS email
FROM numbers;


INSERT INTO users(email,username)
VALUES ('bshafer93@gmail.com','Bert');

