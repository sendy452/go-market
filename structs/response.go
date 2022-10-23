package structs

import "time"

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResponseKategori struct {
	Status string     `json:"status"`
	Data   []Kategori `json:"data"`
}

type Kategori struct {
	IdKategori int       `json:"id_kategori"`
	Kategori   string    `json:"kategori"`
	Status     int       `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}

type ResponseProduk struct {
	Status string   `json:"status"`
	Data   []Produk `json:"data"`
}

type Produk struct {
	IdProduk     *int      `json:"id_produk"`
	NamaBarang   string    `json:"nama_barang"`
	Deskripsi    string    `json:"deskripsi"`
	Kategori     string    `json:"kategori"`
	GambarBarang string    `json:"gambar_barang"`
	Stok         int       `json:"stok"`
	Satuan       string    `json:"satuan"`
	Harga        int       `json:"harga"`
	Diskon       int       `json:"diskon"`
	Status       int       `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
}
