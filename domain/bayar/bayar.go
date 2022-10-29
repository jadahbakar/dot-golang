package bayar

import "time"

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) PostBayar(bayar *Bayar) (string, error) {
	bayar.Tanggal = time.Now().Format(time.RFC3339)
	if bayar.IdBayar == 0 {
		bayar.IdBayar = time.Now().Unix()
	}
	res, err := s.repo.PostBayar(bayar)
	if err != nil {
		return "", err
	}
	return res, nil
}

func (s *service) UpdateBayar(bayar *Bayar) (int, error) {
	bayar.Tanggal = time.Now().Format(time.RFC3339)
	res, err := s.repo.PutBayar(bayar)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (s *service) FindById(nis string) (BayarGet, error) {
	res, err := s.repo.GetOneBayar(nis)
	if err != nil {
		return BayarGet{}, err
	}
	return res, nil
}

func (s *service) FindAllBayar() ([]BayarGet, error) {
	res, err := s.repo.GetAllBayar()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *service) DeleteBayar(nis string, idbayar int64) (int, error) {
	res, err := s.repo.DeleteBayar(nis, idbayar)
	if err != nil {
		return 0, err
	}
	return res, nil
}
