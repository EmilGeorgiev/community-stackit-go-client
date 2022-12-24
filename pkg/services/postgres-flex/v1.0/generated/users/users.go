// Package users provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/do87/oapi-codegen version v0.4.0 DO NOT EDIT.
package users

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	common "github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/do87/oapi-codegen/pkg/runtime"
)

const (
	BearerAuthScopes = "BearerAuth.Scopes"
)

// InstanceCreateUserRequest defines model for instance.CreateUserRequest.
type InstanceCreateUserRequest struct {
	Database *string   `json:"database,omitempty"`
	Roles    *[]string `json:"roles,omitempty"`
	Username *string   `json:"username,omitempty"`
}

// InstanceCreateUserResponse defines model for instance.CreateUserResponse.
type InstanceCreateUserResponse struct {
	Item *InstanceUser `json:"item,omitempty"`
}

// InstanceError defines model for instance.Error.
type InstanceError struct {
	Code    *int                 `json:"code,omitempty"`
	Fields  *map[string][]string `json:"fields,omitempty"`
	Message *string              `json:"message,omitempty"`
	Type    *string              `json:"type,omitempty"`
}

// InstanceGetUserResponse defines model for instance.GetUserResponse.
type InstanceGetUserResponse struct {
	Item *InstanceResponseUser `json:"item,omitempty"`
}

// InstanceListUser defines model for instance.ListUser.
type InstanceListUser struct {
	ID       *string `json:"id,omitempty"`
	Username *string `json:"username,omitempty"`
}

// InstanceListUserResponse defines model for instance.ListUserResponse.
type InstanceListUserResponse struct {
	Count *int                `json:"count,omitempty"`
	Items *[]InstanceListUser `json:"items,omitempty"`
}

// InstanceResponseUser defines model for instance.ResponseUser.
type InstanceResponseUser struct {
	Database *string   `json:"database,omitempty"`
	Host     *string   `json:"host,omitempty"`
	ID       *string   `json:"id,omitempty"`
	Port     *int      `json:"port,omitempty"`
	Roles    *[]string `json:"roles,omitempty"`
	Username *string   `json:"username,omitempty"`
}

// InstanceUser defines model for instance.User.
type InstanceUser struct {
	Database *string   `json:"database,omitempty"`
	Host     *string   `json:"host,omitempty"`
	ID       *string   `json:"id,omitempty"`
	Password *string   `json:"password,omitempty"`
	Port     *int      `json:"port,omitempty"`
	Roles    *[]string `json:"roles,omitempty"`
	Uri      *string   `json:"uri,omitempty"`
	Username *string   `json:"username,omitempty"`
}

// PostInstanceUsersJSONRequestBody defines body for PostInstanceUsers for application/json ContentType.
type PostInstanceUsersJSONRequestBody = InstanceCreateUserRequest

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client common.Client
}

// Creates a new Client, with reasonable defaults
func NewClient(server string, httpClient common.Client) *Client {
	// create a client with sane default values
	client := Client{
		Server: server,
		Client: httpClient,
	}
	return &client
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetUsers request
	GetUsers(ctx context.Context, projectID string, instanceID string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostInstanceUsers request with any body
	PostInstanceUsersWithBody(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostInstanceUsers(ctx context.Context, projectID string, instanceID string, body PostInstanceUsersJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetUser request
	GetUser(ctx context.Context, projectID string, instanceID string, userID string, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetUsers(ctx context.Context, projectID string, instanceID string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUsersRequest(ctx, c.Server, projectID, instanceID)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostInstanceUsersWithBody(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostInstanceUsersRequestWithBody(ctx, c.Server, projectID, instanceID, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostInstanceUsers(ctx context.Context, projectID string, instanceID string, body PostInstanceUsersJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostInstanceUsersRequest(ctx, c.Server, projectID, instanceID, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetUser(ctx context.Context, projectID string, instanceID string, userID string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUserRequest(ctx, c.Server, projectID, instanceID, userID)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetUsersRequest generates requests for GetUsers
func NewGetUsersRequest(ctx context.Context, server string, projectID string, instanceID string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "projectID", runtime.ParamLocationPath, projectID)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "instanceID", runtime.ParamLocationPath, instanceID)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/projects/%s/instances/%s/users", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPostInstanceUsersRequest calls the generic PostInstanceUsers builder with application/json body
func NewPostInstanceUsersRequest(ctx context.Context, server string, projectID string, instanceID string, body PostInstanceUsersJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostInstanceUsersRequestWithBody(ctx, server, projectID, instanceID, "application/json", bodyReader)
}

// NewPostInstanceUsersRequestWithBody generates requests for PostInstanceUsers with any type of body
func NewPostInstanceUsersRequestWithBody(ctx context.Context, server string, projectID string, instanceID string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "projectID", runtime.ParamLocationPath, projectID)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "instanceID", runtime.ParamLocationPath, instanceID)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/projects/%s/instances/%s/users", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetUserRequest generates requests for GetUser
func NewGetUserRequest(ctx context.Context, server string, projectID string, instanceID string, userID string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "projectID", runtime.ParamLocationPath, projectID)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "instanceID", runtime.ParamLocationPath, instanceID)
	if err != nil {
		return nil, err
	}

	var pathParam2 string

	pathParam2, err = runtime.StyleParamWithLocation("simple", false, "userID", runtime.ParamLocationPath, userID)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/projects/%s/instances/%s/users/%s", pathParam0, pathParam1, pathParam2)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, httpClient common.Client) *ClientWithResponses {
	return &ClientWithResponses{NewClient(server, httpClient)}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetUsers request
	GetUsersWithResponse(ctx context.Context, projectID string, instanceID string, reqEditors ...RequestEditorFn) (*GetUsersResponse, error)

	// PostInstanceUsers request with any body
	PostInstanceUsersWithBodyWithResponse(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostInstanceUsersResponse, error)

	PostInstanceUsersWithResponse(ctx context.Context, projectID string, instanceID string, body PostInstanceUsersJSONRequestBody, reqEditors ...RequestEditorFn) (*PostInstanceUsersResponse, error)

	// GetUser request
	GetUserWithResponse(ctx context.Context, projectID string, instanceID string, userID string, reqEditors ...RequestEditorFn) (*GetUserResponse, error)
}

type GetUsersResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *InstanceListUserResponse
	JSON400      *InstanceError
	HasError     error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r GetUsersResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetUsersResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostInstanceUsersResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *InstanceCreateUserResponse
	JSON400      *InstanceError
	HasError     error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r PostInstanceUsersResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostInstanceUsersResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetUserResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *InstanceGetUserResponse
	JSON400      *InstanceError
	HasError     error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r GetUserResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetUserResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetUsersWithResponse request returning *GetUsersResponse
func (c *ClientWithResponses) GetUsersWithResponse(ctx context.Context, projectID string, instanceID string, reqEditors ...RequestEditorFn) (*GetUsersResponse, error) {
	rsp, err := c.GetUsers(ctx, projectID, instanceID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseGetUsersResponse(rsp)
}

// PostInstanceUsersWithBodyWithResponse request with arbitrary body returning *PostInstanceUsersResponse
func (c *ClientWithResponses) PostInstanceUsersWithBodyWithResponse(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostInstanceUsersResponse, error) {
	rsp, err := c.PostInstanceUsersWithBody(ctx, projectID, instanceID, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParsePostInstanceUsersResponse(rsp)
}

func (c *ClientWithResponses) PostInstanceUsersWithResponse(ctx context.Context, projectID string, instanceID string, body PostInstanceUsersJSONRequestBody, reqEditors ...RequestEditorFn) (*PostInstanceUsersResponse, error) {
	rsp, err := c.PostInstanceUsers(ctx, projectID, instanceID, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParsePostInstanceUsersResponse(rsp)
}

// GetUserWithResponse request returning *GetUserResponse
func (c *ClientWithResponses) GetUserWithResponse(ctx context.Context, projectID string, instanceID string, userID string, reqEditors ...RequestEditorFn) (*GetUserResponse, error) {
	rsp, err := c.GetUser(ctx, projectID, instanceID, userID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseGetUserResponse(rsp)
}

// ParseGetUsersResponse parses an HTTP response from a GetUsersWithResponse call
func (c *ClientWithResponses) ParseGetUsersResponse(rsp *http.Response) (*GetUsersResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetUsersResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest InstanceListUserResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest InstanceError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	}

	return response, nil
}

// ParsePostInstanceUsersResponse parses an HTTP response from a PostInstanceUsersWithResponse call
func (c *ClientWithResponses) ParsePostInstanceUsersResponse(rsp *http.Response) (*PostInstanceUsersResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostInstanceUsersResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest InstanceCreateUserResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest InstanceError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	}

	return response, nil
}

// ParseGetUserResponse parses an HTTP response from a GetUserWithResponse call
func (c *ClientWithResponses) ParseGetUserResponse(rsp *http.Response) (*GetUserResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetUserResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest InstanceGetUserResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest InstanceError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	}

	return response, nil
}