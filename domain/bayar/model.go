package bayar

type Bayar struct {
	Nis     string `json:"nis"`
	IdBayar int64  `json:"id_bayar"`
	Tanggal string `json:"tanggal"`
	Nominal int64  `json:"nominal"`
}
