package lts

import (
	"context"

	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	ieb2700c3462f1a9e73e3b7d2d4d35314fa43b94ba0380cc504a8f60edfdf3455 "gitlab.com/mayachain/mayanode/bifrost/pkg/chainclients/radix/coreapi/client/models"
)

// StateAccountFungibleResourceBalanceRequestBuilder builds and executes requests for operations under \lts\state\account-fungible-resource-balance
type StateAccountFungibleResourceBalanceRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// StateAccountFungibleResourceBalanceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type StateAccountFungibleResourceBalanceRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewStateAccountFungibleResourceBalanceRequestBuilderInternal instantiates a new StateAccountFungibleResourceBalanceRequestBuilder and sets the default values.
func NewStateAccountFungibleResourceBalanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *StateAccountFungibleResourceBalanceRequestBuilder {
	m := &StateAccountFungibleResourceBalanceRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/lts/state/account-fungible-resource-balance", pathParameters),
	}
	return m
}

// NewStateAccountFungibleResourceBalanceRequestBuilder instantiates a new StateAccountFungibleResourceBalanceRequestBuilder and sets the default values.
func NewStateAccountFungibleResourceBalanceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *StateAccountFungibleResourceBalanceRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewStateAccountFungibleResourceBalanceRequestBuilderInternal(urlParams, requestAdapter)
}

// Post returns balance of a single fungible resource in an account
// returns a LtsStateAccountFungibleResourceBalanceResponseable when successful
// returns a BasicErrorResponse error when the service returns a 500 status code
func (m *StateAccountFungibleResourceBalanceRequestBuilder) Post(ctx context.Context, body ieb2700c3462f1a9e73e3b7d2d4d35314fa43b94ba0380cc504a8f60edfdf3455.LtsStateAccountFungibleResourceBalanceRequestable, requestConfiguration *StateAccountFungibleResourceBalanceRequestBuilderPostRequestConfiguration) (ieb2700c3462f1a9e73e3b7d2d4d35314fa43b94ba0380cc504a8f60edfdf3455.LtsStateAccountFungibleResourceBalanceResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"500": ieb2700c3462f1a9e73e3b7d2d4d35314fa43b94ba0380cc504a8f60edfdf3455.CreateBasicErrorResponseFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ieb2700c3462f1a9e73e3b7d2d4d35314fa43b94ba0380cc504a8f60edfdf3455.CreateLtsStateAccountFungibleResourceBalanceResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ieb2700c3462f1a9e73e3b7d2d4d35314fa43b94ba0380cc504a8f60edfdf3455.LtsStateAccountFungibleResourceBalanceResponseable), nil
}

// ToPostRequestInformation returns balance of a single fungible resource in an account
// returns a *RequestInformation when successful
func (m *StateAccountFungibleResourceBalanceRequestBuilder) ToPostRequestInformation(ctx context.Context, body ieb2700c3462f1a9e73e3b7d2d4d35314fa43b94ba0380cc504a8f60edfdf3455.LtsStateAccountFungibleResourceBalanceRequestable, requestConfiguration *StateAccountFungibleResourceBalanceRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *StateAccountFungibleResourceBalanceRequestBuilder when successful
func (m *StateAccountFungibleResourceBalanceRequestBuilder) WithUrl(rawUrl string) *StateAccountFungibleResourceBalanceRequestBuilder {
	return NewStateAccountFungibleResourceBalanceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
