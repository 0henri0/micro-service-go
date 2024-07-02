CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    slug VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    origin_url VARCHAR(255) UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX products_slug_index ON products (slug);

-- Path: product-service/internal/database/migrations/V0.0.2__create_product_images_table.sql

CREATE TABLE product_images (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX product_images_product_id_index ON product_images (product_id);

-- Path: product-service/internal/database/migrations/V0.0.3__create_product_categories_table.sql

CREATE TABLE product_categories (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    category_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX product_categories_product_id_index ON product_categories (product_id);
CREATE INDEX product_categories_category_id_index ON product_categories (category_id);

-- Path: product-service/internal/database/migrations/V0.0.4__create_categories_table.sql

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX categories_slug_index ON categories (slug);
