CREATE TABLE "user" (
    "id" bigserial NOT NULL,
    "username" varchar PRIMARY KEY,
    "password" varchar NOT NULL,
    "full_name" varchar NOT NULL,
    "email" varchar UNIQUE NOT NULL,
    "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
    "created_at" timestamptz NOT NULL DEFAULT (now())
);


CREATE TABLE post (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "description" varchar NOT NULL,
  "content" varchar NOT NULL,
  "user_name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_user FOREIGN KEY(user_name) REFERENCES "user"(username)
);

-- CREATE INDEX ON "post" ("user_name");

-- ALTER TABLE "post" ADD CONSTRAINT "fk_user" UNIQUE ("user_name)

-- ALTER TABLE "post" ADD FOREIGN KEY ("user_name") REFERENCES "user" ("username");