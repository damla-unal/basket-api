BEGIN;

drop table customer;

drop table product;

drop table "order";

drop type if exists cart_status cascade;

drop table cart;

COMMIT;