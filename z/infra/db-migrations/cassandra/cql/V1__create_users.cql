CREATE KEYSPACE IF NOT EXISTS sandbox_db
  WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}
  AND durable_writes = true;

USE sandbox_db;

CREATE TABLE IF NOT EXISTS users (
  id          UUID      PRIMARY KEY,
  name        TEXT      NOT NULL,
  email       TEXT      NOT NULL,
  password    TEXT      NOT NULL,
  created_at  TIMESTAMP NOT NULL,
  updated_at  TIMESTAMP NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_users_email ON users (email);