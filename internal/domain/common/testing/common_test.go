package testing

import (
	"regexp"
	"strings"
	"testing"

	"github.com/NovanHsiu/go-demo-api-server/internal/domain/common"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func TestArrayContainsString(t *testing.T) {
	t.Parallel()
	// Args
	type Args struct {
		MatchString string
	}
	var args Args
	_ = faker.FakeData(&args)
	stringArray := []string{}
	assert.False(t, common.ArrayContainsString(stringArray, args.MatchString))
	stringArray = append(stringArray, args.MatchString+"abc")
	assert.True(t, common.ArrayContainsString(stringArray, args.MatchString))
}

func TestArrayHasString(t *testing.T) {
	t.Parallel()
	// Args
	type Args struct {
		MatchString string
	}
	var args Args
	_ = faker.FakeData(&args)
	stringArray := []string{args.MatchString + "abc"}
	assert.False(t, common.ArrayHasString(stringArray, args.MatchString))
	stringArray = append(stringArray, args.MatchString)
	assert.True(t, common.ArrayHasString(stringArray, args.MatchString))
}

func TestGetRandomString(t *testing.T) {
	t.Parallel()
	// Args
	stringLength := 0
	rngStr := common.GetRandomString(stringLength)
	assert.Equal(t, 0, len(rngStr))
	stringLength = 101
	rngStr = common.GetRandomString(stringLength)
	assert.Equal(t, 0, len(rngStr))
	stringLength = 12
	rngStr = common.GetRandomString(stringLength)
	assert.Equal(t, stringLength, len(rngStr))
}

func TestGetRandomTempDirName(t *testing.T) {
	t.Parallel()
	// Args
	basePath := common.GetExecutionDir()
	tmpNameLength := 12
	rngDirPath := common.GetRandomTempDirName(basePath, tmpNameLength)
	assert.True(t, strings.HasPrefix(rngDirPath, basePath))
	assert.Equal(t, len(basePath)+tmpNameLength+1, len(rngDirPath))
}

func TestMatchDatePattern(t *testing.T) {
	t.Parallel()
	// Args
	testDate := "2034-03-12"
	assert.True(t, common.MatchDatePattern(testDate))
	testDate = "99998-13-57"
	assert.False(t, common.MatchDatePattern(testDate))
	testDate = "9999-13-57"
	assert.False(t, common.MatchDatePattern(testDate))
	testDate = "abcde123"
	assert.False(t, common.MatchDatePattern(testDate))
}

func TestToRocDate(t *testing.T) {
	t.Parallel()
	// Args
	testDate := "2034-03-12"
	rocDate := common.ToRocDate(testDate)
	expectedRocDate := "1230312"
	assert.Equal(t, expectedRocDate, rocDate)
}

func TestRocToADYear(t *testing.T) {
	t.Parallel()
	// Args
	rocDate := "1230312"
	actualDate := common.RocToADYear(rocDate)
	expectedRocDate := "2034-03-12"
	assert.Equal(t, expectedRocDate, actualDate)
}

func TestGetMacAddress(t *testing.T) {
	t.Parallel()
	mac := common.GetMacAddress()
	macRegex := `^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$`
	re := regexp.MustCompile(macRegex)
	assert.True(t, re.MatchString(mac))
}

func TestIsImage(t *testing.T) {
	t.Parallel()
	assert.True(t, common.IsImage("abc.png"))
	assert.True(t, common.IsImage("abc.jpg"))
	assert.True(t, common.IsImage("abc.jpeg"))
	assert.False(t, common.IsImage("abc.txt"))
	assert.False(t, common.IsImage("abc.xlsx"))
}
