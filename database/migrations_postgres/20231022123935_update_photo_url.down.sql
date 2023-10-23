-- First, change the data type back to its original type (if it was not originally TEXT)
-- Assuming the original data type was VARCHAR, but you should replace it with the correct original data type
ALTER TABLE photos ALTER COLUMN photo_base64 TYPE VARCHAR;

-- Then, rename the column back to its original name
ALTER TABLE photos RENAME COLUMN photo_base64 TO photo_url;
