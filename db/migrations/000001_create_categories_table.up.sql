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

INSERT INTO categories (id, name, url, path) VALUES (0, '-', '-', '00');
INSERT INTO categories (name, url, path, parent_id) VALUES ('Any Department', 'https://www.amazon.com/Best-Sellers/zgbs', '01', 0);
