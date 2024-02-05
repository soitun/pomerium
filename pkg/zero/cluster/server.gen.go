// Package cluster provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package cluster

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
	strictnethttp "github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /bootstrap)
	GetClusterBootstrapConfig(w http.ResponseWriter, r *http.Request)

	// (GET /bundles)
	GetClusterResourceBundles(w http.ResponseWriter, r *http.Request)

	// (GET /bundles/{bundleId}/download)
	DownloadClusterResourceBundle(w http.ResponseWriter, r *http.Request, bundleId BundleId)

	// (POST /bundles/{bundleId}/status)
	ReportClusterResourceBundleStatus(w http.ResponseWriter, r *http.Request, bundleId BundleId)

	// (POST /exchangeToken)
	ExchangeClusterIdentityToken(w http.ResponseWriter, r *http.Request)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// (GET /bootstrap)
func (_ Unimplemented) GetClusterBootstrapConfig(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /bundles)
func (_ Unimplemented) GetClusterResourceBundles(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /bundles/{bundleId}/download)
func (_ Unimplemented) DownloadClusterResourceBundle(w http.ResponseWriter, r *http.Request, bundleId BundleId) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (POST /bundles/{bundleId}/status)
func (_ Unimplemented) ReportClusterResourceBundleStatus(w http.ResponseWriter, r *http.Request, bundleId BundleId) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (POST /exchangeToken)
func (_ Unimplemented) ExchangeClusterIdentityToken(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetClusterBootstrapConfig operation middleware
func (siw *ServerInterfaceWrapper) GetClusterBootstrapConfig(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetClusterBootstrapConfig(w, r)
	}))

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetClusterResourceBundles operation middleware
func (siw *ServerInterfaceWrapper) GetClusterResourceBundles(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetClusterResourceBundles(w, r)
	}))

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DownloadClusterResourceBundle operation middleware
func (siw *ServerInterfaceWrapper) DownloadClusterResourceBundle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "bundleId" -------------
	var bundleId BundleId

	err = runtime.BindStyledParameterWithOptions("simple", "bundleId", chi.URLParam(r, "bundleId"), &bundleId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "bundleId", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DownloadClusterResourceBundle(w, r, bundleId)
	}))

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// ReportClusterResourceBundleStatus operation middleware
func (siw *ServerInterfaceWrapper) ReportClusterResourceBundleStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "bundleId" -------------
	var bundleId BundleId

	err = runtime.BindStyledParameterWithOptions("simple", "bundleId", chi.URLParam(r, "bundleId"), &bundleId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "bundleId", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ReportClusterResourceBundleStatus(w, r, bundleId)
	}))

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// ExchangeClusterIdentityToken operation middleware
func (siw *ServerInterfaceWrapper) ExchangeClusterIdentityToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ExchangeClusterIdentityToken(w, r)
	}))

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/bootstrap", wrapper.GetClusterBootstrapConfig)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/bundles", wrapper.GetClusterResourceBundles)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/bundles/{bundleId}/download", wrapper.DownloadClusterResourceBundle)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/bundles/{bundleId}/status", wrapper.ReportClusterResourceBundleStatus)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/exchangeToken", wrapper.ExchangeClusterIdentityToken)
	})

	return r
}

type GetClusterBootstrapConfigRequestObject struct {
}

type GetClusterBootstrapConfigResponseObject interface {
	VisitGetClusterBootstrapConfigResponse(w http.ResponseWriter) error
}

type GetClusterBootstrapConfig200JSONResponse GetBootstrapConfigResponse

func (response GetClusterBootstrapConfig200JSONResponse) VisitGetClusterBootstrapConfigResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetClusterBootstrapConfig400JSONResponse ErrorResponse

func (response GetClusterBootstrapConfig400JSONResponse) VisitGetClusterBootstrapConfigResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type GetClusterBootstrapConfig500JSONResponse ErrorResponse

func (response GetClusterBootstrapConfig500JSONResponse) VisitGetClusterBootstrapConfigResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type GetClusterResourceBundlesRequestObject struct {
}

type GetClusterResourceBundlesResponseObject interface {
	VisitGetClusterResourceBundlesResponse(w http.ResponseWriter) error
}

type GetClusterResourceBundles200JSONResponse GetBundlesResponse

func (response GetClusterResourceBundles200JSONResponse) VisitGetClusterResourceBundlesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetClusterResourceBundles400JSONResponse ErrorResponse

func (response GetClusterResourceBundles400JSONResponse) VisitGetClusterResourceBundlesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type GetClusterResourceBundles500JSONResponse ErrorResponse

func (response GetClusterResourceBundles500JSONResponse) VisitGetClusterResourceBundlesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type DownloadClusterResourceBundleRequestObject struct {
	BundleId BundleId `json:"bundleId"`
}

type DownloadClusterResourceBundleResponseObject interface {
	VisitDownloadClusterResourceBundleResponse(w http.ResponseWriter) error
}

type DownloadClusterResourceBundle200JSONResponse DownloadBundleResponse

func (response DownloadClusterResourceBundle200JSONResponse) VisitDownloadClusterResourceBundleResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type DownloadClusterResourceBundle400JSONResponse ErrorResponse

func (response DownloadClusterResourceBundle400JSONResponse) VisitDownloadClusterResourceBundleResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type DownloadClusterResourceBundle404JSONResponse ErrorResponse

func (response DownloadClusterResourceBundle404JSONResponse) VisitDownloadClusterResourceBundleResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type DownloadClusterResourceBundle500JSONResponse ErrorResponse

func (response DownloadClusterResourceBundle500JSONResponse) VisitDownloadClusterResourceBundleResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type ReportClusterResourceBundleStatusRequestObject struct {
	BundleId BundleId `json:"bundleId"`
	Body     *ReportClusterResourceBundleStatusJSONRequestBody
}

type ReportClusterResourceBundleStatusResponseObject interface {
	VisitReportClusterResourceBundleStatusResponse(w http.ResponseWriter) error
}

type ReportClusterResourceBundleStatus204Response struct {
}

func (response ReportClusterResourceBundleStatus204Response) VisitReportClusterResourceBundleStatusResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type ReportClusterResourceBundleStatus400JSONResponse ErrorResponse

func (response ReportClusterResourceBundleStatus400JSONResponse) VisitReportClusterResourceBundleStatusResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type ReportClusterResourceBundleStatus500JSONResponse ErrorResponse

func (response ReportClusterResourceBundleStatus500JSONResponse) VisitReportClusterResourceBundleStatusResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type ExchangeClusterIdentityTokenRequestObject struct {
	Body *ExchangeClusterIdentityTokenJSONRequestBody
}

type ExchangeClusterIdentityTokenResponseObject interface {
	VisitExchangeClusterIdentityTokenResponse(w http.ResponseWriter) error
}

type ExchangeClusterIdentityToken200JSONResponse ExchangeTokenResponse

func (response ExchangeClusterIdentityToken200JSONResponse) VisitExchangeClusterIdentityTokenResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type ExchangeClusterIdentityToken400JSONResponse ErrorResponse

func (response ExchangeClusterIdentityToken400JSONResponse) VisitExchangeClusterIdentityTokenResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type ExchangeClusterIdentityToken500JSONResponse ErrorResponse

func (response ExchangeClusterIdentityToken500JSONResponse) VisitExchangeClusterIdentityTokenResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (GET /bootstrap)
	GetClusterBootstrapConfig(ctx context.Context, request GetClusterBootstrapConfigRequestObject) (GetClusterBootstrapConfigResponseObject, error)

	// (GET /bundles)
	GetClusterResourceBundles(ctx context.Context, request GetClusterResourceBundlesRequestObject) (GetClusterResourceBundlesResponseObject, error)

	// (GET /bundles/{bundleId}/download)
	DownloadClusterResourceBundle(ctx context.Context, request DownloadClusterResourceBundleRequestObject) (DownloadClusterResourceBundleResponseObject, error)

	// (POST /bundles/{bundleId}/status)
	ReportClusterResourceBundleStatus(ctx context.Context, request ReportClusterResourceBundleStatusRequestObject) (ReportClusterResourceBundleStatusResponseObject, error)

	// (POST /exchangeToken)
	ExchangeClusterIdentityToken(ctx context.Context, request ExchangeClusterIdentityTokenRequestObject) (ExchangeClusterIdentityTokenResponseObject, error)
}

type StrictHandlerFunc = strictnethttp.StrictHTTPHandlerFunc
type StrictMiddlewareFunc = strictnethttp.StrictHTTPMiddlewareFunc

type StrictHTTPServerOptions struct {
	RequestErrorHandlerFunc  func(w http.ResponseWriter, r *http.Request, err error)
	ResponseErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}}
}

func NewStrictHandlerWithOptions(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc, options StrictHTTPServerOptions) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: options}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
	options     StrictHTTPServerOptions
}

// GetClusterBootstrapConfig operation middleware
func (sh *strictHandler) GetClusterBootstrapConfig(w http.ResponseWriter, r *http.Request) {
	var request GetClusterBootstrapConfigRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetClusterBootstrapConfig(ctx, request.(GetClusterBootstrapConfigRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetClusterBootstrapConfig")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetClusterBootstrapConfigResponseObject); ok {
		if err := validResponse.VisitGetClusterBootstrapConfigResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetClusterResourceBundles operation middleware
func (sh *strictHandler) GetClusterResourceBundles(w http.ResponseWriter, r *http.Request) {
	var request GetClusterResourceBundlesRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetClusterResourceBundles(ctx, request.(GetClusterResourceBundlesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetClusterResourceBundles")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetClusterResourceBundlesResponseObject); ok {
		if err := validResponse.VisitGetClusterResourceBundlesResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// DownloadClusterResourceBundle operation middleware
func (sh *strictHandler) DownloadClusterResourceBundle(w http.ResponseWriter, r *http.Request, bundleId BundleId) {
	var request DownloadClusterResourceBundleRequestObject

	request.BundleId = bundleId

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.DownloadClusterResourceBundle(ctx, request.(DownloadClusterResourceBundleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DownloadClusterResourceBundle")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(DownloadClusterResourceBundleResponseObject); ok {
		if err := validResponse.VisitDownloadClusterResourceBundleResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// ReportClusterResourceBundleStatus operation middleware
func (sh *strictHandler) ReportClusterResourceBundleStatus(w http.ResponseWriter, r *http.Request, bundleId BundleId) {
	var request ReportClusterResourceBundleStatusRequestObject

	request.BundleId = bundleId

	var body ReportClusterResourceBundleStatusJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.ReportClusterResourceBundleStatus(ctx, request.(ReportClusterResourceBundleStatusRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ReportClusterResourceBundleStatus")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(ReportClusterResourceBundleStatusResponseObject); ok {
		if err := validResponse.VisitReportClusterResourceBundleStatusResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// ExchangeClusterIdentityToken operation middleware
func (sh *strictHandler) ExchangeClusterIdentityToken(w http.ResponseWriter, r *http.Request) {
	var request ExchangeClusterIdentityTokenRequestObject

	var body ExchangeClusterIdentityTokenJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.ExchangeClusterIdentityToken(ctx, request.(ExchangeClusterIdentityTokenRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ExchangeClusterIdentityToken")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(ExchangeClusterIdentityTokenResponseObject); ok {
		if err := validResponse.VisitExchangeClusterIdentityTokenResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}
