package paymentinvoice

import (
	"github.com/shinochiha/gotoko/app"
)

// PaymentInvoice is the main model of PaymentInvoice data. It provides a convenient interface for app.ModelInterface
type PaymentInvoice struct {
	app.Model
	ID                               app.NullUUID     `json:"id"                                   db:"m.id"                       gorm:"column:id;primaryKey"`
	PaymentPaymentType               app.NullString   `json:"payment.payment_type"                 db:"m.payment_type"             gorm:"column:payment_type"`
	PaymentCreditCardBank            app.NullString   `json:"payment.credit_card.bank"             db:"m.payment_credit_card_bank" gorm:"column:payment_credit_card_bank"`
	PaymentCreditCardInstallmentType app.NullString   `json:"payment.credit_card.installment.type" db:"m.credit_installment_type"  gorm:"column:credit_installment_type"`
	PaymentCreditCardInstallmentTerm app.NullInt64    `json:"payment.credit_card.installment.term" db:"m.credit_installment_term"  gorm:"column:credit_installment_term"`
	CustomerName                     app.NullString   `json:"customer.name"                        db:"m.customer_name"            gorm:"column:customer_name"`
	CustomerEmail                    app.NullString   `json:"customer.email"                       db:"m.customer_email"           gorm:"column:customer_email"`
	CustomerPhoneNumber              app.NullString   `json:"customer.phone_number"                db:"m.customer_phone_number"    gorm:"column:customer_phone_number"`
	Item                             []Item           `json:"items"                                db:"a.data_id=id"               gorm:"-"`
	CreatedAt                        app.NullDateTime `json:"created_at"                           db:"m.created_at"               gorm:"column:created_at"`
	UpdatedAt                        app.NullDateTime `json:"updated_at"                           db:"m.updated_at"               gorm:"column:updated_at"`
	DeletedAt                        app.NullDateTime `json:"deleted_at"                           db:"m.deleted_at,hide"          gorm:"column:deleted_at"`
}

// EndPoint returns the PaymentInvoice end point, it used for cache key, etc.
func (PaymentInvoice) EndPoint() string {
	return "payment_invoices"
}

// TableVersion returns the versions of the PaymentInvoice table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (PaymentInvoice) TableVersion() string {
	return "28.06.291152"
}

// TableName returns the name of the PaymentInvoice table in the database.
func (PaymentInvoice) TableName() string {
	return "payment_invoices"
}

// TableAliasName returns the table alias name of the PaymentInvoice table, used for querying.
func (PaymentInvoice) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the PaymentInvoice data in the database, used for querying.
func (m *PaymentInvoice) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the PaymentInvoice data in the database, used for querying.
func (m *PaymentInvoice) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the PaymentInvoice data in the database, used for querying.
func (m *PaymentInvoice) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the PaymentInvoice data in the database, used for querying.
func (m *PaymentInvoice) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the PaymentInvoice schema, used for querying.
func (m *PaymentInvoice) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the PaymentInvoice schema in the open api documentation.
func (PaymentInvoice) OpenAPISchemaName() string {
	return "PaymentInvoice"
}

// GetOpenAPISchema returns the Open API Schema of the PaymentInvoice in the open api documentation.
func (m *PaymentInvoice) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type Item struct {
	app.Model
	ID          app.NullUUID     `json:"id"          db:"a.id"              gorm:"column:id;primaryKey"`
	DataId      app.NullUUID     `json:"-"           db:"a.data_id"         gorm:"column:data_id"`
	Name        app.NullString   `json:"name"        db:"a.name"            gorm:"column:name"`
	Category    app.NullString   `json:"category"    db:"a.category"        gorm:"column:category"`
	Merchant    app.NullString   `json:"merchant"    db:"a.merchant"        gorm:"column:merchant"`
	Description app.NullString   `json:"description" db:"a.description"     gorm:"column:description"`
	Qty         app.NullFloat64  `json:"qty"         db:"a.qty"             gorm:"column:qty"`
	Price       app.NullFloat64  `json:"price"       db:"a.price"           gorm:"column:price"`
	Currency    app.NullFloat64  `json:"currency"    db:"a.currency"        gorm:"column:currency"`
	CreatedAt   app.NullDateTime `json:"created_at"  db:"a.created_at"      gorm:"column:created_at"`
	UpdatedAt   app.NullDateTime `json:"updated_at"  db:"a.updated_at"      gorm:"column:updated_at"`
	DeletedAt   app.NullDateTime `json:"deleted_at"  db:"a.deleted_at,hide" gorm:"column:deleted_at"`
}

func (Item) TableVersion() string {
	return "28.06.291152"
}

// TableName returns the name of the PaymentInvoice table in the database.
func (Item) TableName() string {
	return "items"
}

// TableAliasName returns the table alias name of the PaymentInvoice table, used for querying.
func (Item) TableAliasName() string {
	return "a"
}

// GetRelations returns the relations of the PaymentInvoice data in the database, used for querying.
func (m *Item) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the PaymentInvoice data in the database, used for querying.
func (m *Item) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "a.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the PaymentInvoice data in the database, used for querying.
func (m *Item) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "a.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the PaymentInvoice data in the database, used for querying.
func (m *Item) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the PaymentInvoice schema, used for querying.
func (m *Item) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the PaymentInvoice schema in the open api documentation.
func (Item) OpenAPISchemaName() string {
	return "PaymentInvoice"
}

// GetOpenAPISchema returns the Open API Schema of the PaymentInvoice in the open api documentation.
func (m *Item) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type PaymentInvoiceList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the PaymentInvoiceList schema in the open api documentation.
func (PaymentInvoiceList) OpenAPISchemaName() string {
	return "PaymentInvoiceList"
}

// GetOpenAPISchema returns the Open API Schema of the PaymentInvoiceList in the open api documentation.
func (p *PaymentInvoiceList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&PaymentInvoice{})
}

// ParamCreate is the expected parameters for create a new PaymentInvoice data.
type ParamCreate struct {
	UseCaseHandler
}

// ParamUpdate is the expected parameters for update the PaymentInvoice data.
type ParamUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamPartiallyUpdate is the expected parameters for partially update the PaymentInvoice data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamDelete is the expected parameters for delete the PaymentInvoice data.
type ParamDelete struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}
