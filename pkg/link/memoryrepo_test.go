package link

import (
	"github.com/stretchr/testify/require"
	"ozonTask/shorter"
	"testing"
)

type TestCase struct {
	originalURL string
	shortURL    string
}

var TestCases = []TestCase{
	{
		originalURL: "https://www.ozon.ru/",
		shortURL:    shorter.GetShort("https://www.ozon.ru/"),
	},
	{
		originalURL: "https://www.ozon.ru/",
		shortURL:    shorter.GetShort("https://www.ozon.ru/"),
	},
	{
		originalURL: "https://www.youtube.com/watch?v=h0zxh2TPN_I&t=6433s",
		shortURL:    shorter.GetShort("https://www.youtube.com/watch?v=h0zxh2TPN_I&t=6433s"),
	},
	{
		originalURL: "https://ru.wikipedia.org/wiki/Go#%D0%9D%D0%B0%D0%B7%D0%BD%D0%B0%D1%87%D0%B5%D0%BD%D0%B8%D0%B5," +
			"_%D0%B8%D0%B4%D0%B5%D0%BE%D0%BB%D0%BE%D0%B3%D0%B8%D1%8F",
		shortURL: shorter.GetShort("https://ru.wikipedia.org/wiki/Go#%D0%9D%D0%B0%D0%B7%D0%BD%D0%B0%D1%87%D0%B5%D0%BD%D0%B8%D0%B5," +
			"_%D0%B8%D0%B4%D0%B5%D0%BE%D0%BB%D0%BE%D0%B3%D0%B8%D1%8F"),
	},
	{
		originalURL: "https://ya.ru/",
		shortURL:    shorter.GetShort("https://ya.ru/"),
	},
}

func TestLinkMemory_Add(t *testing.T) {

	req := require.New(t)
	memory := NewLinkMemory()
	for _, testCase := range TestCases {
		result, err := memory.Add(testCase.originalURL)
		req.Equal(testCase.shortURL, result)
		req.NoError(err)
	}
}

func TestLinkMemory_Get(t *testing.T) {
	req := require.New(t)
	memory := NewLinkMemory()
	for _, testCase := range TestCases {
		_, err := memory.Add(testCase.originalURL)
		req.NoError(err)
	}

	for _, testCase := range TestCases {
		result, err := memory.Get(testCase.shortURL)
		req.Equal(testCase.originalURL, result)
		req.NoError(err)
	}
}

func TestLinkMemory_Error(t *testing.T) {

	errorCases := []TestCase{
		{
			originalURL: "",
			shortURL:    shorter.GetShort("https://www.ozon.ru/"),
		},
		{
			originalURL: "",
			shortURL:    shorter.GetShort("https://www.ozon.ru/"),
		},
		{
			originalURL: "",
			shortURL:    shorter.GetShort("https://www.youtube.com/watch?v=h0zxh2TPN_I&t=6433s"),
		},
	}
	req := require.New(t)
	memory := NewLinkMemory()
	for _, errorCase := range errorCases {
		_, err := memory.Get(errorCase.shortURL)
		req.Error(err)
	}
}
