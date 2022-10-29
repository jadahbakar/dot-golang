package siswa

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) PostSiswa(siswa *Siswa) (string, error) {
	res, err := s.repo.PostSiswa(siswa)
	if err != nil {
		return "", err
	}
	return res, nil
}
func (s *service) FindByNIS(nis string) (Siswa, error) {
	res, err := s.repo.GetOneSiswa(nis)
	if err != nil {
		return Siswa{}, err
	}
	return res, nil
}

func (s *service) FindAllSiswa() ([]Siswa, error) {
	res, err := s.repo.GetAllSiswa()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *service) DeleteSiswa(nis string) (res int, err error) {
	res, err = s.repo.DeleteSiswa(nis)
	if err != nil {
		return 0, err
	}
	return res, nil
}
func (s *service) UpdateSiswa(nis string, siswa *Siswa) (int, error) {
	res, err := s.repo.PutSiswa(nis, siswa)
	if err != nil {
		return 0, err
	}
	return res, nil
}
