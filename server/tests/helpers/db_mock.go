package test_helpers

import "github.com/pashagolub/pgxmock/v4"

func SetupPgxMock() (pgxmock.PgxConnIface, error) {
	mock, err := pgxmock.NewConn()
	if err != nil {
		return nil, err
	}
	return mock, nil
}

func MockGetItemsQuery(mock pgxmock.PgxConnIface) {
	mock.ExpectQuery("SELECT id, category_id, name_eng, price FROM items").WillReturnRows(mock.NewRows([]string{"id", "category_id", "name_eng", "price"}).AddRow(1, 1, "item1", float64(100)))
}
