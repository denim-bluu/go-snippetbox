CREATE TABLE sessions (
    token VARCHAR(43) PRIMARY KEY,
    data bytea NOT NULL,
    expiry TIMESTAMP(0) WITH TIME ZONE NOT NULL
);

CREATE INDEX sessions_expiry_idx ON sessions (expiry);