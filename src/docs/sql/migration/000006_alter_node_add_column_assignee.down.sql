ALTER TABLE nodes
DROP CONSTRAINT FK_nodes_assignee;

ALTER TABLE nodes
DROP COLUMN assignee_id;
