package mock

import (
	"database/sql"
	"errors"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

type Domain struct {
	DB *sql.DB
}

type Data struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

func (domain Domain) GetDataByID(id int) (Data, error) {
	result := Data{}
	// write code here
	return result, nil
}

func Test_GetDataByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	tests := []struct {
		name    string
		id      int
		want    Data
		wantErr bool
		mock    func(id int)
	}{
		{
			name: "test 1",
			id:   1,
			want: Data{
				ID:   1,
				Name: "data 1",
			},
			wantErr: false,
			mock: func(id int) {
				row := sqlmock.NewRows([]string{"id", "name"}).AddRow(id, "data 1")
				mock.ExpectQuery(regexp.QuoteMeta("select id, name from data where id = $1")).WithArgs(1).WillReturnRows(row)

			},
		},
		{
			name:    "test 2",
			id:      1,
			want:    Data{},
			wantErr: true,
			mock: func(id int) {
				mock.ExpectQuery(regexp.QuoteMeta("select id, name from data where id = $1")).WithArgs(1).WillReturnError(errors.New("some errors"))

			},
		},
	}
	for _, tt := range tests {
		domain := Domain{DB: db}
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.id)
			got, err := domain.GetDataByID(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want %v", got, tt.want)
				return
			}
		})
	}
}
