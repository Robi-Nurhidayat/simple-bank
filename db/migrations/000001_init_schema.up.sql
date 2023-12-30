-- Tabel accounts
CREATE TABLE accounts (
                          id SERIAL PRIMARY KEY,
                          owner VARCHAR NOT NULL,
                          balance BIGINT NOT NULL,
                          currency VARCHAR NOT NULL,
                          created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Tabel entries
CREATE TABLE entries (
                         id SERIAL PRIMARY KEY,
                         account_id BIGINT NOT NULL, -- Menambahkan NOT NULL di sini
                         amount BIGINT NOT NULL,
                         created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                         FOREIGN KEY (account_id) REFERENCES accounts(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- Tabel transfers
CREATE TABLE transfers (
                           id SERIAL PRIMARY KEY,
                           from_account_id BIGINT NOT NULL, -- Menambahkan NOT NULL di sini
                           to_account_id BIGINT NOT NULL, -- Menambahkan NOT NULL di sini
                           amount BIGINT NOT NULL,
                           created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                           FOREIGN KEY (from_account_id) REFERENCES accounts(id) ON DELETE CASCADE ON UPDATE CASCADE,
                           FOREIGN KEY (to_account_id) REFERENCES accounts(id) ON DELETE CASCADE ON UPDATE CASCADE
);
