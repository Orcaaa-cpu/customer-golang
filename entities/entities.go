package entities

type Customer struct {
	Id      int    `json:"ID"`
	Nama    string `json:"NAMA LENGKAP"`
	Alamat  string `json:"ALAMAT"`
	Telpon  string `json:"NOMOR HP"`
	Tanggal string `json:"TANGGAL"`
}
