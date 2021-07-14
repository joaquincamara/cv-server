package postgres

import (
	"context"
	"log"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joaquincamara/cv-server/internal/coolfeatures"
)

type coolFeaturesRepository struct {
	CoolFeaturesQuerys coolfeatures.Repository
	Pool               *pgxpool.Pool
}

func NewCoolFeaturesRepository(pool *pgxpool.Pool) coolfeatures.Repository {
	return &coolFeaturesRepository{Pool: pool}
}

func (c *coolFeaturesRepository) Add(coolfeatures *coolfeatures.Coolfeatures) error {
	sqlStatement := `INSERT INTO coolfeatures (title, info) VALUES ($1, $2)`
	_, err := c.Pool.Exec(context.Background(), sqlStatement, coolfeatures.Feature)
	if err != nil {
		panic(err)
	}
	return nil
}

func (c *coolFeaturesRepository) Delete(id int) error {

	sqlStatement := `DELETE FROM coolfeatures WHERE id=($1)`
	_, err := c.Pool.Exec(context.Background(), sqlStatement, id)
	if err != nil {
		panic(err)
	}
	return nil
}

func (c *coolFeaturesRepository) FindAll() ([]*coolfeatures.Coolfeatures, error) {
	result := make([]*coolfeatures.Coolfeatures, 0)
	sqlStatement := `SELECT * FROM coolfeatures`
	err := pgxscan.Select(context.Background(), c.Pool, &result, sqlStatement)

	if err != nil {
		panic(err)
	}

	return result, nil
}

func (c *coolFeaturesRepository) Update(coolfeatures *coolfeatures.Coolfeatures) error {
	sqlStatement := `UPDATE coolfeatures SET feature=($1), )  WHERE id=($2)`
	_, err := c.Pool.Exec(context.Background(), sqlStatement, coolfeatures.Feature, coolfeatures.Id)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
