CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE mst_user(
 id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
 name VARCHAR(100) NOT NULL,
 username VARCHAR(100) NOT NULL UNIQUE,
 password VARCHAR(100) NOT NULL,
 role VARCHAR(100) NOT NULL,
 email VARCHAR(100) NOT NULL UNIQUE, 
 phone_number VARCHAR(100) NOT NULL UNIQUE,
 created_at TIMESTAMP,
 updated_at TIMESTAMP
);

CREATE TABLE mst_user_datas(
 id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
 user_id UUID NOT NULl UNIQUE,
 nik VARCHAR(100) UNIQUE NOT NULL,
 jenis_kelamin VARCHAR(10),
 tanggal_lahir TIMESTAMP,
 umur INTEGER,
 photo VARCHAR(100),
 FOREIGN KEY(user_id) REFERENCES mst_user(id)
);

CREATE TABLE mst_saldo(
 id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
 user_id UUID NOT NULL UNIQUE,
 saldo BIGINT   NOT NULL DEFAULT 0,
 pin VARCHAR(6) NOT NULL,
 FOREIGN KEY(user_id) REFERENCES mst_user(id)
);

CREATE TABLE mst_admin(
 id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
 name VARCHAR(50) NOT NULL,
 username VARCHAR(100) NOT NULL,
 password VARCHAR(100) NOT NULL,
 role VARCHAR(5) NOT NULL DEFAULT 'admin',
 email VARCHAR(100) NOT NULL,
 created_at TIMESTAMP NOT NULL,
 updated_at TIMESTAMP NOT NULL
);

CREATE TABLE trx_send_transfer(
 id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
 user_id UUID NOT NULL,
 tujuan_transfer UUID NOT NULL,
 jumlah_transfer BIGINT NOT NULL,
 jenis_transfer VARCHAR(100) NOT NULL,
 transfer_at VARCHAR(100) NOT NULL,
 FOREIGN KEY(tujuan_transfer) REFERENCES mst_user(id),
 FOREIGN KEY(user_id) REFERENCES mst_user(id)
);
CREATE TABLE trx_receive_transfer(
 id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
 trx_id UUID NOT NULL,
 user_id UUID NOT NULL,
 tujuan_transfer UUID NOT NULL,
 jumlah_transfer BIGINT NOT NULL,
 jenis_transfer VARCHAR(100) NOT NULL,
 transfer_at VARCHAR(100) NOT NULL,
 FOREIGN KEY(tujuan_transfer) REFERENCES mst_user(id),
 FOREIGN KEY(user_id) REFERENCES mst_user(id),
 FOREIGN KEY(trx_id) REFERENCES trx_send_transfer(id)
);

CREATE TABLE trx_topup_method_payment(
 id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
 user_id UUID NOT NULL,
 token_midtrans VARCHAR(255),
 ammount BIGINT NOT NULL,
 deskripsi VARCHAR(250) ,
 status VARCHAR(100) NOT NULL,
 url_payment VARCHAR(255),
 created_at TIMESTAMP,
 updated_at TIMESTAMP,
 FOREIGN KEY(user_id) REFERENCES mst_user(id)
);
