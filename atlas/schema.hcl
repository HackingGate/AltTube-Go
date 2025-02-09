table "access_tokens" {
  schema = schema.public
  column "id" {
    null = false
    type = bigint
    identity {
      generated = BY_DEFAULT
    }
  }
  column "create_time" {
    null = false
    type = timestamptz
  }
  column "update_time" {
    null = false
    type = timestamptz
  }
  column "token" {
    null = true
    type = character_varying
  }
  column "expiry" {
    null = true
    type = timestamptz
  }
  column "refresh_token_id" {
    null = false
    type = bigint
  }
  column "user_id" {
    null = false
    type = uuid
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "access_tokens_refresh_tokens_access_tokens" {
    columns     = [column.refresh_token_id]
    ref_columns = [table.refresh_tokens.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  foreign_key "access_tokens_users_access_tokens" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
}
table "like_videos" {
  schema = schema.public
  column "id" {
    null = false
    type = bigint
    identity {
      generated = BY_DEFAULT
    }
  }
  column "create_time" {
    null = false
    type = timestamptz
  }
  column "update_time" {
    null = false
    type = timestamptz
  }
  column "user_id" {
    null = false
    type = uuid
  }
  column "video_id" {
    null = false
    type = character_varying
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "like_videos_users_like_videos" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  foreign_key "like_videos_videos_like_videos" {
    columns     = [column.video_id]
    ref_columns = [table.videos.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
}
table "refresh_tokens" {
  schema = schema.public
  column "id" {
    null = false
    type = bigint
    identity {
      generated = BY_DEFAULT
    }
  }
  column "create_time" {
    null = false
    type = timestamptz
  }
  column "update_time" {
    null = false
    type = timestamptz
  }
  column "token" {
    null = true
    type = character_varying
  }
  column "expiry" {
    null = true
    type = timestamptz
  }
  column "user_agent" {
    null = true
    type = character_varying
  }
  column "ip_address" {
    null = true
    type = character_varying
  }
  column "user_id" {
    null = false
    type = uuid
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "refresh_tokens_users_refresh_tokens" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
}
table "users" {
  schema = schema.public
  column "id" {
    null = false
    type = uuid
  }
  column "create_time" {
    null = false
    type = timestamptz
  }
  column "update_time" {
    null = false
    type = timestamptz
  }
  column "email" {
    null = true
    type = character_varying
  }
  column "password" {
    null = false
    type = character_varying
  }
  primary_key {
    columns = [column.id]
  }
  index "users_email_key" {
    unique  = true
    columns = [column.email]
  }
}
table "videos" {
  schema = schema.public
  column "id" {
    null = false
    type = character_varying
  }
  column "create_time" {
    null = false
    type = timestamptz
  }
  column "update_time" {
    null = false
    type = timestamptz
  }
  column "title" {
    null = false
    type = character_varying
  }
  column "description" {
    null = false
    type = character_varying
  }
  column "upload_date" {
    null = false
    type = timestamptz
  }
  column "uploader" {
    null = false
    type = character_varying
  }
  column "uploader_url" {
    null = false
    type = character_varying
  }
  column "thumbnail_url" {
    null = false
    type = character_varying
  }
  primary_key {
    columns = [column.id]
  }
}
schema "public" {
  comment = "standard public schema"
}
