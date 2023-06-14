package contact

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
	app.DB().RegisterTable("main", Contact{})
	app.DB().MigrateTable(tx, "main", app.Setting{})
	tx.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Contact{})

	app.Server().AddMiddleware(app.Test().NewCtx([]string{
		"contacts.detail",
		"contacts.list",
		"contacts.create",
		"contacts.edit",
		"contacts.delete",
	}))
	app.Server().AddRoute("/contacts", "POST", REST().Create, nil)
	app.Server().AddRoute("/contacts", "GET", REST().Get, nil)
	app.Server().AddRoute("/contacts/:id", "GET", REST().GetByID, nil)
	app.Server().AddRoute("/contacts/:id", "PUT", REST().UpdateByID, nil)
	app.Server().AddRoute("/contacts/:id", "PATCH", REST().PartiallyUpdateByID, nil)
	app.Server().AddRoute("/contacts/:id", "DELETE", REST().DeleteByID, nil)
}

// getTestContactID returns an available Contact ID.
func getTestContactID() string {
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
		description:  "Get empty list of Contact",
		method:       "GET",
		path:         "/contacts",
		token:        app.TestFullAccessToken,
		expectedCode: http.StatusOK,
		expectedBody: `{"count":0,"results":[]}`,
	},
	{
		description:  "Create Contact with minimum payload",
		method:       "POST",
		path:         "/contacts",
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"name":"Kilogram"}`,
		expectedCode: http.StatusCreated,
		expectedBody: `{"name":"Kilogram"}`,
	},
	{
		description:  "Get Contact by ID",
		method:       "GET",
		path:         "/contacts/" + getTestContactID(),
		token:        app.TestFullAccessToken,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"Kilogram"}`,
	},
	{
		description:  "Update Contact by ID",
		method:       "PUT",
		path:         "/contacts/" + getTestContactID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Update Contact by ID","name":"KG"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"KG"}`,
	},
	{
		description:  "Partially update Contact by ID",
		method:       "PATCH",
		path:         "/contacts/" + getTestContactID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Partially Update Contact by ID","name":"Kilo Gram"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"Kilo Gram"}`,
	},
	{
		description:  "Delete Contact by ID",
		method:       "DELETE",
		path:         "/contacts/" + getTestContactID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Delete Contact by ID"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"code":200}`,
	},
}

// TestContactREST tests the REST API of Contact data with specified scenario.
func TestContactREST(t *testing.T) {
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

// BenchmarkContactREST tests the REST API of Contact data with specified scenario.
func BenchmarkContactREST(b *testing.B) {
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
