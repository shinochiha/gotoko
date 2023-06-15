package paymentsubscription

import "github.com/shinochiha/gotoko/app"

// OpenAPI is constructor for *openAPI, to autogenerate open api document.
func OpenAPI() *OpenAPIOperation {
	return &OpenAPIOperation{}
}

// OpenAPIOperation embed from app.OpenAPIOperation for simplicity, used for autogenerate open api document.
type OpenAPIOperation struct {
	app.OpenAPIOperation
}

// Base is common detail of payment_subscriptions open api document component.
func (o *OpenAPIOperation) Base() {
	o.Tags = []string{"Subscription"}
	o.HeaderParams = []map[string]any{{"$ref": "#/components/parameters/headerParam.Accept-Language"}}
	o.Responses = map[string]map[string]any{
		"200": {
			"description": "Success",
			"content":     map[string]any{"application/json": &PaymentSubscription{}}, // will auto create schema $ref: '#/components/schemas/PaymentSubscription' if not exists
		},
		"400": app.OpenAPIError().BadRequest(),
		"401": app.OpenAPIError().Unauthorized(),
		"403": app.OpenAPIError().Forbidden(),
	}
	o.Securities = []map[string][]string{}
}

// Get is detail of `GET /api/v3/payment_subscriptions` open api document component.
func (o *OpenAPIOperation) Get() *OpenAPIOperation {
	if !app.IS_GENERATE_OPEN_API_DOC {
		return o // skip for efficiency
	}

	o.Base()
	o.Summary = "Get Subscription"
	o.Description = "Use this method to get list of Subscription"
	o.QueryParams = []map[string]any{{"$ref": "#/components/parameters/queryParam.Any"}}
	o.Responses = map[string]map[string]any{
		"200": {
			"description": "Success",
			"content":     map[string]any{"application/json": &PaymentSubscriptionList{}}, // will auto create schema $ref: '#/components/schemas/PaymentSubscription.List' if not exists
		},
		"400": app.OpenAPIError().BadRequest(),
		"401": app.OpenAPIError().Unauthorized(),
		"403": app.OpenAPIError().Forbidden(),
	}
	return o
}

// Create is detail of `POST /api/v3/payment_subscriptions` open api document component.
func (o *OpenAPIOperation) Create() *OpenAPIOperation {
	if !app.IS_GENERATE_OPEN_API_DOC {
		return o // skip for efficiency
	}

	o.Base()
	o.Summary = "Create Subscriptions"
	o.Description = "Use this method to create Subscription"
	o.Body = map[string]any{"application/json": &ParamCreate{}}
	return o
}

func (o *OpenAPIOperation) PauseSubscription() *OpenAPIOperation {
	if !app.IS_GENERATE_OPEN_API_DOC {
		return o // skip for efficiency
	}

	o.Base()
	o.Summary = "Pause Subscriptions"
	o.Description = "Use this method to pause Subscription"
	o.Body = map[string]any{"application/json": &ParamCreate{}}
	return o
}

func (o *OpenAPIOperation) ResumeSubscription() *OpenAPIOperation {
	if !app.IS_GENERATE_OPEN_API_DOC {
		return o // skip for efficiency
	}

	o.Base()
	o.Summary = "Resume Subscriptions"
	o.Description = "Use this method to resume Subscription"
	o.Body = map[string]any{"application/json": &ParamCreate{}}
	return o
}

func (o *OpenAPIOperation) StopSubscription() *OpenAPIOperation {
	if !app.IS_GENERATE_OPEN_API_DOC {
		return o // skip for efficiency
	}

	o.Base()
	o.Summary = "Stop Subscriptions"
	o.Description = "Use this method to stop Subscription"
	o.Body = map[string]any{"application/json": &ParamCreate{}}
	return o
}
