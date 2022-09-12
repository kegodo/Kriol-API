-- Filename: migrations/000001_create_kriol_table.up.sql
CREATE TABLE IF NOT EXISTS kriol(
    Kriol_Word text NOT NULL,
    English_Word text NOT NULL
);