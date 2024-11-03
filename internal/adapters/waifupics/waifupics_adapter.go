package waifupics

import (
	"encoding/json"
	"fmt"
	"github.com/gojek/heimdall/v7"
	"github.com/gojek/heimdall/v7/httpclient"
	"go.uber.org/fx"
	"math/rand"
	"time"
	"turubot/infra/logger"
	"turubot/internal/ports"
)

type waifupics struct {
}

var apiEndpoint = "https://api.waifu.pics"

func NewWaifuPicsAPI() ports.WaifuPics {
	return &waifupics{}
}

func (w *waifupics) GetRandomAnime(category ports.WaifuCategory) (string, error) {
	waifuType := w.randomType(category)

	return w.getImageUrl(category, waifuType)
}

func (w *waifupics) getImageUrl(category ports.WaifuCategory, waifuType string) (string, error) {
	url := fmt.Sprintf("%s/%s/%s", apiEndpoint, category, waifuType)

	backoffInterval := 2 * time.Millisecond
	maximumJitterInterval := 5 * time.Millisecond
	backoff := heimdall.NewConstantBackoff(backoffInterval, maximumJitterInterval)
	retrier := heimdall.NewRetrier(backoff)
	timeout := 3000 * time.Millisecond

	client := httpclient.NewClient(
		httpclient.WithHTTPTimeout(timeout),
		httpclient.WithRetryCount(3),
		httpclient.WithRetrier(retrier),
	)

	response, err := client.Get(url, nil)
	if err != nil {
		logger.C.Errorw("heimdall returning an error", "error", err)
		return "", nil
	}
	defer func() {
		_ = response.Body.Close()
	}()

	var apiResponse struct {
		URL string `json:"url"`
	}

	if err := json.NewDecoder(response.Body).Decode(&apiResponse); err != nil {
		logger.C.Errorw("json decoder returning an error", "error", err)
		return "", nil
	}

	return apiResponse.URL, nil
}

func (w *waifupics) randomType(category ports.WaifuCategory) string {
	nsfw := []string{
		"waifu", "neko", "trap", "blowjob",
	}
	sfw := []string{
		"waifu", "neko", "shinobu", "megumin", "bully", "cuddle", "cry", "hug", "awoo", "kiss", "lick", "pat", "smug", "bonk", "yeet", "blush", "smile", "wave", "highfive", "handhold",
		"nom", "byte", "glomp", "slap", "kill", "kick", "happy", "wink", "poke", "dance", "cringe",
	}

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	switch category {
	case ports.WaifuNSFW:
		return nsfw[r.Intn(len(nsfw))]
	case ports.WaifuSFW:
		return sfw[r.Intn(len(sfw))]
	default:
		return ""
	}
}

var Module = fx.Options(
	fx.Provide(NewWaifuPicsAPI),
)
