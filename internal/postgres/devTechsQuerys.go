package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joaquincamara/cv-server/internal/devTechs"
)

type devTechsRepository struct {
	DevTechQuerys devTechs.Repository
	Pool          *pgxpool.Pool
}

func NewDevTechsRepository(pool *pgxpool.Pool) devTechs.Repository {
	return &devTechsRepository{Pool: pool}
}

func (d *devTechsRepository) Add(devTech *devTechs.DevTech) error {
	sqlStatement := `INSERT INTO devtech (name, rank) VALUES ($1, $2)`
	_, err := d.Pool.Exec(context.Background(), sqlStatement, devTech.Name, devTech.Rank)
	if err != nil {
		panic(err)
	}
	return nil
}

func (d *devTechsRepository) Delete(id int) error {

	sqlStatement := `DELETE FROM devtech WHERE id=($1)`
	_, err := d.Pool.Exec(context.Background(), sqlStatement, id)
	if err != nil {
		panic(err)
	}
	return nil
}

func (d *devTechsRepository) FindAll() ([]*devTechs.DevTech, error) {
	result := make([]*devTechs.DevTech, 0)
	sqlStatement := `SELECT * FROM devtech`
	rows, err := d.Pool.Query(context.Background(), sqlStatement)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var tn *devTechs.DevTech
		if err := rows.Scan(&tn); err != nil {
			log.Fatal(err)
		}
		result = append(result, tn)
	}

	return result, nil
}
