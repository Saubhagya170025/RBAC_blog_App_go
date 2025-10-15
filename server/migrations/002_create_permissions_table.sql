CREATE TABLE IF NOT EXISTS permissions (
    permission_id SERIAL PRIMARY KEY,
    role_id INTEGER NOT NULL UNIQUE,
    create_blog BOOLEAN DEFAULT FALSE,
    create_user BOOLEAN DEFAULT FALSE,
    create_category BOOLEAN DEFAULT FALSE,
    create_role BOOLEAN DEFAULT FALSE,
    update_blog BOOLEAN DEFAULT FALSE,
    update_user BOOLEAN DEFAULT FALSE,
    update_category BOOLEAN DEFAULT FALSE,
    update_role BOOLEAN DEFAULT FALSE,
    delete_blog BOOLEAN DEFAULT FALSE,
    delete_user BOOLEAN DEFAULT FALSE,
    delete_category BOOLEAN DEFAULT FALSE,
    delete_role BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_permissions_role FOREIGN KEY (role_id) 
        REFERENCES roles(role_id) 
        ON DELETE CASCADE
);