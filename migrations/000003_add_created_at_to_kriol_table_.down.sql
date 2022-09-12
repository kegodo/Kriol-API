-- Filename: migrations/000003_add_created_at_to_kriol_table.down.sql
ALTER TABLE kriol
DROP created_at TYPE;