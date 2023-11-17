package input_test

import (
	"foxyfy/internal/input"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYoutubeUrlValidation(t *testing.T) {
	assert.Equal(t, true, input.IsValidYoutubeUrl("https://www.youtube.com/watch?v=abc123"))
	assert.Equal(t, true, input.IsValidYoutubeUrl("https://www.youtube.com/watch?v=4-123Abc&list=xyz123&index=79"))

	assert.Equal(t, false, input.IsValidYoutubeUrl("https://www.notyoutube.com/watch?v=abc123"))
	assert.Equal(t, false, input.IsValidYoutubeUrl("http://www.youtube.com/watch?v=abc123"))
	assert.Equal(t, false, input.IsValidYoutubeUrl("https://www.youtubecom/watch?v=abc123"))
}
