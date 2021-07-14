package postgres

import (
	"context"
	"log"

	"github.com/georgysavva/scany/pgxscan"
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

func (a *aboutMeRepository) Add(aboutMe *aboutMe.AboutMe) error {
	sqlStatement := `INSERT INTO aboutme (title, info) VALUES ($1, $2)`
	_, err := a.Pool.Exec(context.Background(), sqlStatement, aboutMe.Title, aboutMe.Info)
	if err != nil {
		panic(err)
	}
	return nil
}

func (a *aboutMeRepository) Delete(id int) error {

	sqlStatement := `DELETE FROM aboutme WHERE id=($1)`
	_, err := a.Pool.Exec(context.Background(), sqlStatement, id)
	if err != nil {
		panic(err)
	}
	return nil
}

func (a *aboutMeRepository) FindAll() ([]*aboutMe.AboutMe, error) {
	result := make([]*aboutMe.AboutMe, 0)
	sqlStatement := `SELECT * FROM aboutme`
	err := pgxscan.Select(context.Background(), a.Pool, &result, sqlStatement)

	if err != nil {
		panic(err)
	}

	return result, nil
}

func (a *aboutMeRepository) Update(aboutMe *aboutMe.AboutMe) error {
	sqlStatement := `UPDATE aboutme SET title=($1), info=($2)  WHERE id=($3)`
	_, err := a.Pool.Exec(context.Background(), sqlStatement, aboutMe.Title, aboutMe.Info, aboutMe.Id)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
