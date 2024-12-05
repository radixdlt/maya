package lts

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TransactionRequestBuilder builds and executes requests for operations under \lts\transaction
type TransactionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Construction the construction property
// returns a *TransactionConstructionRequestBuilder when successful
func (m *TransactionRequestBuilder) Construction()(*TransactionConstructionRequestBuilder) {
    return NewTransactionConstructionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewTransactionRequestBuilderInternal instantiates a new TransactionRequestBuilder and sets the default values.
func NewTransactionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TransactionRequestBuilder) {
    m := &TransactionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/lts/transaction", pathParameters),
    }
    return m
}
// NewTransactionRequestBuilder instantiates a new TransactionRequestBuilder and sets the default values.
func NewTransactionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TransactionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTransactionRequestBuilderInternal(urlParams, requestAdapter)
}
// Status the status property
// returns a *TransactionStatusRequestBuilder when successful
func (m *TransactionRequestBuilder) Status()(*TransactionStatusRequestBuilder) {
    return NewTransactionStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Submit the submit property
// returns a *TransactionSubmitRequestBuilder when successful
func (m *TransactionRequestBuilder) Submit()(*TransactionSubmitRequestBuilder) {
    return NewTransactionSubmitRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
