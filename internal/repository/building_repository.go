package repository

import (
	"database/sql"
	"fmt"
	"leadgen-test-task/internal/model"
	"strings"
)

type BuildingRepository struct {
	db *sql.DB
}

func NewBuildingRepository(db *sql.DB) *BuildingRepository {
	return &BuildingRepository{db: db}
}

func (r *BuildingRepository) Create(b *model.Building) error {
	query := `
    INSERT INTO buildings (name, city, year, floor_count)
    VALUES ($1, $2, $3, $4) RETURNING id`

	err := r.db.QueryRow(query, b.Name, b.City, b.Year, b.FloorCount).Scan(&b.ID)
	return err
}

func (r *BuildingRepository) FindAll(city string, year, floorCount int) ([]model.Building, error) {
	query := `
    SELECT id, name, city, year, floor_count FROM buildings
    WHERE 1=1
    %s
    ORDER BY id`

	var args []interface{}
	var conditions []string
	argCounter := 1

	if city != "" {
		conditions = append(conditions, fmt.Sprintf("AND city = $%d", argCounter))
		args = append(args, city)
		argCounter++
	}

	if year > 0 {
		conditions = append(conditions, fmt.Sprintf("AND year = $%d", argCounter))
		args = append(args, year)
		argCounter++
	}

	if floorCount > 0 {
		conditions = append(conditions, fmt.Sprintf("AND floor_count = $%d", argCounter))
		args = append(args, floorCount)
		argCounter++
	}

	query = fmt.Sprintf(query, strings.Join(conditions, " "))

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var buildings []model.Building
	for rows.Next() {
		var b model.Building
		if err := rows.Scan(&b.ID, &b.Name, &b.City, &b.Year, &b.FloorCount); err != nil {
			return nil, err
		}
		buildings = append(buildings, b)
	}

	return buildings, nil
}
