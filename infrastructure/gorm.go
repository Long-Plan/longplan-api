package infrastructure

import (
	"context"
	"fmt"
	"time"

	"github.com/Long-Plan/longplan-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type SQLlogger struct {
	logger.Interface
}

func (l SQLlogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n==============================\n", sql)
}

func InitDB() {
	var err error
	cfg := config.Config.Database

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable TimeZone=Asia/Bangkok",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Pass,
		cfg.Name,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		// DryRun: true,
		// Logger: &SQLlogger{},
	})
	if err != nil {
		panic(err)
	}
}
