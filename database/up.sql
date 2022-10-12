CREATE TABLE IF NOT EXISTS "users"
(
    "id"         serial PRIMARY KEY,
    "email"      varchar(255) NOT NULL,
    "password"   varchar(255) NOT NULL,
    "first_name" varchar(255) NOT NULL,
    "last_name"  varchar(255) NOT NULL,
    "role"       varchar(255) NOT NULL,
    "photo"      varchar(255) NOT NULL,
    "country"    varchar(255) NOT NULL,
    "created_at" timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "is_deleted" boolean      NOT NULL DEFAULT false
);

CREATE TABLE IF NOT EXISTS "products"
(
    "id"          serial PRIMARY KEY,
    "name"        varchar(255)   NOT NULL,
    "description" varchar(255)   NOT NULL,
    "price"       decimal(10, 2) NOT NULL,
    "photo"       varchar(255)   NOT NULL,
    "creator_id"  int            NOT NULL,
    "stock"       int            NOT NULL,
    "created_at"  timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at"  timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at"  timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "is_deleted"  boolean        NOT NULL DEFAULT false,
    FOREIGN KEY ("creator_id") REFERENCES "users" ("id")
);

CREATE TABLE IF NOT EXISTS "payments"
(
    "id"         serial PRIMARY KEY,
    "amount"     decimal(10, 2) NOT NULL,
    "status"     varchar(255)   NOT NULL,
    "created_at" timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "order_id"   int            NOT NULL,
    "is_deleted" boolean        NOT NULL DEFAULT false
);

CREATE TABLE IF NOT EXISTS "orders"
(
    "id"         serial PRIMARY KEY,
    "payment_id" int            NOT NULL,
    "buyer_id"   int            NOT NULL,
    "amount"     decimal(10, 2) NOT NULL,
    "status"     varchar(255)   NOT NULL,
    "created_at" timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "is_deleted" boolean        NOT NULL DEFAULT false,
    FOREIGN KEY ("payment_id") REFERENCES "payments" ("id"),
    FOREIGN KEY ("buyer_id") REFERENCES "users" ("id")
);

CREATE TABLE IF NOT EXISTS "order_items"
(
    "id"         serial PRIMARY KEY,
    "order_id"   int       NOT NULL,
    "product_id" int       NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "is_deleted" boolean   NOT NULL DEFAULT false,
    FOREIGN KEY ("order_id") REFERENCES "orders" ("id"),
    FOREIGN KEY ("product_id") REFERENCES "products" ("id")
);


