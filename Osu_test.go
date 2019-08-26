package osu

import (
	"os"
	"testing"

	"github.com/akyoto/assert"
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
	assert.NotEqual(t, user.PPRaw, "")
	assert.NotEqual(t, user.Level, "")
	assert.NotEqual(t, user.Accuracy, "")
	assert.NotEqual(t, user.PlayCount, "")
}
