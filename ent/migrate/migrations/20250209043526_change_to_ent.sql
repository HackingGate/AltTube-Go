--
-- MIGRATION SCRIPT: from old schema (GORM) to new schema (Ent)
--

begin;

--------------------------------------------------------------------------
-- STEP 0: Drop existing foreign keys and constraints that we will recreate
--------------------------------------------------------------------------

-- access_tokens
alter table if exists public.access_tokens drop constraint if exists fk_access_tokens_refresh_token;
alter table if exists public.access_tokens drop constraint if exists fk_access_tokens_user;
-- The old PK name is the same in both schemas (access_tokens_pkey), so we leave it
-- unless we decide to drop/recreate it. We'll handle the IDENTITY update separately.

-- like_videos
alter table if exists public.like_videos drop constraint if exists fk_like_videos_user;
alter table if exists public.like_videos drop constraint if exists fk_like_videos_video;
-- old PK is like_videos_pkey, same name in new schema, so we keep it.

-- refresh_tokens
alter table if exists public.refresh_tokens drop constraint if exists fk_refresh_tokens_user;
-- old PK is refresh_tokens_pkey, same name in new schema, so we keep it.

-- users
-- Drop the old unique constraint on email:
alter table if exists public.users drop constraint if exists uni_users_email;
-- Drop the old PK named uni_users_id; new schema calls it users_pkey:
alter table if exists public.users drop constraint if exists uni_users_id;

-- videos
-- Drop the old PK named uni_videos_id; new schema calls it videos_pkey:
alter table if exists public.videos drop constraint if exists uni_videos_id;

commit;


begin;

--------------------------------------------------------------------------------
-- STEP 1: Drop old/obsolete indexes that do not appear in the new schema
--------------------------------------------------------------------------------

-- Access tokens indexes
drop index if exists public.idx_access_tokens_deleted_at;
drop index if exists public.idx_access_tokens_refresh_token_id;
drop index if exists public.idx_access_tokens_user_id;

-- Like videos indexes
drop index if exists public.idx_like_videos_deleted_at;
drop index if exists public.idx_like_videos_video_id;

-- Refresh tokens indexes
drop index if exists public.idx_refresh_tokens_deleted_at;
drop index if exists public.idx_refresh_tokens_user_id;
drop index if exists public.idx_token;  -- on refresh_tokens(token)

-- Users
drop index if exists public.idx_users_deleted_at;

-- Videos
drop index if exists public.idx_videos_deleted_at;

commit;


begin;

--------------------------------------------------------------------------------
-- STEP 2: Remove deleted_at columns, rename created_at/updated_at -> create_time/update_time, set NOT NULL
--------------------------------------------------------------------------------

-- 2.1: access_tokens
alter table public.access_tokens
drop column if exists deleted_at;
alter table public.access_tokens
rename column created_at to create_time;
alter table public.access_tokens
rename column updated_at to update_time;
alter table public.access_tokens
alter column create_time set not null;
alter table public.access_tokens
alter column update_time set not null;

-- 2.2: like_videos
alter table public.like_videos
drop column if exists deleted_at;
alter table public.like_videos
rename column created_at to create_time;
alter table public.like_videos
rename column updated_at to update_time;
alter table public.like_videos
alter column create_time set not null;
alter table public.like_videos
alter column update_time set not null;

-- 2.3: refresh_tokens
alter table public.refresh_tokens
drop column if exists deleted_at;
alter table public.refresh_tokens
rename column created_at to create_time;
alter table public.refresh_tokens
rename column updated_at to update_time;
alter table public.refresh_tokens
alter column create_time set not null;
alter table public.refresh_tokens
alter column update_time set not null;

-- 2.4: users
alter table public.users
drop column if exists deleted_at;
alter table public.users
rename column created_at to create_time;
alter table public.users
rename column updated_at to update_time;
alter table public.users
alter column create_time set not null;
alter table public.users
alter column update_time set not null;

-- 2.5: videos
alter table public.videos
drop column if exists deleted_at;
alter table public.videos
rename column created_at to create_time;
alter table public.videos
rename column updated_at to update_time;
alter table public.videos
alter column create_time set not null;
alter table public.videos
alter column update_time set not null;

commit;


begin;

--------------------------------------------------------------------------------
-- STEP 3: Convert column types, enforce NOT NULL
--------------------------------------------------------------------------------

-- access_tokens:
--   user_id: text -> uuid
--   token: text -> character varying
--   refresh_token_id: stays bigint, but enforce NOT NULL
--   expiry: stays timestamp with time zone
alter table public.access_tokens
alter column user_id type uuid using user_id::uuid;
alter table public.access_tokens
alter column token type character varying;
alter table public.access_tokens
alter column refresh_token_id set not null;
alter table public.access_tokens
alter column user_id set not null;

-- like_videos:
--   user_id: text -> uuid
--   video_id: text -> character varying
alter table public.like_videos
alter column user_id type uuid using user_id::uuid;
alter table public.like_videos
alter column video_id type character varying;
alter table public.like_videos
alter column user_id set not null;
alter table public.like_videos
alter column video_id set not null;

-- refresh_tokens:
--   user_id: text -> uuid
--   token, user_agent, ip_address: text -> character varying
alter table public.refresh_tokens
alter column user_id type uuid using user_id::uuid;
alter table public.refresh_tokens
alter column token type character varying;
alter table public.refresh_tokens
alter column user_agent type character varying;
alter table public.refresh_tokens
alter column ip_address type character varying;
alter table public.refresh_tokens
alter column user_id set not null;

-- users:
--   id: text -> uuid
--   email: text -> character varying NOT NULL
--   password: text -> character varying NOT NULL
alter table public.users
alter column id type uuid using id::uuid;
alter table public.users
alter column email type character varying;
alter table public.users
alter column email set not null;
alter table public.users
alter column password type character varying;
alter table public.users
alter column password set not null;

-- videos:
--   id: text -> character varying
--   title, description, uploader, uploader_url, thumbnail_url: text -> character varying
alter table public.videos
alter column id type character varying;
alter table public.videos
alter column title type character varying;
alter table public.videos
alter column description type character varying;
alter table public.videos
alter column uploader type character varying;
alter table public.videos
alter column uploader_url type character varying;
alter table public.videos
alter column thumbnail_url type character varying;

commit;


begin;

--------------------------------------------------------------------------------
-- STEP 4: Convert numeric ID columns to Postgres identity
--         (replacing old sequences/SET DEFAULT nextval)
--------------------------------------------------------------------------------

-- 4.1: access_tokens.id
-- Drop default nextval if it exists, then add IDENTITY
alter table public.access_tokens
alter column id drop default;

alter table public.access_tokens
alter column id
add generated by default as identity (
    sequence name public.access_tokens_id_seq
    start with 1
    increment by 1
    no minvalue
    no maxvalue
    cache 1
);

-- 4.2: like_videos.id
alter table public.like_videos
alter column id drop default;

alter table public.like_videos
alter column id
add generated by default as identity (
    sequence name public.like_videos_id_seq
    start with 1
    increment by 1
    no minvalue
    no maxvalue
    cache 1
);

-- 4.3: refresh_tokens.id
alter table public.refresh_tokens
alter column id drop default;

alter table public.refresh_tokens
alter column id
add generated by default as identity (
    sequence name public.refresh_tokens_id_seq
    start with 1
    increment by 1
    no minvalue
    no maxvalue
    cache 1
);

-- users.id and videos.id remain text/uuid/character varying, so no identity needed.

commit;


begin;

--------------------------------------------------------------------------------
-- STEP 5: Recreate new PK constraints or unique indexes as needed
--------------------------------------------------------------------------------

-- 5.1: users -> new PK name = users_pkey, plus unique index on email
alter table only public.users
add constraint users_pkey primary key (id);

create unique index users_email_key on public.users using btree (email);

-- 5.2: videos -> new PK name = videos_pkey
alter table only public.videos
add constraint videos_pkey primary key (id);

-- For access_tokens, like_videos, refresh_tokens the old PK names (xxx_pkey) match the new schema,
-- so we don't need to re-add them unless we explicitly dropped them in Step 0.

commit;


begin;

--------------------------------------------------------------------------------
-- STEP 6: Re-add foreign key constraints with new constraint names
--------------------------------------------------------------------------------

-- access_tokens -> refresh_tokens
alter table only public.access_tokens
add constraint access_tokens_refresh_tokens_access_tokens
foreign key (refresh_token_id) references public.refresh_tokens (id);

-- access_tokens -> users
alter table only public.access_tokens
add constraint access_tokens_users_access_tokens
foreign key (user_id) references public.users (id);

-- like_videos -> users
alter table only public.like_videos
add constraint like_videos_users_like_videos
foreign key (user_id) references public.users (id);

-- like_videos -> videos
alter table only public.like_videos
add constraint like_videos_videos_like_videos
foreign key (video_id) references public.videos (id);

-- refresh_tokens -> users
alter table only public.refresh_tokens
add constraint refresh_tokens_users_refresh_tokens
foreign key (user_id) references public.users (id);

commit;


begin;

-------------------------------------------------------------------------
-- STEP 7: Recreate indexes that are in the new schema but not the old
-------------------------------------------------------------------------

-- Drop the old sequence so we can reuse its name in the identity column.
drop sequence if exists public.access_tokens_id_seq cascade;

-- Now convert the column to an identity column, giving it the same sequence name.
alter table public.access_tokens
alter column id drop default;  -- drop any "nextval" default if still present
alter table public.access_tokens
alter column id
add generated by default as identity (
    sequence name public.access_tokens_id_seq
    start with 1
    increment by 1
    no minvalue
    no maxvalue
    cache 1
);

commit;

begin;

-- like_videos
drop sequence if exists public.like_videos_id_seq cascade;

alter table public.like_videos
alter column id drop default;
alter table public.like_videos
alter column id
add generated by default as identity (
    sequence name public.like_videos_id_seq
    start with 1
    increment by 1
    no minvalue
    no maxvalue
    cache 1
);

-- refresh_tokens
drop sequence if exists public.refresh_tokens_id_seq cascade;

alter table public.refresh_tokens
alter column id drop default;
alter table public.refresh_tokens
alter column id
add generated by default as identity (
    sequence name public.refresh_tokens_id_seq
    start with 1
    increment by 1
    no minvalue
    no maxvalue
    cache 1
);

commit;

-- Done!
