
CREATE TABLE
IF NOT EXISTS users
(
		"key" TEXT PRIMARY KEY,
		"value" TEXT
	);
DELETE FROM users WHERE rowid NOT IN (SELECT max(rowid)
FROM users
GROUP BY "key");
CREATE UNIQUE INDEX
IF NOT EXISTS idx_users_key ON users
("key");