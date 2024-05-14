package status

import (
	"context"

	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	ieb2700c3462f1a9e73e3b7d2d4d35314fa43b94ba0380cc504a8f60edfdf3455 "github.com/radixdlt/maya/radix_core_api_client/models"
)

// NetworkConfigurationRequestBuilder builds and executes requests for operations under \status\network-configuration
type NetworkConfigurationRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NetworkConfigurationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type NetworkConfigurationRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewNetworkConfigurationRequestBuilderInternal instantiates a new NetworkConfigurationRequestBuilder and sets the default values.
func NewNetworkConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *NetworkConfigurationRequestBuilder {
	m := &NetworkConfigurationRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/status/network-configuration", pathParameters),
	}
	return m
}

// NewNetworkConfigurationRequestBuilder instantiates a new NetworkConfigurationRequestBuilder and sets the default values.
func NewNetworkConfigurationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *NetworkConfigurationRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewNetworkConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}

// Post returns the network configuration of the network the node is connected to.
// returns a NetworkConfigurationResponseable when successful
// returns a BasicErrorResponse error when the service returns a 500 status code
func (m *NetworkConfigurationRequestBuilder) Post(ctx context.Context, requestConfiguration *NetworkConfigurationRequestBuilderPostRequestConfiguration) (ieb2700c3462f1a9e73e3b7d2d4d35314fa43b94ba0380cc504a8f60edfdf3455.NetworkConfigurationResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"500": ieb2700c3462f1a9e73e3b7d2d4d35314fa43b94ba0380cc504a8f60edfdf3455.CreateBasicErrorResponseFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ieb2700c3462f1a9e73e3b7d2d4d35314fa43b94ba0380cc504a8f60edfdf3455.CreateNetworkConfigurationResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ieb2700c3462f1a9e73e3b7d2d4d35314fa43b94ba0380cc504a8f60edfdf3455.NetworkConfigurationResponseable), nil
}

// ToPostRequestInformation returns the network configuration of the network the node is connected to.
// returns a *RequestInformation when successful
func (m *NetworkConfigurationRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *NetworkConfigurationRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *NetworkConfigurationRequestBuilder when successful
func (m *NetworkConfigurationRequestBuilder) WithUrl(rawUrl string) *NetworkConfigurationRequestBuilder {
	return NewNetworkConfigurationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
