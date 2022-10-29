package bayar

import (
	"time"

	"github.com/jadahbakar/dot-golang/domain/siswa"
)

type service struct {
	repo      Repository
	siswaRepo siswa.Repository
}

func NewService(r Repository, s siswa.Repository) Service {
	return &service{repo: r, siswaRepo: s}
}

func (s *service) PostBayar(bayar *Bayar) (string, error) {
	siswa, err := s.siswaRepo.GetOneSiswa(bayar.Nis)
	if err != nil {
		return "", err
	}
	bayar.Nis = siswa.Nis
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
