package postgres

import (
	"context"
	"log"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joaquincamara/cv-server/internal/coolFeatures"
)

type coolFeaturesRepository struct {
	CoolFeaturesQuerys coolFeatures.Repository
	Pool               *pgxpool.Pool
}

func NewCoolFeaturesRepository(pool *pgxpool.Pool) coolFeatures.Repository {
	return &coolFeaturesRepository{Pool: pool}
}

func (c *coolFeaturesRepository) Add(coolfeatures *coolFeatures.Coolfeatures) error {
	sqlStatement := `INSERT INTO coolfeatures (feature) VALUES ($1)`
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

func (c *coolFeaturesRepository) FindAll() ([]*coolFeatures.Coolfeatures, error) {
	result := make([]*coolFeatures.Coolfeatures, 0)
	sqlStatement := `SELECT * FROM coolfeatures`
	err := pgxscan.Select(context.Background(), c.Pool, &result, sqlStatement)

	if err != nil {
		panic(err)
	}

	return result, nil
}

func (c *coolFeaturesRepository) Update(coolfeatures *coolFeatures.Coolfeatures) error {
	sqlStatement := `UPDATE coolfeatures SET feature=($1)  WHERE id=($2)`
	_, err := c.Pool.Exec(context.Background(), sqlStatement, coolfeatures.Feature, coolfeatures.Id)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
