package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jadahbakar/dot-golang/domain/bayar"
	"github.com/jadahbakar/dot-golang/domain/siswa"
	"github.com/jadahbakar/dot-golang/util/logger"
)

type pgRepository struct {
	db *pgxpool.Pool
}

func newPostgresClient(ctx context.Context, url string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return pool, nil
}

func NewRepository(ctx context.Context, url string) (siswa.Repository, bayar.Repository, error) {
	dbClient, err := newPostgresClient(context.Background(), url)
	if err != nil {
		return nil, nil, errors.New(err.Error())
	}
	siswaRepo := &pgRepository{db: dbClient}
	bayarRepo := &pgRepository{db: dbClient}
	return siswaRepo, bayarRepo, nil
}

func (r *pgRepository) PostSiswa(s *siswa.Siswa) (string, error) {
	ctx := context.Background()
	query := fmt.Sprintf(`INSERT INTO mst.siswa(nis,nama) VALUES ('%s', '%s') RETURNING nis`, s.Nis, s.Nama)
	var id string
	err := r.db.QueryRow(ctx, query).Scan(&id)
	if err != nil {
		logger.Error(err)
		return "", err
	}
	return id, nil
}

func (r *pgRepository) PutSiswa(nis string, s *siswa.Siswa) (int, error) {
	ctx := context.Background()
	query := fmt.Sprintf(`UPDATE mst.siswa SET nama = '%s' WHERE nis = '%s'`, s.Nama, nis)
	res, err := r.db.Exec(ctx, query)
	if err != nil {
		return 0, err
	}
	if res.RowsAffected() != 1 {
		logger.Error(err)
		return 0, err
	}
	return int(res.RowsAffected()), nil
}

func (r *pgRepository) GetOneSiswa(nis string) (siswa.Siswa, error) {
	ctx := context.Background()
	var t siswa.Siswa
	query := fmt.Sprintf(`SELECT nis, nama FROM mst.siswa WHERE nis = '%s'`, nis)
	err := r.db.QueryRow(ctx, query).Scan(&t.Nis, &t.Nama)
	if err != nil {
		logger.Errorf("repo:%v", err)
		return siswa.Siswa{}, err
	}
	return t, nil
}

func (r *pgRepository) GetAllSiswa() ([]siswa.Siswa, error) {
	ctx := context.Background()
	result := make([]siswa.Siswa, 0)
	t := siswa.Siswa{}
	query := `SELECT nis, nama FROM mst.siswa`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		logger.Errorf("repo:%v", err)
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&t.Nis, &t.Nama)
		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}
	if rows.Err() != nil {
		logger.Errorf("Error Reading Rows: \n", err)
		return nil, rows.Err()
	}
	return result, nil
}

func (r *pgRepository) DeleteSiswa(nis string) (int, error) {
	query := fmt.Sprintf(`DELETE FROM mst.siswa WHERE nis = '%s'`, nis)
	ctx := context.Background()
	commandTag, err := r.db.Exec(ctx, query)
	if err != nil {
		logger.Error(err)
		return 0, err
	}
	if commandTag.RowsAffected() != 1 {
		logger.Error(err)
		return 0, err
	}
	return int(commandTag.RowsAffected()), nil
}

func (r *pgRepository) PostBayar(b *bayar.Bayar) (string, error) {
	ctx := context.Background()
	query := fmt.Sprintf(`INSERT INTO mst.bayar(nis,idbayar,tanggal,nominal) 
				VALUES ('%s', %d,'%s', %d) RETURNING nis`, b.Nis, b.IdBayar, b.Tanggal, b.Nominal)
	var id string
	err := r.db.QueryRow(ctx, query).Scan(&id)
	if err != nil {
		logger.Error(err)
		return "", err
	}
	return id, nil
}

func (r *pgRepository) PutBayar(b *bayar.Bayar) (int, error) {
	ctx := context.Background()
	query := fmt.Sprintf(`UPDATE mst.bayar SET tanggal = '%s', nominal= %d  WHERE nis = '%s' AND idbayar = %d`, b.Tanggal, b.Nominal, b.Nis, b.IdBayar)
	res, err := r.db.Exec(ctx, query)
	if err != nil {
		return 0, err
	}
	if res.RowsAffected() != 1 {
		logger.Error(err)
		return 0, err
	}
	return int(res.RowsAffected()), nil
}

func (r *pgRepository) GetOneBayar(nis string) (bayar.Bayar, error) {
	ctx := context.Background()
	var t bayar.Bayar
	query := fmt.Sprintf(`SELECT nis, idbayar, tanggal, nominal FROM mst.bayar WHERE nis = '%s'`, nis)
	err := r.db.QueryRow(ctx, query).Scan(&t.Nis, &t.IdBayar, &t.Tanggal, &t.Nominal)
	if err != nil {
		logger.Errorf("repo:%v", err)
		return bayar.Bayar{}, err
	}
	return t, nil
}

func (r *pgRepository) GetAllBayar() ([]bayar.Bayar, error) {
	ctx := context.Background()
	result := make([]bayar.Bayar, 0)
	t := bayar.Bayar{}
	query := `SELECT nis, idbayar, tanggal, nominal FROM mst.bayar`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		logger.Errorf("repo:%v", err)
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&t.Nis, &t.IdBayar, &t.Tanggal, &t.Nominal)
		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}
	if rows.Err() != nil {
		logger.Errorf("Error Reading Rows: \n", err)
		return nil, rows.Err()
	}
	return result, nil
}

func (r *pgRepository) DeleteBayar(nis string, idbayar int64) (int, error) {
	query := fmt.Sprintf(`DELETE FROM mst.bayar  WHERE nis = '%s' AND idbayar = %d`, nis, idbayar)
	ctx := context.Background()
	commandTag, err := r.db.Exec(ctx, query)
	if err != nil {
		logger.Error(err)
		return 0, err
	}
	if commandTag.RowsAffected() != 1 {
		logger.Error(err)
		return 0, err
	}
	return int(commandTag.RowsAffected()), nil
}
