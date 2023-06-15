package paymentsubscription

import "github.com/shinochiha/gotoko/app"

// PaymentSubscription is the main model of PaymentSubscription data. It provides a convenient interface for app.ModelInterface
type PaymentSubscription struct {
	app.Model
	ID                   app.NullUUID     `json:"id"                 db:"m.id"                 gorm:"column:id;primaryKey"`
	Name                 app.NullString   `json:"name"               db:"m.name"               gorm:"column:name"`
	Description          app.NullString   `json:"description"        db:"m.description"        gorm:"column:description"`
	Amount               app.NullFloat64  `json:"amount"             db:"m.amount"             gorm:"column:amount"`
	UserId               app.NullString   `json:"user_id"            db:"m.user_id"            gorm:"column:user_id"`
	Currency             app.NullString   `json:"currency"           db:"m.currency"           gorm:"column:currency"`
	TotalRecurrence      app.NullInt64    `json:"total_recurrence"   db:"m.total_recurrence"   gorm:"column:total_recurrence"`
	CardToken            app.NullString   `json:"card_token"         db:"m.card_token"         gorm:"column:card_token"`
	ChargeToken          app.NullString   `json:"charge_token"       db:"m.charge_token"       gorm:"column:charge_token"`
	ChargeImmediately    app.NullBool     `json:"charge_immediately" db:"m.charge_immediately" gorm:"column:charge_immediately"`
	ScheduleInterval     app.NullInt64    `json:"schedule.interval"           db:"m.schedule_interval"           gorm:"column:schedule_interval"`
	ScheduleIntervalUnit app.NullString   `json:"schedule.interval_unit"      db:"m.schedule_interval_unit"      gorm:"column:schedule_interval_unit"`
	ScheduleStartAt      app.NullDateTime `json:"schedule.start_at"           db:"m.schedule_start_at"           gorm:"column:schedule_start_at"`
	CreatedAt            app.NullDateTime `json:"created_at"         db:"m.created_at"         gorm:"column:created_at"`
	UpdatedAt            app.NullDateTime `json:"updated_at"         db:"m.updated_at"         gorm:"column:updated_at"`
	DeletedAt            app.NullDateTime `json:"deleted_at"         db:"m.deleted_at,hide"    gorm:"column:deleted_at"`
}

// EndPoint returns the PaymentSubscription end point, it used for cache key, etc.
func (PaymentSubscription) EndPoint() string {
	return "payment_subscriptions"
}

// TableVersion returns the versions of the PaymentSubscription table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (PaymentSubscription) TableVersion() string {
	return "28.06.291152"
}

// TableName returns the name of the PaymentSubscription table in the database.
func (PaymentSubscription) TableName() string {
	return "payment_subscriptions"
}

// TableAliasName returns the table alias name of the PaymentSubscription table, used for querying.
func (PaymentSubscription) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the PaymentSubscription data in the database, used for querying.
func (m *PaymentSubscription) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the PaymentSubscription data in the database, used for querying.
func (m *PaymentSubscription) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the PaymentSubscription data in the database, used for querying.
func (m *PaymentSubscription) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the PaymentSubscription data in the database, used for querying.
func (m *PaymentSubscription) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the PaymentSubscription schema, used for querying.
func (m *PaymentSubscription) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the PaymentSubscription schema in the open api documentation.
func (PaymentSubscription) OpenAPISchemaName() string {
	return "PaymentSubscription"
}

// GetOpenAPISchema returns the Open API Schema of the PaymentSubscription in the open api documentation.
func (m *PaymentSubscription) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type PaymentSubscriptionList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the PaymentSubscriptionList schema in the open api documentation.
func (PaymentSubscriptionList) OpenAPISchemaName() string {
	return "PaymentSubscriptionList"
}

// GetOpenAPISchema returns the Open API Schema of the PaymentSubscriptionList in the open api documentation.
func (p *PaymentSubscriptionList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&PaymentSubscription{})
}

// ParamCreate is the expected parameters for create a new PaymentSubscription data.
type ParamCreate struct {
	UseCaseHandler
}

// ParamUpdate is the expected parameters for update the PaymentSubscription data.
type ParamUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamPartiallyUpdate is the expected parameters for partially update the PaymentSubscription data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamDelete is the expected parameters for delete the PaymentSubscription data.
type ParamDelete struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}
