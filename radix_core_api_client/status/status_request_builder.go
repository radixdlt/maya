package status

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// StatusRequestBuilder builds and executes requests for operations under \status
type StatusRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewStatusRequestBuilderInternal instantiates a new StatusRequestBuilder and sets the default values.
func NewStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StatusRequestBuilder) {
    m := &StatusRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/status", pathParameters),
    }
    return m
}
// NewStatusRequestBuilder instantiates a new StatusRequestBuilder and sets the default values.
func NewStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// NetworkConfiguration the networkConfiguration property
// returns a *NetworkConfigurationRequestBuilder when successful
func (m *StatusRequestBuilder) NetworkConfiguration()(*NetworkConfigurationRequestBuilder) {
    return NewNetworkConfigurationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NetworkStatus the networkStatus property
// returns a *NetworkStatusRequestBuilder when successful
func (m *StatusRequestBuilder) NetworkStatus()(*NetworkStatusRequestBuilder) {
    return NewNetworkStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Scenarios the scenarios property
// returns a *ScenariosRequestBuilder when successful
func (m *StatusRequestBuilder) Scenarios()(*ScenariosRequestBuilder) {
    return NewScenariosRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
