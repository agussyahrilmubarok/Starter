CREATE TABLE users (
  id          UNIQUEIDENTIFIER  DEFAULT NEWID() PRIMARY KEY,
  name        NVARCHAR(100)     NOT NULL,
  email       NVARCHAR(150)     NOT NULL,
  password    NVARCHAR(255)     NOT NULL,
  created_at  DATETIME2         DEFAULT GETDATE(),
  updated_at  DATETIME2         DEFAULT GETDATE()
);

CREATE UNIQUE INDEX uq_users_email ON users (email);

CREATE TRIGGER trg_users_updated_at
ON users
AFTER UPDATE
AS
BEGIN
  SET NOCOUNT ON;
  UPDATE users
  SET updated_at = GETDATE()
  FROM users u
  INNER JOIN inserted i ON u.id = i.id;
END;