package contact

import "github.com/shinochiha/gotoko/app"

// Contact is the main model of Contact data. It provides a convenient interface for app.ModelInterface
type Contact struct {
	app.Model
	ID          app.NullUUID     `json:"id"           db:"m.id"              gorm:"column:id;primaryKey"`
	Name        app.NullString   `json:"name"         db:"m.name"            gorm:"column:name"`
	Code        app.NullString   `json:"code"         db:"m.code"            gorm:"column:code"`
	Email       app.NullString   `json:"email"        db:"m.email"           gorm:"column:email"`
	PhoneNumber app.NullString   `json:"phone_number" db:"m.phone_number"    gorm:"column:phone_number"`
	IsActive    app.NullBool     `json:"is_active"    db:"m.is_active"       gorm:"column:is_active"`
	CreatedAt   app.NullDateTime `json:"created_at"   db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt   app.NullDateTime `json:"updated_at"   db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt   app.NullDateTime `json:"deleted_at"   db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

// EndPoint returns the Contact end point, it used for cache key, etc.
func (Contact) EndPoint() string {
	return "contacts"
}

// TableVersion returns the versions of the Contact table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (Contact) TableVersion() string {
	return "28.06.291152"
}

// TableName returns the name of the Contact table in the database.
func (Contact) TableName() string {
	return "contacts"
}

// TableAliasName returns the table alias name of the Contact table, used for querying.
func (Contact) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the Contact data in the database, used for querying.
func (m *Contact) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the Contact data in the database, used for querying.
func (m *Contact) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the Contact data in the database, used for querying.
func (m *Contact) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the Contact data in the database, used for querying.
func (m *Contact) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the Contact schema, used for querying.
func (m *Contact) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the Contact schema in the open api documentation.
func (Contact) OpenAPISchemaName() string {
	return "Contact"
}

// GetOpenAPISchema returns the Open API Schema of the Contact in the open api documentation.
func (m *Contact) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type ContactList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the ContactList schema in the open api documentation.
func (ContactList) OpenAPISchemaName() string {
	return "ContactList"
}

// GetOpenAPISchema returns the Open API Schema of the ContactList in the open api documentation.
func (p *ContactList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&Contact{})
}

// ParamCreate is the expected parameters for create a new Contact data.
type ParamCreate struct {
	UseCaseHandler
}

// ParamUpdate is the expected parameters for update the Contact data.
type ParamUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamPartiallyUpdate is the expected parameters for partially update the Contact data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamDelete is the expected parameters for delete the Contact data.
type ParamDelete struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}
