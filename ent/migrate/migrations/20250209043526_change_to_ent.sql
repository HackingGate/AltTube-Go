--
-- MIGRATION SCRIPT: from old schema (GORM) to new schema (Ent)
--

BEGIN;

--------------------------------------------------------------------------
-- STEP 0: Drop existing foreign keys and constraints that we will recreate
--------------------------------------------------------------------------

-- access_tokens
ALTER TABLE IF EXISTS public.access_tokens DROP CONSTRAINT IF EXISTS fk_access_tokens_refresh_token;
ALTER TABLE IF EXISTS public.access_tokens DROP CONSTRAINT IF EXISTS fk_access_tokens_user;
-- The old PK name is the same in both schemas (access_tokens_pkey), so we leave it
-- unless we decide to drop/recreate it. We'll handle the IDENTITY update separately.

-- like_videos
ALTER TABLE IF EXISTS public.like_videos DROP CONSTRAINT IF EXISTS fk_like_videos_user;
ALTER TABLE IF EXISTS public.like_videos DROP CONSTRAINT IF EXISTS fk_like_videos_video;
-- old PK is like_videos_pkey, same name in new schema, so we keep it.

-- refresh_tokens
ALTER TABLE IF EXISTS public.refresh_tokens DROP CONSTRAINT IF EXISTS fk_refresh_tokens_user;
-- old PK is refresh_tokens_pkey, same name in new schema, so we keep it.

-- users
-- Drop the old unique constraint on email:
ALTER TABLE IF EXISTS public.users DROP CONSTRAINT IF EXISTS uni_users_email;
-- Drop the old PK named uni_users_id; new schema calls it users_pkey:
ALTER TABLE IF EXISTS public.users DROP CONSTRAINT IF EXISTS uni_users_id;

-- videos
-- Drop the old PK named uni_videos_id; new schema calls it videos_pkey:
ALTER TABLE IF EXISTS public.videos DROP CONSTRAINT IF EXISTS uni_videos_id;

COMMIT;


BEGIN;

--------------------------------------------------------------------------------
-- STEP 1: Delete row if deleted_at is not null
--------------------------------------------------------------------------------

-- access_tokens
DELETE FROM public.access_tokens WHERE deleted_at IS NOT NULL;

-- like_videos
DELETE FROM public.like_videos WHERE deleted_at IS NOT NULL;

-- refresh_tokens
DELETE FROM public.refresh_tokens WHERE deleted_at IS NOT NULL;

-- users
DELETE FROM public.users WHERE deleted_at IS NOT NULL;

-- videos
DELETE FROM public.videos WHERE deleted_at IS NOT NULL;

COMMIT;


BEGIN;

--------------------------------------------------------------------------------
-- STEP 2: Drop old/obsolete indexes that do not appear in the new schema
--------------------------------------------------------------------------------

-- Access tokens indexes
DROP INDEX IF EXISTS public.idx_access_tokens_deleted_at;
DROP INDEX IF EXISTS public.idx_access_tokens_refresh_token_id;
DROP INDEX IF EXISTS public.idx_access_tokens_user_id;

-- Like videos indexes
DROP INDEX IF EXISTS public.idx_like_videos_deleted_at;
DROP INDEX IF EXISTS public.idx_like_videos_video_id;

-- Refresh tokens indexes
DROP INDEX IF EXISTS public.idx_refresh_tokens_deleted_at;
DROP INDEX IF EXISTS public.idx_refresh_tokens_user_id;
DROP INDEX IF EXISTS public.idx_token;  -- on refresh_tokens(token)

-- Users
DROP INDEX IF EXISTS public.idx_users_deleted_at;

-- Videos
DROP INDEX IF EXISTS public.idx_videos_deleted_at;

COMMIT;


BEGIN;

--------------------------------------------------------------------------------
-- STEP 3: Remove deleted_at columns, rename created_at/updated_at -> create_time/update_time, set NOT NULL
--------------------------------------------------------------------------------

-- 3.1: access_tokens
ALTER TABLE public.access_tokens
    DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE public.access_tokens
    RENAME COLUMN created_at TO create_time;
ALTER TABLE public.access_tokens
    RENAME COLUMN updated_at TO update_time;
ALTER TABLE public.access_tokens
    ALTER COLUMN create_time SET NOT NULL;
ALTER TABLE public.access_tokens
    ALTER COLUMN update_time SET NOT NULL;

-- 3.2: like_videos
ALTER TABLE public.like_videos
    DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE public.like_videos
    RENAME COLUMN created_at TO create_time;
ALTER TABLE public.like_videos
    RENAME COLUMN updated_at TO update_time;
ALTER TABLE public.like_videos
    ALTER COLUMN create_time SET NOT NULL;
ALTER TABLE public.like_videos
    ALTER COLUMN update_time SET NOT NULL;

-- 3.3: refresh_tokens
ALTER TABLE public.refresh_tokens
    DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE public.refresh_tokens
    RENAME COLUMN created_at TO create_time;
ALTER TABLE public.refresh_tokens
    RENAME COLUMN updated_at TO update_time;
ALTER TABLE public.refresh_tokens
    ALTER COLUMN create_time SET NOT NULL;
ALTER TABLE public.refresh_tokens
    ALTER COLUMN update_time SET NOT NULL;

-- 3.4: users
ALTER TABLE public.users
    DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE public.users
    RENAME COLUMN created_at TO create_time;
ALTER TABLE public.users
    RENAME COLUMN updated_at TO update_time;
ALTER TABLE public.users
    ALTER COLUMN create_time SET NOT NULL;
ALTER TABLE public.users
    ALTER COLUMN update_time SET NOT NULL;

-- 3.5: videos
ALTER TABLE public.videos
    DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE public.videos
    RENAME COLUMN created_at TO create_time;
ALTER TABLE public.videos
    RENAME COLUMN updated_at TO update_time;
ALTER TABLE public.videos
    ALTER COLUMN create_time SET NOT NULL;
ALTER TABLE public.videos
    ALTER COLUMN update_time SET NOT NULL;

COMMIT;


BEGIN;

--------------------------------------------------------------------------------
-- STEP 4: Convert column types, enforce NOT NULL
--------------------------------------------------------------------------------

-- access_tokens:
--   user_id: text -> uuid
--   token: text -> character varying
--   refresh_token_id: stays bigint, but enforce NOT NULL
--   expiry: stays timestamp with time zone
ALTER TABLE public.access_tokens
    ALTER COLUMN user_id TYPE uuid USING user_id::uuid;
ALTER TABLE public.access_tokens
    ALTER COLUMN token TYPE character varying;
ALTER TABLE public.access_tokens
    ALTER COLUMN refresh_token_id SET NOT NULL;
ALTER TABLE public.access_tokens
    ALTER COLUMN user_id SET NOT NULL;

-- like_videos:
--   user_id: text -> uuid
--   video_id: text -> character varying
ALTER TABLE public.like_videos
    ALTER COLUMN user_id TYPE uuid USING user_id::uuid;
ALTER TABLE public.like_videos
    ALTER COLUMN video_id TYPE character varying;
ALTER TABLE public.like_videos
    ALTER COLUMN user_id SET NOT NULL;
ALTER TABLE public.like_videos
    ALTER COLUMN video_id SET NOT NULL;

-- refresh_tokens:
--   user_id: text -> uuid
--   token, user_agent, ip_address: text -> character varying
ALTER TABLE public.refresh_tokens
    ALTER COLUMN user_id TYPE uuid USING user_id::uuid;
ALTER TABLE public.refresh_tokens
    ALTER COLUMN token TYPE character varying;
ALTER TABLE public.refresh_tokens
    ALTER COLUMN user_agent TYPE character varying;
ALTER TABLE public.refresh_tokens
    ALTER COLUMN ip_address TYPE character varying;
ALTER TABLE public.refresh_tokens
    ALTER COLUMN user_id SET NOT NULL;

-- users:
--   id: text -> uuid
--   email: text -> character varying NOT NULL
--   password: text -> character varying NOT NULL
ALTER TABLE public.users
    ALTER COLUMN id TYPE uuid USING id::uuid;
ALTER TABLE public.users
    ALTER COLUMN email TYPE character varying;
ALTER TABLE public.users
    ALTER COLUMN email SET NOT NULL;
ALTER TABLE public.users
    ALTER COLUMN password TYPE character varying;
ALTER TABLE public.users
    ALTER COLUMN password SET NOT NULL;

-- videos:
--   id: text -> character varying
--   title, description, uploader, uploader_url, thumbnail_url: text -> character varying
ALTER TABLE public.videos
    ALTER COLUMN id TYPE character varying;
ALTER TABLE public.videos
    ALTER COLUMN title TYPE character varying;
ALTER TABLE public.videos
    ALTER COLUMN description TYPE character varying;
ALTER TABLE public.videos
    ALTER COLUMN uploader TYPE character varying;
ALTER TABLE public.videos
    ALTER COLUMN uploader_url TYPE character varying;
ALTER TABLE public.videos
    ALTER COLUMN thumbnail_url TYPE character varying;

COMMIT;


BEGIN;

--------------------------------------------------------------------------------
-- STEP 5: Convert numeric ID columns to Postgres identity
--         (replacing old sequences/SET DEFAULT nextval)
--------------------------------------------------------------------------------

-- 5.1: access_tokens.id
-- Drop default nextval if it exists, then add IDENTITY
ALTER TABLE public.access_tokens
    ALTER COLUMN id DROP DEFAULT;

ALTER TABLE public.access_tokens
    ALTER COLUMN id
        ADD GENERATED BY DEFAULT AS IDENTITY (
        SEQUENCE NAME public.access_tokens_id_seq
        START WITH 1
        INCREMENT BY 1
        NO MINVALUE
        NO MAXVALUE
        CACHE 1
        );

-- 5.2: like_videos.id
ALTER TABLE public.like_videos
    ALTER COLUMN id DROP DEFAULT;

ALTER TABLE public.like_videos
    ALTER COLUMN id
        ADD GENERATED BY DEFAULT AS IDENTITY (
        SEQUENCE NAME public.like_videos_id_seq
        START WITH 1
        INCREMENT BY 1
        NO MINVALUE
        NO MAXVALUE
        CACHE 1
        );

-- 5.3: refresh_tokens.id
ALTER TABLE public.refresh_tokens
    ALTER COLUMN id DROP DEFAULT;

ALTER TABLE public.refresh_tokens
    ALTER COLUMN id
        ADD GENERATED BY DEFAULT AS IDENTITY (
        SEQUENCE NAME public.refresh_tokens_id_seq
        START WITH 1
        INCREMENT BY 1
        NO MINVALUE
        NO MAXVALUE
        CACHE 1
        );

-- users.id and videos.id remain text/uuid/character varying, so no identity needed.

COMMIT;


BEGIN;

--------------------------------------------------------------------------------
-- STEP 6: Recreate new PK constraints or unique indexes as needed
--------------------------------------------------------------------------------

-- 6.1: users -> new PK name = users_pkey, plus unique index on email
ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);

CREATE UNIQUE INDEX users_email_key ON public.users USING btree (email);

-- 6.2: videos -> new PK name = videos_pkey
ALTER TABLE ONLY public.videos
    ADD CONSTRAINT videos_pkey PRIMARY KEY (id);

-- For access_tokens, like_videos, refresh_tokens the old PK names (xxx_pkey) match the new schema,
-- so we don't need to re-add them unless we explicitly dropped them in Step 0.

COMMIT;


BEGIN;

--------------------------------------------------------------------------------
-- STEP 7: Re-add foreign key constraints with new constraint names
--------------------------------------------------------------------------------

-- access_tokens -> refresh_tokens
ALTER TABLE ONLY public.access_tokens
    ADD CONSTRAINT access_tokens_refresh_tokens_access_tokens
        FOREIGN KEY (refresh_token_id) REFERENCES public.refresh_tokens(id);

-- access_tokens -> users
ALTER TABLE ONLY public.access_tokens
    ADD CONSTRAINT access_tokens_users_access_tokens
        FOREIGN KEY (user_id) REFERENCES public.users(id);

-- like_videos -> users
ALTER TABLE ONLY public.like_videos
    ADD CONSTRAINT like_videos_users_like_videos
        FOREIGN KEY (user_id) REFERENCES public.users(id);

-- like_videos -> videos
ALTER TABLE ONLY public.like_videos
    ADD CONSTRAINT like_videos_videos_like_videos
        FOREIGN KEY (video_id) REFERENCES public.videos(id);

-- refresh_tokens -> users
ALTER TABLE ONLY public.refresh_tokens
    ADD CONSTRAINT refresh_tokens_users_refresh_tokens
        FOREIGN KEY (user_id) REFERENCES public.users(id);

COMMIT;


BEGIN;

-------------------------------------------------------------------------
-- STEP 8: Recreate indexes that are in the new schema but not the old
-------------------------------------------------------------------------

-- Drop the old sequence so we can reuse its name in the identity column.
DROP SEQUENCE IF EXISTS public.access_tokens_id_seq CASCADE;

-- Now convert the column to an identity column, giving it the same sequence name.
ALTER TABLE public.access_tokens
    ALTER COLUMN id DROP DEFAULT;  -- drop any "nextval" default if still present
ALTER TABLE public.access_tokens
    ALTER COLUMN id
        ADD GENERATED BY DEFAULT AS IDENTITY (
        SEQUENCE NAME public.access_tokens_id_seq
        START WITH 1
        INCREMENT BY 1
        NO MINVALUE
        NO MAXVALUE
        CACHE 1
        );

COMMIT;

BEGIN;

-- like_videos
DROP SEQUENCE IF EXISTS public.like_videos_id_seq CASCADE;

ALTER TABLE public.like_videos
    ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public.like_videos
    ALTER COLUMN id
        ADD GENERATED BY DEFAULT AS IDENTITY (
        SEQUENCE NAME public.like_videos_id_seq
        START WITH 1
        INCREMENT BY 1
        NO MINVALUE
        NO MAXVALUE
        CACHE 1
        );

-- refresh_tokens
DROP SEQUENCE IF EXISTS public.refresh_tokens_id_seq CASCADE;

ALTER TABLE public.refresh_tokens
    ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public.refresh_tokens
    ALTER COLUMN id
        ADD GENERATED BY DEFAULT AS IDENTITY (
        SEQUENCE NAME public.refresh_tokens_id_seq
        START WITH 1
        INCREMENT BY 1
        NO MINVALUE
        NO MAXVALUE
        CACHE 1
        );

COMMIT;

-- Done!
