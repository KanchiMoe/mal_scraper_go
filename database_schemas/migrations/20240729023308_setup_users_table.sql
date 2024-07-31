-- Create "change_tracking" table
CREATE TABLE "change_tracking" ("uuid" uuid NOT NULL, "timestamp" uuid NOT NULL, "category" text NOT NULL, "item_id" bigint NOT NULL, "field" text NOT NULL, "old_value" text NOT NULL, "new_value" text NOT NULL);
-- Create "username_changes" table
CREATE TABLE "username_changes" ("uuid" uuid NOT NULL, "id" bigint NOT NULL, "old_username" text NOT NULL, "new_username" text NOT NULL, "timestamp" timestamptz NOT NULL, PRIMARY KEY ("uuid"));
-- Create "users" table
CREATE TABLE "users" ("id" bigint NOT NULL, "username" text NOT NULL, "last_interacted" timestamptz NOT NULL, PRIMARY KEY ("id"));
