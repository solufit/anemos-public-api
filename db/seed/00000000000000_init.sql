-- migrate:up
INSERT INTO users (id, name, email, created_At, updated_at)
VALUES ('01F8MECHZX3TBDSZ7XY8E9ZHDH', '須藤史郎', 'sudo.shiron@solufit.net', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)

-- migrate:down