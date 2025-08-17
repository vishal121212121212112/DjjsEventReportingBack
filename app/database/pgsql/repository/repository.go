package database

import (
	connection "event-reporting/app/database/pgsql/connection"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository() *Repository {
	if connection.Db == nil {
		panic("Database connection is not initialized")
	}
	return &Repository{DB: connection.Db}
}

func (r *Repository) Find(model interface{}, conditions map[string]interface{}) error {
	query := connection.Db.Where(conditions)
	if err := query.First(model).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) Create(model interface{}) error {
	if err := connection.Db.Create(model).Error; err != nil {
		return fmt.Errorf("error creating data: %w", err)
	}
	return nil
}

func (r *Repository) Update(model interface{}) error {
	if err := connection.Db.Save(model).Error; err != nil {
		return fmt.Errorf("error updating data: %w", err)
	}
	return nil
}

func (r *Repository) Delete(model interface{}) error {
	if err := connection.Db.Delete(model).Error; err != nil {
		return fmt.Errorf("error deleting data: %w", err)
	}
	return nil
}
func (r *Repository) DeleteWhere(model interface{}, conditions map[string]interface{}) error {
	if err := connection.Db.Where(conditions).Delete(model).Error; err != nil {
		return fmt.Errorf("error deleting data: %w", err)
	}
	return nil
}

func (r *Repository) Count(model interface{}, conditions map[string]interface{}) (int64, error) {
	var count int64
	query := connection.Db.Model(model).Where(conditions)
	if err := query.Count(&count).Error; err != nil {
		return 0, fmt.Errorf("error counting data: %w", err)
	}
	fmt.Println(count)
	return count, nil
}
func (r *Repository) CountBetween(model interface{}, conditions map[string]interface{}, column string, start, end time.Time) (int64, error) {
	db := connection.Db.Model(model)

	for k, v := range conditions {
		switch vv := v.(type) {
		case []string:
			db = db.Where(fmt.Sprintf("%s IN ?", k), vv)
		default:
			db = db.Where(fmt.Sprintf("%s = ?", k), v)
		}
	}

	db = db.Where(fmt.Sprintf("%s >= ? AND %s < ?", column, column), start, end)

	var count int64
	if err := db.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *Repository) FindAll(model interface{}, conditions map[string]interface{}, limit, offset int) error {
	query := connection.Db.Where(conditions)
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	if err := query.Find(model).Error; err != nil {
		return fmt.Errorf("error fetching data: %w", err)
	}
	return nil
}

// UpdateFields updates specific fields in the record matching conditions.
func (r *Repository) UpdateFields(model interface{}, conditions map[string]interface{}, updates map[string]interface{}) error {
	return connection.Db.Model(model).Where(conditions).Updates(updates).Error
}

func (r *Repository) FindFieldValues(model interface{}, conditions map[string]interface{}, fieldName string, pageSize, offset int) ([]string, error) {
	var result []string

	query := connection.Db.Model(model).Select(fieldName).Where(conditions)

	if pageSize > 0 {
		query = query.Limit(pageSize)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	if err := query.Pluck(fieldName, &result).Error; err != nil {
		return nil, fmt.Errorf("error fetching %s values: %w", fieldName, err)
	}

	return result, nil
}

func (r *Repository) GetAllAndGroupBy(model interface{}, conditions map[string]interface{}, groupByField string) (map[string]int, error) {
	type Result struct {
		GroupField string `gorm:"column:group_field"`
		Count      int    `gorm:"column:count"`
	}

	var results []Result

	err := connection.Db.Model(model).
		Select(fmt.Sprintf("%s as group_field, COUNT(*) as count", groupByField)).
		Where(conditions).
		Group(groupByField).
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	groupedCounts := make(map[string]int)
	for _, res := range results {
		groupedCounts[res.GroupField] = res.Count
	}

	return groupedCounts, nil
}

func (r *Repository) GetAllAndGroupByWithRawCondition(
	model interface{},
	conditions map[string]interface{},
	rawCondition string,
	args []interface{},
	groupByField string,
) (map[string]int, error) {
	type Result struct {
		GroupField string `gorm:"column:group_field"`
		Count      int    `gorm:"column:count"`
	}
	var results []Result

	db := connection.Db.Model(model).Select(fmt.Sprintf("%s as group_field, COUNT(*) as count", groupByField))

	if len(conditions) > 0 {
		db = db.Where(conditions)
	}

	if rawCondition != "" {
		db = db.Where(rawCondition, args...)
	}

	err := db.Group(groupByField).Scan(&results).Error
	if err != nil {
		return nil, err
	}

	groupedCounts := make(map[string]int)
	for _, res := range results {
		groupedCounts[res.GroupField] = res.Count
	}

	return groupedCounts, nil
}
func (r *Repository) FindFieldValuesWithRawCondition(
	model interface{},
	conditions map[string]interface{},
	rawCondition string,
	args []interface{},
	fieldName string,
	pageSize, offset int,
) ([]string, error) {
	var result []string

	db := connection.Db.Model(model).Select(fieldName)

	if len(conditions) > 0 {
		db = db.Where(conditions)
	}

	if rawCondition != "" {
		db = db.Where(rawCondition, args...)
	}

	if pageSize > 0 {
		db = db.Limit(pageSize)
	}

	if offset > 0 {
		db = db.Offset(offset)
	}

	if err := db.Pluck(fieldName, &result).Error; err != nil {
		return nil, fmt.Errorf("error fetching %s values: %w", fieldName, err)
	}

	return result, nil
}
func (r *Repository) CountWithRawCondition(model interface{}, conditions map[string]interface{}, rawCondition string, args []interface{}) (int64, error) {
	var count int64
	query := connection.Db.Model(model).Where(conditions)
	if rawCondition != "" {
		query = query.Where(rawCondition, args...)
	}
	err := query.Count(&count).Error
	return count, err
}

func (r *Repository) FindAllWithRawCondition(model interface{}, conditions map[string]interface{}, rawCondition string, args []interface{}, limit, offset int) error {
	query := connection.Db.Model(model).Where(conditions)
	if rawCondition != "" {
		query = query.Where(rawCondition, args...)
	}
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	return query.Find(model).Error
}

// ExecuteRawFunction executes a PostgreSQL function and scans results into the provided model
func (r *Repository) ExecuteRawFunction(functionName string, args []interface{}, result interface{}) error {
	// Build the function call with parameters
	query := fmt.Sprintf("SELECT * FROM %s(", functionName)

	// Add placeholders for parameters
	placeholders := make([]string, len(args))
	for i := range args {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}
	query += strings.Join(placeholders, ", ") + ")"

	// Execute the raw SQL
	return connection.Db.Raw(query, args...).Scan(result).Error
}

// FindAllWithRawConditionAndOrder fetches rows with optional raw WHERE, ORDER BY, LIMIT/OFFSET.
func (r *Repository) FindAllWithRawConditionAndOrder(
	model interface{},
	conditions map[string]interface{},
	rawCondition string,
	args []interface{},
	orderBy string, // e.g. "name ASC"
	limit, offset int,
) error {
	query := connection.Db.Model(model).Where(conditions)
	if rawCondition != "" {
		query = query.Where(rawCondition, args...)
	}
	if orderBy != "" {
		query = query.Order(orderBy)
	}
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	return query.Find(model).Error
}
