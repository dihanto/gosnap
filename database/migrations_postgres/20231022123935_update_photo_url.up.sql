ALTER TABLE if exists photos RENAME COLUMN photo_url TO photo_base64;

ALTER TABLE photos ALTER COLUMN photo_base64 TYPE TEXT;
