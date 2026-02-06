package configs

import (
	"context"
	"time"

	"github.com/Toppira-Official/backend/internal/domain/entities"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB(lc fx.Lifecycle, envs Environments, log *zap.Logger) *gorm.DB {
	var sqliteFileName string

	switch envs.MODE.String() {
	case "production":
		sqliteFileName = "production.db"
	default:
		sqliteFileName = "dev.db"
	}

	gormLogger := logger.New(
		zap.NewStdLog(log),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Error,
			Colorful:      false,
		},
	)

	db, err := gorm.Open(sqlite.Open(sqliteFileName+"?_journal_mode=WAL&_foreign_keys=ON&_busy_timeout=5000"),
		&gorm.Config{
			Logger:                                   gormLogger,
			PrepareStmt:                              true,
			DisableForeignKeyConstraintWhenMigrating: true,
		})
	if err != nil {
		log.Fatal("failed to connect to db", zap.Error(err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("failed to get sql db", zap.Error(err))
	}

	sqlDB.SetMaxOpenConns(1)
	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetConnMaxLifetime(time.Hour)

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				log.Info("running database migrations")
				return db.AutoMigrate(
					&entities.User{},
					&entities.Reminder{},
				)
			},
			OnStop: func(ctx context.Context) error {
				log.Info("closing database connection")
				return sqlDB.Close()
			},
		},
	)

	return db
}
