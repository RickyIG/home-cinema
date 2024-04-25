-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS roles(
    id_role SERIAL PRIMARY KEY,
    role_name VARCHAR(255) UNIQUE NOT NULL,
    role_desc TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS users (
    id_user SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    phone_number VARCHAR(15),
    balance INTEGER NOT NULL DEFAULT 0,
    id_role INT REFERENCES roles(id_role),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE films (
    id_film SERIAL PRIMARY KEY,
    judul_film VARCHAR(255) NOT NULL,
    genre VARCHAR(50) NOT NULL,
    sinopsis TEXT,
    durasi INTEGER NOT NULL,
    rating VARCHAR(10) NOT NULL,
    image_url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE studios (
    id_studio SERIAL PRIMARY KEY,
    nama_studio VARCHAR(50) NOT NULL,
    jumlah_kursi INTEGER NOT NULL,
    tipe_studio VARCHAR(10) NOT NULL
);

CREATE TABLE kursis (
    id_kursi SERIAL PRIMARY KEY,
    id_studio INTEGER NOT NULL,
    nomor_kursi VARCHAR(10) NOT NULL,
    status VARCHAR(10) NOT NULL DEFAULT 'tersedia',
    FOREIGN KEY (id_studio) REFERENCES studios(id_studio)
);

CREATE TABLE jadwals (
  id_jadwal VARCHAR(50) PRIMARY KEY,
  id_film INTEGER NOT NULL,
  id_studio INTEGER NOT NULL,
  tanggal_tayang DATE NOT NULL,
  jam_tayang TIME NOT NULL,
  harga_tiket INTEGER NOT NULL,
  FOREIGN KEY (id_film) REFERENCES films(id_film),
  FOREIGN KEY (id_studio) REFERENCES studios(id_studio)
);

CREATE TABLE transaksis (
  id_transaksi SERIAL PRIMARY KEY,
  id_user INTEGER NOT NULL,
  id_jadwal VARCHAR(50) NOT NULL,
  total_bayar INTEGER NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (id_user) REFERENCES users(id_user),
  FOREIGN KEY (id_jadwal) REFERENCES jadwals(id_jadwal)
);

CREATE TABLE tickets (
  id_ticket SERIAL PRIMARY KEY,
  id_jadwal VARCHAR(50) NOT NULL,
  id_kursi INTEGER NOT NULL,
  id_user INTEGER NOT NULL,
  id_transaksi INTEGER NOT NULL,
  ticket_status VARCHAR(20) NOT NULL DEFAULT 'belum_dibayar',
  FOREIGN KEY (id_user) REFERENCES users(id_user),
  FOREIGN KEY (id_jadwal) REFERENCES jadwals(id_jadwal),
  FOREIGN KEY (id_kursi) REFERENCES kursis(id_kursi),
  FOREIGN KEY (id_transaksi) REFERENCES transaksis(id_transaksi)
);

INSERT INTO roles (role_name, role_desc) VALUES ('user', 'Normal user.') ON CONFLICT DO NOTHING;
INSERT INTO roles (role_name, role_desc) VALUES ('admin', 'Admin.') ON CONFLICT DO NOTHING;
INSERT INTO roles (role_name, role_desc) VALUES ('fintech', 'Fintech.') ON CONFLICT DO NOTHING;

-- +migrate StatementEnd