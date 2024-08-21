package testing

import (
	"testing"

	"github.com/NovanHsiu/go-demo-api-server/internal/domain/cipher"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func TestEncodePassword(t *testing.T) {
	t.Parallel()
	// Args
	type Args struct {
		PlainText string
	}
	var args Args
	_ = faker.FakeData(&args)
	dcipher := cipher.DefaultCipher()
	encodedPassword := dcipher.EncodePassword(args.PlainText)
	assert.Greater(t, len(encodedPassword), len(args.PlainText))
}

func TestComparePassword(t *testing.T) {
	// Args
	type Args struct {
		PlainText string
	}
	var args Args
	_ = faker.FakeData(&args)
	dcipher := cipher.DefaultCipher()
	encodedPassword := dcipher.EncodePassword(args.PlainText)
	equalPassword := dcipher.ComparePassword(encodedPassword, args.PlainText)
	assert.True(t, equalPassword)
}

func TestGetJWT(t *testing.T) {
	t.Parallel()
	// Args
	type Args struct {
		PlainText string
	}
	var args Args
	_ = faker.FakeData(&args)
	dcipher := cipher.DefaultCipher()
	jwt := dcipher.GetJWT(args.PlainText)
	assert.Greater(t, len(jwt), len(args.PlainText))
}

func TestParseJWT(t *testing.T) {
	t.Parallel()
	// Args
	type Args struct {
		PlainText string
	}
	var args Args
	_ = faker.FakeData(&args)
	dcipher := cipher.DefaultCipher()
	jwt := dcipher.GetJWT(args.PlainText)
	parsedPlainText, err := dcipher.ParseJWT(jwt)
	assert.NoError(t, err)
	assert.Equal(t, args.PlainText, parsedPlainText)
}
