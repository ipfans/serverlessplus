package tencent

import (
	"context"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"

	"github.com/akutz/memconn"
	"github.com/tencentyun/scf-go-lib/events"
	"github.com/tencentyun/scf-go-lib/functioncontext"
)

var defaultClient = &http.Client{
	Transport: &http.Transport{
		DialContext: func(ctx context.Context, _, _ string) (net.Conn, error) {
			return memconn.DialContext(ctx, "memu", serverMemAddr)
		},
	},
}

func translateRequest(ctx context.Context, r *events.APIGatewayRequest) (req *http.Request) {
	u := &url.URL{
		Scheme:   "http",
		Host:     serverMemAddr,
		Path:     r.Path,
		RawPath:  r.Path,
		RawQuery: url.Values(r.QueryString).Encode(),
	}
	rBody := strings.NewReader(r.Body)
	req, _ = http.NewRequest(r.Method, u.String(), rBody)
	for name, value := range r.Headers {
		req.Header.Set(http.CanonicalHeaderKey(name), value)
	}
	funcCtx, ok := functioncontext.FromContext(ctx)
	if ok {
		req.Header.Add(http.CanonicalHeaderKey("x-request-id"), funcCtx.RequestID)
		req.Header.Add(http.CanonicalHeaderKey("x-scf-requestid"), funcCtx.RequestID)
	}
	req.Header.Add(http.CanonicalHeaderKey("x-apigateway-serviceid"), r.Context.ServiceID)
	req.Header.Add(http.CanonicalHeaderKey("x-apigateway-requestid"), r.Context.RequestID)
	req.Header.Add(http.CanonicalHeaderKey("x-apigateway-method"), r.Context.Method)
	req.Header.Add(http.CanonicalHeaderKey("x-apigateway-path"), r.Context.Path)
	req.Header.Add(http.CanonicalHeaderKey("x-apigateway-sourceip"), r.Context.SourceIP)
	req.Header.Add(http.CanonicalHeaderKey("x-forwarded-for"), r.Context.SourceIP)
	req.Header.Add(http.CanonicalHeaderKey("x-apigateway-stage"), r.Context.Stage)
	if r.Context.Identity.SecretID != nil {
		req.Header.Add(http.CanonicalHeaderKey("x-apigateway-secretid"), *r.Context.Identity.SecretID)
	}
	return
}

func translateResponse(r *http.Response) (resp *events.APIGatewayResponse) {
	resp = &events.APIGatewayResponse{
		StatusCode: r.StatusCode,
		Headers:    make(map[string]string),
	}
	resp.Headers[http.CanonicalHeaderKey("x-powered-by")] = "serverlessplus"
	for name, values := range r.Header {
		resp.Headers[name] = values[0]
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resp.StatusCode = 500
		return
	}
	resp.Body = string(body)
	return
}

func toHTTPRequest(ctx context.Context, apiReq *events.APIGatewayRequest) (apiResp *events.APIGatewayResponse, err error) {
	req := translateRequest(ctx, apiReq)
	resp, err := defaultClient.Do(req)
	if err != nil {
		return &events.APIGatewayResponse{StatusCode: 500}, err
	}
	return translateResponse(resp), err

}
