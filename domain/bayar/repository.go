package bayar

type Repository interface {
	PostBayar(*Bayar) (string, error)
	PutBayar(*Bayar) (int, error)
	GetOneBayar(string) (Bayar, error)
	GetAllBayar() ([]Bayar, error)
	DeleteBayar(string, int64) (int, error)
}
