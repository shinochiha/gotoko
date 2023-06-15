package callback

import "github.com/shinochiha/gotoko/app"

// Callback is the main model of Callback data. It provides a convenient interface for app.ModelInterface
type Callback struct {
	app.Model
	ID        app.NullUUID     `json:"id"         db:"m.id"              gorm:"column:id;primaryKey"`
	CreatedAt app.NullDateTime `json:"created_at" db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt app.NullDateTime `json:"updated_at" db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt app.NullDateTime `json:"deleted_at" db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

// EndPoint returns the Callback end point, it used for cache key, etc.
func (Callback) EndPoint() string {
	return "callback"
}

// TableVersion returns the versions of the Callback table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (Callback) TableVersion() string {
	return "28.06.291152"
}

// TableName returns the name of the Callback table in the database.
func (Callback) TableName() string {
	return "callback"
}

// TableAliasName returns the table alias name of the Callback table, used for querying.
func (Callback) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the Callback data in the database, used for querying.
func (m *Callback) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the Callback data in the database, used for querying.
func (m *Callback) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the Callback data in the database, used for querying.
func (m *Callback) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the Callback data in the database, used for querying.
func (m *Callback) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the Callback schema, used for querying.
func (m *Callback) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the Callback schema in the open api documentation.
func (Callback) OpenAPISchemaName() string {
	return "Callback"
}

// GetOpenAPISchema returns the Open API Schema of the Callback in the open api documentation.
func (m *Callback) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type CallbackList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the CallbackList schema in the open api documentation.
func (CallbackList) OpenAPISchemaName() string {
	return "CallbackList"
}

// GetOpenAPISchema returns the Open API Schema of the CallbackList in the open api documentation.
func (p *CallbackList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&Callback{})
}

// ParamCreate is the expected parameters for create a new Callback data.
type ParamCreate struct {
	UseCaseHandler
}

// ParamUpdate is the expected parameters for update the Callback data.
type ParamUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamPartiallyUpdate is the expected parameters for partially update the Callback data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamDelete is the expected parameters for delete the Callback data.
type ParamDelete struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}
