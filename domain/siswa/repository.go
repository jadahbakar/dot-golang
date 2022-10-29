package siswa

type Repository interface {
	PostSiswa(*Siswa) (string, error)
	PutSiswa(string, *Siswa) (int, error)
	GetOneSiswa(string) (Siswa, error)
	GetAllSiswa() ([]Siswa, error)
	DeleteSiswa(nis string) (int, error)
}
