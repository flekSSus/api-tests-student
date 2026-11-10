package assertions

import (
	"testing"
	"api-tests-template/internal/managers/advertisements/models"
	"github.com/stretchr/testify/require"
)

func ContainsAdvertisementByTitle(t *testing.T, items []models.Advertisement, expectedTitle string) {
	found := false
	for _, item := range items {
		if item.Title == expectedTitle {
			found = true
			break
		}
	}
	require.True(t, found, "Объявление с названием '%s' не найдено", expectedTitle)
}
