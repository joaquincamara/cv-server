package postgres

import (
	"context"
	"log"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joaquincamara/cv-server/internal/personalProjects"
)

type personalProjectsRepository struct {
	PersonalProjectsQuerys personalProjects.Repository
	Pool                   *pgxpool.Pool
}

func NewPersonalProjectsRepository(pool *pgxpool.Pool) personalProjects.Repository {
	return &personalProjectsRepository{Pool: pool}
}

func (p *personalProjectsRepository) Add(personalprojects *personalProjects.PersonalProjects) error {
	sqlStatement := `INSERT INTO personalprojects (title, description, url) VALUES ($1, $2, $3)`
	_, err := p.Pool.Exec(context.Background(), sqlStatement, personalprojects.Title, personalprojects.Description, personalprojects.Url)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (p *personalProjectsRepository) Delete(id int) error {
	sqlStatement := `DELETE FROM personalprojects WHERE id=($1)`
	_, err := p.Pool.Exec(context.Background(), sqlStatement, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (p *personalProjectsRepository) FindAll() ([]*personalProjects.PersonalProjects, error) {
	result := make([]*personalProjects.PersonalProjects, 0)
	sqlStatement := `SELECT * FROM personalprojects`
	err := pgxscan.Select(context.Background(), p.Pool, &result, sqlStatement)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (p *personalProjectsRepository) Update(personalprojects *personalProjects.PersonalProjects) error {
	sqlStatement := `UPDATE personalprojects SET title=($1), description=($2), url=($3)  WHERE id=($4)`
	_, err := p.Pool.Exec(context.Background(), sqlStatement, personalprojects.Title, personalprojects.Description, personalprojects.Url, personalprojects.Id)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
