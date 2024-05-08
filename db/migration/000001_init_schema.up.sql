CREATE TABLE "users" (
 "username" varchar UNIQUE PRIMARY KEY,
 "password_hash" varchar NOT NULL,
 "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "profiles" (
 "id" bigserial PRIMARY KEY,
 "username" varchar NOT NULL,
 "first_name" varchar NOT NULL,
 "last_name" varchar NOT NULL,
 "email" varchar NOT NULL,
 "phone_number" varchar NOT NULL,
 "birthday" varchar NOT NULL,
 "nickname" varchar NOT NULL
);

CREATE TABLE "food" (
 "id" bigserial PRIMARY KEY,
 "name" varchar NOT NULL,
 "description" varchar NOT NULL,
 "price" integer NOT NULL,
 "rate" integer NOT NULL,
 "discount" integer NOT NULL,
 "food_tag" varchar NOT NULL
);

CREATE TABLE "addresses" (
 "id" bigserial PRIMARY KEY,
 "username" varchar NOT NULL,
 "address_line" varchar NOT NULL,
 "address_tag" varchar NOT NULL,
 "phone_number" varchar NOT NULL
);

CREATE TABLE "favorites" (
 "id" bigserial PRIMARY KEY,
 "username" varchar NOT NULL,
 "food_id" bigint NOT NULL,
 "added_at" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "profiles" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "addresses" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "favorites" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "favorites" ADD FOREIGN KEY ("food_id") REFERENCES "food" ("id");
