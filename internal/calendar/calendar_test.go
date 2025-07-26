package calendar

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPokeCalendar(t *testing.T) {
	// @@@ html 내용이 테스트 시점에 따라 바뀌므로 다른 테스트 방식 고려 필요 @@@
	// Test: Good url with 5 days duration
	err := PokeCalendar()
	// require.NoError(t, err)
	require.Error(t, err)
	assert.Equal(t, ErrTemp, err)
}
