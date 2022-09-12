-- Filename: migrations/000003_add_created_at_to_kriol_table.down.sql
ALTER TABLE kriol
ADD created_at timestamp(0) WITH time zone NOT NULL DEFAULT NOW();