package siswa

type Repository interface {
	Post(*Siswa) (string, error)
	Put(string, *Siswa) (int, error)
	GetOne(string) (Siswa, error)
	GetAll() ([]Siswa, error)
	Delete(nis string) (int, error)
}
