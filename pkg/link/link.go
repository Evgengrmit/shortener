package link

type Link struct {
	Data string `json:"link"`
}

type LinkStorage interface {
	Add(originalURL string) (string, error)
	Get(shortURL string) (string, error)
}
