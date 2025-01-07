package futures

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDo(t *testing.T) {
	inputs := []string{"a", "b", "c", "d", "e", "f", "g"}
	expected := []string{"aa", "bb", "cc", "dd", "ee", "ff", "gg"}

	start := time.Now()
	futures := make([]Future[string], len(inputs))
	for i, input := range inputs {
		input := input
		futures[i] = Do(func() (string, error) {
			time.Sleep(100 * time.Millisecond)
			return input + input, nil
		})
	}

	results := make([]string, len(futures))
	for i, future := range futures {
		res, err := future.Await()
		assert.NoError(t, err)
		results[i] = res
	}
	assert.WithinDuration(t, time.Now(), start, 300*time.Millisecond)

	assert.Equal(t, len(expected), len(results), "wrong number of results")
	for i, result := range results {
		assert.Equal(t, expected[i], result, "wrong result at index %d", i)
	}
}
