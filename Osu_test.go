package osu

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	APIKey = os.Getenv("OSU_API_KEY")
}

func TestUser(t *testing.T) {
	userName := "Aky"
	user, err := GetUser(userName)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, user.UserName, userName)
	assert.NotEmpty(t, user.PPRaw)
	assert.NotEmpty(t, user.Level)
	assert.NotEmpty(t, user.Accuracy)
	assert.NotEmpty(t, user.PlayCount)
}
