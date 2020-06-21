CREATE TABLE products (
  id bigserial PRIMARY KEY NOT NULL,
  name text NOT NULL,
  rank smallserial NOT NULL,
  image_path varchar(255),
  page smallserial NOT NULL,
  category_id bigserial NOT NULL REFERENCES categories (id) ON UPDATE CASCADE,
  created_at timestamp DEFAULT current_timestamp,
  updated_at timestamp,
  deleted_at timestamp
);

CREATE UNIQUE INDEX products_name_rank_category_id_idx ON products (name, rank, category_id);
