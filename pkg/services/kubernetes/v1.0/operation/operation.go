// Package operation provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/do87/stackit-client-generator version v0.0.2 DO NOT EDIT.
package operation

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"

	contracts "github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/helpers/runtime"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Defines values for RuntimeErrorCode.
const (
	SKE_API_SERVER_ERROR         RuntimeErrorCode = "SKE_API_SERVER_ERROR"
	SKE_ARGUS_INSTANCE_NOT_FOUND RuntimeErrorCode = "SKE_ARGUS_INSTANCE_NOT_FOUND"
	SKE_CONFIGURATION_PROBLEM    RuntimeErrorCode = "SKE_CONFIGURATION_PROBLEM"
	SKE_INFRA_ERROR              RuntimeErrorCode = "SKE_INFRA_ERROR"
	SKE_QUOTA_EXCEEDED           RuntimeErrorCode = "SKE_QUOTA_EXCEEDED"
	SKE_RATE_LIMITS              RuntimeErrorCode = "SKE_RATE_LIMITS"
	SKE_REMAINING_RESOURCES      RuntimeErrorCode = "SKE_REMAINING_RESOURCES"
	SKE_TMP_AUTH_ERROR           RuntimeErrorCode = "SKE_TMP_AUTH_ERROR"
	SKE_UNREADY_NODES            RuntimeErrorCode = "SKE_UNREADY_NODES"
	SKE_UNSPECIFIED              RuntimeErrorCode = "SKE_UNSPECIFIED"
)

// RuntimeError defines model for RuntimeError.
type RuntimeError struct {
	// Code - Code:    "SKE_UNSPECIFIED"
	//   Message: "An error occurred. Please open a support ticket if this error persists."
	// - Code:    "SKE_TMP_AUTH_ERROR"
	//   Message: "Authentication failed. This is a temporary error. Please wait while the system recovers."
	// - Code:    "SKE_QUOTA_EXCEEDED"
	//   Message: "Your project's resource quotas are exhausted. Please make sure your quota is sufficient for the ordered cluster."
	// - Code:    "SKE_ARGUS_INSTANCE_NOT_FOUND"
	//   Message: "The provided Argus instance could not be found."
	// - Code:    "SKE_RATE_LIMITS"
	//   Message: "While provisioning your cluster, request rate limits where incurred. Please wait while the system recovers."
	// - Code:    "SKE_INFRA_ERROR"
	//   Message: "An error occurred with the underlying infrastructure. Please open a support ticket if this error persists."
	// - Code:    "SKE_REMAINING_RESOURCES"
	//   Message: "There are remaining Kubernetes resources in your cluster that prevent deletion. Please make sure to remove them."
	// - Code:    "SKE_CONFIGURATION_PROBLEM"
	//   Message: "A configuration error occurred. Please open a support ticket if this error persists."
	// - Code:    "SKE_UNREADY_NODES"
	//   Message: "Not all worker nodes are ready. Please open a support ticket if this error persists."
	// - Code:    "SKE_API_SERVER_ERROR"
	//   Message: "The Kubernetes API server is not reporting readiness. Please open a support ticket if this error persists."
	Code    *RuntimeErrorCode `json:"code,omitempty"`
	Details *string           `json:"details,omitempty"`
	Message *string           `json:"message,omitempty"`
}

// RuntimeErrorCode - Code:    "SKE_UNSPECIFIED"
//
//		Message: "An error occurred. Please open a support ticket if this error persists."
//	  - Code:    "SKE_TMP_AUTH_ERROR"
//	    Message: "Authentication failed. This is a temporary error. Please wait while the system recovers."
//	  - Code:    "SKE_QUOTA_EXCEEDED"
//	    Message: "Your project's resource quotas are exhausted. Please make sure your quota is sufficient for the ordered cluster."
//	  - Code:    "SKE_ARGUS_INSTANCE_NOT_FOUND"
//	    Message: "The provided Argus instance could not be found."
//	  - Code:    "SKE_RATE_LIMITS"
//	    Message: "While provisioning your cluster, request rate limits where incurred. Please wait while the system recovers."
//	  - Code:    "SKE_INFRA_ERROR"
//	    Message: "An error occurred with the underlying infrastructure. Please open a support ticket if this error persists."
//	  - Code:    "SKE_REMAINING_RESOURCES"
//	    Message: "There are remaining Kubernetes resources in your cluster that prevent deletion. Please make sure to remove them."
//	  - Code:    "SKE_CONFIGURATION_PROBLEM"
//	    Message: "A configuration error occurred. Please open a support ticket if this error persists."
//	  - Code:    "SKE_UNREADY_NODES"
//	    Message: "Not all worker nodes are ready. Please open a support ticket if this error persists."
//	  - Code:    "SKE_API_SERVER_ERROR"
//	    Message: "The Kubernetes API server is not reporting readiness. Please open a support ticket if this error persists."
type RuntimeErrorCode string

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
	Client contracts.BaseClientInterface
}

// NewRawClient Creates a new Client, with reasonable defaults
func NewRawClient(server string, httpClient contracts.BaseClientInterface) *Client {
	// create a client with sane default values
	client := Client{
		Server: server,
		Client: httpClient,
	}
	return &client
}

// The interface specification for the client above.
type rawClientInterface interface {
	// TriggerHibernation request
	TriggerHibernationRaw(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// TriggerMaintenance request
	TriggerMaintenanceRaw(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// TriggerReconciliation request
	TriggerReconciliationRaw(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// TriggerRotation request
	TriggerRotationRaw(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// TriggerWakeup request
	TriggerWakeupRaw(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) TriggerHibernationRaw(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewTriggerHibernationRequest(ctx, c.Server, projectID, clusterName)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) TriggerMaintenanceRaw(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewTriggerMaintenanceRequest(ctx, c.Server, projectID, clusterName)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) TriggerReconciliationRaw(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewTriggerReconciliationRequest(ctx, c.Server, projectID, clusterName)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) TriggerRotationRaw(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewTriggerRotationRequest(ctx, c.Server, projectID, clusterName)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) TriggerWakeupRaw(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewTriggerWakeupRequest(ctx, c.Server, projectID, clusterName)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewTriggerHibernationRequest generates requests for TriggerHibernation
func NewTriggerHibernationRequest(ctx context.Context, server string, projectID string, clusterName string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "projectID", runtime.ParamLocationPath, projectID)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "clusterName", runtime.ParamLocationPath, clusterName)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/projects/%s/clusters/%s/hibernate", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewTriggerMaintenanceRequest generates requests for TriggerMaintenance
func NewTriggerMaintenanceRequest(ctx context.Context, server string, projectID string, clusterName string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "projectID", runtime.ParamLocationPath, projectID)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "clusterName", runtime.ParamLocationPath, clusterName)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/projects/%s/clusters/%s/maintenance", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewTriggerReconciliationRequest generates requests for TriggerReconciliation
func NewTriggerReconciliationRequest(ctx context.Context, server string, projectID string, clusterName string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "projectID", runtime.ParamLocationPath, projectID)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "clusterName", runtime.ParamLocationPath, clusterName)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/projects/%s/clusters/%s/reconcile", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewTriggerRotationRequest generates requests for TriggerRotation
func NewTriggerRotationRequest(ctx context.Context, server string, projectID string, clusterName string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "projectID", runtime.ParamLocationPath, projectID)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "clusterName", runtime.ParamLocationPath, clusterName)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/projects/%s/clusters/%s/rotate-credentials", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewTriggerWakeupRequest generates requests for TriggerWakeup
func NewTriggerWakeupRequest(ctx context.Context, server string, projectID string, clusterName string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "projectID", runtime.ParamLocationPath, projectID)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "clusterName", runtime.ParamLocationPath, clusterName)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/projects/%s/clusters/%s/wakeup", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", queryURL.String(), nil)
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

// ClientWithResponses builds on rawClientInterface to offer response payloads
type ClientWithResponses struct {
	rawClientInterface
}

// NewClient creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClient(server string, httpClient contracts.BaseClientInterface) *ClientWithResponses {
	return &ClientWithResponses{NewRawClient(server, httpClient)}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// TriggerHibernation request
	TriggerHibernation(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*TriggerHibernationResponse, error)

	// TriggerMaintenance request
	TriggerMaintenance(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*TriggerMaintenanceResponse, error)

	// TriggerReconciliation request
	TriggerReconciliation(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*TriggerReconciliationResponse, error)

	// TriggerRotation request
	TriggerRotation(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*TriggerRotationResponse, error)

	// TriggerWakeup request
	TriggerWakeup(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*TriggerWakeupResponse, error)
}

type TriggerHibernationResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *map[string]interface{}
	JSON202      *map[string]interface{}
	JSON404      *map[string]interface{}
	JSONDefault  *RuntimeError
	Error        error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r TriggerHibernationResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r TriggerHibernationResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type TriggerMaintenanceResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *map[string]interface{}
	JSON202      *map[string]interface{}
	JSON404      *map[string]interface{}
	JSONDefault  *RuntimeError
	Error        error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r TriggerMaintenanceResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r TriggerMaintenanceResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type TriggerReconciliationResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *map[string]interface{}
	JSON202      *map[string]interface{}
	JSON404      *map[string]interface{}
	JSONDefault  *RuntimeError
	Error        error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r TriggerReconciliationResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r TriggerReconciliationResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type TriggerRotationResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *map[string]interface{}
	JSON202      *map[string]interface{}
	JSON404      *map[string]interface{}
	JSONDefault  *RuntimeError
	Error        error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r TriggerRotationResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r TriggerRotationResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type TriggerWakeupResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *map[string]interface{}
	JSON202      *map[string]interface{}
	JSON404      *map[string]interface{}
	JSONDefault  *RuntimeError
	Error        error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r TriggerWakeupResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r TriggerWakeupResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// TriggerHibernation request returning *TriggerHibernationResponse
func (c *ClientWithResponses) TriggerHibernation(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*TriggerHibernationResponse, error) {
	rsp, err := c.TriggerHibernationRaw(ctx, projectID, clusterName, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseTriggerHibernationResponse(rsp)
}

// TriggerMaintenance request returning *TriggerMaintenanceResponse
func (c *ClientWithResponses) TriggerMaintenance(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*TriggerMaintenanceResponse, error) {
	rsp, err := c.TriggerMaintenanceRaw(ctx, projectID, clusterName, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseTriggerMaintenanceResponse(rsp)
}

// TriggerReconciliation request returning *TriggerReconciliationResponse
func (c *ClientWithResponses) TriggerReconciliation(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*TriggerReconciliationResponse, error) {
	rsp, err := c.TriggerReconciliationRaw(ctx, projectID, clusterName, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseTriggerReconciliationResponse(rsp)
}

// TriggerRotation request returning *TriggerRotationResponse
func (c *ClientWithResponses) TriggerRotation(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*TriggerRotationResponse, error) {
	rsp, err := c.TriggerRotationRaw(ctx, projectID, clusterName, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseTriggerRotationResponse(rsp)
}

// TriggerWakeup request returning *TriggerWakeupResponse
func (c *ClientWithResponses) TriggerWakeup(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*TriggerWakeupResponse, error) {
	rsp, err := c.TriggerWakeupRaw(ctx, projectID, clusterName, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseTriggerWakeupResponse(rsp)
}

// ParseTriggerHibernationResponse parses an HTTP response from a TriggerHibernation call
func (c *ClientWithResponses) ParseTriggerHibernationResponse(rsp *http.Response) (*TriggerHibernationResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &TriggerHibernationResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	response.Error = validate.DefaultResponseErrorHandler(rsp)

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 202:
		var dest map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON202 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest RuntimeError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSONDefault = &dest

	}

	return response, validate.ResponseObject(response)
}

// ParseTriggerMaintenanceResponse parses an HTTP response from a TriggerMaintenance call
func (c *ClientWithResponses) ParseTriggerMaintenanceResponse(rsp *http.Response) (*TriggerMaintenanceResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &TriggerMaintenanceResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	response.Error = validate.DefaultResponseErrorHandler(rsp)

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 202:
		var dest map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON202 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest RuntimeError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSONDefault = &dest

	}

	return response, validate.ResponseObject(response)
}

// ParseTriggerReconciliationResponse parses an HTTP response from a TriggerReconciliation call
func (c *ClientWithResponses) ParseTriggerReconciliationResponse(rsp *http.Response) (*TriggerReconciliationResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &TriggerReconciliationResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	response.Error = validate.DefaultResponseErrorHandler(rsp)

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 202:
		var dest map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON202 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest RuntimeError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSONDefault = &dest

	}

	return response, validate.ResponseObject(response)
}

// ParseTriggerRotationResponse parses an HTTP response from a TriggerRotation call
func (c *ClientWithResponses) ParseTriggerRotationResponse(rsp *http.Response) (*TriggerRotationResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &TriggerRotationResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	response.Error = validate.DefaultResponseErrorHandler(rsp)

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 202:
		var dest map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON202 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest RuntimeError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSONDefault = &dest

	}

	return response, validate.ResponseObject(response)
}

// ParseTriggerWakeupResponse parses an HTTP response from a TriggerWakeup call
func (c *ClientWithResponses) ParseTriggerWakeupResponse(rsp *http.Response) (*TriggerWakeupResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &TriggerWakeupResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	response.Error = validate.DefaultResponseErrorHandler(rsp)

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 202:
		var dest map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON202 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest RuntimeError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSONDefault = &dest

	}

	return response, validate.ResponseObject(response)
}
