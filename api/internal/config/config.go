package config

import (
	"os"
	"strings"
)

type Config struct {
	DatabaseURL string
	SupabaseJWT string
	CORSOrigins []string
}

// Load reads configuration from environment variables.
func Load() (*Config, error) {
	origins := strings.Split(strings.TrimSpace(os.Getenv("CORS_ALLOW_ORIGINS")), ",")
	if len(origins) == 1 && origins[0] == "" {
		origins = []string{"*"}
	}
	return &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		SupabaseJWT: os.Getenv("SUPABASE_JWT_SECRET"),
		CORSOrigins: origins,
	}, nil
}