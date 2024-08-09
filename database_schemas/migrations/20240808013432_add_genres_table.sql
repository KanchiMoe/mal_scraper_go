-- Create "genres" table
CREATE TABLE "genres" ("uuid" uuid NOT NULL, "id" smallint NOT NULL, "name" text NOT NULL, "description" text NULL, "count" integer NOT NULL, "last_interacted" timestamptz NOT NULL, PRIMARY KEY ("uuid"));
