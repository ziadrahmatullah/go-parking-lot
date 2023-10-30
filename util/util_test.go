package util_test

import (
	"testing"

	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/util"
	"github.com/stretchr/testify/assert"
)
func TestDeleteElement(t *testing.T) {
	t.Run("should delete value in the array", func(t *testing.T) {
		arr := []int{9,8,7,6}
		expected := []int{9,7,6}

		arr = util.DeleteElement[int](arr, 1)

		assert.Equal(t, expected, arr)
	})
}