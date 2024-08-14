-- Drop the unique constraint on accounts for owner and currency
ALTER TABLE "accounts" DROP CONSTRAINT "owner_currency_key";

-- Drop the foreign key from accounts to users
ALTER TABLE "accounts" DROP CONSTRAINT IF EXISTS accounts_owner_fkey;

-- Drop the users table
DROP TABLE IF EXISTS "users";