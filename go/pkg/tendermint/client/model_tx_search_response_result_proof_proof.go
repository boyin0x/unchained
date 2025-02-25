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

// TxSearchResponseResultProofProof struct for TxSearchResponseResultProofProof
type TxSearchResponseResultProofProof struct {
	Total string `json:"total"`
	Index string `json:"index"`
	LeafHash string `json:"leaf_hash"`
	Aunts []string `json:"aunts"`
}

// NewTxSearchResponseResultProofProof instantiates a new TxSearchResponseResultProofProof object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTxSearchResponseResultProofProof(total string, index string, leafHash string, aunts []string) *TxSearchResponseResultProofProof {
	this := TxSearchResponseResultProofProof{}
	this.Total = total
	this.Index = index
	this.LeafHash = leafHash
	this.Aunts = aunts
	return &this
}

// NewTxSearchResponseResultProofProofWithDefaults instantiates a new TxSearchResponseResultProofProof object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTxSearchResponseResultProofProofWithDefaults() *TxSearchResponseResultProofProof {
	this := TxSearchResponseResultProofProof{}
	return &this
}

// GetTotal returns the Total field value
func (o *TxSearchResponseResultProofProof) GetTotal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *TxSearchResponseResultProofProof) GetTotalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *TxSearchResponseResultProofProof) SetTotal(v string) {
	o.Total = v
}

// GetIndex returns the Index field value
func (o *TxSearchResponseResultProofProof) GetIndex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *TxSearchResponseResultProofProof) GetIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *TxSearchResponseResultProofProof) SetIndex(v string) {
	o.Index = v
}

// GetLeafHash returns the LeafHash field value
func (o *TxSearchResponseResultProofProof) GetLeafHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LeafHash
}

// GetLeafHashOk returns a tuple with the LeafHash field value
// and a boolean to check if the value has been set.
func (o *TxSearchResponseResultProofProof) GetLeafHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LeafHash, true
}

// SetLeafHash sets field value
func (o *TxSearchResponseResultProofProof) SetLeafHash(v string) {
	o.LeafHash = v
}

// GetAunts returns the Aunts field value
func (o *TxSearchResponseResultProofProof) GetAunts() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Aunts
}

// GetAuntsOk returns a tuple with the Aunts field value
// and a boolean to check if the value has been set.
func (o *TxSearchResponseResultProofProof) GetAuntsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Aunts, true
}

// SetAunts sets field value
func (o *TxSearchResponseResultProofProof) SetAunts(v []string) {
	o.Aunts = v
}

func (o TxSearchResponseResultProofProof) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["index"] = o.Index
	}
	if true {
		toSerialize["leaf_hash"] = o.LeafHash
	}
	if true {
		toSerialize["aunts"] = o.Aunts
	}
	return json.Marshal(toSerialize)
}

type NullableTxSearchResponseResultProofProof struct {
	value *TxSearchResponseResultProofProof
	isSet bool
}

func (v NullableTxSearchResponseResultProofProof) Get() *TxSearchResponseResultProofProof {
	return v.value
}

func (v *NullableTxSearchResponseResultProofProof) Set(val *TxSearchResponseResultProofProof) {
	v.value = val
	v.isSet = true
}

func (v NullableTxSearchResponseResultProofProof) IsSet() bool {
	return v.isSet
}

func (v *NullableTxSearchResponseResultProofProof) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTxSearchResponseResultProofProof(val *TxSearchResponseResultProofProof) *NullableTxSearchResponseResultProofProof {
	return &NullableTxSearchResponseResultProofProof{value: val, isSet: true}
}

func (v NullableTxSearchResponseResultProofProof) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTxSearchResponseResultProofProof) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


