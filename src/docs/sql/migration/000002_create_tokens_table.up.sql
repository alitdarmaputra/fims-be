CREATE TABLE tokens (    
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,    
    token VARCHAR(128) NOT NULL UNIQUE,    
    token_expiry DATETIME NOT NULL,    
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
    deleted_at DATETIME,
    CONSTRAINT FK_users_token FOREIGN KEY (user_id)
    REFERENCES users(id)
)ENGINE=InnoDB; 

