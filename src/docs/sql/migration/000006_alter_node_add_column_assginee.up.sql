ALTER TABLE nodes
ADD assignee_id int;

ALTER TABLE nodes
ADD CONSTRAINT FK_nodes_assignee
FOREIGN KEY (assignee_id) REFERENCES users(id);
