CREATE TABLE histories (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    history_type VARCHAR(50) NOT NULL,
    status_from_id INT,
    status_to_id INT,
    node_id INT NOT NULL,
    updated_by INT NOT NULL,
    figma_url VARCHAR(200),
    figma_version VARCHAR(50),
    snapshot_path VARCHAR(200),
    description VARCHAR(200),
    CONSTRAINT FK_histories_users FOREIGN KEY (updated_by)
    REFERENCES users(id),
    CONSTRAINT FK_histories_nodes FOREIGN KEY (node_id)
    REFERENCES nodes(id),
    CONSTRAINT FK_histories_status_from FOREIGN KEY (status_from_id)
    REFERENCES status(id),
    CONSTRAINT FK_histories_status_to FOREIGN KEY (status_to_id)
    REFERENCES status(id),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
)ENGINE=InnoDB
