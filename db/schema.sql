CREATE TABLE IF NOT EXISTS persons (
  id INT generated always as identity, 
  info JSON NOT NULL,
  PRIMARY KEY(id)
);
