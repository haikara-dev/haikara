package database

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/haikara-dev/haikara/config"
	"github.com/haikara-dev/haikara/ent"
	"log"
	"time"
)

var Client *ent.Client

func dbConnect(dsn string, retryCount uint) (err error) {
	for retryCount > 1 {
		if Client, err = ent.Open("mysql", dsn); err != nil {
			time.Sleep(time.Second * 2)
			retryCount--
			log.Printf("retry... count:%v\n", retryCount)
			continue
		}
		break
	}
	return err
}

func dbMigration(retryCount uint) (err error) {
	for retryCount > 1 {
		if err = Client.Schema.Create(context.Background()); err != nil {
			time.Sleep(time.Second * 2)
			retryCount--
			log.Printf("retry... count:%v\n", retryCount)
			continue
		}
		break
	}
	return err
}

func Init() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.DbUser,
		config.Config.DbPassword,
		config.Config.DbHost,
		config.Config.DbName,
	)
	//なぜか、mysqlが起動していなくても errにならない
	if err = dbConnect(dsn, 100); err != nil {
		log.Fatalf("action: db Open, %s", err.Error())
	}

	if err = dbMigration(100); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func Close() {
	_ = Client.Close()
}
