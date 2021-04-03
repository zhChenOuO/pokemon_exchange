-- +goose Up
CREATE SCHEMA IF NOT EXISTS pokemon;

DROP TABLE IF EXISTS pokemon."identity_accounts";

CREATE TABLE pokemon."identity_accounts" (
    "id" serial NOT NULL PRIMARY KEY,
    "name" varchar(50) DEFAULT '' :: character varying,
    "password" varchar(30) DEFAULT '' :: character varying,
    "created_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON COLUMN pokemon."identity_accounts"."name" IS '姓名';

--- 
DROP TABLE IF EXISTS pokemon."cards";

CREATE TABLE pokemon."cards"(
    "id" serial NOT NULL PRIMARY KEY,
    "name" varchar(200) DEFAULT '' :: character varying
);

COMMENT ON COLUMN pokemon."cards"."name" IS '卡片名稱';

---
DROP TABLE IF EXISTS pokemon."spot_orders";

CREATE TABLE pokemon."spot_orders"(
    "id" serial NOT NULL PRIMARY KEY,
    "uuid" varchar(25) DEFAULT '' :: character varying,
    "card_id" BIGINT NOT NULL,
    "user_id" BIGINT NOT NULL,
    "status" SMALLINT NOT NULL,
    "type" SMALLINT NOT NULL,
    "trade_side" SMALLINT NOT NULL,
    "expected_amount" decimal NOT NULL,
    "card_quantity" decimal NOT NULL,
    "created_at" timestamp DEFAULT now() NOT NULL,
    "updated_at" timestamp DEFAULT now() NOT NULL
);

COMMENT ON COLUMN pokemon."spot_orders"."expected_amount" IS '預期金額';

COMMENT ON COLUMN pokemon."spot_orders"."card_quantity" IS '卡片數量';

COMMENT ON COLUMN pokemon."spot_orders"."trade_side" IS '交易方向, 1:買,2:賣';

COMMENT ON COLUMN pokemon."spot_orders"."type" IS '1:掛單者, 2:吃單者';

--- 
DROP TABLE IF EXISTS pokemon."users";

CREATE TABLE pokemon."users"(
    "id" serial NOT NULL PRIMARY KEY,
    "created_at" timestamp DEFAULT now() NOT NULL,
    "updated_at" timestamp DEFAULT now() NOT NULL
);

---
DROP TABLE IF EXISTS pokemon."trade_orders";

CREATE TABLE pokemon."trade_orders"(
    "id" serial NOT NULL PRIMARY KEY,
    "turnover" decimal NOT NULL,
    "taker_order_id" BIGINT NOT NULL,
    "maker_order_id" BIGINT NOT NULL,
    "created_at" timestamp DEFAULT now() NOT NULL
);

COMMENT ON COLUMN pokemon."trade_orders"."turnover" IS '成交金額';

-- +goose Down