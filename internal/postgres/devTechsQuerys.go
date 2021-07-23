package postgres

import (
	"context"
	"log"

	"github.com/georgysavva/scany/pgxscan"
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
		log.Println(err)
		return err
	}
	return nil
}

func (d *devTechsRepository) Delete(id int) error {
	sqlStatement := `DELETE FROM devtech WHERE id=($1)`
	_, err := d.Pool.Exec(context.Background(), sqlStatement, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (d *devTechsRepository) FindAll() ([]*devTechs.DevTech, error) {
	result := make([]*devTechs.DevTech, 0)
	sqlStatement := `SELECT * FROM devtech`
	err := pgxscan.Select(context.Background(), d.Pool, &result, sqlStatement)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (d *devTechsRepository) Update(devTech *devTechs.DevTech) error {
	log.Println(devTech)
	sqlStatement := `UPDATE devtech SET name=($1), rank=($2)  WHERE id=($3)`
	_, err := d.Pool.Exec(context.Background(), sqlStatement, devTech.Name, devTech.Rank, devTech.Id)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
