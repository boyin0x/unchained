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

// BroadcastTxCommitResponseResultDeliverTx struct for BroadcastTxCommitResponseResultDeliverTx
type BroadcastTxCommitResponseResultDeliverTx struct {
	Log string `json:"log"`
	Data string `json:"data"`
	Code string `json:"code"`
}

// NewBroadcastTxCommitResponseResultDeliverTx instantiates a new BroadcastTxCommitResponseResultDeliverTx object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBroadcastTxCommitResponseResultDeliverTx(log string, data string, code string) *BroadcastTxCommitResponseResultDeliverTx {
	this := BroadcastTxCommitResponseResultDeliverTx{}
	this.Log = log
	this.Data = data
	this.Code = code
	return &this
}

// NewBroadcastTxCommitResponseResultDeliverTxWithDefaults instantiates a new BroadcastTxCommitResponseResultDeliverTx object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBroadcastTxCommitResponseResultDeliverTxWithDefaults() *BroadcastTxCommitResponseResultDeliverTx {
	this := BroadcastTxCommitResponseResultDeliverTx{}
	return &this
}

// GetLog returns the Log field value
func (o *BroadcastTxCommitResponseResultDeliverTx) GetLog() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Log
}

// GetLogOk returns a tuple with the Log field value
// and a boolean to check if the value has been set.
func (o *BroadcastTxCommitResponseResultDeliverTx) GetLogOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Log, true
}

// SetLog sets field value
func (o *BroadcastTxCommitResponseResultDeliverTx) SetLog(v string) {
	o.Log = v
}

// GetData returns the Data field value
func (o *BroadcastTxCommitResponseResultDeliverTx) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BroadcastTxCommitResponseResultDeliverTx) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BroadcastTxCommitResponseResultDeliverTx) SetData(v string) {
	o.Data = v
}

// GetCode returns the Code field value
func (o *BroadcastTxCommitResponseResultDeliverTx) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *BroadcastTxCommitResponseResultDeliverTx) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *BroadcastTxCommitResponseResultDeliverTx) SetCode(v string) {
	o.Code = v
}

func (o BroadcastTxCommitResponseResultDeliverTx) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["log"] = o.Log
	}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["code"] = o.Code
	}
	return json.Marshal(toSerialize)
}

type NullableBroadcastTxCommitResponseResultDeliverTx struct {
	value *BroadcastTxCommitResponseResultDeliverTx
	isSet bool
}

func (v NullableBroadcastTxCommitResponseResultDeliverTx) Get() *BroadcastTxCommitResponseResultDeliverTx {
	return v.value
}

func (v *NullableBroadcastTxCommitResponseResultDeliverTx) Set(val *BroadcastTxCommitResponseResultDeliverTx) {
	v.value = val
	v.isSet = true
}

func (v NullableBroadcastTxCommitResponseResultDeliverTx) IsSet() bool {
	return v.isSet
}

func (v *NullableBroadcastTxCommitResponseResultDeliverTx) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBroadcastTxCommitResponseResultDeliverTx(val *BroadcastTxCommitResponseResultDeliverTx) *NullableBroadcastTxCommitResponseResultDeliverTx {
	return &NullableBroadcastTxCommitResponseResultDeliverTx{value: val, isSet: true}
}

func (v NullableBroadcastTxCommitResponseResultDeliverTx) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBroadcastTxCommitResponseResultDeliverTx) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


