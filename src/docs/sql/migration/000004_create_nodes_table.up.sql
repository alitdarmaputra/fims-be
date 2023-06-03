CREATE TABLE nodes (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(40) NOT NULL,
    figma_key VARCHAR(50) NOT NULL,
    node_id VARCHAR(20) NOT NULL,
    description TEXT,
    user_id INT NOT NULL,
    status_id INT NOT NULL,
    CONSTRAINT FK_users_nodes FOREIGN KEY (user_id)
    REFERENCES users(id),
    CONSTRAINT FK_status_nodes FOREIGN KEY (status_id)
    REFERENCES status(id),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
    deleted_at DATETIME
)ENGINE=InnoDB;
