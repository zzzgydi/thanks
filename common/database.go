package common

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
	slogGorm "github.com/zzzgydi/slog-gorm"
	"github.com/zzzgydi/thanks/common/initializer"
	L "github.com/zzzgydi/thanks/common/logger"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var MDB *gorm.DB

func initDatabase() error {
	dsn := viper.GetString("DATABASE_DSN")
	if dsn == "" {
		return fmt.Errorf("database dsn error")
	}

	gormLogger := slogGorm.New(
		slogGorm.WithHandler(L.Logger.Handler()),
		slogGorm.WithParameterizedQueries(true),
		slogGorm.WithSlowThreshold(200*time.Millisecond),
	)

	var dialector gorm.Dialector

	if strings.HasPrefix(dsn, "mysql://") {
		dialector = mysql.Open(dsn)
	} else if strings.HasPrefix(dsn, "postgres://") {
		dialector = postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true,
		})
	} else if strings.HasPrefix(dsn, "sqlite://") {
		dsn = strings.TrimPrefix(dsn, "sqlite://")
		dialector = sqlite.Open(dsn)
	} else {
		return fmt.Errorf("unknown database dsn: %s", dsn)
	}

	db, err := gorm.Open(dialector, &gorm.Config{Logger: gormLogger, PrepareStmt: true})
	if err != nil {
		return err
	}

	MDB = db
	return nil
}

func init() {
	initializer.Register("database", initDatabase)
}
