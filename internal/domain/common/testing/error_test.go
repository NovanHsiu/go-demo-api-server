package testing

import (
	"fmt"
	"testing"

	"github.com/NovanHsiu/go-demo-api-server/internal/domain/common"
	"github.com/stretchr/testify/assert"
)

func TestStatusCode(t *testing.T) {
	t.Parallel()
	// Args
	err := common.NewError(common.ErrorCodeParameterInvalid, fmt.Errorf("error test message")).(common.DomainError)
	assert.Equal(t, err.StatusCode(), common.ErrorCodeParameterInvalid.StatusCode)
}

func TestError(t *testing.T) {
	t.Parallel()
	// Args
	errMsg := "error test message"
	err := common.NewError(common.ErrorCodeParameterInvalid, fmt.Errorf(errMsg)).(common.DomainError)
	assert.Equal(t, err.Error(), errMsg)
	errMsg2 := "error test message 2"
	err.ErrorMessage = errMsg2
	assert.Equal(t, err.Error(), errMsg2)
}
