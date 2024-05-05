CREATE TABLE "users" (
 "username" varchar UNIQUE PRIMARY KEY,
 "password_hash" varchar NOT NULL,
 "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "profiles" (
 "id" bigserial PRIMARY KEY,
 "user_id" varchar NOT NULL,
 "first_name" varchar,
 "last_name" varchar,
 "email" varchar UNIQUE NOT NULL,
 "phone_number" varchar,
 "birthday" date,
 "nickname" varchar
);

CREATE TABLE "food" (
 "id" bigserial PRIMARY KEY,
 "name" varchar NOT NULL,
 "description" text NOT NULL,
 "price" numeric NOT NULL,
 "rate" integer NOT NULL,
 "discount" integer NOT NULL
);

CREATE TABLE "addresses" (
 "id" bigserial PRIMARY KEY,
 "user_id" varchar NOT NULL,
 "address_line" varchar,
 "address_tag" varchar,
 "phone_number" varchar
);

CREATE TABLE "favorites" (
 "id" bigserial PRIMARY KEY,
 "user_id" varchar NOT NULL,
 "food_id" bigint NOT NULL,
 "added_at" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "profiles" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("username");

ALTER TABLE "addresses" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("username");

ALTER TABLE "favorites" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("username");

ALTER TABLE "favorites" ADD FOREIGN KEY ("food_id") REFERENCES "food" ("id");
