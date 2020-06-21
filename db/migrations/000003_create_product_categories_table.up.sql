CREATE TABLE product_categories (
  product_id int REFERENCES products (id) ON UPDATE CASCADE,
  category_id int REFERENCES categories (id) ON UPDATE CASCADE,
  created_at timestamp DEFAULT current_timestamp,
  updated_at timestamp,
  deleted_at timestamp,
  CONSTRAINT product_category_pkey PRIMARY KEY (product_id, category_id)
);
