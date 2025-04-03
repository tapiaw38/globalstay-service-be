package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tapiaw38/globalstay-service-be/internal/platform/config"
)

func initConfig() error {
	configService, err := readConfig()
	if err != nil {
		return err
	}
	config.InitConfigService(configService)
	return nil
}

func readConfig() (*config.ConfigurationService, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error load env file")
	}

	configService := &config.ConfigurationService{
		ServerConfig: config.ServerConfig{
			GinMode:   config.GinModeServer(getEnv("GIN_MODE", "release")),
			Port:      getEnv("PORT", "8080"),
			Host:      getEnv("HOST", "localhost"),
			JWTSecret: getEnv("JWT_SECRET", "secret"),
		},
		NoSQLConfig: config.NoSQLConfig{
			Migrations: &config.NoSQLCollectionConfig{
				DatabaseURI: getEnv("DATABASE_URI", "mongodb://localhost:27017"),
				Database:    getEnv("DB_NAME", "hotel-reservation-db"),
				Collection:  getEnv("DB_COLLECTION_MIGRATIONS", "migrations"),
			},
			Hotels: &config.NoSQLCollectionConfig{
				DatabaseURI: getEnv("DATABASE_URI", "mongodb://localhost:27017"),
				Database:    getEnv("DB_NAME", "hotel-reservation-db"),
				Collection:  getEnv("DB_COLLECTION_HOTELS", "hotels"),
			},
			Services: &config.NoSQLCollectionConfig{
				DatabaseURI: getEnv("DATABASE_URI", "mongodb://localhost:27017"),
				Database:    getEnv("DB_NAME", "hotel-reservation-db"),
				Collection:  getEnv("DB_COLLECTION_SERVICES", "services"),
			},
			Reservations: &config.NoSQLCollectionConfig{
				DatabaseURI: getEnv("DATABASE_URI", "mongodb://localhost:27017"),
				Database:    getEnv("DB_NAME", "hotel-reservation-db"),
				Collection:  getEnv("DB_COLLECTION_RESERVATIONS", "reservations"),
			},
		},
		GCP: config.GCPConfig{
			MapsKey: getEnv("GOOGLE_MAPS_KEY", "AIzaSyD-0x9a6-8-9"),
		},
	}

	return configService, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
