package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joaquincamara/cv-server/internal/aboutMe"
)

type aboutMeRepository struct {
	DevTechQuerys aboutMe.Repository
	Pool          *pgxpool.Pool
}

func NewAboutMeRepository(pool *pgxpool.Pool) aboutMe.Repository {
	return &aboutMeRepository{Pool: pool}
}

func (d *aboutMeRepository) Add(aboutMe *aboutMe.AboutMe) error {
	sqlStatement := `INSERT INTO devtech (name, rank) VALUES ($1, $2)`
	_, err := d.Pool.Exec(context.Background(), sqlStatement, aboutMe.Title, aboutMe.Info)
	if err != nil {
		panic(err)
	}
	return nil
}

func (d *aboutMeRepository) Delete(id int) error {

	sqlStatement := `DELETE FROM devtech WHERE id=($1)`
	_, err := d.Pool.Exec(context.Background(), sqlStatement, id)
	if err != nil {
		panic(err)
	}
	return nil
}

func (d *aboutMeRepository) FindAll() ([]*aboutMe.AboutMe, error) {
	result := make([]*aboutMe.AboutMe, 0)
	sqlStatement := `SELECT * FROM devtech`
	rows, err := d.Pool.Query(context.Background(), sqlStatement)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var tn *aboutMe.AboutMe
		if err := rows.Scan(&tn); err != nil {
			log.Fatal(err)
		}
		result = append(result, tn)
	}

	return result, nil
}
