package categorymodel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

func GetAll() []entities.Category {
	rows, err := config.DB.Query(`SELECT * FROM categories`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var categories []entities.Category
	for rows.Next() {
		var category entities.Category
		err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			panic(err)
		}

		categories = append(categories, category)
	}

	return categories
}

func Create(category entities.Category) bool {
	result, err := config.DB.Exec(`
		INSERT INTO categories (name, created_at, updated_at)
		VALUES (?, ?, ?)`,
		category.Name, category.CreatedAt, category.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

func Detil(id int) entities.Category {
	row := config.DB.QueryRow(`SELECT id, name FROM categories WHERE id = ?`, id)

	var categories entities.Category
	if err := row.Scan(&categories.Id, &categories.Name); err != nil {
		panic(err)
	}

	return categories
}

func Update(id int, category entities.Category) bool {
	query, err := config.DB.Exec(`
		UPDATE categories Set name = ?, updated_at = ? where id = ?`,
		category.Name, category.UpdatedAt, id)

	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec(`DELETE FROM categories WHERE id = ?`, id)
	return err
}
