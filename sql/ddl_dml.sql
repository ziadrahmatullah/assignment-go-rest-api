CREATE DATABASE wallet_db;

CREATE TABLE IF NOT EXISTS users(
    id BIGSERIAL,
    name VARCHAR NOT NULL,
    birthdate DATE NOT NULL,
    email VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS reset_password_tokens(
    id BIGSERIAL,
    token VARCHAR NOT NULL,
    expire TIMESTAMP NOT NULL,
    is_used BOOLEAN NOT NULL DEFAULT FALSE,
    user_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE SEQUENCE wallet_number_seq;

CREATE TABLE IF NOT EXISTS wallets(
    id BIGSERIAL,
    wallet_number VARCHAR NOT NULL DEFAULT (concat('700',lpad(nextval('wallet_number_seq')::VARCHAR,10,'0'))),
    balance DECIMAL NOT NULL DEFAULT 0 CHECK (balance >= 0),
    user_id BIGINT NOT NULL,
    chance INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TYPE source_of_funds AS ENUM ('Bank Transfer', 'Credit Card', 'Cash', 'Reward');
CREATE TYPE transaction_types AS ENUM ('Transfer', 'Top up', 'Game Reward');

CREATE TABLE IF NOT EXISTS transactions(
    id BIGSERIAL,
    wallet_id BIGINT NOT NULL,
    transaction_type transaction_types NOT NULL,
    source_of_fund source_of_funds,
    recipient_id BIGINT,
    amount DECIMAL NOT NULL,
    description VARCHAR,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (wallet_id) REFERENCES wallets(id),
    FOREIGN KEY (recipient_id) REFERENCES wallets(id)
);

CREATE TABLE IF NOT EXISTS boxs(
    id BIGSERIAL,
    reward_amount DECIMAL NOT NULL CHECK (reward_amount > 0),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS games(
    id BIGSERIAL,
    box_id BIGINT NOT NULL,
    wallet_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    PRIMARY KEY (box_id, wallet_id),
    FOREIGN KEY (box_id) REFERENCES boxs(id),
    FOREIGN KEY (wallet_id) REFERENCES wallets(id)
);


INSERT INTO users (name, email, birthdate, password, created_at, updated_at)
VALUES 
    ('Alice', 'alice@gmail.com', '2001-04-14', 'alice123', NOW(), NOW()),
    ('Bob', 'bob@gmail.com', '2000-08-15', 'bob123', NOW(), NOW()),
    ('Celine', 'celine@gmail.com', '1999-07-22', 'celine123', NOW(), NOW()),
    ('Denis', 'denis@gmail.com', '2000-03-10', 'denis123', NOW(), NOW()),
    ('Ekil', 'ekil@gmail.com', '2001-01-19', 'ekil123', NOW(), NOW());

INSERT INTO wallets (wallet_number, balance, user_id, chance, created_at, updated_at)
VALUES 
    ('7000000000001',10000, 1, 0,'2021-11-01', '2021-11-01'),
    ('7000000000002',20000, 2, 0, '2021-11-01', '2021-11-01'),
    ('7000000000003',30000, 3, 0, '2021-11-01', '2021-11-01'),
    ('7000000000004',40000, 4, 0, '2021-11-01', '2021-11-01'),
    ('7000000000005',50000, 5, 0, '2021-11-01', '2021-11-01');

INSERT INTO transactions (wallet_id,transaction_type,source_of_fund, amount, description, created_at, updated_at)
VALUES 
    (2,'Top up','Bank Transfer', 100000, 'ML',  '2022-11-01', '2022-11-01'),
    (1,'Top up', 'Credit Card',200000, 'PUBG', '2022-10-15', '2022-10-15'),
    (2,'Top up', 'Bank Transfer',700000, 'ML',  '2022-09-20', '2022-09-20'),
    (3,'Top up','Credit Card', 400000, 'ML',  '2022-08-05', '2022-08-05'),
    (3,'Top up', 'Credit Card',1000000, 'Genshin',  '2022-06-01', '2022-06-01'),
    (1,'Top up','Bank Transfer', 600000, 'Genshin',  '2022-06-15', '2022-06-15'),
    (4,'Top up','Cash', 200000, 'ML',  '2022-05-10', '2022-05-10'),
    (4,'Top up','Bank Transfer', 300000, 'PUBG', '2023-04-20', '2023-04-20'),
    (2,'Top up','Cash', 400000, 'Genshin',  '2023-03-05', '2023-03-05'),
    (5,'Top up','Bank Transfer', 600000, 'Genshin', '2023-02-01', '2023-02-01');

INSERT INTO transactions (wallet_id,transaction_type, amount, description,recipient_id, created_at, updated_at)
VALUES 
    (1, 'Transfer', 100000, 'Ngasih', 2, '2022-11-03', '2022-11-03'),
    (2,'Transfer', 100000, 'Ngasih', 3, '2022-10-25', '2022-10-25'),
    (2,'Transfer', 100000, 'Ngasih', 1, '2022-09-10', '2022-09-10'),
    (3, 'Transfer',100000, 'Ngasih', 5, '2022-08-15', '2022-08-15'),
    (1, 'Transfer',100000, 'Ngasih', 3, '2022-07-05', '2022-07-05'),
    (5, 'Transfer',100000, 'Ngasih', 1, '2022-06-20', '2022-06-20'),
    (5, 'Transfer',100000, 'Ngasih', 4, '2022-05-10', '2022-05-10'),
    (5, 'Transfer',100000, 'Ngasih', 3, '2023-04-15', '2023-04-15'),
    (4, 'Transfer',100000, 'Ngasih', 2, '2023-03-20', '2023-03-20'),
    (3,'Transfer', 100000, 'Ngasih', 2, '2023-02-05', '2023-02-05');

INSERT INTO boxs (reward_amount, created_at, updated_at)
VALUES
    (300000, NOW(), NOW()),
    (400000, NOW(), NOW()),
    (500000, NOW(), NOW()),
    (600000, NOW(), NOW()),
    (700000, NOW(), NOW()),
    (800000, NOW(), NOW()),
    (900000, NOW(), NOW()),
    (1000000, NOW(), NOW()),
    (1100000, NOW(), NOW());