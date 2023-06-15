create table users (
   id bigint unsigned primary key auto_increment,
   customer_id bigint unsigned not null,
   username varchar(255) not null,
   password varchar(255) not null, 
   role_id int not null default 1,
   created_at datetime not null,
   updated_at datetime not null, 
   deleted_at datetime
);

create table customers (
   id bigint unsigned primary key auto_increment, 
   legal_name text not null,
   address text not null,
   phone varchar(255) not null,
   birthplace varchar(255) not null,
   birthdate datetime not null,
   nik varchar(50) not null,
   occupation varchar(50) not null,
   ktp_url varchar(255) not null,
   created_at datetime not null,
   updated_at datetime not null, 
   deleted_at datetime
);

create table accounts (
   id bigint unsigned primary key auto_increment,
   customer_id bigint unsigned not null,
   account_no varchar(255) not null,
   account_type tinyint not null,
   card_number varchar(20),
   cvv varchar(5),
   pin varchar(50),
   balance decimal(18,2) not null default 0.00,
   created_at datetime not null,
   updated_at datetime not null, 
   deleted_at datetime
);

create table transactions (
   id bigint unsigned primary key auto_increment,
   account_id bigint unsigned not null,
   recipient_id bigint unsigned not null, 
   trx_type tinyint not null,
   trx_datetime datetime not null,
   trx_status tinyint not null,
   nominal decimal(18,2),
   transaction_fee decimal(18,2),
   created_at datetime not null,
   updated_at datetime not null, 
   deleted_at datetime
);

create table settlements (
   id bigint unsigned primary key auto_increment,
   transaction_id bigint unsigned not null,
   merchant_id bigint unsigned not null,
   nominal decimal(18,2) not null,
   status tinyint not null,
   created_at datetime not null,
   updated_at datetime not null, 
   deleted_at datetime
);

create table merchants (
   id bigint unsigned primary key auto_increment,
   name varchar(255) not null,
   address varchar(255) not null,
   phone varchar(255) not null,
   merchant_code varchar(255) not null,
   created_at datetime not null,
   updated_at datetime not null, 
   deleted_at datetime
);