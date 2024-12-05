package transaction

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e "github.com/radixdlt/maya/radix_core_api_client/models"
)

// ParseRequestBuilder builds and executes requests for operations under \transaction\parse
type ParseRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ParseRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ParseRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewParseRequestBuilderInternal instantiates a new ParseRequestBuilder and sets the default values.
func NewParseRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ParseRequestBuilder) {
    m := &ParseRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/transaction/parse", pathParameters),
    }
    return m
}
// NewParseRequestBuilder instantiates a new ParseRequestBuilder and sets the default values.
func NewParseRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ParseRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewParseRequestBuilderInternal(urlParams, requestAdapter)
}
// Post extracts the contents and hashes of various types of transaction payloads, or sub-payloads.
// returns a TransactionParseResponseable when successful
// returns a BasicErrorResponse error when the service returns a 400 status code
// returns a BasicErrorResponse error when the service returns a 500 status code
func (m *ParseRequestBuilder) Post(ctx context.Context, body i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.TransactionParseRequestable, requestConfiguration *ParseRequestBuilderPostRequestConfiguration)(i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.TransactionParseResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.CreateBasicErrorResponseFromDiscriminatorValue,
        "500": i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.CreateBasicErrorResponseFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.CreateTransactionParseResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.TransactionParseResponseable), nil
}
// ToPostRequestInformation extracts the contents and hashes of various types of transaction payloads, or sub-payloads.
// returns a *RequestInformation when successful
func (m *ParseRequestBuilder) ToPostRequestInformation(ctx context.Context, body i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.TransactionParseRequestable, requestConfiguration *ParseRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ParseRequestBuilder when successful
func (m *ParseRequestBuilder) WithUrl(rawUrl string)(*ParseRequestBuilder) {
    return NewParseRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
