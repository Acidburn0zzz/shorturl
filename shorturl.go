package shorturl

import (
	"errors"
	"github.com/subosito/shorturl/bitly"
	"github.com/subosito/shorturl/gitio"
	"github.com/subosito/shorturl/isgd"
	"github.com/subosito/shorturl/lns"
	"github.com/subosito/shorturl/tinyurl"
	"os"
)

type Client struct {
	Provider string
	Params   map[string]string
}

func NewClient(provider string) *Client {
	return &Client{Provider: provider}
}

func (c *Client) Shorten(u string) ([]byte, error) {
	switch c.Provider {
	case "tinyurl":
		s := tinyurl.New()
		return s.Shorten(u)
	case "isgd":
		s := isgd.New()
		return s.Shorten(u)
	case "gitio":
		s := gitio.New()
		return s.Shorten(u)
	case "bitly":
		s := bitly.New()
		s.Params["login"] = os.Getenv("BITLY_LOGIN")
		s.Params["apiKey"] = os.Getenv("BITLY_API_KEY")
		return s.Shorten(u)
	case "lns":
		s := lns.New()
		return s.Shorten(u)
	}

	err := errors.New("You should not see this :P")
	return nil, err
}
