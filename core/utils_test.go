package core

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateSurveyHashURL(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		customerName string
	}{
		{name: "typical name", customerName: "1004035530"},
		{name: "empty name", customerName: "1004035530"},
		{name: "special characters", customerName: "1004035530"},
		{name: "unicode", customerName: "1004035530"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := GenerateSurveyHashURL(tt.customerName)
			t.Logf("survey hash url: customer=%q input=%q url=%s", tt.customerName, "?hash_id=MB"+tt.customerName, got)

			wantInput := "?hash_id=MB" + tt.customerName
			sum := sha256.Sum256([]byte(wantInput))
			want := hex.EncodeToString(sum[:])

			require.NotEmpty(t, got)
			assert.Equal(t, want, got)
			assert.Len(t, got, 64) // SHA-256 hex digest
		})
	}
}

func TestGenerateSurveyHashURL_Deterministic(t *testing.T) {
	t.Parallel()

	name := "Jane Smith"
	first := GenerateSurveyHashURL(name)
	second := GenerateSurveyHashURL(name)
	t.Logf("survey hash url: customer=%q url=%s", name, first)

	assert.Equal(t, first, second)
}

func TestGenerateSurveyHashURL_DifferentCustomersDiffer(t *testing.T) {
	t.Parallel()

	a := GenerateSurveyHashURL("Alice")
	b := GenerateSurveyHashURL("Bob")
	t.Logf("survey hash url: Alice=%s Bob=%s", a, b)

	assert.NotEqual(t, a, b)
}

func TestGenerateSHA256(t *testing.T) {
	t.Parallel()

	input := "?hash_id=MBAlice"
	got := GenerateSHA256(input)
	t.Logf("survey hash url: input=%q url=%s", input, got)

	sum := sha256.Sum256([]byte(input))
	want := hex.EncodeToString(sum[:])

	assert.Equal(t, want, got)
}
