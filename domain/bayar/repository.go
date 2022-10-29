package bayar

type Repository interface {
	PostBayar(*Bayar) (string, error)
	PutBayar(*Bayar) (int, error)
	GetOneBayar(string) (BayarGet, error)
	GetAllBayar() ([]BayarGet, error)
	DeleteBayar(string, int64) (int, error)
}
