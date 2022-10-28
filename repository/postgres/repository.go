package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jadahbakar/dot-golang/siswa"
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

func NewRepository(ctx context.Context, url string) (siswa.Repository, error) {
	dbClient, err := newPostgresClient(context.Background(), url)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	repo := &pgRepository{db: dbClient}
	return repo, nil
}

func (r *pgRepository) Post(s *siswa.Siswa) (string, error) {
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

func (r *pgRepository) Put(nis string, s *siswa.Siswa) (int, error) {
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

func (r *pgRepository) GetOne(nis string) (siswa.Siswa, error) {
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

func (r *pgRepository) GetAll() ([]siswa.Siswa, error) {
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

func (r *pgRepository) Delete(nis string) (int, error) {
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
