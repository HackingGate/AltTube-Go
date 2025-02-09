-- Create "users" table
CREATE TABLE "users" ("id" text NOT NULL, "email" text NULL, "password" text NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "deleted_at" timestamptz NULL, CONSTRAINT "uni_users_id" PRIMARY KEY ("id"), CONSTRAINT "uni_users_email" UNIQUE ("email"));
-- Create index "idx_users_deleted_at" to table: "users"
CREATE INDEX "idx_users_deleted_at" ON "users" ("deleted_at");
-- Create "refresh_tokens" table
CREATE TABLE "refresh_tokens" ("id" bigserial NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "deleted_at" timestamptz NULL, "token" text NULL, "expiry" timestamptz NULL, "user_agent" text NULL, "ip_address" text NULL, "user_id" text NULL, PRIMARY KEY ("id"), CONSTRAINT "fk_refresh_tokens_user" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create index "idx_refresh_tokens_deleted_at" to table: "refresh_tokens"
CREATE INDEX "idx_refresh_tokens_deleted_at" ON "refresh_tokens" ("deleted_at");
-- Create index "idx_refresh_tokens_user_id" to table: "refresh_tokens"
CREATE INDEX "idx_refresh_tokens_user_id" ON "refresh_tokens" ("user_id");
-- Create index "idx_token" to table: "refresh_tokens"
CREATE INDEX "idx_token" ON "refresh_tokens" ("token");
-- Create "access_tokens" table
CREATE TABLE "access_tokens" ("id" bigserial NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "deleted_at" timestamptz NULL, "token" text NULL, "user_id" text NULL, "expiry" timestamptz NULL, "refresh_token_id" bigint NULL, PRIMARY KEY ("id"), CONSTRAINT "fk_access_tokens_refresh_token" FOREIGN KEY ("refresh_token_id") REFERENCES "refresh_tokens" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "fk_access_tokens_user" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create index "idx_access_tokens_deleted_at" to table: "access_tokens"
CREATE INDEX "idx_access_tokens_deleted_at" ON "access_tokens" ("deleted_at");
-- Create index "idx_access_tokens_refresh_token_id" to table: "access_tokens"
CREATE INDEX "idx_access_tokens_refresh_token_id" ON "access_tokens" ("refresh_token_id");
-- Create index "idx_access_tokens_user_id" to table: "access_tokens"
CREATE INDEX "idx_access_tokens_user_id" ON "access_tokens" ("user_id");
-- Create "videos" table
CREATE TABLE "videos" ("id" text NOT NULL, "title" text NOT NULL, "description" text NOT NULL, "upload_date" timestamptz NOT NULL, "uploader" text NOT NULL, "uploader_url" text NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "deleted_at" timestamptz NULL, "thumbnail_url" text NOT NULL, CONSTRAINT "uni_videos_id" PRIMARY KEY ("id"));
-- Create index "idx_videos_deleted_at" to table: "videos"
CREATE INDEX "idx_videos_deleted_at" ON "videos" ("deleted_at");
-- Create "like_videos" table
CREATE TABLE "like_videos" ("id" bigserial NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "deleted_at" timestamptz NULL, "user_id" text NULL, "video_id" text NULL, PRIMARY KEY ("id"), CONSTRAINT "fk_like_videos_user" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "fk_like_videos_video" FOREIGN KEY ("video_id") REFERENCES "videos" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create index "idx_like_videos_deleted_at" to table: "like_videos"
CREATE INDEX "idx_like_videos_deleted_at" ON "like_videos" ("deleted_at");
-- Create index "idx_like_videos_video_id" to table: "like_videos"
CREATE INDEX "idx_like_videos_video_id" ON "like_videos" ("video_id");
