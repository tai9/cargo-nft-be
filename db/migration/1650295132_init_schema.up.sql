CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "bio" varchar NOT NULL DEFAULT '',
  "email" varchar NOT NULL DEFAULT '',
  "wallet_address" varchar UNIQUE NOT NULL,
  "avatar" varchar NOT NULL DEFAULT '',
  "banner_img" varchar NOT NULL DEFAULT '',
  "ins_link" varchar NOT NULL DEFAULT '',
  "twitter_link" varchar NOT NULL DEFAULT '',
  "website_link" varchar NOT NULL DEFAULT '',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "collections" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "name" varchar NOT NULL DEFAULT '',
  "description" varchar NOT NULL DEFAULT '',
  "blockchain" varchar NOT NULL DEFAULT '',
  "owners" varchar NOT NULL DEFAULT '',
  "payment_token" varchar NOT NULL DEFAULT '',
  "creator_earning" varchar NOT NULL DEFAULT '',
  "featured_img" varchar NOT NULL DEFAULT '',
  "banner_img" varchar NOT NULL DEFAULT '',
  "ins_link" varchar NOT NULL DEFAULT '',
  "twitter_link" varchar NOT NULL DEFAULT '',
  "website_link" varchar NOT NULL DEFAULT '',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "categories" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL DEFAULT '',
  "featured_img" varchar NOT NULL DEFAULT '',
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
  "quantity" float NOT NULL DEFAULT 0,
  "usd_price" float NOT NULL DEFAULT 0,
  "floor_difference" float NOT NULL DEFAULT 0,
  "token" varchar NOT NULL DEFAULT '',
  "expiration" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "listings" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "nft_id" bigint NOT NULL,
  "from_user_id" bigint NOT NULL,
  "quantity" float NOT NULL DEFAULT 0,
  "usd_price" float NOT NULL DEFAULT 0,
  "usd_unit_price" float NOT NULL DEFAULT 0,
  "token" varchar NOT NULL DEFAULT '',
  "expiration" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "nfts" (
  "id" bigserial PRIMARY KEY,
  "owner_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  "collection_id" bigint NOT NULL,
  "name" varchar NOT NULL DEFAULT '',
  "description" varchar NOT NULL DEFAULT '',
  "featured_img" varchar NOT NULL DEFAULT '',
  "supply" bigint NOT NULL DEFAULT 0,
  "views" bigint NOT NULL DEFAULT 0,
  "favorites" varchar NOT NULL DEFAULT '',
  "contract_address" varchar NOT NULL DEFAULT '',
  "token_id" varchar NOT NULL DEFAULT '',
  "token_standard" varchar NOT NULL DEFAULT '',
  "blockchain" varchar NOT NULL DEFAULT '',
  "metadata" varchar NOT NULL DEFAULT '',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transactions" (
  "id" bigserial PRIMARY KEY,
  "nft_id" bigint NOT NULL,
  "event" varchar NOT NULL DEFAULT '',
  "token" varchar NOT NULL DEFAULT '',
  "quantity" float NOT NULL DEFAULT 0,
  "from_user_id" bigint NOT NULL,
  "to_user_id" bigint NOT NULL,
  "transaction_hash" varchar NOT NULL DEFAULT '',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "crypto_currencies" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL DEFAULT '',
  "code" varchar NOT NULL DEFAULT '',
  "price" float NOT NULL DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "collections" ("user_id");

CREATE INDEX ON "cate_collections" ("collection_id");

CREATE INDEX ON "cate_collections" ("category_id");

CREATE INDEX ON "cate_collections" ("collection_id", "category_id");

CREATE INDEX ON "offers" ("user_id");

CREATE INDEX ON "offers" ("nft_id");

CREATE INDEX ON "listings" ("user_id");

CREATE INDEX ON "listings" ("nft_id");

CREATE INDEX ON "listings" ("from_user_id");

CREATE INDEX ON "nfts" ("user_id");

CREATE INDEX ON "nfts" ("collection_id");

CREATE INDEX ON "nfts" ("owner_id");

CREATE INDEX ON "transactions" ("from_user_id");

CREATE INDEX ON "transactions" ("to_user_id");

CREATE INDEX ON "transactions" ("from_user_id", "to_user_id");

CREATE INDEX ON "transactions" ("id", "nft_id");

ALTER TABLE "collections" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "cate_collections" ADD FOREIGN KEY ("collection_id") REFERENCES "collections" ("id");

ALTER TABLE "cate_collections" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "offers" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "offers" ADD FOREIGN KEY ("nft_id") REFERENCES "nfts" ("id");

ALTER TABLE "listings" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "listings" ADD FOREIGN KEY ("nft_id") REFERENCES "nfts" ("id");

ALTER TABLE "listings" ADD FOREIGN KEY ("from_user_id") REFERENCES "users" ("id");

ALTER TABLE "nfts" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("id");

ALTER TABLE "nfts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "nfts" ADD FOREIGN KEY ("collection_id") REFERENCES "collections" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("nft_id") REFERENCES "nfts" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("from_user_id") REFERENCES "users" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("to_user_id") REFERENCES "users" ("id");
