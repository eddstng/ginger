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

var mockItems = []models.Item{
	{ID: models.PtrInt(1), MenuID: models.PtrInt(1), CategoryID: models.PtrInt(1), Price: models.PtrFloat64(5.99), NameEng: models.PtrString("Spring Rolls"), NameOth: models.PtrString("春卷"), Special: models.PtrBool(false), Alcohol: models.PtrBool(false), Custom: models.PtrBool(false), Variant: models.PtrString(""), VariantDefault: models.PtrBool(false), VariantPriceCharge: models.PtrFloat64(0)},
	{ID: models.PtrInt(2), MenuID: models.PtrInt(2), CategoryID: models.PtrInt(2), Price: models.PtrFloat64(4.99), NameEng: models.PtrString("Hot and Sour Soup"), NameOth: models.PtrString("酸辣汤"), Special: models.PtrBool(false), Alcohol: models.PtrBool(false), Custom: models.PtrBool(false), Variant: models.PtrString("Small"), VariantDefault: models.PtrBool(true), VariantPriceCharge: models.PtrFloat64(0)},
	{ID: models.PtrInt(3), MenuID: models.PtrInt(2), CategoryID: models.PtrInt(2), Price: models.PtrFloat64(4.99), NameEng: models.PtrString("Hot and Sour Soup"), NameOth: models.PtrString("酸辣汤"), Special: models.PtrBool(false), Alcohol: models.PtrBool(false), Custom: models.PtrBool(false), Variant: models.PtrString("Large"), VariantDefault: models.PtrBool(false), VariantPriceCharge: models.PtrFloat64(4)},
	{ID: models.PtrInt(4), MenuID: models.PtrInt(3), CategoryID: models.PtrInt(3), Price: models.PtrFloat64(6.99), NameEng: models.PtrString("Chicken Egg Foo Yung"), NameOth: models.PtrString("雞芙蓉蛋"), Special: models.PtrBool(false), Alcohol: models.PtrBool(false), Custom: models.PtrBool(false), Variant: models.PtrString(""), VariantDefault: models.PtrBool(false), VariantPriceCharge: models.PtrFloat64(0)},
	{ID: models.PtrInt(5), MenuID: models.PtrInt(4), CategoryID: models.PtrInt(4), Price: models.PtrFloat64(7.99), NameEng: models.PtrString("Stir-fried Bok Choy"), NameOth: models.PtrString("炒青菜"), Special: models.PtrBool(false), Alcohol: models.PtrBool(false), Custom: models.PtrBool(false), Variant: models.PtrString(""), VariantDefault: models.PtrBool(false), VariantPriceCharge: models.PtrFloat64(0)},
	{ID: models.PtrInt(6), MenuID: models.PtrInt(5), CategoryID: models.PtrInt(5), Price: models.PtrFloat64(12.99), NameEng: models.PtrString("Salt and Pepper Shrimp"), NameOth: models.PtrString("椒鹽蝦"), Special: models.PtrBool(false), Alcohol: models.PtrBool(false), Custom: models.PtrBool(false), Variant: models.PtrString(""), VariantDefault: models.PtrBool(false), VariantPriceCharge: models.PtrFloat64(0)},
	{ID: models.PtrInt(7), MenuID: models.PtrInt(6), CategoryID: models.PtrInt(6), Price: models.PtrFloat64(13.99), NameEng: models.PtrString("Stir-fried Scallops"), NameOth: models.PtrString("炒帶子"), Special: models.PtrBool(false), Alcohol: models.PtrBool(false), Custom: models.PtrBool(false), Variant: models.PtrString(""), VariantDefault: models.PtrBool(false), VariantPriceCharge: models.PtrFloat64(0)},
	{ID: models.PtrInt(8), MenuID: models.PtrInt(7), CategoryID: models.PtrInt(7), Price: models.PtrFloat64(14.99), NameEng: models.PtrString("Beef Hot Pot"), NameOth: models.PtrString("牛肉煲"), Special: models.PtrBool(false), Alcohol: models.PtrBool(false), Custom: models.PtrBool(false), Variant: models.PtrString(""), VariantDefault: models.PtrBool(false), VariantPriceCharge: models.PtrFloat64(0)},
	{ID: models.PtrInt(9), MenuID: models.PtrInt(8), CategoryID: models.PtrInt(8), Price: models.PtrFloat64(9.99), NameEng: models.PtrString("Sweet and Sour Pork"), NameOth: models.PtrString("糖醋排骨"), Special: models.PtrBool(false), Alcohol: models.PtrBool(false), Custom: models.PtrBool(false), Variant: models.PtrString(""), VariantDefault: models.PtrBool(false), VariantPriceCharge: models.PtrFloat64(0)},
	{ID: models.PtrInt(10), MenuID: models.PtrInt(9), CategoryID: models.PtrInt(9), Price: models.PtrFloat64(10.99), NameEng: models.PtrString("Beef with Broccoli"), NameOth: models.PtrString("西蘭花牛肉"), Special: models.PtrBool(false), Alcohol: models.PtrBool(false), Custom: models.PtrBool(false), Variant: models.PtrString(""), VariantDefault: models.PtrBool(false), VariantPriceCharge: models.PtrFloat64(0)},
	{ID: models.PtrInt(11), MenuID: models.PtrInt(10), CategoryID: models.PtrInt(10), Price: models.PtrFloat64(8.99), NameEng: models.PtrString("Kung Pao Chicken"), NameOth: models.PtrString("宫保鸡丁"), Special: models.PtrBool(false), Alcohol: models.PtrBool(false), Custom: models.PtrBool(false), Variant: models.PtrString(""), VariantDefault: models.PtrBool(false), VariantPriceCharge: models.PtrFloat64(0)},
	{ID: models.PtrInt(12), MenuID: models.PtrInt(11), CategoryID: models.PtrInt(11), Price: models.PtrFloat64(7.99), NameEng: models.PtrString("BBQ Pork Over Rice"), NameOth: models.PtrString("叉燒飯"), Special: models.PtrBool(false), Alcohol: models.PtrBool(false), Custom: models.PtrBool(false), Variant: models.PtrString(""), VariantDefault: models.PtrBool(false), VariantPriceCharge: models.PtrFloat64(0)},
	{ID: models.PtrInt(13), MenuID: models.PtrInt(12), CategoryID: models.PtrInt(12), Price: models.PtrFloat64(8.99), NameEng: models.PtrString("Yangzhou Fried Rice"), NameOth: models.PtrString("扬州炒饭"), Special: models.PtrBool(false), Alcohol: models.PtrBool(false), Custom: models.PtrBool(false), Variant: models.PtrString(""), VariantDefault: models.PtrBool(false), VariantPriceCharge: models.PtrFloat64(0)},
	{ID: models.PtrInt(14), MenuID: models.PtrInt(13), CategoryID: models.PtrInt(13), Price: models.PtrFloat64(9.99), NameEng: models.PtrString("Chicken Chow Mein"), NameOth: models.PtrString("雞肉炒麵"), Special: models.PtrBool(false), Alcohol: models.PtrBool(false), Custom: models.PtrBool(false), Variant: models.PtrString(""), VariantDefault: models.PtrBool(false), VariantPriceCharge: models.PtrFloat64(0)},
	{ID: models.PtrInt(15), MenuID: models.PtrInt(14), CategoryID: models.PtrInt(14), Price: models.PtrFloat64(6.99), NameEng: models.PtrString("Beef Noodle Soup"), NameOth: models.PtrString("牛肉湯麵"), Special: models.PtrBool(false), Alcohol: models.PtrBool(false), Custom: models.PtrBool(false), Variant: models.PtrString(""), VariantDefault: models.PtrBool(false), VariantPriceCharge: models.PtrFloat64(0)},
	{ID: models.PtrInt(16), MenuID: models.PtrInt(15), CategoryID: models.PtrInt(15), Price: models.PtrFloat64(5.99), NameEng: models.PtrString("Pork Congee"), NameOth: models.PtrString("豬肉粥"), Special: models.PtrBool(false), Alcohol: models.PtrBool(false), Custom: models.PtrBool(false), Variant: models.PtrString(""), VariantDefault: models.PtrBool(false), VariantPriceCharge: models.PtrFloat64(0)},
	{ID: models.PtrInt(17), MenuID: models.PtrInt(16), CategoryID: models.PtrInt(16), Price: models.PtrFloat64(15.99), NameEng: models.PtrString("General Tso's Chicken"), NameOth: models.PtrString("左宗棠雞"), Special: models.PtrBool(false), Alcohol: models.PtrBool(false), Custom: models.PtrBool(false), Variant: models.PtrString(""), VariantDefault: models.PtrBool(false), VariantPriceCharge: models.PtrFloat64(0)},
	{ID: models.PtrInt(18), MenuID: models.PtrInt(17), CategoryID: models.PtrInt(17), Price: models.PtrFloat64(2.99), NameEng: models.PtrString("Local Beer"), NameOth: models.PtrString("本地啤酒"), Special: models.PtrBool(false), Alcohol: models.PtrBool(true), Custom: models.PtrBool(false), Variant: models.PtrString(""), VariantDefault: models.PtrBool(false), VariantPriceCharge: models.PtrFloat64(0)},
}

func MockGetItemQuery(mock pgxmock.PgxConnIface, id int) {
	rows := mock.NewRows([]string{"id", "menu_id", "category_id", "price", "name_eng", "name_oth", "special", "alcohol", "custom", "variant", "variant_default", "variant_price_charge"})
	if id > 0 && id <= len(mockItems) {
		item := mockItems[id-1]
		rows.AddRow(item.ID, item.MenuID, item.CategoryID, item.Price, item.NameEng, item.NameOth, item.Special, item.Alcohol, item.Custom, item.Variant, item.VariantDefault, item.VariantPriceCharge)

	} else {
		fmt.Println("MockGetItemQuery: Invalid id", id, "for mockItems. There are only", len(mockItems), "items.")
	}

	mock.ExpectQuery("SELECT id, menu_id, category_id, price, name_eng, name_oth, special, alcohol, custom, variant, variant_default, variant_price_charge FROM items WHERE id = \\$1").WithArgs(id).WillReturnRows(rows)
}

func MockGetItemsQuery(mock pgxmock.PgxConnIface) {
	rows := mock.NewRows([]string{"id", "menu_id", "category_id", "price", "name_eng", "name_oth", "special", "alcohol", "custom", "variant", "variant_default", "variant_price_charge"})

	for _, item := range mockItems {
		rows.AddRow(item.ID, item.MenuID, item.CategoryID, item.Price, item.NameEng, item.NameOth, item.Special, item.Alcohol, item.Custom, item.Variant, item.VariantDefault, item.VariantPriceCharge)
	}
	mock.ExpectQuery("SELECT id, menu_id, category_id, price, name_eng, name_oth, special, alcohol, custom, variant, variant_default, variant_price_charge FROM items ORDER BY id ASC").
		WillReturnRows(rows)
}

func MockInsertItemQuery(mock pgxmock.PgxConnIface, item models.Item) {
	mock.ExpectQuery("INSERT INTO items \\(menu_id, category_id, price, name_eng, name_oth, special, alcohol, custom, variant, variant_default, variant_price_charge\\) VALUES \\(\\$1, \\$2, \\$3, \\$4, \\$5, \\$6, \\$7, \\$8, \\$9, \\$10, \\$11\\) RETURNING id, menu_id, category_id, price, name_eng, name_oth, special, alcohol, custom, variant, variant_default, variant_price_charge").
		WithArgs(*item.MenuID, *item.CategoryID, *item.Price, *item.NameEng, *item.NameOth, *item.Special, *item.Alcohol, *item.Custom, *item.Variant, *item.VariantDefault, *item.VariantPriceCharge).
		WillReturnRows(mock.NewRows([]string{"id", "menu_id", "category_id", "price", "name_eng", "name_oth", "special", "alcohol", "custom", "variant", "variant_default", "variant_price_charge"}).
			AddRow(item.ID, item.MenuID, item.CategoryID, item.Price, item.NameEng, item.NameOth, item.Special, item.Alcohol, item.Custom, item.Variant, item.VariantDefault, item.VariantPriceCharge))
}

func MockUpdateItemQuery(mock pgxmock.PgxConnIface, item models.Item) {
	mock.ExpectQuery("UPDATE items SET menu_id = \\$1, category_id = \\$2, price = \\$3, name_eng = \\$4, name_oth = \\$5, special = \\$6, alcohol = \\$7, custom = \\$8, variant = \\$9, variant_default = \\$10, variant_price_charge = \\$11 WHERE id = \\$12 RETURNING id, menu_id, category_id, price, name_eng, name_oth, special, alcohol, custom, variant, variant_default, variant_price_charge").
		WithArgs(*item.MenuID, *item.CategoryID, *item.Price, *item.NameEng, *item.NameOth, *item.Special, *item.Alcohol, *item.Custom, *item.Variant, *item.VariantDefault, *item.VariantPriceCharge, *item.ID).
		WillReturnRows(mock.NewRows([]string{"id", "menu_id", "category_id", "price", "name_eng", "name_oth", "special", "alcohol", "custom", "variant", "variant_default", "variant_price_charge"}).
			AddRow(item.ID, item.MenuID, item.CategoryID, item.Price, item.NameEng, item.NameOth, item.Special, item.Alcohol, item.Custom, item.Variant, item.VariantDefault, item.VariantPriceCharge))
}

var mockCustomers = []models.Customer{
	{
		ID:           models.PtrInt(1),
		Name:         models.PtrString("John Doe"),
		Phone:        models.PtrString("604-123-1234"),
		UnitNumber:   models.PtrString(""),
		StreetNumber: models.PtrString("5555"),
		StreetName:   models.PtrString("Powel St"),
		BuzzerNumber: models.PtrString(""),
		Note:         models.PtrString(""),
	},
	{
		ID:           models.PtrInt(2),
		Name:         models.PtrString("Christine StClaire"),
		Phone:        models.PtrString("123-456-7890"),
		UnitNumber:   models.PtrString("A12"),
		StreetNumber: models.PtrString("123"),
		StreetName:   models.PtrString("Maple St"),
		BuzzerNumber: models.PtrString("A12"),
		Note:         models.PtrString("good tips"),
	},
	{
		ID:           models.PtrInt(3),
		Name:         models.PtrString("David Hogan"),
		Phone:        models.PtrString("778-123-1234"),
		UnitNumber:   models.PtrString("BSM"),
		StreetNumber: models.PtrString("5555"),
		StreetName:   models.PtrString("Powel St"),
		BuzzerNumber: models.PtrString(""),
		Note:         models.PtrString(""),
	},
}

func MockGetCustomerQuery(mock pgxmock.PgxConnIface, id int) {
	rows := mock.NewRows([]string{"id", "name", "phone", "unit_number", "street_number", "street_name", "buzzer_number", "note"})
	if id > 0 && id <= len(mockCustomers) {
		customer := mockCustomers[id-1]
		rows.AddRow(customer.ID, customer.Name, customer.Phone, customer.UnitNumber, customer.StreetNumber, customer.StreetName, customer.BuzzerNumber, customer.Note)
	} else {
		fmt.Println("MockGetCustomerQuery: Invalid id", id, "for mockCustomers. There are only", len(mockCustomers), "customers.")
	}
	mock.ExpectQuery("SELECT id, name, phone, unit_number, street_number, street_name, buzzer_number, note FROM customers WHERE id = \\$1").WithArgs(id).WillReturnRows(rows)
}

func MockGetCustomersQuery(mock pgxmock.PgxConnIface) {
	rows := mock.NewRows([]string{"id", "name", "phone", "unit_number", "street_number", "street_name", "buzzer_number", "note"})
	for _, customer := range mockCustomers {
		rows.AddRow(customer.ID, customer.Name, customer.Phone, customer.UnitNumber, customer.StreetNumber, customer.StreetName, customer.BuzzerNumber, customer.Note)
	}
	mock.ExpectQuery("SELECT id, name, phone, unit_number, street_number, street_name, buzzer_number, note FROM customers ORDER BY id ASC").
		WillReturnRows(rows)
}

func MockInsertCustomerQuery(mock pgxmock.PgxConnIface, customer models.Customer) {
	mock.ExpectQuery("INSERT INTO customers \\(name, phone, unit_number, street_number, street_name, buzzer_number, note\\) VALUES \\(\\$1, \\$2, \\$3, \\$4, \\$5, \\$6, \\$7\\) RETURNING id, name, phone, unit_number, street_number, street_name, buzzer_number, note").
		WithArgs(*customer.Name, *customer.Phone, *customer.UnitNumber, *customer.StreetNumber, *customer.StreetName, *customer.BuzzerNumber, *customer.Note).
		WillReturnRows(mock.NewRows([]string{"id", "name", "phone", "unit_number", "street_number", "street_name", "buzzer_number", "note"}).
			AddRow(customer.ID, customer.Name, customer.Phone, customer.UnitNumber, customer.StreetNumber, customer.StreetName, customer.BuzzerNumber, customer.Note))
}

func MockUpdateCustomerQuery(mock pgxmock.PgxConnIface, customer models.Customer) {
	mock.ExpectQuery("UPDATE customers SET name = \\$1, phone = \\$2, unit_number = \\$3, street_number = \\$4, street_name = \\$5, buzzer_number = \\$6, note = \\$7 WHERE id = \\$8 RETURNING id, name, phone, unit_number, street_number, street_name, buzzer_number, note").
		WithArgs(*customer.Name, *customer.Phone, *customer.UnitNumber, *customer.StreetNumber, *customer.StreetName, *customer.BuzzerNumber, *customer.Note, *customer.ID).
		WillReturnRows(mock.NewRows([]string{"id", "name", "phone", "unit_number", "street_number", "street_name", "buzzer_number", "note"}).
			AddRow(customer.ID, customer.Name, customer.Phone, customer.UnitNumber, customer.StreetNumber, customer.StreetName, customer.BuzzerNumber, customer.Note))
}
