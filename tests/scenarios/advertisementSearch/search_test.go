package advertisementSearch

import (
	"os"
	"strings"
	"testing"

	"api-tests-student/internal/assertions"
	"api-tests-student/internal/managers/auth"
	authModels "api-tests-student/internal/managers/auth/models"
	advList "api-tests-student/internal/managers/advertisements"
	"api-tests-student/internal/utils"
	"api-tests-student/tests/base"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

const (
	existingAdTitle = "Новые сандали Toyota"
	searchBeginning = "Новые сандали"
	searchMiddle    = "сандали"
	searchEnd       = "Toyota"
)

type TestSuite struct {
	base.BaseTestSuite
	loginData authModels.LoginOkResponse
}

func (s *TestSuite) SetupSuite() {
	base.SetupSuite()

	base.Precondition("Проверка заполнения констант для поиска объявления")
	require.NotEmpty(s.T(), existingAdTitle, "задайте existingAdTitle")
	require.NotEmpty(s.T(), searchBeginning, "задайте searchBeginning")
	require.NotEmpty(s.T(), searchMiddle, "задайте searchMiddle")
	require.NotEmpty(s.T(), searchEnd, "задайте searchEnd")
  
	s.loginData = auth.Login(s.T(), os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))
}

func (s *TestSuite) TestSearchByTitle() {
	list := advList.GetAdvertisementsList(s.T(), s.loginData.Token, searchMiddle)

	s.Run("В ответе есть хотя бы одно объявление", func() {
		require.NotEmpty(s.T(), list.Items)
	})

	s.Run("Среди найденных есть наше объявление", func() {
		assertions.ContainsAdvertisementByTitle(s.T(), list.Items, existingAdTitle)
	})
}


func (s *TestSuite) TestSearchByTitle_Substring() {
	cases := []struct {
		name   string
		search string
	}{
		{"начало", searchBeginning},
		{"середина", searchMiddle},
		{"конец", searchEnd},
		{"нижний регистр", strings.ToLower(searchMiddle)},
		{"верхний регистр", strings.ToUpper(searchMiddle)},
		{"смешанный регистр", utils.MixCase(searchMiddle)},
	}

	for _, tc := range cases {
		s.Run(tc.name, func() {
			list := advList.GetAdvertisementsList(s.T(), s.loginData.Token, tc.search)
			assertions.ContainsAdvertisementByTitle(s.T(), list.Items, existingAdTitle)
		})
	}
}

func TestAdvertisementSearchSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
