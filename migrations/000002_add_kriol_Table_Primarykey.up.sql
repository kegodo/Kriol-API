-- Filename: migrations/000002_add_kriol_Table.PrimaryKey.down.sql
ALTER TABLE kriol
ADD id SERIAL PRIMARY KEY;