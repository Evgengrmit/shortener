package link

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"ozonTask/shorter"
	"testing"
)

func TestLinkSQL_Add(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()
	req := require.New(t)
	req.NoError(err)
	memory := NewLinkSQL(db)
	for index, testCase := range TestCases {
		mock.
			ExpectExec("INSERT INTO links").
			WithArgs(testCase.shortURL, testCase.originalURL).
			WillReturnResult(sqlmock.NewResult(int64(index+1), int64(index+1)))
		result, err := memory.Add(testCase.originalURL)
		req.Equal(testCase.shortURL, result)
		req.NoError(err)
	}
}

func TestLinkSQL_Get(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()
	req := require.New(t)
	req.NoError(err)
	memory := NewLinkSQL(db)
	for _, testCase := range TestCases {
		mock.ExpectQuery("SELECT original FROM links WHERE").
			WithArgs(testCase.shortURL).
			WillReturnRows(func() *sqlmock.Rows {
				result := sqlmock.NewRows([]string{"long_URL"})
				result.AddRow(testCase.originalURL)
				return result
			}())
		result, err := memory.Get(testCase.shortURL)
		req.NoError(err)
		req.NoError(mock.ExpectationsWereMet())
		req.Equal(testCase.originalURL, result)
	}

}

func TestLinkSQL_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()
	req := require.New(t)
	req.NoError(err)
	memory := NewLinkSQL(db)
	errorCase := TestCase{
		originalURL: "",
		shortURL:    shorter.GetShort("https://ozon.ru"),
	}
	mock.ExpectQuery("SELECT original FROM links WHERE").
		WithArgs(errorCase.shortURL).
		WillReturnError(fmt.Errorf("sql: no rows in result set"))

	_, err = memory.Get(errorCase.shortURL)
	req.Error(err)
}
