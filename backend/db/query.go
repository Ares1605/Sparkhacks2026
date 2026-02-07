package db

var sqlInitTables = `
  CREATE TABLE IF NOT EXISTS Orders (
    Id         INT NOT NULL PRIMARY KEY,
    ProviderId INT
    Name       TEXT
    Price      REAL
    OrderDate  TEXT
  );

  CREATE TABLE IF NOT EXISTS Providers (
    Id       INT NOT NULL PRIMARY KEY,
    Name     TEXT
    LastSync TEXT
  );
`

var sqlInsertOrder = `
	INSERT INTO Orders
	VALUES (?, ?, ?, ?, ?)
`

var sqlDeleteByProvider = `
  DELETE FROM Orders
  WHERE ProviderId = (
    SELECT Id FROM Providers WHERE Id = ?
  )
`

var sqlGetProviderId = `SELECT Id FROM Providers WHERE Id = ?`
