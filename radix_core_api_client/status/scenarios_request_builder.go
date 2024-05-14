package status

import (
	"context"

	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	ieb2700c3462f1a9e73e3b7d2d4d35314fa43b94ba0380cc504a8f60edfdf3455 "github.com/radixdlt/maya/radix_core_api_client/models"
)

// ScenariosRequestBuilder builds and executes requests for operations under \status\scenarios
type ScenariosRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ScenariosRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ScenariosRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewScenariosRequestBuilderInternal instantiates a new ScenariosRequestBuilder and sets the default values.
func NewScenariosRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ScenariosRequestBuilder {
	m := &ScenariosRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/status/scenarios", pathParameters),
	}
	return m
}

// NewScenariosRequestBuilder instantiates a new ScenariosRequestBuilder and sets the default values.
func NewScenariosRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ScenariosRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewScenariosRequestBuilderInternal(urlParams, requestAdapter)
}

// Post get results of test "Scenarios" executed on this Network.Note: these Scenarios are meant to only be executed on test Networks; on a production Node,the response is expected to be empty.
// returns a ScenariosResponseable when successful
// returns a BasicErrorResponse error when the service returns a 500 status code
func (m *ScenariosRequestBuilder) Post(ctx context.Context, body ieb2700c3462f1a9e73e3b7d2d4d35314fa43b94ba0380cc504a8f60edfdf3455.ScenariosRequestable, requestConfiguration *ScenariosRequestBuilderPostRequestConfiguration) (ieb2700c3462f1a9e73e3b7d2d4d35314fa43b94ba0380cc504a8f60edfdf3455.ScenariosResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"500": ieb2700c3462f1a9e73e3b7d2d4d35314fa43b94ba0380cc504a8f60edfdf3455.CreateBasicErrorResponseFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ieb2700c3462f1a9e73e3b7d2d4d35314fa43b94ba0380cc504a8f60edfdf3455.CreateScenariosResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ieb2700c3462f1a9e73e3b7d2d4d35314fa43b94ba0380cc504a8f60edfdf3455.ScenariosResponseable), nil
}

// ToPostRequestInformation get results of test "Scenarios" executed on this Network.Note: these Scenarios are meant to only be executed on test Networks; on a production Node,the response is expected to be empty.
// returns a *RequestInformation when successful
func (m *ScenariosRequestBuilder) ToPostRequestInformation(ctx context.Context, body ieb2700c3462f1a9e73e3b7d2d4d35314fa43b94ba0380cc504a8f60edfdf3455.ScenariosRequestable, requestConfiguration *ScenariosRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ScenariosRequestBuilder when successful
func (m *ScenariosRequestBuilder) WithUrl(rawUrl string) *ScenariosRequestBuilder {
	return NewScenariosRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
