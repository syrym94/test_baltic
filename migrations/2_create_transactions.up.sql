CREATE TABLE transactions (
                              transaction_id TEXT PRIMARY KEY,
                              user_id BIGINT NOT NULL REFERENCES users(user_id),
                              amount NUMERIC(20, 2),
                              state TEXT NOT NULL
);
