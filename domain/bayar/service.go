package bayar

type Service interface {
	PostBayar(*Bayar) (string, error)
	UpdateBayar(*Bayar) (int, error)
	FindById(string) (Bayar, error)
	FindAllBayar() ([]Bayar, error)
	DeleteBayar(string, int64) (int, error)
}
