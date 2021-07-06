package postgres

import (
	"context"

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

func (d *devTechsRepository) Delete(devTech *devTechs.DevTech) error {
	sqlStatement := `INSERT INTO devtech (name, rank) VALUES ($1, $2)`
	_, err := d.Pool.Exec(context.Background(), sqlStatement, devTech.Name, devTech.Rank)
	if err != nil {
		panic(err)
	}
	return nil
}
