package stream

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// StreamRequestBuilder builds and executes requests for operations under \stream
type StreamRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewStreamRequestBuilderInternal instantiates a new StreamRequestBuilder and sets the default values.
func NewStreamRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StreamRequestBuilder) {
    m := &StreamRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/stream", pathParameters),
    }
    return m
}
// NewStreamRequestBuilder instantiates a new StreamRequestBuilder and sets the default values.
func NewStreamRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StreamRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewStreamRequestBuilderInternal(urlParams, requestAdapter)
}
// Proofs the proofs property
// returns a *ProofsRequestBuilder when successful
func (m *StreamRequestBuilder) Proofs()(*ProofsRequestBuilder) {
    return NewProofsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Transactions the transactions property
// returns a *TransactionsRequestBuilder when successful
func (m *StreamRequestBuilder) Transactions()(*TransactionsRequestBuilder) {
    return NewTransactionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
