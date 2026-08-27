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
		{name: "transaction one", customerName: "FT262137QPZ281154281"},
		{name: "transaction two", customerName: "FT26237HY8GJ95578843"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := GenerateSurveyHashURL("https://feedback.cbe.com.et/M/81c8db50", tt.customerName)
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
	first := GenerateSurveyHashURL("https://feedback.cbe.com.et/M/81c8db50", name)
	second := GenerateSurveyHashURL("https://feedback.cbe.com.et/M/81c8db50", name)
	t.Logf("survey hash url: customer=%q url=%s", name, first)

	assert.Equal(t, first, second)
}

func TestGenerateSurveyHashURL_DifferentCustomersDiffer(t *testing.T) {
	t.Parallel()

	a := GenerateSurveyHashURL("https://feedback.cbe.com.et/M/81c8db50", "Alice")
	b := GenerateSurveyHashURL("https://feedback.cbe.com.et/M/81c8db50", "Bob")
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
