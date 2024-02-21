CREATE EXTENSION citext;

CREATE TABLE IF NOT EXISTS users (
                                     user_id uuid PRIMARY KEY default gen_random_uuid(),
                                     created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
                                     name text NOT NULL,
                                     email citext UNIQUE NOT NULL,
                                     password_hash bytea NOT NULL,
                                     activated bool NOT NULL,
                                     version integer NOT NULL DEFAULT 1
);