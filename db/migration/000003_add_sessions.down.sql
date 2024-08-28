-- Drop the unique constraint on sessions for owner and currency
ALTER TABLE "sessions" DROP CONSTRAINT "owner_currency_key";

-- Drop the foreign key from sessions to users
ALTER TABLE "sessions" DROP CONSTRAINT IF EXISTS sessions_owner_fkey;

-- Drop the users table
DROP TABLE IF EXISTS "users";