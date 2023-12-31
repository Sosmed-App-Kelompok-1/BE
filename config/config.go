package config

import (
	"os"
	"reflect"
	"strconv"

	"github.com/joho/godotenv"
)

type Database struct {
	Host     string
	Port     uint16
	Username string
	Password string
	Database string
}

func (cfg *Database) LoadFromEnv(file ...string) error {
	if host, ok := os.LookupEnv("DB_HOST"); ok {
		cfg.Host = host
	}

	if port, ok := os.LookupEnv("DB_PORT"); ok {
		if cnv, err := strconv.Atoi(port); err != nil {
			return err
		} else {
			cfg.Port = uint16(cnv)
		}
	}

	if username, ok := os.LookupEnv("DB_USERNAME"); ok {
		cfg.Username = username
	}

	if pasword, ok := os.LookupEnv("DB_PASSWORD"); ok {
		cfg.Password = pasword
	}

	if database, ok := os.LookupEnv("DB_DATABASE"); ok {
		cfg.Database = database
	}

	if reflect.ValueOf(*cfg).IsZero() {
		if err := godotenv.Load(file...); err != nil {
			return err
		}

		return cfg.LoadFromEnv(file...)
	}

	return nil
}

type Cloudinary struct {
	CloudName string
	ApiKey    string
	ApiSecret string
}

func (cfg *Cloudinary) LoadFromEnv(file ...string) error {
	if name, ok := os.LookupEnv("CLOUDINARY_NAME"); ok {
		cfg.CloudName = name
	}

	if key, ok := os.LookupEnv("CLOUDINARY_KEY"); ok {
		cfg.ApiKey = key
	}

	if secret, ok := os.LookupEnv("CLOUDINARY_SECRET"); ok {
		cfg.ApiSecret = secret
	}

	if reflect.ValueOf(*cfg).IsZero() {
		if err := godotenv.Load(file...); err != nil {
			return err
		}

		return cfg.LoadFromEnv(file...)
	}

	return nil
}
