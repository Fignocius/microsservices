CREATE TABLE trackings  (
    id uuid PRIMARY KEY NOT NULL,
    code varchar(128) NOT NULL,
    status varchar(128) NOT NULL,
    description varchar(256),
    created_at timestamptz NOT NULL DEFAULT NOW(),
    updated_at timestamptz
);