package bayar

type Bayar struct {
	Nis     string `json:"nis"`
	IdBayar int64  `json:"id_bayar"`
	Tanggal string `json:"tanggal"`
	Nominal int64  `json:"nominal"`
}

type BayarGet struct {
	Nis     string `json:"nis"`
	Nama    string `json:"nama"`
	IdBayar int64  `json:"id_bayar"`
	Tanggal string `json:"tanggal"`
	Nominal int64  `json:"nominal"`
}
