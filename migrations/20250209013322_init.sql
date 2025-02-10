-- Create "users" table
create table users (
    id text not null,
    email text null,
    password text not null,
    created_at timestamptz null,
    updated_at timestamptz null,
    deleted_at timestamptz null,
    constraint uni_users_id primary key (id),
    constraint uni_users_email unique (email)
);
-- Create index "idx_users_deleted_at" to table: "users"
create index idx_users_deleted_at on users (deleted_at);
-- Create "refresh_tokens" table
create table refresh_tokens (
    id bigserial not null,
    created_at timestamptz null,
    updated_at timestamptz null,
    deleted_at timestamptz null,
    token text null,
    expiry timestamptz null,
    user_agent text null,
    ip_address text null,
    user_id text null,
    primary key (id),
    constraint fk_refresh_tokens_user foreign key (user_id) references users (
        id
    ) on update no action on delete no action
);
-- Create index "idx_refresh_tokens_deleted_at" to table: "refresh_tokens"
create index idx_refresh_tokens_deleted_at on refresh_tokens (deleted_at);
-- Create index "idx_refresh_tokens_user_id" to table: "refresh_tokens"
create index idx_refresh_tokens_user_id on refresh_tokens (user_id);
-- Create index "idx_token" to table: "refresh_tokens"
create index idx_token on refresh_tokens (token);
-- Create "access_tokens" table
create table access_tokens (
    id bigserial not null,
    created_at timestamptz null,
    updated_at timestamptz null,
    deleted_at timestamptz null,
    token text null,
    user_id text null,
    expiry timestamptz null,
    refresh_token_id bigint null,
    primary key (id),
    constraint fk_access_tokens_refresh_token foreign key (refresh_token_id) references refresh_tokens (
        id
    ) on update no action on delete no action,
    constraint fk_access_tokens_user foreign key (user_id) references users (id) on update no action on delete no action
);
-- Create index "idx_access_tokens_deleted_at" to table: "access_tokens"
create index idx_access_tokens_deleted_at on access_tokens (deleted_at);
-- Create index "idx_access_tokens_refresh_token_id" to table: "access_tokens"
create index idx_access_tokens_refresh_token_id on access_tokens (refresh_token_id);
-- Create index "idx_access_tokens_user_id" to table: "access_tokens"
create index idx_access_tokens_user_id on access_tokens (user_id);
-- Create "videos" table
create table videos (
    id text not null,
    title text not null,
    description text not null,
    upload_date timestamptz not null,
    uploader text not null,
    uploader_url text not null,
    created_at timestamptz null,
    updated_at timestamptz null,
    deleted_at timestamptz null,
    thumbnail_url text not null,
    constraint uni_videos_id primary key (id)
);
-- Create index "idx_videos_deleted_at" to table: "videos"
create index idx_videos_deleted_at on videos (deleted_at);
-- Create "like_videos" table
create table like_videos (
    id bigserial not null,
    created_at timestamptz null,
    updated_at timestamptz null,
    deleted_at timestamptz null,
    user_id text null,
    video_id text null,
    primary key (id),
    constraint fk_like_videos_user foreign key (user_id) references users (id) on update no action on delete no action,
    constraint fk_like_videos_video foreign key (video_id) references videos (
        id
    ) on update no action on delete no action
);
-- Create index "idx_like_videos_deleted_at" to table: "like_videos"
create index idx_like_videos_deleted_at on like_videos (deleted_at);
-- Create index "idx_like_videos_video_id" to table: "like_videos"
create index idx_like_videos_video_id on like_videos (video_id);
