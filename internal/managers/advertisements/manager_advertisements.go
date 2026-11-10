package advertisements

import (
	"encoding/json"
	"io"
	"testing"

	client "api-tests-student/internal/client/http/advertisements"
	"api-tests-student/internal/managers/advertisements/models"
	"github.com/stretchr/testify/require"
)

func GetAdvertisementsList(t *testing.T, token, text string) models.AdvertisementsList {
	raw := client.HttpGetAdvertisements(t, token, text)
	defer raw.Body.Close()

	body, err := io.ReadAll(raw.Body)
	require.NoError(t, err)

	var resp models.AdvertisementsList
	require.NoError(t, json.Unmarshal(body, &resp))

	return resp
}
