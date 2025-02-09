table "access_tokens" {
  schema = schema.public
  column "id" {
    null = false
    type = bigserial
  }
  column "created_at" {
    null = true
    type = timestamptz
  }
  column "updated_at" {
    null = true
    type = timestamptz
  }
  column "deleted_at" {
    null = true
    type = timestamptz
  }
  column "token" {
    null = true
    type = text
  }
  column "user_id" {
    null = true
    type = text
  }
  column "expiry" {
    null = true
    type = timestamptz
  }
  column "refresh_token_id" {
    null = true
    type = bigint
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "fk_access_tokens_refresh_token" {
    columns     = [column.refresh_token_id]
    ref_columns = [table.refresh_tokens.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  foreign_key "fk_access_tokens_user" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "idx_access_tokens_deleted_at" {
    columns = [column.deleted_at]
  }
  index "idx_access_tokens_refresh_token_id" {
    columns = [column.refresh_token_id]
  }
  index "idx_access_tokens_user_id" {
    columns = [column.user_id]
  }
}
table "like_videos" {
  schema = schema.public
  column "id" {
    null = false
    type = bigserial
  }
  column "created_at" {
    null = true
    type = timestamptz
  }
  column "updated_at" {
    null = true
    type = timestamptz
  }
  column "deleted_at" {
    null = true
    type = timestamptz
  }
  column "user_id" {
    null = true
    type = text
  }
  column "video_id" {
    null = true
    type = text
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "fk_like_videos_user" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  foreign_key "fk_like_videos_video" {
    columns     = [column.video_id]
    ref_columns = [table.videos.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "idx_like_videos_deleted_at" {
    columns = [column.deleted_at]
  }
  index "idx_like_videos_video_id" {
    columns = [column.video_id]
  }
}
table "refresh_tokens" {
  schema = schema.public
  column "id" {
    null = false
    type = bigserial
  }
  column "created_at" {
    null = true
    type = timestamptz
  }
  column "updated_at" {
    null = true
    type = timestamptz
  }
  column "deleted_at" {
    null = true
    type = timestamptz
  }
  column "token" {
    null = true
    type = text
  }
  column "expiry" {
    null = true
    type = timestamptz
  }
  column "user_agent" {
    null = true
    type = text
  }
  column "ip_address" {
    null = true
    type = text
  }
  column "user_id" {
    null = true
    type = text
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "fk_refresh_tokens_user" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "idx_refresh_tokens_deleted_at" {
    columns = [column.deleted_at]
  }
  index "idx_refresh_tokens_user_id" {
    columns = [column.user_id]
  }
  index "idx_token" {
    columns = [column.token]
  }
}
table "users" {
  schema = schema.public
  column "id" {
    null = false
    type = text
  }
  column "email" {
    null = true
    type = text
  }
  column "password" {
    null = false
    type = text
  }
  column "created_at" {
    null = true
    type = timestamptz
  }
  column "updated_at" {
    null = true
    type = timestamptz
  }
  column "deleted_at" {
    null = true
    type = timestamptz
  }
  primary_key "uni_users_id" {
    columns = [column.id]
  }
  index "idx_users_deleted_at" {
    columns = [column.deleted_at]
  }
  unique "uni_users_email" {
    columns = [column.email]
  }
}
table "videos" {
  schema = schema.public
  column "id" {
    null = false
    type = text
  }
  column "title" {
    null = false
    type = text
  }
  column "description" {
    null = false
    type = text
  }
  column "upload_date" {
    null = false
    type = timestamptz
  }
  column "uploader" {
    null = false
    type = text
  }
  column "uploader_url" {
    null = false
    type = text
  }
  column "created_at" {
    null = true
    type = timestamptz
  }
  column "updated_at" {
    null = true
    type = timestamptz
  }
  column "deleted_at" {
    null = true
    type = timestamptz
  }
  column "thumbnail_url" {
    null = false
    type = text
  }
  primary_key "uni_videos_id" {
    columns = [column.id]
  }
  index "idx_videos_deleted_at" {
    columns = [column.deleted_at]
  }
}
schema "public" {
  comment = "standard public schema"
}
