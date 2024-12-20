INSERT INTO users (user_id, balance) VALUES (1, 0.00) ON CONFLICT (user_id) DO NOTHING;
INSERT INTO users (user_id, balance) VALUES (2, 0.00) ON CONFLICT (user_id) DO NOTHING;
INSERT INTO users (user_id, balance) VALUES (3, 0.00) ON CONFLICT (user_id) DO NOTHING;
