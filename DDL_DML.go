package database

const ManagersDDL = `
CREATE TABLE IF NOT EXISTS managers
(
    id       INTEGER PRIMARY KEY AUTOINCREMENT,
    name     TEXT    NOT NULL,
    surname  TEXT    NOT NULL,
    login    TEXT    NOT NULL UNIQUE,
    password TEXT    NOT NULL
);`

const ClientsDDL = `
CREATE TABLE IF NOT EXISTS clients
(
    id       INTEGER PRIMARY KEY AUTOINCREMENT,
    name     TEXT    NOT NULL,
    surname  TEXT    NOT NULL,
    login    TEXT    NOT NULL UNIQUE,
    password TEXT    NOT NULL
);`

//TODO add managersCardsDDL
const ClientsCardsDDL = `
CREATE TABLE IF NOT EXISTS clients_cards
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    pan        INTEGER NOT NULL UNIQUE,
    pin        INTEGER NOT NULL,
    balance    INTEGER NOT NULL,
    holderName TEXT    NOT NULL,
    cvv        INTEGER NOT NULL,
    validity   INTEGER NOT NULL,
    client_id  INTEGER NOT NULL REFERENCES clients
);`

const AtmsDDL = `
CREATE TABLE IF NOT EXISTS atms
(
    id       INTEGER PRIMARY KEY AUTOINCREMENT,
    city     TEXT NOT NULL,
    district TEXT NOT NULL,
    street   TEXT NOT NULL
);`

const ServicesDDL = `
CREATE TABLE IF NOT EXISTS services
(
    id      INTEGER PRIMARY KEY AUTOINCREMENT,
    service TEXT    NOT NULL,
    balance INTEGER
);`

const ManagersDML = `
INSERT INTO managers
VALUES (1, 'Admin', 'Administrator', 'adminM', 'adminM')
ON CONFLICT DO NOTHING;`

const ClientsDML = `
INSERT INTO clients
VALUES (1, 'Admin', 'Administrator', 'adminC', 'adminC')
ON CONFLICT DO NOTHING;`

const ClientsCardsDML = `
INSERT INTO clients_cards
VALUES (1, 2021600000000000, 1994, 1000000, 'ADMIN CLIENT', 333, 0222, 1)
ON CONFLICT DO NOTHING;`

const AtmsDML = `
INSERT INTO atms
VALUES (1, 'Dushanbe', 'Somoni', 'Foteh51')
ON CONFLICT DO NOTHING;`

const ServicesDML = `
INSERT INTO services
VALUES (1, 'internet', 1500)
ON CONFLICT DO NOTHING;`
