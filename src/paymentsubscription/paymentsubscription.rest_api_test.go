package paymentsubscription

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2/utils"
	"gorm.io/gorm"

	"github.com/shinochiha/gotoko/app"
)

// prepareTest prepares the test.
func prepareTest(tb testing.TB) {
	app.Test()
	tx := app.Test().Tx
	app.DB().RegisterTable("main", PaymentSubscription{})
	app.DB().MigrateTable(tx, "main", app.Setting{})
	tx.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&PaymentSubscription{})

	app.Server().AddMiddleware(app.Test().NewCtx([]string{
		"payment_subscriptions.detail",
		"payment_subscriptions.list",
		"payment_subscriptions.create",
		"payment_subscriptions.edit",
		"payment_subscriptions.delete",
	}))
	app.Server().AddRoute("/payment_subscriptions", "POST", REST().Create, nil)
	app.Server().AddRoute("/payment_subscriptions", "GET", REST().Get, nil)
	app.Server().AddRoute("/payment_subscriptions/:id", "GET", REST().GetByID, nil)
	app.Server().AddRoute("/payment_subscriptions/:id", "PUT", REST().UpdateByID, nil)
	app.Server().AddRoute("/payment_subscriptions/:id", "PATCH", REST().PartiallyUpdateByID, nil)
	app.Server().AddRoute("/payment_subscriptions/:id", "DELETE", REST().DeleteByID, nil)
}

// getTestPaymentSubscriptionID returns an available PaymentSubscription ID.
func getTestPaymentSubscriptionID() string {
	return "todo"
}

// tests is test scenario.
var tests = []struct {
	description  string // description of the test case
	method       string // method to test
	path         string // route path to test
	token        string // token to test
	bodyRequest  string // body to test
	expectedCode int    // expected HTTP status code
	expectedBody string // expected body response
}{
	{
		description:  "Get empty list of PaymentSubscription",
		method:       "GET",
		path:         "/payment_subscriptions",
		token:        app.TestFullAccessToken,
		expectedCode: http.StatusOK,
		expectedBody: `{"count":0,"results":[]}`,
	},
	{
		description:  "Create PaymentSubscription with minimum payload",
		method:       "POST",
		path:         "/payment_subscriptions",
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"name":"Kilogram"}`,
		expectedCode: http.StatusCreated,
		expectedBody: `{"name":"Kilogram"}`,
	},
	{
		description:  "Get PaymentSubscription by ID",
		method:       "GET",
		path:         "/payment_subscriptions/" + getTestPaymentSubscriptionID(),
		token:        app.TestFullAccessToken,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"Kilogram"}`,
	},
	{
		description:  "Update PaymentSubscription by ID",
		method:       "PUT",
		path:         "/payment_subscriptions/" + getTestPaymentSubscriptionID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Update PaymentSubscription by ID","name":"KG"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"KG"}`,
	},
	{
		description:  "Partially update PaymentSubscription by ID",
		method:       "PATCH",
		path:         "/payment_subscriptions/" + getTestPaymentSubscriptionID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Partially Update PaymentSubscription by ID","name":"Kilo Gram"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"Kilo Gram"}`,
	},
	{
		description:  "Delete PaymentSubscription by ID",
		method:       "DELETE",
		path:         "/payment_subscriptions/" + getTestPaymentSubscriptionID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Delete PaymentSubscription by ID"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"code":200}`,
	},
}

// TestPaymentSubscriptionREST tests the REST API of PaymentSubscription data with specified scenario.
func TestPaymentSubscriptionREST(t *testing.T) {
	prepareTest(t)

	// Iterate through test single test cases
	for _, test := range tests {

		// Create a new http request with the route from the test case
		req := httptest.NewRequest(test.method, test.path, strings.NewReader(test.bodyRequest))
		req.Header.Add("Authorization", "Bearer "+test.token)
		req.Header.Add("Content-Type", "application/json")

		// Perform the request plain with the app, the second argument is a request latency (set to -1 for no latency)
		res, err := app.Server().Test(req)

		// Verify if the status code is as expected
		utils.AssertEqual(t, nil, err, "app.Server().Test(req)")
		utils.AssertEqual(t, test.expectedCode, res.StatusCode, test.description)

		// Verify if the body response is as expected
		body, err := io.ReadAll(res.Body)
		utils.AssertEqual(t, nil, err, "io.ReadAll(res.Body)")
		app.Test().AssertMatchJSONElement(t, []byte(test.expectedBody), body, test.description)
		res.Body.Close()
	}
}

// BenchmarkPaymentSubscriptionREST tests the REST API of PaymentSubscription data with specified scenario.
func BenchmarkPaymentSubscriptionREST(b *testing.B) {
	b.ReportAllocs()
	prepareTest(b)
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			req := httptest.NewRequest(test.method, test.path, strings.NewReader(test.bodyRequest))
			req.Header.Add("Authorization", "Bearer "+test.token)
			req.Header.Add("Content-Type", "application/json")
			app.Server().Test(req)
		}
	}
}
