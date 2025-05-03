package config

import "os"

type Config struct {
	DBPath string
}

func Load() Config {
	dbPath := os.Getenv("BOOK_DB_PATH")
	if dbPath == "" {
		dbPath = "./books.db"
	}
	return Config{DBPath: dbPath}
}
