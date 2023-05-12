-- +migrate Up
INSERT INTO materials(id,title)
VALUES (1, 'gold'),(2, 'silver'),(3, 'diamond'),(4,'platinum');

-- +migrate Down