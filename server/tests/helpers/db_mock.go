package test_helpers

import (
	"fmt"
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

var mockItems = [][]interface{}{
	{1, 1, 1, float64(5.99), "Spring Rolls", "春卷", false, false, false, "", false, float64(0)},                // Appetizers
	{2, 2, 2, float64(4.99), "Hot and Sour Soup", "酸辣汤", false, false, false, "Small", true, float64(0)},      // Soups
	{3, 2, 2, float64(4.99), "Hot and Sour Soup", "酸辣汤", false, false, false, "Large", false, float64(4)},     // Soups
	{4, 3, 3, float64(6.99), "Chicken Egg Foo Yung", "雞芙蓉蛋", false, false, false, "", false, float64(0)},      // Egg Foo Yung
	{5, 4, 4, float64(7.99), "Stir-fried Bok Choy", "炒青菜", false, false, false, "", false, float64(0)},        // Vegetables
	{6, 5, 5, float64(12.99), "Salt and Pepper Shrimp", "椒鹽蝦", false, false, false, "", false, float64(0)},    // Seafood
	{7, 6, 6, float64(13.99), "Stir-fried Scallops", "炒帶子", false, false, false, "", false, float64(0)},       // Oysters/Scallops
	{8, 7, 7, float64(14.99), "Beef Hot Pot", "牛肉煲", false, false, false, "", false, float64(0)},              // Hot Pot
	{9, 8, 8, float64(9.99), "Sweet and Sour Pork", "糖醋排骨", false, false, false, "", false, float64(0)},       // Pork
	{10, 9, 9, float64(10.99), "Beef with Broccoli", "西蘭花牛肉", false, false, false, "", false, float64(0)},     // Beef
	{11, 10, 10, float64(8.99), "Kung Pao Chicken", "宫保鸡丁", false, false, false, "", false, float64(0)},       // Chicken
	{12, 11, 11, float64(7.99), "BBQ Pork Over Rice", "叉燒飯", false, false, false, "", false, float64(0)},      // Over Rice
	{13, 12, 12, float64(8.99), "Yangzhou Fried Rice", "扬州炒饭", false, false, false, "", false, float64(0)},    // Fried Rice
	{14, 13, 13, float64(9.99), "Chicken Chow Mein", "雞肉炒麵", false, false, false, "", false, float64(0)},      // Chow Mein
	{15, 14, 14, float64(6.99), "Beef Noodle Soup", "牛肉湯麵", false, false, false, "", false, float64(0)},       // Noodle Soup
	{16, 15, 15, float64(5.99), "Pork Congee", "豬肉粥", false, false, false, "", false, float64(0)},             // Congee
	{17, 16, 16, float64(15.99), "General Tso's Chicken", "左宗棠雞", false, false, false, "", false, float64(0)}, // Specials
	{18, 17, 17, float64(2.99), "Local Beer", "本地啤酒", false, true, false, "", false, float64(0)},              // Drinks
}

func MockGetItemQuery(mock pgxmock.PgxConnIface, id int) {
	rows := mock.NewRows([]string{"id", "menu_id", "category_id", "price", "name_eng", "name_oth", "special", "alcohol", "custom", "variant", "variant_default", "variant_price_charge"})
	if id > 0 && id <= len(mockItems) {
		rows.AddRow(mockItems[id-1]...)
	} else {
		fmt.Println("MockGetItemQuery: Invalid id", id, "for mockItems. There are only", len(mockItems), "items.")
	}

	mock.ExpectQuery("SELECT id, menu_id, category_id, price, name_eng, name_oth, special, alcohol, custom, variant, variant_default, variant_price_charge FROM items WHERE id = \\$1").WithArgs(id).WillReturnRows(rows)
}

func MockGetItemsQuery(mock pgxmock.PgxConnIface) {
	rows := mock.NewRows([]string{"id", "menu_id", "category_id", "price", "name_eng", "name_oth", "special", "alcohol", "custom", "variant", "variant_default", "variant_price_charge"})

	for _, item := range mockItems {
		rows.AddRow(item...)
	}
	mock.ExpectQuery("SELECT id, menu_id, category_id, price, name_eng, name_oth, special, alcohol, custom, variant, variant_default, variant_price_charge FROM items ORDER BY id ASC").
		WillReturnRows(rows)
}

func MockInsertItemQuery(mock pgxmock.PgxConnIface, item *models.Item) {
	mock.ExpectQuery("INSERT INTO items \\(menu_id, category_id, price, name_eng, name_oth, special, alcohol, custom, variant, variant_default, variant_price_charge\\) VALUES \\(\\$1, \\$2, \\$3, \\$4, \\$5, \\$6, \\$7, \\$8, \\$9, \\$10, \\$11\\) RETURNING id, menu_id, category_id, price, name_eng, name_oth, special, alcohol, custom, variant, variant_default, variant_price_charge").
		WithArgs(item.MenuID, item.CategoryID, item.Price, item.NameEng, item.NameOth, item.Special, item.Alcohol, item.Custom, item.Variant, item.VariantDefault, item.VariantPriceCharge).
		WillReturnRows(mock.NewRows([]string{"id", "menu_id", "category_id", "price", "name_eng", "name_oth", "special", "alcohol", "custom", "variant", "variant_default", "variant_price_charge"}).
			AddRow(19, item.MenuID, item.CategoryID, item.Price, item.NameEng, item.NameOth, item.Special, item.Alcohol, item.Custom, item.Variant, item.VariantDefault, item.VariantPriceCharge))
}

func MockUpdateItemQuery(mock pgxmock.PgxConnIface, item *models.Item) {
	mock.ExpectQuery("UPDATE items SET menu_id = \\$1, category_id = \\$2, price = \\$3, name_eng = \\$4, name_oth = \\$5, special = \\$6, alcohol = \\$7, custom = \\$8, variant = \\$9, variant_default = \\$10, variant_price_charge = \\$11 WHERE id = \\$12 RETURNING id, menu_id, category_id, price, name_eng, name_oth, special, alcohol, custom, variant, variant_default, variant_price_charge").
		WithArgs(item.MenuID, item.CategoryID, item.Price, item.NameEng, item.NameOth, item.Special, item.Alcohol, item.Custom, item.Variant, item.VariantDefault, item.VariantPriceCharge, item.ID).
		WillReturnRows(mock.NewRows([]string{"id", "menu_id", "category_id", "price", "name_eng", "name_oth", "special", "alcohol", "custom", "variant", "variant_default", "variant_price_charge"}).
			AddRow(item.ID, item.MenuID, item.CategoryID, item.Price, item.NameEng, item.NameOth, item.Special, item.Alcohol, item.Custom, item.Variant, item.VariantDefault, item.VariantPriceCharge))
}
