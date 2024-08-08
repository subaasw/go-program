package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"snippetbox-app/internal/assert"
	"testing"
)

func TestPing(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	ping(rr, r)

	rs := rr.Result()
	assert.Equal(t, rs.StatusCode, http.StatusOK)

	defer rs.Body.Close()

	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	bytes.TrimSpace(body)

	assert.Equal(t, string(body), "OK")
}

func TestPing2(t *testing.T) {
	app := &application{
		errorLog: log.New(io.Discard, "", 0),
		infoLog:  log.New(io.Discard, "", 0),
	}

	ts := httptest.NewTLSServer(app.routes())
	defer ts.Close()

	rs, err := ts.Client().Get(ts.URL + "/ping")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, rs.StatusCode, http.StatusOK)
	defer rs.Body.Close()

	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	bytes.TrimSpace(body)
	assert.Equal(t, string(body), "OK")
}

func TestPing3(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/ping")

	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, body, "OK")
}

func TestSnippetView(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody string
	}{
		{
			name:     "Valid Id",
			urlPath:  "/snippet/view/1",
			wantCode: http.StatusOK,
			wantBody: "An old silent pod...",
		},
		{
			name:     "Non-existent Id",
			urlPath:  "/snippet/view/20",
			wantCode: http.StatusInternalServerError,
		},
		{
			name:     "Negative Id",
			urlPath:  "/snippet/view/-1",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Decimal Id",
			urlPath:  "/snippet/view/1.23",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "String Id",
			urlPath:  "/snippet/view/foo",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Empty Id",
			urlPath:  "/snippet/view/",
			wantCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, body := ts.get(t, tt.urlPath)
			assert.Equal(t, code, tt.wantCode)

			if tt.wantBody != "" {
				assert.StringContains(t, body, tt.wantBody)
			}
		},
		)
	}
}

func TestUserSignup(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	_, _, body := ts.get(t, "/user/signup")
	validCSRFToken := extractCSRFToken(t, body)

	const (
		validName     = "Bobg"
		validEmail    = "bobg@mail.com"
		validPassword = "validPa$$word"
		formAction    = "/user/signup"
	)

	tests := []struct {
		name           string
		userName       string
		userEmail      string
		userPassword   string
		csrfToken      string
		wantCode       int
		wantFormAction string
	}{
		{
			name:         "Valid Submission",
			userName:     validName,
			userEmail:    validEmail,
			userPassword: validPassword,
			csrfToken:    validCSRFToken,
			wantCode:     http.StatusSeeOther,
		},
		{
			name:         "Invalid CSRF Token",
			userName:     validName,
			userEmail:    validEmail,
			userPassword: validPassword,
			csrfToken:    "wrongToken",
			wantCode:     http.StatusBadRequest,
		},
		{
			name:           "Empty name",
			userName:       "",
			userEmail:      validEmail,
			userPassword:   validPassword,
			csrfToken:      validCSRFToken,
			wantCode:       http.StatusUnprocessableEntity,
			wantFormAction: formAction,
		},
		{
			name:           "Empty email",
			userName:       validName,
			userEmail:      "",
			userPassword:   validPassword,
			csrfToken:      validCSRFToken,
			wantCode:       http.StatusUnprocessableEntity,
			wantFormAction: formAction,
		},
		{
			name:           "Empty password",
			userName:       validName,
			userEmail:      validEmail,
			userPassword:   "",
			csrfToken:      validCSRFToken,
			wantCode:       http.StatusUnprocessableEntity,
			wantFormAction: formAction,
		},
		{
			name:           "Invalid email",
			userName:       validName,
			userEmail:      "bob@example.",
			userPassword:   validPassword,
			csrfToken:      validCSRFToken,
			wantCode:       http.StatusUnprocessableEntity,
			wantFormAction: formAction,
		},
		{
			name:           "Short password",
			userName:       validName,
			userEmail:      validEmail,
			userPassword:   "pa$$",
			csrfToken:      validCSRFToken,
			wantCode:       http.StatusUnprocessableEntity,
			wantFormAction: formAction,
		},
		{
			name:           "Duplicate email",
			userName:       validName,
			userEmail:      "dupe@example.com",
			userPassword:   validPassword,
			csrfToken:      validCSRFToken,
			wantCode:       http.StatusUnprocessableEntity,
			wantFormAction: formAction,
		},
		{
			name:           "Empty password",
			userName:       validName,
			userEmail:      validEmail,
			userPassword:   "",
			csrfToken:      validCSRFToken,
			wantCode:       http.StatusUnprocessableEntity,
			wantFormAction: formAction,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form := url.Values{}
			form.Add("name", tt.userName)
			form.Add("email", tt.userEmail)
			form.Add("password", tt.userPassword)
			form.Add("csrf_token", tt.csrfToken)
			// println(tt.csrfToken)

			code, _, body := ts.postForm(t, "/user/signup", form)

			assert.Equal(t, code, tt.wantCode)

			if tt.wantFormAction != "" {
				assert.StringContains(t, body, tt.wantFormAction)
			}
		})
	}

}
