package config

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

// Config menyimpan seluruh konfigurasi aplikasi dari environment variables.
type Config struct {
	AppEnv             string
	AppHost            string
	AppPort            string
	DBHost             string
	DBPort             string
	DBName             string
	DBUser             string
	DBPassword         string
	DBMaxOpenConns     int
	DBMaxIdleConns     int
	DBConnMaxLifetime  time.Duration
	CORSAllowedOrigins []string
}

// Load membaca file .env (jika ada di environment development) lalu membaca variabel environment.
func Load() *Config {
	loadDotEnv(".env")

	originsRaw := getEnv("CORS_ALLOWED_ORIGINS", "http://localhost:5173,http://127.0.0.1:5173,http://localhost:5174,http://127.0.0.1:5174,http://localhost:5175,http://127.0.0.1:5175,https://trusmi.riskiahmadiilham.my.id")
	var origins []string
	for _, o := range strings.Split(originsRaw, ",") {
		trimmed := strings.TrimSpace(o)
		if trimmed != "" {
			origins = append(origins, trimmed)
		}
	}

	return &Config{
		AppEnv:             getEnv("APP_ENV", "development"),
		AppHost:            getEnv("APP_HOST", "0.0.0.0"),
		AppPort:            getEnv("APP_PORT", "8090"),
		DBHost:             getEnv("DB_HOST", "127.0.0.1"),
		DBPort:             getEnv("DB_PORT", "3306"),
		DBName:             getEnv("DB_NAME", "db_test_programmer_trusmi"),
		DBUser:             getEnv("DB_USER", "root"),
		DBPassword:         getEnv("DB_PASSWORD", ""),
		DBMaxOpenConns:     getEnvInt("DB_MAX_OPEN_CONNS", 15),
		DBMaxIdleConns:     getEnvInt("DB_MAX_IDLE_CONNS", 5),
		DBConnMaxLifetime:  time.Duration(getEnvInt("DB_CONN_MAX_LIFETIME_MIN", 15)) * time.Minute,
		CORSAllowedOrigins: origins,
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	valStr, ok := os.LookupEnv(key)
	if !ok || strings.TrimSpace(valStr) == "" {
		return fallback
	}
	val, err := strconv.Atoi(strings.TrimSpace(valStr))
	if err != nil {
		return fallback
	}
	return val
}

// loadDotEnv membaca file .env secara ringan tanpa menimpa env yang sudah diset di OS/Container.
func loadDotEnv(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])
		val = strings.Trim(val, `"'`)
		if _, exists := os.LookupEnv(key); !exists {
			_ = os.Setenv(key, val)
		}
	}
}
