package services

import (
	"context"
	"dgut-icourse/config"
	"dgut-icourse/ent"
	"dgut-icourse/ent/migrate"
	_ "dgut-icourse/ent/runtime"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// Client 全局数据库连接
var Client = InitEntClient()

// InitEntClient
// @Description: 初始化数据库连接
// @return *ent.Client
func InitEntClient() *ent.Client {
	// Create ent.Client and run the schema migration.
	dsn := config.Config.Sql.User + ":" +
		config.Config.Sql.Passwd + "@tcp(" +
		config.Config.Sql.Host + ":" +
		config.Config.Sql.Port + ")/" +
		config.Config.Sql.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		log.Fatal("opening ent client", err)
	}
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}
	return client
}

func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			_ = tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
