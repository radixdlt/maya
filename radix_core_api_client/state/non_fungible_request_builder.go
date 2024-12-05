package state

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e "github.com/radixdlt/maya/radix_core_api_client/models"
)

// NonFungibleRequestBuilder builds and executes requests for operations under \state\non-fungible
type NonFungibleRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NonFungibleRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type NonFungibleRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewNonFungibleRequestBuilderInternal instantiates a new NonFungibleRequestBuilder and sets the default values.
func NewNonFungibleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*NonFungibleRequestBuilder) {
    m := &NonFungibleRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/state/non-fungible", pathParameters),
    }
    return m
}
// NewNonFungibleRequestBuilder instantiates a new NonFungibleRequestBuilder and sets the default values.
func NewNonFungibleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*NonFungibleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewNonFungibleRequestBuilderInternal(urlParams, requestAdapter)
}
// Post reads the data associated with a single Non-Fungible Unit under a Non-Fungible Resource.
// returns a StateNonFungibleResponseable when successful
// returns a BasicErrorResponse error when the service returns a 400 status code
// returns a BasicErrorResponse error when the service returns a 404 status code
// returns a BasicErrorResponse error when the service returns a 500 status code
func (m *NonFungibleRequestBuilder) Post(ctx context.Context, body i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.StateNonFungibleRequestable, requestConfiguration *NonFungibleRequestBuilderPostRequestConfiguration)(i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.StateNonFungibleResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.CreateBasicErrorResponseFromDiscriminatorValue,
        "404": i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.CreateBasicErrorResponseFromDiscriminatorValue,
        "500": i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.CreateBasicErrorResponseFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.CreateStateNonFungibleResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.StateNonFungibleResponseable), nil
}
// ToPostRequestInformation reads the data associated with a single Non-Fungible Unit under a Non-Fungible Resource.
// returns a *RequestInformation when successful
func (m *NonFungibleRequestBuilder) ToPostRequestInformation(ctx context.Context, body i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.StateNonFungibleRequestable, requestConfiguration *NonFungibleRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *NonFungibleRequestBuilder when successful
func (m *NonFungibleRequestBuilder) WithUrl(rawUrl string)(*NonFungibleRequestBuilder) {
    return NewNonFungibleRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
