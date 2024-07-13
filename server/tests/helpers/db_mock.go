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
	mock.ExpectQuery("SELECT id, category_id, name_eng, price FROM items ORDER BY id ASC").WillReturnRows(mock.NewRows([]string{"id", "category_id", "name_eng", "price"}).AddRow(1, 1, "item1", float64(100)))
}

func MockInsertItemQuery(mock pgxmock.PgxConnIface, item models.Item) {
	mock.ExpectQuery(`INSERT INTO items \(name_eng, price, category_id\) VALUES \(\$1, \$2, \$3\) RETURNING id, category_id, name_eng, price`).
		WithArgs(item.NameEng, item.Price, item.CategoryID).
		WillReturnRows(mock.NewRows([]string{"id", "category_id", "name_eng", "price"}).AddRow(item.ID, item.CategoryID, item.NameEng, item.Price))
}

func MockUpdateItemQuery(mock pgxmock.PgxConnIface, item models.Item) {
	mock.ExpectQuery(`UPDATE items SET name_eng = \$1, price = \$2, category_id = \$3 WHERE id = \$4 RETURNING id, category_id, name_eng, price`).
		WithArgs(item.NameEng, item.Price, item.CategoryID, item.ID).
		WillReturnRows(mock.NewRows([]string{"id", "category_id", "name_eng", "price"}).AddRow(item.ID, item.CategoryID, item.NameEng, item.Price))
}
