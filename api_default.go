/*
 * Jupiter API v6
 *
 * The core of [jup.ag](https://jup.ag). Easily get a quote and swap through Jupiter API.  ### Rate Limit We update our rate limit from time to time depending on the load of our servers. We recommend running your own instance of the API if you want to have high rate limit, here to learn how to run the [self-hosted API](https://station.jup.ag/docs/apis/self-hosted).  ### API Wrapper - Typescript [@jup-ag/api](https://github.com/jup-ag/jupiter-quote-api-node)  ### Data types - Public keys are base58 encoded strings - raw data such as Vec<u8\\> are base64 encoded strings
 *
 * API version: 6.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package jupiter

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type DefaultApiService service

/*
DefaultApiService GET /indexed-route-map
DEPRECATED, please use /tokens for tradable mints. Returns a hash map, input mint as key and an array of valid output mint as values, token mints are indexed to reduce the file size
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *DefaultApiIndexedRouteMapGetOpts - Optional Parameters:
     * @param "OnlyDirectRoutes" (optional.Bool) -  Default is false. Direct Routes limits Jupiter routing to single hop routes only.
@return IndexedRouteMapResponse
*/

type DefaultApiIndexedRouteMapGetOpts struct {
	OnlyDirectRoutes optional.Bool
}

func (a *DefaultApiService) IndexedRouteMapGet(ctx context.Context, localVarOptionals *DefaultApiIndexedRouteMapGetOpts) (IndexedRouteMapResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue IndexedRouteMapResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/indexed-route-map"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.OnlyDirectRoutes.IsSet() {
		localVarQueryParams.Add("onlyDirectRoutes", parameterToString(localVarOptionals.OnlyDirectRoutes.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v IndexedRouteMapResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
DefaultApiService GET /program-id-to-label
Returns a hash, which key is the program id and value is the label. This is used to help map error from transaction by identifying the fault program id. With that, we can use the &#x60;excludeDexes&#x60; or &#x60;dexes&#x60; parameter.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return map[string]string
*/
func (a *DefaultApiService) ProgramIdToLabelGet(ctx context.Context) (map[string]string, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue map[string]string
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/program-id-to-label"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v map[string]string
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
DefaultApiService GET /quote
Sends a GET request to the Jupiter API to get the best priced quote.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param inputMint Input token mint address
 * @param outputMint Output token mint address
 * @param amount The amount to swap, have to factor in the token decimals.
 * @param optional nil or *DefaultApiQuoteGetOpts - Optional Parameters:
     * @param "SlippageBps" (optional.Int32) -  The slippage in basis points, 1 basis point is 0.01%. If the output token amount exceeds the slippage then the swap transaction will fail.
     * @param "AutoSlippage" (optional.Bool) -  Automatically calculate the slippage based on pairs.
     * @param "AutoSlippageCollisionUsdValue" (optional.Int32) -  Automatic slippage collision value.
     * @param "ComputeAutoSlippage" (optional.Bool) -  Compute auto slippage value without using it.
     * @param "MaxAutoSlippageBps" (optional.Int32) -  Max slippage in basis points for auto slippage calculation. Default is 400.
     * @param "SwapMode" (optional.String) -  (ExactIn or ExactOut) Defaults to ExactIn. ExactOut is for supporting use cases where you need an exact token amount, like payments. In this case the slippage is on the input token.
     * @param "Dexes" (optional.Interface of []string) -  Default is that all DEXes are included. You can pass in the DEXes that you want to include only and separate them by &#x60;,&#x60;. You can check out the full list [here](https://api.jup.ag/swap/v1/program-id-to-label).
     * @param "ExcludeDexes" (optional.Interface of []string) -  Default is that all DEXes are included. You can pass in the DEXes that you want to exclude and separate them by &#x60;,&#x60;. You can check out the full list [here](https://api.jup.ag/swap/v1/program-id-to-label).
     * @param "RestrictIntermediateTokens" (optional.Bool) -  Restrict intermediate tokens to a top token set that has stable liquidity. This will help to ease potential high slippage error rate when swapping with minimal impact on pricing.
     * @param "OnlyDirectRoutes" (optional.Bool) -  Default is false. Direct Routes limits Jupiter routing to single hop routes only.
     * @param "AsLegacyTransaction" (optional.Bool) -  Default is false. Instead of using versioned transaction, this will use the legacy transaction.
     * @param "PlatformFeeBps" (optional.Int32) -  If you want to charge the user a fee, you can specify the fee in BPS. Fee % is taken out of the output token.
     * @param "MaxAccounts" (optional.Int32) -  Rough estimate of the max accounts to be used for the quote, so that you can compose with your own accounts
     * @param "MinimizeSlippage" (optional.Bool) -  Default is false. Miminize slippage attempts to find routes with lower slippage.
     * @param "PreferLiquidDexes" (optional.Bool) -  Default is false. Enabling it would only consider markets with high liquidity to reduce slippage.
@return QuoteResponse
*/

type DefaultApiQuoteGetOpts struct {
	SlippageBps                   optional.Int32
	AutoSlippage                  optional.Bool
	AutoSlippageCollisionUsdValue optional.Int32
	ComputeAutoSlippage           optional.Bool
	MaxAutoSlippageBps            optional.Int32
	SwapMode                      optional.String
	Dexes                         optional.Interface
	ExcludeDexes                  optional.Interface
	RestrictIntermediateTokens    optional.Bool
	OnlyDirectRoutes              optional.Bool
	AsLegacyTransaction           optional.Bool
	PlatformFeeBps                optional.Int32
	MaxAccounts                   optional.Int32
	MinimizeSlippage              optional.Bool
	PreferLiquidDexes             optional.Bool
}

func (a *DefaultApiService) QuoteGet(ctx context.Context, inputMint string, outputMint string, amount uint64, localVarOptionals *DefaultApiQuoteGetOpts) (QuoteResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue QuoteResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/quote"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("inputMint", parameterToString(inputMint, ""))
	localVarQueryParams.Add("outputMint", parameterToString(outputMint, ""))
	localVarQueryParams.Add("amount", parameterToString(amount, ""))
	if localVarOptionals != nil && localVarOptionals.SlippageBps.IsSet() {
		localVarQueryParams.Add("slippageBps", parameterToString(localVarOptionals.SlippageBps.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AutoSlippage.IsSet() {
		localVarQueryParams.Add("autoSlippage", parameterToString(localVarOptionals.AutoSlippage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AutoSlippageCollisionUsdValue.IsSet() {
		localVarQueryParams.Add("autoSlippageCollisionUsdValue", parameterToString(localVarOptionals.AutoSlippageCollisionUsdValue.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ComputeAutoSlippage.IsSet() {
		localVarQueryParams.Add("computeAutoSlippage", parameterToString(localVarOptionals.ComputeAutoSlippage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MaxAutoSlippageBps.IsSet() {
		localVarQueryParams.Add("maxAutoSlippageBps", parameterToString(localVarOptionals.MaxAutoSlippageBps.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SwapMode.IsSet() {
		localVarQueryParams.Add("swapMode", parameterToString(localVarOptionals.SwapMode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Dexes.IsSet() {
		localVarQueryParams.Add("dexes", parameterToString(localVarOptionals.Dexes.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.ExcludeDexes.IsSet() {
		localVarQueryParams.Add("excludeDexes", parameterToString(localVarOptionals.ExcludeDexes.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.RestrictIntermediateTokens.IsSet() {
		localVarQueryParams.Add("restrictIntermediateTokens", parameterToString(localVarOptionals.RestrictIntermediateTokens.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OnlyDirectRoutes.IsSet() {
		localVarQueryParams.Add("onlyDirectRoutes", parameterToString(localVarOptionals.OnlyDirectRoutes.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AsLegacyTransaction.IsSet() {
		localVarQueryParams.Add("asLegacyTransaction", parameterToString(localVarOptionals.AsLegacyTransaction.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PlatformFeeBps.IsSet() {
		localVarQueryParams.Add("platformFeeBps", parameterToString(localVarOptionals.PlatformFeeBps.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MaxAccounts.IsSet() {
		localVarQueryParams.Add("maxAccounts", parameterToString(localVarOptionals.MaxAccounts.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MinimizeSlippage.IsSet() {
		localVarQueryParams.Add("minimizeSlippage", parameterToString(localVarOptionals.MinimizeSlippage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PreferLiquidDexes.IsSet() {
		localVarQueryParams.Add("preferLiquidDexes", parameterToString(localVarOptionals.PreferLiquidDexes.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v QuoteResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
DefaultApiService POST /swap-instructions
Returns instructions that you can use from the quote you get from &#x60;/quote&#x60;.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body

@return SwapInstructionsResponse
*/
func (a *DefaultApiService) SwapInstructionsPost(ctx context.Context, body SwapRequest) (SwapInstructionsResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue SwapInstructionsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/swap-instructions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v SwapInstructionsResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
DefaultApiService POST /swap
Returns a transaction that you can use from the quote you get from &#x60;/quote&#x60;.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body

@return SwapResponse
*/
func (a *DefaultApiService) SwapPost(ctx context.Context, body SwapRequest) (SwapResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue SwapResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/swap"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v SwapResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
DefaultApiService GET /tokens
Returns a list of all the tradable mints
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return []string
*/
func (a *DefaultApiService) TokensGet(ctx context.Context) ([]string, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []string
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/tokens"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []string
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
