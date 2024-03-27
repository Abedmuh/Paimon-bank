CREATE TABLE IF NOT EXISTS balances (
  id VARCHAR(255) PRIMARY KEY,
  owner VARCHAR(255) NOT NULL,
  name VARCHAR(255),
  acc_number VARCHAR(255),
  currency VARCHAR(10),
  balance INTEGER,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (owner) REFERENCES users(id),
  CONSTRAINT unique_owner_currency UNIQUE (owner, currency)
);
