package main

import (
	"testing"

	"github.com/zan8in/afrog/v3/pkg/config"
)

func TestShouldReportFingerprintHit(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		options  *config.Options
		severity string
		want     bool
	}{
		{
			name:     "nil options reports all",
			options:  nil,
			severity: "info",
			want:     true,
		},
		{
			name:     "empty filter reports all",
			options:  &config.Options{},
			severity: "info",
			want:     true,
		},
		{
			name:     "matching severity is reported",
			options:  &config.Options{Severity: "high"},
			severity: "high",
			want:     true,
		},
		{
			name:     "non matching severity is hidden",
			options:  &config.Options{Severity: "high"},
			severity: "info",
			want:     false,
		},
		{
			name:     "multiple severities are supported",
			options:  &config.Options{Severity: "high, critical"},
			severity: "critical",
			want:     true,
		},
		{
			name:     "matching is case insensitive",
			options:  &config.Options{Severity: "HiGh"},
			severity: "HIGH",
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := shouldReportFingerprintHit(tt.options, tt.severity); got != tt.want {
				t.Fatalf("shouldReportFingerprintHit(%q) = %v, want %v", tt.severity, got, tt.want)
			}
		})
	}
}
