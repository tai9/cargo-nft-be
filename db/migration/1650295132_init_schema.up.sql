CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "bio" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "wallet_address" varchar UNIQUE NOT NULL,
  "avatar" varchar NOT NULL,
  "banner_img" varchar NOT NULL,
  "ins_link" varchar NOT NULL,
  "twitter_link" varchar NOT NULL,
  "website_link" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "collections" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "blockchain" varchar NOT NULL,
  "owners" varchar NOT NULL,
  "payment_token" varchar NOT NULL,
  "creator_earning" varchar NOT NULL,
  "featured_img" varchar NOT NULL,
  "banner_img" varchar NOT NULL,
  "ins_link" varchar NOT NULL,
  "twitter_link" varchar NOT NULL,
  "website_link" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "categories" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "featured_img" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "cate_collections" (
  "id" bigserial PRIMARY KEY,
  "collection_id" bigint NOT NULL,
  "category_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "offers" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "nft_id" bigint NOT NULL,
  "quantity" double precision NOT NULL,
  "usd_price" double precision NOT NULL,
  "token" varchar NOT NULL,
  "floor_difference" varchar NOT NULL,
  "expiration" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "nfts" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "collection_id" bigint NOT NULL,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "featured_img" varchar NOT NULL,
  "supply" int NOT NULL,
  "views" varchar NOT NULL,
  "favorites" varchar NOT NULL,
  "contract_address" varchar NOT NULL,
  "token_id" varchar NOT NULL,
  "token_standard" varchar NOT NULL,
  "blockchain" varchar NOT NULL,
  "metadata" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transactions" (
  "id" bigserial PRIMARY KEY,
  "nft_id" bigint NOT NULL,
  "event" varchar NOT NULL,
  "quantity" double precision NOT NULL,
  "token" varchar NOT NULL,
  "from_user_id" bigint NOT NULL,
  "to_user_id" bigint NOT NULL,
  "transaction_hash" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "crypto_currencies" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "code" varchar UNIQUE NOT NULL,
  "price" double precision NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "collections" ("user_id");

CREATE INDEX ON "cate_collections" ("collection_id");

CREATE INDEX ON "cate_collections" ("category_id");

CREATE INDEX ON "cate_collections" ("collection_id", "category_id");

CREATE INDEX ON "offers" ("user_id");

CREATE INDEX ON "offers" ("nft_id");

CREATE INDEX ON "nfts" ("user_id");

CREATE INDEX ON "nfts" ("collection_id");

CREATE INDEX ON "transactions" ("from_user_id");

CREATE INDEX ON "transactions" ("to_user_id");

CREATE INDEX ON "transactions" ("from_user_id", "to_user_id");

ALTER TABLE "collections" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "cate_collections" ADD FOREIGN KEY ("collection_id") REFERENCES "collections" ("id");

ALTER TABLE "cate_collections" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "offers" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "offers" ADD FOREIGN KEY ("nft_id") REFERENCES "nfts" ("id");

ALTER TABLE "nfts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "nfts" ADD FOREIGN KEY ("collection_id") REFERENCES "collections" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("nft_id") REFERENCES "nfts" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("from_user_id") REFERENCES "users" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("to_user_id") REFERENCES "users" ("id");
