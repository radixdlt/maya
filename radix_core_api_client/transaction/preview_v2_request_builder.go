package transaction

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e "github.com/radixdlt/maya/radix_core_api_client/models"
)

// PreviewV2RequestBuilder builds and executes requests for operations under \transaction\preview-v2
type PreviewV2RequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PreviewV2RequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PreviewV2RequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPreviewV2RequestBuilderInternal instantiates a new PreviewV2RequestBuilder and sets the default values.
func NewPreviewV2RequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PreviewV2RequestBuilder) {
    m := &PreviewV2RequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/transaction/preview-v2", pathParameters),
    }
    return m
}
// NewPreviewV2RequestBuilder instantiates a new PreviewV2RequestBuilder and sets the default values.
func NewPreviewV2RequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PreviewV2RequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPreviewV2RequestBuilderInternal(urlParams, requestAdapter)
}
// Post previews a transaction against the latest network state, and returns the preview receipt.If the node has enabled it, you may be able to also preview against recent network state.This endpoint supports V2 transactions (and beyond). If you still need to preview V1transactions, you should use the `/preview` endpoint instead.
// returns a TransactionPreviewV2Responseable when successful
// returns a TransactionPreviewV2ErrorResponse error when the service returns a 400 status code
// returns a TransactionPreviewV2ErrorResponse error when the service returns a 500 status code
func (m *PreviewV2RequestBuilder) Post(ctx context.Context, body i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.TransactionPreviewV2Requestable, requestConfiguration *PreviewV2RequestBuilderPostRequestConfiguration)(i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.TransactionPreviewV2Responseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.CreateTransactionPreviewV2ErrorResponseFromDiscriminatorValue,
        "500": i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.CreateTransactionPreviewV2ErrorResponseFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.CreateTransactionPreviewV2ResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.TransactionPreviewV2Responseable), nil
}
// ToPostRequestInformation previews a transaction against the latest network state, and returns the preview receipt.If the node has enabled it, you may be able to also preview against recent network state.This endpoint supports V2 transactions (and beyond). If you still need to preview V1transactions, you should use the `/preview` endpoint instead.
// returns a *RequestInformation when successful
func (m *PreviewV2RequestBuilder) ToPostRequestInformation(ctx context.Context, body i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.TransactionPreviewV2Requestable, requestConfiguration *PreviewV2RequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PreviewV2RequestBuilder when successful
func (m *PreviewV2RequestBuilder) WithUrl(rawUrl string)(*PreviewV2RequestBuilder) {
    return NewPreviewV2RequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
