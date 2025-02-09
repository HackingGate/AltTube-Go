-- change_to_ent.sql
-- This migration converts your old schema into the new Atlas target schema.
-- It performs:
--   - Column type changes (e.g. converting text IDs to UUID)
--   - Renaming of timestamp columns (created_at/updated_at to create_time/update_time)
--   - Dropping of unused columns/indexes (deleted_at)
--   - Re-adding foreign key constraints with explicit names

------------------------------------------------------------
-- STEP 1: Drop Existing Foreign Key Constraints
------------------------------------------------------------
BEGIN;
-- Drop all FKs that reference columns that will be altered.
ALTER TABLE refresh_tokens DROP CONSTRAINT IF EXISTS fk_refresh_tokens_user;
ALTER TABLE access_tokens DROP CONSTRAINT IF EXISTS fk_access_tokens_user;
ALTER TABLE access_tokens DROP CONSTRAINT IF EXISTS fk_access_tokens_refresh_token;
ALTER TABLE like_videos DROP CONSTRAINT IF EXISTS fk_like_videos_user;
ALTER TABLE like_videos DROP CONSTRAINT IF EXISTS fk_like_videos_video;
COMMIT;

------------------------------------------------------------
-- STEP 2: Migrate the "users" Table
------------------------------------------------------------
BEGIN;
-- Convert id from text to uuid (assumes existing values are valid UUIDs).
ALTER TABLE users
ALTER COLUMN id TYPE uuid USING id::uuid;

-- Rename timestamp columns.
ALTER TABLE users RENAME COLUMN created_at TO create_time;
ALTER TABLE users RENAME COLUMN updated_at TO update_time;

-- Drop the deleted_at column and its index.
DROP INDEX IF EXISTS idx_users_deleted_at;
ALTER TABLE users DROP COLUMN IF EXISTS deleted_at;

-- Enforce NOT NULL on the new timestamp columns.
ALTER TABLE users ALTER COLUMN create_time SET NOT NULL;
ALTER TABLE users ALTER COLUMN update_time SET NOT NULL;
COMMIT;

------------------------------------------------------------
-- STEP 3: Migrate the "refresh_tokens" Table
------------------------------------------------------------
BEGIN;
-- Convert user_id from text to uuid.
ALTER TABLE refresh_tokens
ALTER COLUMN user_id TYPE uuid USING user_id::uuid;

-- Rename timestamp columns.
ALTER TABLE refresh_tokens RENAME COLUMN created_at TO create_time;
ALTER TABLE refresh_tokens RENAME COLUMN updated_at TO update_time;

-- Drop the deleted_at column and its index.
DROP INDEX IF EXISTS idx_refresh_tokens_deleted_at;
ALTER TABLE refresh_tokens DROP COLUMN IF EXISTS deleted_at;

-- Enforce NOT NULL on user_id.
ALTER TABLE refresh_tokens ALTER COLUMN user_id SET NOT NULL;
COMMIT;

------------------------------------------------------------
-- STEP 4: Migrate the "access_tokens" Table
------------------------------------------------------------
BEGIN;
-- Convert user_id from text to uuid.
ALTER TABLE access_tokens
ALTER COLUMN user_id TYPE uuid USING user_id::uuid;

-- Rename timestamp columns.
ALTER TABLE access_tokens RENAME COLUMN created_at TO create_time;
ALTER TABLE access_tokens RENAME COLUMN updated_at TO update_time;

-- Drop the deleted_at column and its index.
DROP INDEX IF EXISTS idx_access_tokens_deleted_at;
ALTER TABLE access_tokens DROP COLUMN IF EXISTS deleted_at;

-- Enforce NOT NULL on user_id.
ALTER TABLE access_tokens ALTER COLUMN user_id SET NOT NULL;

-- Enforce NOT NULL on refresh_token_id.
ALTER TABLE access_tokens ALTER COLUMN refresh_token_id SET NOT NULL;
COMMIT;

------------------------------------------------------------
-- STEP 5: Migrate the "videos" Table
------------------------------------------------------------
BEGIN;
-- Rename timestamp columns.
ALTER TABLE videos RENAME COLUMN created_at TO create_time;
ALTER TABLE videos RENAME COLUMN updated_at TO update_time;

-- Drop the deleted_at column and its index.
DROP INDEX IF EXISTS idx_videos_deleted_at;
ALTER TABLE videos DROP COLUMN IF EXISTS deleted_at;

-- Enforce NOT NULL on the new timestamp columns.
ALTER TABLE videos ALTER COLUMN create_time SET NOT NULL;
ALTER TABLE videos ALTER COLUMN update_time SET NOT NULL;
COMMIT;

------------------------------------------------------------
-- STEP 6: Migrate the "like_videos" Table
------------------------------------------------------------
BEGIN;
-- Convert user_id from text to uuid.
ALTER TABLE like_videos
ALTER COLUMN user_id TYPE uuid USING user_id::uuid;

-- Rename timestamp columns.
ALTER TABLE like_videos RENAME COLUMN created_at TO create_time;
ALTER TABLE like_videos RENAME COLUMN updated_at TO update_time;

-- Drop the deleted_at column and its index.
DROP INDEX IF EXISTS idx_like_videos_deleted_at;
ALTER TABLE like_videos DROP COLUMN IF EXISTS deleted_at;

-- Enforce NOT NULL on user_id and video_id.
ALTER TABLE like_videos ALTER COLUMN user_id SET NOT NULL;
ALTER TABLE like_videos ALTER COLUMN video_id SET NOT NULL;
COMMIT;

------------------------------------------------------------
-- STEP 7: Re-add Foreign Key Constraints
------------------------------------------------------------
BEGIN;
-- (Extra drop commands make this step idempotent in case the migration is re-run.)
ALTER TABLE refresh_tokens DROP CONSTRAINT IF EXISTS fk_refresh_tokens_user;
ALTER TABLE access_tokens DROP CONSTRAINT IF EXISTS fk_access_tokens_user;
ALTER TABLE access_tokens DROP CONSTRAINT IF EXISTS fk_access_tokens_refresh_token;
ALTER TABLE like_videos DROP CONSTRAINT IF EXISTS fk_like_videos_user;
ALTER TABLE like_videos DROP CONSTRAINT IF EXISTS fk_like_videos_video;

-- Recreate foreign key from refresh_tokens.user_id to users.id.
ALTER TABLE refresh_tokens ADD CONSTRAINT fk_refresh_tokens_user
    FOREIGN KEY (user_id) REFERENCES users(id)
        ON UPDATE NO ACTION ON DELETE NO ACTION;

-- Recreate foreign key from access_tokens.user_id to users.id.
ALTER TABLE access_tokens ADD CONSTRAINT fk_access_tokens_user
    FOREIGN KEY (user_id) REFERENCES users(id)
        ON UPDATE NO ACTION ON DELETE NO ACTION;

-- Recreate foreign key from access_tokens.refresh_token_id to refresh_tokens.id.
ALTER TABLE access_tokens ADD CONSTRAINT fk_access_tokens_refresh_token
    FOREIGN KEY (refresh_token_id) REFERENCES refresh_tokens(id)
        ON UPDATE NO ACTION ON DELETE NO ACTION;

-- Recreate foreign key from like_videos.user_id to users.id.
ALTER TABLE like_videos ADD CONSTRAINT fk_like_videos_user
    FOREIGN KEY (user_id) REFERENCES users(id)
        ON UPDATE NO ACTION ON DELETE NO ACTION;

-- Recreate foreign key from like_videos.video_id to videos.id.
ALTER TABLE like_videos ADD CONSTRAINT fk_like_videos_video
    FOREIGN KEY (video_id) REFERENCES videos(id)
        ON UPDATE NO ACTION ON DELETE NO ACTION;
COMMIT;
