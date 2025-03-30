CREATE TABLE customers (
                           id INT AUTO_INCREMENT PRIMARY KEY,
                           name VARCHAR(255) NOT NULL,
                           email VARCHAR(255) UNIQUE NOT NULL,
                           phone VARCHAR(20) NOT NULL,
                           gender VARCHAR(10) NOT NULL,
                           is_active BOOLEAN DEFAULT TRUE,
                           created_at BIGINT NOT NULL,
                           updated_at BIGINT NOT NULL,
                           deleted_at BIGINT DEFAULT 0
);
