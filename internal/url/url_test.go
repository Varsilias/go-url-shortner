package url

import "testing"

func TestShorten(t *testing.T) {

	tests := []struct {
		name           string
		originalUrl    string
		expectedResult string
	}{
		{
			name:           "",
			originalUrl:    "twitter.com",
			expectedResult: "",
		},
		{
			name:           "",
			originalUrl:    "https://github.com",
			expectedResult: "",
		},
		{
			name:           "",
			originalUrl:    "https://danielokoronkwo.tech",
			expectedResult: "",
		},
		{
			name:           "",
			originalUrl:    "https://danielokoronkwo.me",
			expectedResult: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Shorten(tt.originalUrl); got != tt.expectedResult {
				t.Errorf("Shorten() = %v, want %v", got, tt.expectedResult)
			}
		})
	}
}
