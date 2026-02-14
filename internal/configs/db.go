package configs

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB(lc fx.Lifecycle, envs Environments, log *zap.Logger) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		envs.POSTGRES_HOST.String(),
		envs.POSTGRES_USER.String(),
		envs.POSTGRES_PASSWORD.String(),
		envs.POSTGRES_DB.String(),
		envs.POSTGRES_PORT.String(),
	)

	gormLogger := logger.New(
		zap.NewStdLog(log),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      false,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn),
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

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(time.Hour)

	lc.Append(
		fx.Hook{
			OnStop: func(ctx context.Context) error {
				log.Info("closing database connection")
				return sqlDB.Close()
			},
		},
	)

	return db
}
