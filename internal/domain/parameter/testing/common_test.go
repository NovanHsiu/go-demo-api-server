package testing

import (
	"testing"

	"github.com/NovanHsiu/go-demo-api-server/internal/domain/parameter"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func TestPageGetOffset(t *testing.T) {
	t.Parallel()
	// Args
	type Args struct {
		PageNumber int
		PageSize   int
	}
	var args Args
	_ = faker.FakeData(&args)
	page := parameter.Page{
		PageNumber: args.PageNumber,
		PageSize:   args.PageSize,
	}
	offset := page.GetOffset()
	expectedOffset := (page.PageNumber - 1) * page.PageSize
	assert.Equal(t, expectedOffset, offset)
}

func TestPageGetPages(t *testing.T) {
	t.Parallel()
	page := parameter.Page{
		PageSize: 0,
	}
	elementCount := 100
	pages := page.GetPages(elementCount)
	expectedPages := elementCount / 10
	assert.Equal(t, expectedPages, pages)
	elementCount2 := 101
	pages = page.GetPages(elementCount2)
	expectedPages = (elementCount / 10) + 1
	assert.Equal(t, expectedPages, pages)
}
