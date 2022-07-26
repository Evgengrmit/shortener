package link

type Link struct {
	Data string
}

type LinkStorage interface {
	Add(originalURL string) (string, error)
	Get(shortURL string) (string, error)
}
