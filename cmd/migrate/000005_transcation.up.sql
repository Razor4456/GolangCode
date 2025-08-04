CREATE TABLE IF NOT EXISTS transactions(
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    userid INT NOT NULL,
    idbarang INT NOT NULL,
    nama_barang varchar (200) NOT NULL,
    jumlah_barang INT NOT NULL,
    harga INT NOT NULL
);