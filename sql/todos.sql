DROP TABLE IF EXISTS todos;

CREATE TABLE IF NOT EXISTS todos (
  id SERIAL NOT NULL,
  content TEXT NOT NULL,
  user_id INT NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  is_deleted BOOLEAN DEFAULT FALSE,
  deleted_at TIMESTAMP,
  PRIMARY KEY     (id)
);