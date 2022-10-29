package bayar

type Service interface {
	PostBayar(*Bayar) (string, error)
	UpdateBayar(*Bayar) (int, error)
	FindById(string) (BayarGet, error)
	FindAllBayar() ([]BayarGet, error)
	DeleteBayar(string, int64) (int, error)
}
