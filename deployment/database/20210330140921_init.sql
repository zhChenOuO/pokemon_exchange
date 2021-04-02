-- +goose Up
DROP TABLE IF EXISTS "identity_accounts";

CREATE TABLE "identity_accounts" (
    "id" serial NOT NULL PRIMARY KEY,
    "name" TEXT UNIQUE,
    "password" TEXT,
    "created_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON COLUMN "identity_accounts"."name" IS '姓名';

--- 
DROP TABLE IF EXISTS "cards";

CREATE TABLE "cards"(
    "id" serial NOT NULL PRIMARY KEY,
    "name" TEXT
);

COMMENT ON COLUMN "cards"."name" IS '卡片名稱';
---
DROP TABLE IF EXISTS "spot_orders";

CREATE TABLE "spot_orders"(
    "id" serial NOT NULL PRIMARY KEY,
    "uuid" TEXT,
    "card_id" BIGINT,
    "user_id" BIGINT,
    "status" SMALLINT,
    "type" SMALLINT,
    "trade_side" SMALLINT,
    "expected_amount" decimal,
    "card_quantity" decimal,
    "created_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON COLUMN "spot_orders"."expected_amount" IS '預期金額';

COMMENT ON COLUMN "spot_orders"."card_quantity" IS '卡片數量';

COMMENT ON COLUMN "spot_orders"."trade_side" IS '交易方向, 1:買,2:賣';

COMMENT ON COLUMN "spot_orders"."type" IS '1:掛單者, 2:吃單者';

--- 
DROP TABLE IF EXISTS "user";

CREATE TABLE "users"(
    "id" serial NOT NULL PRIMARY KEY,
    "created_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
);

---
DROP TABLE IF EXISTS "trade_orders";

CREATE TABLE "trade_orders"(
    "id" serial NOT NULL PRIMARY KEY,
    "turnover" decimal,
    "taker_order_id" BIGINT,
    "maker_order_id" BIGINT,
    "created_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON COLUMN "trade_orders"."turnover" IS '成交金額';

-- +goose Down