-- Create "email_notification" table
CREATE TABLE "email_notification" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "email" character varying(100) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_email_notification_deleted_at" to table: "email_notification"
CREATE INDEX "idx_email_notification_deleted_at" ON "email_notification" ("deleted_at");
