// Package instances provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/do87/oapi-codegen version v0.5.1 DO NOT EDIT.
package instances

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"

	common "github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
	"github.com/do87/oapi-codegen/pkg/runtime"
)

// Defines values for LastOperationState.
const (
	FAILED      LastOperationState = "failed"
	IN_PROGRESS LastOperationState = "in progress"
	SUCCEEDED   LastOperationState = "succeeded"
)

// Defines values for LastOperationType.
const (
	CREATE LastOperationType = "create"
	DELETE LastOperationType = "delete"
	UPDATE LastOperationType = "update"
)

// Error defines model for Error.
type Error struct {
	Description string `json:"description"`
	Error       string `json:"error"`
}

// Instance defines model for Instance.
type Instance struct {
	CfGuid           string        `json:"cfGuid"`
	CfSpaceGuid      string        `json:"cfSpaceGuid"`
	DashboardUrl     string        `json:"dashboardUrl"`
	ImageUrl         string        `json:"imageUrl"`
	InstanceID       *string       `json:"instanceId,omitempty"`
	LastOperation    LastOperation `json:"lastOperation"`
	Name             string        `json:"name"`
	OrganizationGuid *string       `json:"organizationGuid,omitempty"`
	Parameters       Object        `json:"parameters"`
	PlanID           string        `json:"planId"`
}

// InstanceID defines model for InstanceID.
type InstanceID struct {
	InstanceID string `json:"instanceId"`
}

// InstanceList defines model for InstanceList.
type InstanceList struct {
	Instances []Instance `json:"instances"`
}

// InstanceProvisionRequest defines model for InstanceProvisionRequest.
type InstanceProvisionRequest struct {
	InstanceName string  `json:"instanceName"`
	Parameters   *Object `json:"parameters,omitempty"`
	PlanID       string  `json:"planId"`
}

// InstanceUpdateRequest defines model for InstanceUpdateRequest.
type InstanceUpdateRequest struct {
	Parameters *Object `json:"parameters,omitempty"`
	PlanID     string  `json:"planId"`
}

// Object defines model for Object.
type Object = map[string]interface{}

// LastOperation defines model for lastOperation.
type LastOperation struct {
	Description string             `json:"description"`
	State       LastOperationState `json:"state"`
	Type        LastOperationType  `json:"type"`
}

// LastOperationState defines model for LastOperation.State.
type LastOperationState string

// LastOperationType defines model for LastOperation.Type.
type LastOperationType string

// InstanceId defines model for instanceId.
type InstanceId = string

// BadRequest defines model for BadRequest.
type BadRequest = Error

// NotFound defines model for NotFound.
type NotFound = Error

// CreateJSONRequestBody defines body for Create for application/json ContentType.
type CreateJSONRequestBody = InstanceProvisionRequest

// UpdateJSONRequestBody defines body for Update for application/json ContentType.
type UpdateJSONRequestBody = InstanceUpdateRequest

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
	// List request
	List(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// Create request with any body
	CreateWithBody(ctx context.Context, projectID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	Create(ctx context.Context, projectID string, body CreateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// Delete request
	Delete(ctx context.Context, projectID string, instanceID string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// Get request
	Get(ctx context.Context, projectID string, instanceID InstanceId, reqEditors ...RequestEditorFn) (*http.Response, error)

	// Update request with any body
	UpdateWithBody(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	Update(ctx context.Context, projectID string, instanceID string, body UpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) List(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListRequest(ctx, c.Server, projectID)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateWithBody(ctx context.Context, projectID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateRequestWithBody(ctx, c.Server, projectID, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) Create(ctx context.Context, projectID string, body CreateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateRequest(ctx, c.Server, projectID, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) Delete(ctx context.Context, projectID string, instanceID string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteRequest(ctx, c.Server, projectID, instanceID)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) Get(ctx context.Context, projectID string, instanceID InstanceId, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetRequest(ctx, c.Server, projectID, instanceID)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateWithBody(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateRequestWithBody(ctx, c.Server, projectID, instanceID, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) Update(ctx context.Context, projectID string, instanceID string, body UpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateRequest(ctx, c.Server, projectID, instanceID, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewListRequest generates requests for List
func NewListRequest(ctx context.Context, server string, projectID string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "projectID", runtime.ParamLocationPath, projectID)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/projects/%s/instances", pathParam0)
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

// NewCreateRequest calls the generic Create builder with application/json body
func NewCreateRequest(ctx context.Context, server string, projectID string, body CreateJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateRequestWithBody(ctx, server, projectID, "application/json", bodyReader)
}

// NewCreateRequestWithBody generates requests for Create with any type of body
func NewCreateRequestWithBody(ctx context.Context, server string, projectID string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "projectID", runtime.ParamLocationPath, projectID)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/projects/%s/instances", pathParam0)
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

// NewDeleteRequest generates requests for Delete
func NewDeleteRequest(ctx context.Context, server string, projectID string, instanceID string) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/v1/projects/%s/instances/%s", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetRequest generates requests for Get
func NewGetRequest(ctx context.Context, server string, projectID string, instanceID InstanceId) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/v1/projects/%s/instances/%s", pathParam0, pathParam1)
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

// NewUpdateRequest calls the generic Update builder with application/json body
func NewUpdateRequest(ctx context.Context, server string, projectID string, instanceID string, body UpdateJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateRequestWithBody(ctx, server, projectID, instanceID, "application/json", bodyReader)
}

// NewUpdateRequestWithBody generates requests for Update with any type of body
func NewUpdateRequestWithBody(ctx context.Context, server string, projectID string, instanceID string, contentType string, body io.Reader) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/v1/projects/%s/instances/%s", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "PATCH", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

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
	// List request
	ListWithResponse(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*ListResponse, error)

	// Create request with any body
	CreateWithBodyWithResponse(ctx context.Context, projectID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateResponse, error)

	CreateWithResponse(ctx context.Context, projectID string, body CreateJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateResponse, error)

	// Delete request
	DeleteWithResponse(ctx context.Context, projectID string, instanceID string, reqEditors ...RequestEditorFn) (*DeleteResponse, error)

	// Get request
	GetWithResponse(ctx context.Context, projectID string, instanceID InstanceId, reqEditors ...RequestEditorFn) (*GetResponse, error)

	// Update request with any body
	UpdateWithBodyWithResponse(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateResponse, error)

	UpdateWithResponse(ctx context.Context, projectID string, instanceID string, body UpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateResponse, error)
}

type ListResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *InstanceList
	JSON404      *Error
	HasError     error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r ListResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON202      *InstanceID
	JSON400      *Error
	JSON409      *Error
	HasError     error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r CreateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON400      *Error
	JSON404      *Error
	HasError     error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r DeleteResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Instance
	JSON404      *Error
	JSON410      *Error
	HasError     error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r GetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON400      *Error
	JSON404      *Error
	HasError     error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r UpdateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ListWithResponse request returning *ListResponse
func (c *ClientWithResponses) ListWithResponse(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*ListResponse, error) {
	rsp, err := c.List(ctx, projectID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseListResponse(rsp)
}

// CreateWithBodyWithResponse request with arbitrary body returning *CreateResponse
func (c *ClientWithResponses) CreateWithBodyWithResponse(ctx context.Context, projectID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateResponse, error) {
	rsp, err := c.CreateWithBody(ctx, projectID, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseCreateResponse(rsp)
}

func (c *ClientWithResponses) CreateWithResponse(ctx context.Context, projectID string, body CreateJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateResponse, error) {
	rsp, err := c.Create(ctx, projectID, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseCreateResponse(rsp)
}

// DeleteWithResponse request returning *DeleteResponse
func (c *ClientWithResponses) DeleteWithResponse(ctx context.Context, projectID string, instanceID string, reqEditors ...RequestEditorFn) (*DeleteResponse, error) {
	rsp, err := c.Delete(ctx, projectID, instanceID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseDeleteResponse(rsp)
}

// GetWithResponse request returning *GetResponse
func (c *ClientWithResponses) GetWithResponse(ctx context.Context, projectID string, instanceID InstanceId, reqEditors ...RequestEditorFn) (*GetResponse, error) {
	rsp, err := c.Get(ctx, projectID, instanceID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseGetResponse(rsp)
}

// UpdateWithBodyWithResponse request with arbitrary body returning *UpdateResponse
func (c *ClientWithResponses) UpdateWithBodyWithResponse(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateResponse, error) {
	rsp, err := c.UpdateWithBody(ctx, projectID, instanceID, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseUpdateResponse(rsp)
}

func (c *ClientWithResponses) UpdateWithResponse(ctx context.Context, projectID string, instanceID string, body UpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateResponse, error) {
	rsp, err := c.Update(ctx, projectID, instanceID, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseUpdateResponse(rsp)
}

// ParseListResponse parses an HTTP response from a ListWithResponse call
func (c *ClientWithResponses) ParseListResponse(rsp *http.Response) (*ListResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	response.HasError = validate.DefaultResponseErrorHandler(rsp)

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest InstanceList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ParseCreateResponse parses an HTTP response from a CreateWithResponse call
func (c *ClientWithResponses) ParseCreateResponse(rsp *http.Response) (*CreateResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	response.HasError = validate.DefaultResponseErrorHandler(rsp)

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 202:
		var dest InstanceID
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON202 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 409:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON409 = &dest

	}

	return response, nil
}

// ParseDeleteResponse parses an HTTP response from a DeleteWithResponse call
func (c *ClientWithResponses) ParseDeleteResponse(rsp *http.Response) (*DeleteResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	response.HasError = validate.DefaultResponseErrorHandler(rsp)

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ParseGetResponse parses an HTTP response from a GetWithResponse call
func (c *ClientWithResponses) ParseGetResponse(rsp *http.Response) (*GetResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	response.HasError = validate.DefaultResponseErrorHandler(rsp)

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Instance
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 410:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON410 = &dest

	}

	return response, nil
}

// ParseUpdateResponse parses an HTTP response from a UpdateWithResponse call
func (c *ClientWithResponses) ParseUpdateResponse(rsp *http.Response) (*UpdateResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	response.HasError = validate.DefaultResponseErrorHandler(rsp)

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON404 = &dest

	}

	return response, nil
}
