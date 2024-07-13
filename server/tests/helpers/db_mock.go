package test_helpers

import (
	"server/models"

	"github.com/pashagolub/pgxmock/v4"
)

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

func MockInsertItemQuery(mock pgxmock.PgxConnIface, item models.Item) {
	mock.ExpectQuery(`INSERT INTO items \(name_eng, price, category_id\) VALUES \(\$1, \$2, \$3\) RETURNING id, category_id, name_eng, price`).
		WithArgs(item.NameEng, item.Price, item.CategoryID).
		WillReturnRows(mock.NewRows([]string{"id", "category_id", "name_eng", "price"}).AddRow(111, item.CategoryID, item.NameEng, item.Price))
}
