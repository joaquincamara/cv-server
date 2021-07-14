package postgres

import (
	"context"
	"log"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joaquincamara/cv-server/internal/experience"
)

type experienceRepository struct {
	ExperienceQuerys experience.Repository
	Pool             *pgxpool.Pool
}

func NewExperienceRepository(pool *pgxpool.Pool) experience.Repository {
	return &experienceRepository{Pool: pool}
}

func (e *experienceRepository) Add(experience *experience.Experience) error {
	sqlStatement := `INSERT INTO experience (jobTitle, date) VALUES ($1, $2)`
	_, err := e.Pool.Exec(context.Background(), sqlStatement, experience.Jobtitle, experience.Date)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (e *experienceRepository) Delete(id int) error {
	sqlStatement := `DELETE FROM experience WHERE id=($1)`
	_, err := e.Pool.Exec(context.Background(), sqlStatement, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (d *experienceRepository) FindAll() ([]*experience.Experience, error) {
	result := make([]*experience.Experience, 0)
	sqlStatement := `SELECT * FROM experience`
	err := pgxscan.Select(context.Background(), d.Pool, &result, sqlStatement)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (d *experienceRepository) Update(experience *experience.Experience) error {
	sqlStatement := `UPDATE experience SET jobtitle=($1), date=($2)  WHERE id=($3)`
	_, err := d.Pool.Exec(context.Background(), sqlStatement, experience.Jobtitle, experience.Date, experience.Id)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
