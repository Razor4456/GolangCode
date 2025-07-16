CREATE TABLE IF NOT EXISTS stuff(
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    nama_barang varchar (200) NOT NULL,
    jumlah_barang INT NOT NULL,
    harga INT NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);