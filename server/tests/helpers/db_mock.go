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

// The reason for using pointers in mockItems is to accurately mimic the behavior of the pgx library,
// which returns data as pointers in pgx.Rows. We then need to use pgx.Scan to populate the struct fields with these values.
var mockItems = [][]interface{}{
	{models.PtrInt(1), models.PtrInt(1), models.PtrInt(1), models.PtrFloat64(5.99), models.PtrString("Spring Rolls"), models.PtrString("春卷"), models.PtrBool(false), models.PtrBool(false), models.PtrBool(false), models.PtrString(""), models.PtrBool(false), models.PtrFloat64(0)},                // Appetizers
	{models.PtrInt(2), models.PtrInt(2), models.PtrInt(2), models.PtrFloat64(4.99), models.PtrString("Hot and Sour Soup"), models.PtrString("酸辣汤"), models.PtrBool(false), models.PtrBool(false), models.PtrBool(false), models.PtrString("Small"), models.PtrBool(true), models.PtrFloat64(0)},      // Soups
	{models.PtrInt(3), models.PtrInt(2), models.PtrInt(2), models.PtrFloat64(4.99), models.PtrString("Hot and Sour Soup"), models.PtrString("酸辣汤"), models.PtrBool(false), models.PtrBool(false), models.PtrBool(false), models.PtrString("Large"), models.PtrBool(false), models.PtrFloat64(4)},     // Soups
	{models.PtrInt(4), models.PtrInt(3), models.PtrInt(3), models.PtrFloat64(6.99), models.PtrString("Chicken Egg Foo Yung"), models.PtrString("雞芙蓉蛋"), models.PtrBool(false), models.PtrBool(false), models.PtrBool(false), models.PtrString(""), models.PtrBool(false), models.PtrFloat64(0)},      // Egg Foo Yung
	{models.PtrInt(5), models.PtrInt(4), models.PtrInt(4), models.PtrFloat64(7.99), models.PtrString("Stir-fried Bok Choy"), models.PtrString("炒青菜"), models.PtrBool(false), models.PtrBool(false), models.PtrBool(false), models.PtrString(""), models.PtrBool(false), models.PtrFloat64(0)},        // Vegetables
	{models.PtrInt(6), models.PtrInt(5), models.PtrInt(5), models.PtrFloat64(12.99), models.PtrString("Salt and Pepper Shrimp"), models.PtrString("椒鹽蝦"), models.PtrBool(false), models.PtrBool(false), models.PtrBool(false), models.PtrString(""), models.PtrBool(false), models.PtrFloat64(0)},    // Seafood
	{models.PtrInt(7), models.PtrInt(6), models.PtrInt(6), models.PtrFloat64(13.99), models.PtrString("Stir-fried Scallops"), models.PtrString("炒帶子"), models.PtrBool(false), models.PtrBool(false), models.PtrBool(false), models.PtrString(""), models.PtrBool(false), models.PtrFloat64(0)},       // Oysters/Scallops
	{models.PtrInt(8), models.PtrInt(7), models.PtrInt(7), models.PtrFloat64(14.99), models.PtrString("Beef Hot Pot"), models.PtrString("牛肉煲"), models.PtrBool(false), models.PtrBool(false), models.PtrBool(false), models.PtrString(""), models.PtrBool(false), models.PtrFloat64(0)},              // Hot Pot
	{models.PtrInt(9), models.PtrInt(8), models.PtrInt(8), models.PtrFloat64(9.99), models.PtrString("Sweet and Sour Pork"), models.PtrString("糖醋排骨"), models.PtrBool(false), models.PtrBool(false), models.PtrBool(false), models.PtrString(""), models.PtrBool(false), models.PtrFloat64(0)},       // Pork
	{models.PtrInt(10), models.PtrInt(9), models.PtrInt(9), models.PtrFloat64(10.99), models.PtrString("Beef with Broccoli"), models.PtrString("西蘭花牛肉"), models.PtrBool(false), models.PtrBool(false), models.PtrBool(false), models.PtrString(""), models.PtrBool(false), models.PtrFloat64(0)},     // Beef
	{models.PtrInt(11), models.PtrInt(10), models.PtrInt(10), models.PtrFloat64(8.99), models.PtrString("Kung Pao Chicken"), models.PtrString("宫保鸡丁"), models.PtrBool(false), models.PtrBool(false), models.PtrBool(false), models.PtrString(""), models.PtrBool(false), models.PtrFloat64(0)},       // Chicken
	{models.PtrInt(12), models.PtrInt(11), models.PtrInt(11), models.PtrFloat64(7.99), models.PtrString("BBQ Pork Over Rice"), models.PtrString("叉燒飯"), models.PtrBool(false), models.PtrBool(false), models.PtrBool(false), models.PtrString(""), models.PtrBool(false), models.PtrFloat64(0)},      // Over Rice
	{models.PtrInt(13), models.PtrInt(12), models.PtrInt(12), models.PtrFloat64(8.99), models.PtrString("Yangzhou Fried Rice"), models.PtrString("扬州炒饭"), models.PtrBool(false), models.PtrBool(false), models.PtrBool(false), models.PtrString(""), models.PtrBool(false), models.PtrFloat64(0)},    // Fried Rice
	{models.PtrInt(14), models.PtrInt(13), models.PtrInt(13), models.PtrFloat64(9.99), models.PtrString("Chicken Chow Mein"), models.PtrString("雞肉炒麵"), models.PtrBool(false), models.PtrBool(false), models.PtrBool(false), models.PtrString(""), models.PtrBool(false), models.PtrFloat64(0)},      // Chow Mein
	{models.PtrInt(15), models.PtrInt(14), models.PtrInt(14), models.PtrFloat64(6.99), models.PtrString("Beef Noodle Soup"), models.PtrString("牛肉湯麵"), models.PtrBool(false), models.PtrBool(false), models.PtrBool(false), models.PtrString(""), models.PtrBool(false), models.PtrFloat64(0)},       // Noodle Soup
	{models.PtrInt(16), models.PtrInt(15), models.PtrInt(15), models.PtrFloat64(5.99), models.PtrString("Pork Congee"), models.PtrString("豬肉粥"), models.PtrBool(false), models.PtrBool(false), models.PtrBool(false), models.PtrString(""), models.PtrBool(false), models.PtrFloat64(0)},             // Congee
	{models.PtrInt(17), models.PtrInt(16), models.PtrInt(16), models.PtrFloat64(15.99), models.PtrString("General Tso's Chicken"), models.PtrString("左宗棠雞"), models.PtrBool(false), models.PtrBool(false), models.PtrBool(false), models.PtrString(""), models.PtrBool(false), models.PtrFloat64(0)}, // Specials
	{models.PtrInt(18), models.PtrInt(17), models.PtrInt(17), models.PtrFloat64(2.99), models.PtrString("Local Beer"), models.PtrString("本地啤酒"), models.PtrBool(false), models.PtrBool(true), models.PtrBool(false), models.PtrString(""), models.PtrBool(false), models.PtrFloat64(0)},              // Drinks
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

func MockInsertItemQuery(mock pgxmock.PgxConnIface, item models.Item) {
	mock.ExpectQuery("INSERT INTO items \\(menu_id, category_id, price, name_eng, name_oth, special, alcohol, custom, variant, variant_default, variant_price_charge\\) VALUES \\(\\$1, \\$2, \\$3, \\$4, \\$5, \\$6, \\$7, \\$8, \\$9, \\$10, \\$11\\) RETURNING id, menu_id, category_id, price, name_eng, name_oth, special, alcohol, custom, variant, variant_default, variant_price_charge").
		WithArgs(*item.MenuID, *item.CategoryID, *item.Price, *item.NameEng, *item.NameOth, *item.Special, *item.Alcohol, *item.Custom, *item.Variant, *item.VariantDefault, *item.VariantPriceCharge).
		WillReturnRows(mock.NewRows([]string{"id", "menu_id", "category_id", "price", "name_eng", "name_oth", "special", "alcohol", "custom", "variant", "variant_default", "variant_price_charge"}).
			AddRow(models.PtrInt(19), item.MenuID, item.CategoryID, item.Price, item.NameEng, item.NameOth, item.Special, item.Alcohol, item.Custom, item.Variant, item.VariantDefault, item.VariantPriceCharge))
}

func MockUpdateItemQuery(mock pgxmock.PgxConnIface, item models.Item) {
	mock.ExpectQuery("UPDATE items SET menu_id = \\$1, category_id = \\$2, price = \\$3, name_eng = \\$4, name_oth = \\$5, special = \\$6, alcohol = \\$7, custom = \\$8, variant = \\$9, variant_default = \\$10, variant_price_charge = \\$11 WHERE id = \\$12 RETURNING id, menu_id, category_id, price, name_eng, name_oth, special, alcohol, custom, variant, variant_default, variant_price_charge").
		WithArgs(*item.MenuID, *item.CategoryID, *item.Price, *item.NameEng, *item.NameOth, *item.Special, *item.Alcohol, *item.Custom, *item.Variant, *item.VariantDefault, *item.VariantPriceCharge, *item.ID).
		WillReturnRows(mock.NewRows([]string{"id", "menu_id", "category_id", "price", "name_eng", "name_oth", "special", "alcohol", "custom", "variant", "variant_default", "variant_price_charge"}).
			AddRow(item.ID, item.MenuID, item.CategoryID, item.Price, item.NameEng, item.NameOth, item.Special, item.Alcohol, item.Custom, item.Variant, item.VariantDefault, item.VariantPriceCharge))
}
