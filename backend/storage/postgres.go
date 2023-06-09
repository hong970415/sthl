package storage

import (
	"context"
	"fmt"
	"sthl/config"
	"sthl/ent"
	"sthl/ent/migrate"
	"strings"

	"entgo.io/ent/dialect"
	"go.uber.org/zap"
)

func NewPostgresDb(l *zap.Logger, c *config.Config) (*ent.Client, error) {
	// db with ent
	dataSourceName := []string{
		fmt.Sprintf("host=%s", c.GetDbDomain()),
		fmt.Sprintf("port=%d", c.GetDbPort()),
		fmt.Sprintf("user=%s", c.GetDbUser()),
		fmt.Sprintf("password=%s", c.GetDbPw()),
		"sslmode=disable",
	}
	joinedDataSourceName := strings.Join(dataSourceName, " ")
	client, err := ent.Open(dialect.Postgres, joinedDataSourceName)
	if err != nil {
		l.Info("ent.Open err", zap.Error(err))
		return nil, err
	}

	// auto ent migrate
	l.Info("start ent auto migrating...")
	err = client.Schema.Create(context.TODO(),
		migrate.WithGlobalUniqueID(true),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		l.Info("ent auto migrate err", zap.Error(err))
		return nil, err
	}
	l.Info("ent auto migration finished")
	return client, nil
}
