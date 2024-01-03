BEGIN;

-- Add promo_price column to promos table
ALTER TABLE promos
ADD COLUMN price DECIMAL(10, 2);

COMMIT;