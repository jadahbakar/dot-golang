package siswa

type Service interface {
	PostSiswa(*Siswa) (string, error)
	FindByNIS(string) (Siswa, error)
	FindAllSiswa() ([]Siswa, error)
	DeleteSiswa(string) (int, error)
	UpdateSiswa(string, *Siswa) (int, error)
}
