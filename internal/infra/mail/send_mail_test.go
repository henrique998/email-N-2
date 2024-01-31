package mail

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Send_Mail(t *testing.T) {
	assert := assert.New(t)

	err := SendMail()

	assert.Nil(err)
}
