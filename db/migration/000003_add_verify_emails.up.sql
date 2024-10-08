CREATE TABLE "verify_emails" (
 "id" bigserial PRIMARY KEY,
 "username" varchar NOT NULL,
 "email" varchar NOT NULL,
 "secret_code" varchar NOT NULL,
 "is_used" bool NOT NULL DEFAULT (false),
 "created_at" timestamp NOT NULL DEFAULT (now()),
 "expire_at" timestamp NOT NULL DEFAULT (now() + interval '15 minute')
);

ALTER TABLE "verify_emails" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "profile" ADD COLUMN "is_email_verified" bool NOT NULL  DEFAULT false;
