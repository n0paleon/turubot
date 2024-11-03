package waifupics

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"turubot/internal/ports"
)

func TestWaifuPicsAPI(t *testing.T) {
	waifu := NewWaifuPicsAPI()

	url, err := waifu.GetRandomAnime(ports.WaifuNSFW)

	assert.NoError(t, err)
	assert.True(t, strings.Contains(url, "https"))

	t.Logf("result: %s", url)
}

func TestWaifuPicsAPISFW(t *testing.T) {
	waifu := NewWaifuPicsAPI()

	url, err := waifu.GetRandomAnime(ports.WaifuSFW)

	assert.NoError(t, err)
	assert.True(t, strings.Contains(url, "https"))

	t.Logf("result: %s", url)
}

func BenchmarkWaifuPicsAPI(b *testing.B) {
	waifu := NewWaifuPicsAPI()

	for i := 0; i < b.N; i++ {
		url, err := waifu.GetRandomAnime(ports.WaifuNSFW)
		if err != nil {
			b.Fatal(err)
		}

		assert.True(b, strings.Contains(url, "https"))
	}
}
