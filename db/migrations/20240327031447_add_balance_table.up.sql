CREATE TABLE IF NOT EXISTS balances (
  id VARCHAR(255) PRIMARY KEY,
  owner VARCHAR(255) NOT NULL,
  currency VARCHAR(10),
  balance INTEGER,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (owner) REFERENCES users(id)
);
