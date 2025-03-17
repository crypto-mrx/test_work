package config

import (
	"fmt"

	"myapp-2/internal/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Host     string `env:"DB_HOST"`
	Port     uint16 `env:"DB_PORT"`
	User     string `env:"DB_USER"`
	Name     string `env:"DB_NAME"`
	Password string `env:"DB_PASSWORD"`
	SSLMode  string `env:"DB_SSLMODE"`
}

func buildDSN(cfg *DB) string {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s  dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name,
	)

	return dsn
}
func NewPostgresDB(cfg *DB) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(buildDSN(cfg)), &gorm.Config{
		TranslateError: true,
	})
}

func MustNewPostgresDB(cfg *DB) *gorm.DB {
	log := logger.GetLogger()
	db, err := NewPostgresDB(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database")
	}
	return db
}
