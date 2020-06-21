CREATE TABLE categories (
  id serial PRIMARY KEY NOT NULL,
  name varchar(255) NOT NULL,
  url varchar(255) NOT NULL,
  path ltree NOT NULL,
  parent_id int REFERENCES categories (id) ON UPDATE CASCADE,
  created_at timestamp DEFAULT current_timestamp,
  updated_at timestamp,
  deleted_at timestamp
);

CREATE INDEX categories_path_gist_idx ON categories USING gist(path);
CREATE INDEX categories_path_idx ON categories (path);
CREATE INDEX categories_parent_id_idx ON categories (parent_id);
CREATE UNIQUE INDEX categories_name_path_idx ON categories (name, path);
