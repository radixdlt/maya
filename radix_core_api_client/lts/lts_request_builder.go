package lts

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// LtsRequestBuilder builds and executes requests for operations under \lts
type LtsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewLtsRequestBuilderInternal instantiates a new LtsRequestBuilder and sets the default values.
func NewLtsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LtsRequestBuilder) {
    m := &LtsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/lts", pathParameters),
    }
    return m
}
// NewLtsRequestBuilder instantiates a new LtsRequestBuilder and sets the default values.
func NewLtsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LtsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLtsRequestBuilderInternal(urlParams, requestAdapter)
}
// State the state property
// returns a *StateRequestBuilder when successful
func (m *LtsRequestBuilder) State()(*StateRequestBuilder) {
    return NewStateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Stream the stream property
// returns a *StreamRequestBuilder when successful
func (m *LtsRequestBuilder) Stream()(*StreamRequestBuilder) {
    return NewStreamRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Transaction the transaction property
// returns a *TransactionRequestBuilder when successful
func (m *LtsRequestBuilder) Transaction()(*TransactionRequestBuilder) {
    return NewTransactionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
