package repositories

import (
	"context"
	"errors"
	"fmt"
	"server/db"
	"server/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

var SelectAllItemsQuery = "SELECT id, menu_id, category_id, price, name_eng, name_oth, special, alcohol, custom, variant, variant_default, variant_price_charge FROM items ORDER BY id ASC"

func QueryAllItems() ([]models.Item, error) {
	rows, err := db.DBClient.Query(context.Background(), SelectAllItemsQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to query all items - Error: %v", err)
	}
	defer rows.Close()

	items, err := scanItems(rows)
	if err != nil {
		return nil, fmt.Errorf("failed to scan items: %w", err)
	}
	return items, nil
}

func scanItems(rows pgx.Rows) ([]models.Item, error) {
	defer rows.Close()
	var items []models.Item

	for rows.Next() {
		var item models.Item
		err := rows.Scan(
			&item.ID, &item.MenuID, &item.CategoryID, &item.Price, &item.NameEng, &item.NameOth,
			&item.Special, &item.Alcohol, &item.Custom, &item.Variant, &item.VariantDefault, &item.VariantPriceCharge,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan items row: %w", err)
		}
		items = append(items, item)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return items, nil
}

func InsertItem(item *models.Item) ([]models.Item, error) {
	query := "INSERT INTO items (menu_id, category_id, price, name_eng, name_oth, special, alcohol, custom, variant, variant_default, variant_price_charge) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id, menu_id, category_id, price, name_eng, name_oth, special, alcohol, custom, variant, variant_default, variant_price_charge"
	rows, err := db.DBClient.Query(context.Background(), query, *item.MenuID, *item.CategoryID, *item.Price, *item.NameEng, *item.NameOth, *item.Special, *item.Alcohol, *item.Custom, *item.Variant, *item.VariantDefault, *item.VariantPriceCharge)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23503":
				return nil, fmt.Errorf("failed to insert item: category ID %v does not exist (error code: %s)", item.CategoryID, pgErr.Code)
			default:
				return nil, fmt.Errorf("failed to insert item: %v (error code: %s)", pgErr.Message, pgErr.Code)
			}
		}
		return nil, fmt.Errorf("failed to insert item: %v", err)
	}
	defer rows.Close()
	items, err := scanItems(rows)
	if err != nil {
		return nil, fmt.Errorf("failed to scan items: %w", err)
	}
	return items, nil
}

func UpdateItem(itemInput *models.Item) ([]models.Item, error) {
	item, err := QueryItem(*itemInput.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to query item: %v", err)
	}
	updateItemFields(item, itemInput)
	query := "UPDATE items SET menu_id = $1, category_id = $2, price = $3, name_eng = $4, name_oth = $5, special = $6, alcohol = $7, custom = $8, variant = $9, variant_default = $10, variant_price_charge = $11 WHERE id = $12 RETURNING id, menu_id, category_id, price, name_eng, name_oth, special, alcohol, custom, variant, variant_default, variant_price_charge"
	rows, err := db.DBClient.Query(context.Background(), query, *item.MenuID, *item.CategoryID, *item.Price, *item.NameEng, *item.NameOth, *item.Special, *item.Alcohol, *item.Custom, *item.Variant, *item.VariantDefault, *item.VariantPriceCharge, *item.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to update item: %v", err)
	}
	defer rows.Close()
	items, err := scanItems(rows)
	if err != nil {
		return nil, fmt.Errorf("failed to scan items: %w", err)
	}
	return items, nil
}

func updateItemFields(item *models.Item, itemInput *models.Item) {
	if itemInput.ID != nil {
		item.ID = itemInput.ID
	}
	if itemInput.MenuID != nil {
		item.MenuID = itemInput.MenuID
	}
	if itemInput.CategoryID != nil {
		item.CategoryID = itemInput.CategoryID
	}
	if itemInput.Price != nil {
		item.Price = itemInput.Price
	}
	if itemInput.NameEng != nil {
		item.NameEng = itemInput.NameEng
	}
	if itemInput.NameOth != nil {
		item.NameOth = itemInput.NameOth
	}
	if itemInput.Special != nil {
		item.Special = itemInput.Special
	}
	if itemInput.Alcohol != nil {
		item.Alcohol = itemInput.Alcohol
	}
	if itemInput.Custom != nil {
		item.Custom = itemInput.Custom
	}
	if itemInput.Variant != nil {
		item.Variant = itemInput.Variant
	}
	if itemInput.VariantDefault != nil {
		item.VariantDefault = itemInput.VariantDefault
	}
	if itemInput.VariantPriceCharge != nil {
		item.VariantPriceCharge = itemInput.VariantPriceCharge
	}
}

func QueryItem(id int) (*models.Item, error) {
	query := "SELECT id, menu_id, category_id, price, name_eng, name_oth, special, alcohol, custom, variant, variant_default, variant_price_charge FROM items WHERE id = $1"
	rows, err := db.DBClient.Query(context.Background(), query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to query item: %v", err)
	}
	defer rows.Close()
	items, err := scanItems(rows)
	if err != nil {
		return nil, fmt.Errorf("failed to scan items: %w", err)
	}
	if len(items) == 0 {
		return nil, nil
	}
	return &items[0], nil
}
