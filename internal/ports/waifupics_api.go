package ports

type WaifuCategory string

const (
	WaifuSFW  = "sfw"
	WaifuNSFW = "nsfw"
)

type WaifuPics interface {
	GetRandomAnime(category WaifuCategory) (string, error)
}
