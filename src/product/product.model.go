package product

import "github.com/shinochiha/gotoko/app"

// Product is the main model of Product data. It provides a convenient interface for app.ModelInterface
type Product struct {
	app.Model
	ID        app.NullUUID     `json:"id"         db:"m.id"              gorm:"column:id;primaryKey"`
	Name      app.NullString   `json:"name"       db:"m.name"            gorm:"column:name"`
	Code      app.NullString   `json:"code"       db:"m.code"            gorm:"column:code"`
	Quantity  app.NullFloat64  `json:"quantity"   db:"m.quantity"        gorm:"column:quantity"`
	IsActive  app.NullBool     `json:"is_active"  db:"m.is_active"       gorm:"column:is_active"`
	CreatedAt app.NullDateTime `json:"created_at" db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt app.NullDateTime `json:"updated_at" db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt app.NullDateTime `json:"deleted_at" db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

// EndPoint returns the Product end point, it used for cache key, etc.
func (Product) EndPoint() string {
	return "products"
}

// TableVersion returns the versions of the Product table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (Product) TableVersion() string {
	return "28.06.291152"
}

// TableName returns the name of the Product table in the database.
func (Product) TableName() string {
	return "products"
}

// TableAliasName returns the table alias name of the Product table, used for querying.
func (Product) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the Product data in the database, used for querying.
func (m *Product) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the Product data in the database, used for querying.
func (m *Product) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the Product data in the database, used for querying.
func (m *Product) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the Product data in the database, used for querying.
func (m *Product) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the Product schema, used for querying.
func (m *Product) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the Product schema in the open api documentation.
func (Product) OpenAPISchemaName() string {
	return "Product"
}

// GetOpenAPISchema returns the Open API Schema of the Product in the open api documentation.
func (m *Product) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type ProductList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the ProductList schema in the open api documentation.
func (ProductList) OpenAPISchemaName() string {
	return "ProductList"
}

// GetOpenAPISchema returns the Open API Schema of the ProductList in the open api documentation.
func (p *ProductList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&Product{})
}

// ParamCreate is the expected parameters for create a new Product data.
type ParamCreate struct {
	UseCaseHandler
}

// ParamUpdate is the expected parameters for update the Product data.
type ParamUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamPartiallyUpdate is the expected parameters for partially update the Product data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamDelete is the expected parameters for delete the Product data.
type ParamDelete struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}
