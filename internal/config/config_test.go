package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	// Save original env and restore after test
	originalPort := os.Getenv("PORT")
	originalHost := os.Getenv("HOST")
	defer func() {
		os.Setenv("PORT", originalPort)
		os.Setenv("HOST", originalHost)
	}()

	tests := []struct {
		name        string
		envPort     string
		envHost     string
		wantPort    int
		wantHost    string
		wantErr     bool
	}{
		{
			name:     "default values",
			envPort:  "",
			envHost:  "",
			wantPort: 8080,
			wantHost: "0.0.0.0",
			wantErr:  false,
		},
		{
			name:     "custom port",
			envPort:  "3000",
			envHost:  "",
			wantPort: 3000,
			wantHost: "0.0.0.0",
			wantErr:  false,
		},
		{
			name:     "custom host and port",
			envPort:  "9000",
			envHost:  "127.0.0.1",
			wantPort: 9000,
			wantHost: "127.0.0.1",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("PORT", tt.envPort)
			os.Setenv("HOST", tt.envHost)

			cfg, err := Load()
			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if cfg.Port != tt.wantPort {
				t.Errorf("Load() Port = %v, want %v", cfg.Port, tt.wantPort)
			}

			if cfg.Host != tt.wantHost {
				t.Errorf("Load() Host = %v, want %v", cfg.Host, tt.wantHost)
			}
		})
	}
}

func TestConfig_ServerAddress(t *testing.T) {
	cfg := &Config{
		Host: "localhost",
		Port: 8080,
	}

	expected := "localhost:8080"
	if got := cfg.ServerAddress(); got != expected {
		t.Errorf("ServerAddress() = %v, want %v", got, expected)
	}
}

func TestConfig_IsDevelopment(t *testing.T) {
	tests := []struct {
		environment string
		want        bool
	}{
		{"development", true},
		{"production", false},
		{"staging", false},
	}

	for _, tt := range tests {
		t.Run(tt.environment, func(t *testing.T) {
			cfg := &Config{Environment: tt.environment}
			if got := cfg.IsDevelopment(); got != tt.want {
				t.Errorf("IsDevelopment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_IsProduction(t *testing.T) {
	tests := []struct {
		environment string
		want        bool
	}{
		{"production", true},
		{"development", false},
		{"staging", false},
	}

	for _, tt := range tests {
		t.Run(tt.environment, func(t *testing.T) {
			cfg := &Config{Environment: tt.environment}
			if got := cfg.IsProduction(); got != tt.want {
				t.Errorf("IsProduction() = %v, want %v", got, tt.want)
			}
		})
	}
}
