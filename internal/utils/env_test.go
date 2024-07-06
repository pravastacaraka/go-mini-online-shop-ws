package utils

import "testing"

type MockEnv struct {
	envs map[string]string
}

func (m MockEnv) Getenv(key string) string {
	return m.envs[key]
}

func TestIsProduction(t *testing.T) {
	tests := []struct {
		name     string
		envs     map[string]string
		expected bool
	}{
		{
			name:     "positive test case: production",
			envs:     map[string]string{"APP_ENV": "production"},
			expected: true,
		},
		{
			name:     "positive test case: development",
			envs:     map[string]string{"APP_ENV": "development"},
			expected: false,
		},
		{
			name:     "positive test case: not set",
			envs:     map[string]string{},
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setEnv(MockEnv{envs: tt.envs})

			if got := IsProduction(); got != tt.expected {
				t.Errorf("IsProduction() = %v, want %v", got, tt.expected)
			}
		})
	}
}
