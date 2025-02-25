/*
Tendermint RPC

Tendermint supports the following RPC protocols:  * URI over HTTP * JSONRPC over HTTP * JSONRPC over websockets  ## Configuration  RPC can be configured by tuning parameters under `[rpc]` table in the `$TMHOME/config/config.toml` file or by using the `--rpc.X` command-line flags.  Default rpc listen address is `tcp://0.0.0.0:26657`. To set another address, set the `laddr` config parameter to desired value. CORS (Cross-Origin Resource Sharing) can be enabled by setting `cors_allowed_origins`, `cors_allowed_methods`, `cors_allowed_headers` config parameters.  ## Arguments  Arguments which expect strings or byte arrays may be passed as quoted strings, like `\"abc\"` or as `0x`-prefixed strings, like `0x616263`.  ## URI/HTTP  A REST like interface.      curl localhost:26657/block?height=5  ## JSONRPC/HTTP  JSONRPC requests can be POST'd to the root RPC endpoint via HTTP.      curl --header \"Content-Type: application/json\" --request POST --data '{\"method\": \"block\", \"params\": [\"5\"], \"id\": 1}' localhost:26657  ## JSONRPC/websockets  JSONRPC requests can be also made via websocket. The websocket endpoint is at `/websocket`, e.g. `localhost:26657/websocket`. Asynchronous RPC functions like event `subscribe` and `unsubscribe` are only available via websockets.  Example using https://github.com/hashrocket/ws:      ws ws://localhost:26657/websocket     > { \"jsonrpc\": \"2.0\", \"method\": \"subscribe\", \"params\": [\"tm.event='NewBlock'\"], \"id\": 1 } 

API version: Master
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ABCIInfoResponseResult struct for ABCIInfoResponseResult
type ABCIInfoResponseResult struct {
	Response ABCIInfoResponseResultResponse `json:"response"`
}

// NewABCIInfoResponseResult instantiates a new ABCIInfoResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewABCIInfoResponseResult(response ABCIInfoResponseResultResponse) *ABCIInfoResponseResult {
	this := ABCIInfoResponseResult{}
	this.Response = response
	return &this
}

// NewABCIInfoResponseResultWithDefaults instantiates a new ABCIInfoResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewABCIInfoResponseResultWithDefaults() *ABCIInfoResponseResult {
	this := ABCIInfoResponseResult{}
	return &this
}

// GetResponse returns the Response field value
func (o *ABCIInfoResponseResult) GetResponse() ABCIInfoResponseResultResponse {
	if o == nil {
		var ret ABCIInfoResponseResultResponse
		return ret
	}

	return o.Response
}

// GetResponseOk returns a tuple with the Response field value
// and a boolean to check if the value has been set.
func (o *ABCIInfoResponseResult) GetResponseOk() (*ABCIInfoResponseResultResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Response, true
}

// SetResponse sets field value
func (o *ABCIInfoResponseResult) SetResponse(v ABCIInfoResponseResultResponse) {
	o.Response = v
}

func (o ABCIInfoResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableABCIInfoResponseResult struct {
	value *ABCIInfoResponseResult
	isSet bool
}

func (v NullableABCIInfoResponseResult) Get() *ABCIInfoResponseResult {
	return v.value
}

func (v *NullableABCIInfoResponseResult) Set(val *ABCIInfoResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableABCIInfoResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableABCIInfoResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableABCIInfoResponseResult(val *ABCIInfoResponseResult) *NullableABCIInfoResponseResult {
	return &NullableABCIInfoResponseResult{value: val, isSet: true}
}

func (v NullableABCIInfoResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableABCIInfoResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


