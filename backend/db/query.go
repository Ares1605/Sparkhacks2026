package db

var sqlInitTables = `
  CREATE TABLE IF NOT EXISTS Orders (
    Id         INT NOT NULL PRIMARY KEY,
    ProviderId INT,
    Name       TEXT,
    Price      REAL,
    OrderDate  TEXT
  );

  CREATE TABLE IF NOT EXISTS Providers (
    Id       INT NOT NULL PRIMARY KEY,
    Name     TEXT,
    LastSync TEXT,
    Username TEXT,
    Password TEXT
  );

  CREATE TABLE IF NOT EXISTS ChatSessions (
    Id       INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    Uuid     TEXT NOT NULL
  );

  CREATE TABLE IF NOT EXISTS ChatMessages (
    Id              INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    SessionUuid     TEXT NOT NULL,
    Message         TEXT NOT NULL,
    Role            TEXT NOT NULL,
    MessageDate     TEXT
  );
`

var sqlInsertOrder = `
	INSERT INTO Orders
	VALUES (?, ?, ?, ?, ?)
`

var sqlDeleteByProvider = `
  DELETE FROM Orders
  WHERE ProviderId = ?
`

var sqlGetProviderId = `SELECT Id FROM Providers WHERE Id = ?`

var sqlGetProviderStatusByID = `
  SELECT Username, LastSync
  FROM Providers
  WHERE Id = ?
  LIMIT 1
`

var sqlUpsertProviderSync = `
  INSERT INTO Providers
  (Id, Name, LastSync, Username, Password)
  VALUES (?, ?, ?, ?, COALESCE((SELECT Password FROM Providers WHERE Id = ?), ''))
  ON CONFLICT(Id) DO UPDATE SET
    Name = excluded.Name,
    LastSync = excluded.LastSync,
    Username = excluded.Username
`

var sqlUpsertProviderCredentials = `
  INSERT INTO Providers
  (Id, Name, LastSync, Username, Password)
  VALUES (?, ?, NULL, ?, ?)
  ON CONFLICT(Id) DO UPDATE SET
    Name = excluded.Name,
    Username = excluded.Username,
    Password = excluded.Password
`

var sqlGetAllProviders = `SELECT * FROM Providers`

var sqlGetAllOrders = `SELECT * FROM Orders`

var sqlInsertChatSession = `
	INSERT INTO ChatSessions
	(Uuid) VALUES (?)
`

var sqlInsertChatMessage = `
  INSERT INTO ChatMessages
  (SessionUuid, Message, Role, MessageDate) VALUES (?, ?, ?, ?)
`
