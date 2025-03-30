INSERT INTO customers (name, email, phone, gender, is_active, created_at, updated_at)
VALUES
    ('John Doe', 'johndoe@example.com', '1234567890', 'Male', true, UNIX_TIMESTAMP() * 1000, UNIX_TIMESTAMP() * 1000),
    ('Jane Smith', 'janesmith@example.com', '0987654321', 'Female', true, UNIX_TIMESTAMP() * 1000, UNIX_TIMESTAMP() * 1000);