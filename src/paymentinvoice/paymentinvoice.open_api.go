package paymentinvoice

import (
	"github.com/shinochiha/gotoko/app"
)

// OpenAPI is constructor for *openAPI, to autogenerate open api document.
func OpenAPI() *OpenAPIOperation {
	return &OpenAPIOperation{}
}

// OpenAPIOperation embed from app.OpenAPIOperation for simplicity, used for autogenerate open api document.
type OpenAPIOperation struct {
	app.OpenAPIOperation
}

// Base is common detail of payment_invoices open api document component.
func (o *OpenAPIOperation) Base() {
	o.Tags = []string{"Payment"}
	o.HeaderParams = []map[string]any{{"$ref": "#/components/parameters/headerParam.Accept-Language"}}
	o.Responses = map[string]map[string]any{
		"200": {
			"description": "Success",
			"content":     map[string]any{"application/json": &PaymentInvoice{}}, // will auto create schema $ref: '#/components/schemas/PaymentInvoice' if not exists
		},
		"400": app.OpenAPIError().BadRequest(),
		"401": app.OpenAPIError().Unauthorized(),
		"403": app.OpenAPIError().Forbidden(),
	}
	o.Securities = []map[string][]string{}
}

// Get is detail of `GET /api/v3/payment_invoices` open api document component.
func (o *OpenAPIOperation) Get() *OpenAPIOperation {
	if !app.IS_GENERATE_OPEN_API_DOC {
		return o // skip for efficiency
	}

	o.Base()
	o.Summary = "Get PaymentInvoice"
	o.Description = "Use this method to get list of PaymentInvoice"
	o.QueryParams = []map[string]any{{"$ref": "#/components/parameters/queryParam.Any"}}
	o.Responses = map[string]map[string]any{
		"200": {
			"description": "Success",
			"content":     map[string]any{"application/json": &PaymentInvoiceList{}}, // will auto create schema $ref: '#/components/schemas/PaymentInvoice.List' if not exists
		},
		"400": app.OpenAPIError().BadRequest(),
		"401": app.OpenAPIError().Unauthorized(),
		"403": app.OpenAPIError().Forbidden(),
	}
	return o
}

// Create is detail of `POST /api/v3/payment_invoices` open api document component.
func (o *OpenAPIOperation) Create() *OpenAPIOperation {
	if !app.IS_GENERATE_OPEN_API_DOC {
		return o // skip for efficiency
	}

	o.Base()
	o.Summary = "Non Credit Card Payment"
	o.Description = "Use this method to create PaymentInvoice Non Credit Card Payment"
	o.Body = map[string]any{"application/json": &ParamCreate{}}
	return o
}

func (o *OpenAPIOperation) CreditCardWithoutInstallment() *OpenAPIOperation {
	if !app.IS_GENERATE_OPEN_API_DOC {
		return o // skip for efficiency
	}

	o.Base()
	o.Summary = "Credit Card without installment"
	o.Description = "Use this method to create PaymentInvoice Credit Card without installment"
	// o.Body = map[string]any{"application/json": &ParamCreate{}}
	return o
}

func (o *OpenAPIOperation) CreditCardWithInstallment() *OpenAPIOperation {
	if !app.IS_GENERATE_OPEN_API_DOC {
		return o // skip for efficiency
	}

	o.Base()
	o.Summary = "Credit Card with installment"
	o.Description = "Use this method to create PaymentInvoice Credit Card with installment"
	// o.Body = map[string]any{"application/json": &ParamCreate{}}
	return o
}
