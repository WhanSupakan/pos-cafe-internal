package postgres

import (
	"cafe-pos/internal/db/models"
	"context"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
)

func NewEntClient(dsn string) (*models.Client, error) {
	client, err := models.Open(dialect.Postgres, dsn)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func EntMigrate(client *models.Client) error {
	if err := client.Schema.Create(context.Background()); err != nil {
		return err
	}
	return nil
}
