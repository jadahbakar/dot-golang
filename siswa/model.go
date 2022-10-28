package siswa

type Siswa struct {
	Nis  string `json:"nis"`
	Nama string `json:"nama"`
}

type Bayar struct {
	Nis     string `json:"nis"`
	IdBayar int64  `json:"id_bayar"`
	Tanggal int64  `json:"tanggal"`
	Nominal int64  `json:"nominal"`
}
