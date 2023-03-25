// Package accesskey provides primitives to interact with the openapi HTTP API.
//
// Code generated by dev.azure.com/schwarzit/schwarzit.odj.core/_git/stackit-client-generator.git version v1.0.23 DO NOT EDIT.
package accesskey

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"

	contracts "github.com/SchwarzIT/community-stackit-go-client/internal/contracts"
	"github.com/SchwarzIT/community-stackit-go-client/internal/helpers/runtime"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
)

// CreateJSONBody defines parameters for Create.
type CreateJSONBody struct {
	// Expires Expiration date. Null means never expires.
	Expires *time.Time `json:"expires,omitempty"`
}

// CreateParams defines parameters for Create.
type CreateParams struct {
	CredentialsGroup *string `form:"credentials-group,omitempty" json:"credentials-group,omitempty"`
}

// DeleteParams defines parameters for Delete.
type DeleteParams struct {
	CredentialsGroup *string `form:"credentials-group,omitempty" json:"credentials-group,omitempty"`
}

// GetParams defines parameters for Get.
type GetParams struct {
	CredentialsGroup *string `form:"credentials-group,omitempty" json:"credentials-group,omitempty"`
}

// CreateJSONRequestBody defines body for Create for application/json ContentType.
type CreateJSONRequestBody CreateJSONBody

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Client which conforms to the OpenAPI3 specification for this service.
type Client[K contracts.ClientFlowConfig] struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client contracts.ClientInterface[K]
}

// NewRawClient Creates a new Client, with reasonable defaults
func NewRawClient[K contracts.ClientFlowConfig](server string, httpClient contracts.ClientInterface[K]) *Client[K] {
	// create a client with sane default values
	client := Client[K]{
		Server: server,
		Client: httpClient,
	}
	return &client
}

// The interface specification for the client above.
type rawClientInterface interface {
	// Create request with any body
	CreateRawWithBody(ctx context.Context, projectID string, params *CreateParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateRaw(ctx context.Context, projectID string, params *CreateParams, body CreateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// Delete request
	DeleteRaw(ctx context.Context, projectID string, keyID string, params *DeleteParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// Get request
	GetRaw(ctx context.Context, projectID string, params *GetParams, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client[K]) CreateRawWithBody(ctx context.Context, projectID string, params *CreateParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateRequestWithBody(ctx, c.Server, projectID, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client[K]) CreateRaw(ctx context.Context, projectID string, params *CreateParams, body CreateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateRequest(ctx, c.Server, projectID, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client[K]) DeleteRaw(ctx context.Context, projectID string, keyID string, params *DeleteParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteRequest(ctx, c.Server, projectID, keyID, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client[K]) GetRaw(ctx context.Context, projectID string, params *GetParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetRequest(ctx, c.Server, projectID, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewCreateRequest calls the generic Create builder with application/json body
func NewCreateRequest(ctx context.Context, server string, projectID string, params *CreateParams, body CreateJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateRequestWithBody(ctx, server, projectID, params, "application/json", bodyReader)
}

// NewCreateRequestWithBody generates requests for Create with any type of body
func NewCreateRequestWithBody(ctx context.Context, server string, projectID string, params *CreateParams, contentType string, body io.Reader) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/v1/project/%s/access-key", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.CredentialsGroup != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "credentials-group", runtime.ParamLocationQuery, *params.CredentialsGroup); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequestWithContext(ctx, "POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewDeleteRequest generates requests for Delete
func NewDeleteRequest(ctx context.Context, server string, projectID string, keyID string, params *DeleteParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "projectID", runtime.ParamLocationPath, projectID)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "keyID", runtime.ParamLocationPath, keyID)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/project/%s/access-key/%s", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.CredentialsGroup != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "credentials-group", runtime.ParamLocationQuery, *params.CredentialsGroup); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequestWithContext(ctx, "DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetRequest generates requests for Get
func NewGetRequest(ctx context.Context, server string, projectID string, params *GetParams) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/v1/project/%s/access-keys", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.CredentialsGroup != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "credentials-group", runtime.ParamLocationQuery, *params.CredentialsGroup); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequestWithContext(ctx, "GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client[K]) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on rawClientInterface to offer response payloads
type ClientWithResponses[K contracts.ClientFlowConfig] struct {
	rawClientInterface
}

// NewClient creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClient[K contracts.ClientFlowConfig](server string, httpClient contracts.ClientInterface[K]) *ClientWithResponses[K] {
	return &ClientWithResponses[K]{NewRawClient(server, httpClient)}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface[K contracts.ClientFlowConfig] interface {
	// Create request with any body
	CreateWithBody(ctx context.Context, projectID string, params *CreateParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateResponse, error)

	Create(ctx context.Context, projectID string, params *CreateParams, body CreateJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateResponse, error)

	// Delete request
	Delete(ctx context.Context, projectID string, keyID string, params *DeleteParams, reqEditors ...RequestEditorFn) (*DeleteResponse, error)

	// Get request
	Get(ctx context.Context, projectID string, params *GetParams, reqEditors ...RequestEditorFn) (*GetResponse, error)
}

type CreateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		// AccessKey Access key
		AccessKey string `json:"accessKey"`

		// DisplayName Obfuscated access key
		DisplayName string `json:"displayName"`

		// Expires Expiration date. Null means never expires.
		Expires string `json:"expires"`

		// KeyId Identifies the pair of access key and secret access key for deletion
		KeyID string `json:"keyId"`

		// Project Project ID
		Project string `json:"project"`

		// SecretAccessKey Secret access key
		SecretAccessKey string `json:"secretAccessKey"`
	}
	JSON403 *struct {
		Details []struct {
			Key string `json:"key"`
			Msg string `json:"msg"`
		} `json:"detail"`
	}
	JSON404 *struct {
		Details []struct {
			Key string `json:"key"`
			Msg string `json:"msg"`
		} `json:"detail"`
	}
	JSON422 *struct {
		Details *[]struct {
			Loc  []string `json:"loc"`
			Msg  string   `json:"msg"`
			Type string   `json:"type"`
		} `json:"detail,omitempty"`
	}
	JSON500 *struct {
		Details []struct {
			Key string `json:"key"`
			Msg string `json:"msg"`
		} `json:"detail"`
	}
	Error error // Aggregated error
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
	JSON200      *struct {
		// KeyId Identifies the pair of access key and secret access key for deletion
		KeyID string `json:"keyId"`

		// Project Project ID
		Project string `json:"project"`
	}
	JSON307 *struct {
		Details []struct {
			Key string `json:"key"`
			Msg string `json:"msg"`
		} `json:"detail"`
	}
	JSON403 *struct {
		Details []struct {
			Key string `json:"key"`
			Msg string `json:"msg"`
		} `json:"detail"`
	}
	JSON404 *struct {
		Details []struct {
			Key string `json:"key"`
			Msg string `json:"msg"`
		} `json:"detail"`
	}
	JSON422 *struct {
		Details *[]struct {
			Loc  []string `json:"loc"`
			Msg  string   `json:"msg"`
			Type string   `json:"type"`
		} `json:"detail,omitempty"`
	}
	JSON500 *struct {
		Details []struct {
			Key string `json:"key"`
			Msg string `json:"msg"`
		} `json:"detail"`
	}
	Error error // Aggregated error
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
	JSON200      *struct {
		AccessKeys []struct {
			DisplayName string `json:"displayName"`
			Expires     string `json:"expires"`

			// KeyId Identifies the pair of access key and secret access key for deletion
			KeyID string `json:"keyId"`
		} `json:"accessKeys"`

		// Project Project ID
		Project string `json:"project"`
	}
	JSON403 *struct {
		Details []struct {
			Key string `json:"key"`
			Msg string `json:"msg"`
		} `json:"detail"`
	}
	JSON404 *struct {
		Details []struct {
			Key string `json:"key"`
			Msg string `json:"msg"`
		} `json:"detail"`
	}
	JSON422 *struct {
		Details *[]struct {
			Loc  []string `json:"loc"`
			Msg  string   `json:"msg"`
			Type string   `json:"type"`
		} `json:"detail,omitempty"`
	}
	JSON500 *struct {
		Details []struct {
			Key string `json:"key"`
			Msg string `json:"msg"`
		} `json:"detail"`
	}
	Error error // Aggregated error
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

// CreateWithBody request with arbitrary body returning *CreateResponse
func (c *ClientWithResponses[K]) CreateWithBody(ctx context.Context, projectID string, params *CreateParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateResponse, error) {
	rsp, err := c.CreateRawWithBody(ctx, projectID, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseCreateResponse(rsp)
}

func (c *ClientWithResponses[K]) Create(ctx context.Context, projectID string, params *CreateParams, body CreateJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateResponse, error) {
	rsp, err := c.CreateRaw(ctx, projectID, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseCreateResponse(rsp)
}

// Delete request returning *DeleteResponse
func (c *ClientWithResponses[K]) Delete(ctx context.Context, projectID string, keyID string, params *DeleteParams, reqEditors ...RequestEditorFn) (*DeleteResponse, error) {
	rsp, err := c.DeleteRaw(ctx, projectID, keyID, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseDeleteResponse(rsp)
}

// Get request returning *GetResponse
func (c *ClientWithResponses[K]) Get(ctx context.Context, projectID string, params *GetParams, reqEditors ...RequestEditorFn) (*GetResponse, error) {
	rsp, err := c.GetRaw(ctx, projectID, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseGetResponse(rsp)
}

// ParseCreateResponse parses an HTTP response from a Create call
func (c *ClientWithResponses[K]) ParseCreateResponse(rsp *http.Response) (*CreateResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	response.Error = validate.DefaultResponseErrorHandler(rsp)

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			// AccessKey Access key
			AccessKey string `json:"accessKey"`

			// DisplayName Obfuscated access key
			DisplayName string `json:"displayName"`

			// Expires Expiration date. Null means never expires.
			Expires string `json:"expires"`

			// KeyId Identifies the pair of access key and secret access key for deletion
			KeyID string `json:"keyId"`

			// Project Project ID
			Project string `json:"project"`

			// SecretAccessKey Secret access key
			SecretAccessKey string `json:"secretAccessKey"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON201 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest struct {
			Details []struct {
				Key string `json:"key"`
				Msg string `json:"msg"`
			} `json:"detail"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest struct {
			Details []struct {
				Key string `json:"key"`
				Msg string `json:"msg"`
			} `json:"detail"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest struct {
			Details *[]struct {
				Loc  []string `json:"loc"`
				Msg  string   `json:"msg"`
				Type string   `json:"type"`
			} `json:"detail,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON422 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest struct {
			Details []struct {
				Key string `json:"key"`
				Msg string `json:"msg"`
			} `json:"detail"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON500 = &dest

	}

	return response, validate.ResponseObject(response)
}

// ParseDeleteResponse parses an HTTP response from a Delete call
func (c *ClientWithResponses[K]) ParseDeleteResponse(rsp *http.Response) (*DeleteResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	response.Error = validate.DefaultResponseErrorHandler(rsp)

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			// KeyId Identifies the pair of access key and secret access key for deletion
			KeyID string `json:"keyId"`

			// Project Project ID
			Project string `json:"project"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 307:
		var dest struct {
			Details []struct {
				Key string `json:"key"`
				Msg string `json:"msg"`
			} `json:"detail"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON307 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest struct {
			Details []struct {
				Key string `json:"key"`
				Msg string `json:"msg"`
			} `json:"detail"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest struct {
			Details []struct {
				Key string `json:"key"`
				Msg string `json:"msg"`
			} `json:"detail"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest struct {
			Details *[]struct {
				Loc  []string `json:"loc"`
				Msg  string   `json:"msg"`
				Type string   `json:"type"`
			} `json:"detail,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON422 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest struct {
			Details []struct {
				Key string `json:"key"`
				Msg string `json:"msg"`
			} `json:"detail"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON500 = &dest

	}

	return response, validate.ResponseObject(response)
}

// ParseGetResponse parses an HTTP response from a Get call
func (c *ClientWithResponses[K]) ParseGetResponse(rsp *http.Response) (*GetResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	response.Error = validate.DefaultResponseErrorHandler(rsp)

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			AccessKeys []struct {
				DisplayName string `json:"displayName"`
				Expires     string `json:"expires"`

				// KeyId Identifies the pair of access key and secret access key for deletion
				KeyID string `json:"keyId"`
			} `json:"accessKeys"`

			// Project Project ID
			Project string `json:"project"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest struct {
			Details []struct {
				Key string `json:"key"`
				Msg string `json:"msg"`
			} `json:"detail"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest struct {
			Details []struct {
				Key string `json:"key"`
				Msg string `json:"msg"`
			} `json:"detail"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest struct {
			Details *[]struct {
				Loc  []string `json:"loc"`
				Msg  string   `json:"msg"`
				Type string   `json:"type"`
			} `json:"detail,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON422 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest struct {
			Details []struct {
				Key string `json:"key"`
				Msg string `json:"msg"`
			} `json:"detail"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON500 = &dest

	}

	return response, validate.ResponseObject(response)
}