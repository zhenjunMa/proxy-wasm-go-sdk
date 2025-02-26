// Copyright 2020-2021 Tetrate
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package proxywasm

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/rawhostcall"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

func GetPluginConfiguration(size int) ([]byte, error) {
	ret, st := getBuffer(types.BufferTypePluginConfiguration, 0, size)
	return ret, types.StatusToError(st)
}

func GetVMConfiguration(size int) ([]byte, error) {
	ret, st := getBuffer(types.BufferTypeVMConfiguration, 0, size)
	return ret, types.StatusToError(st)
}

func SendHttpResponse(statusCode uint32, headers types.Headers, body []byte) error {
	shs := SerializeMap(headers)
	hp := &shs[0]
	hl := len(shs)
	return types.StatusToError(
		rawhostcall.ProxySendLocalResponse(
			statusCode, nil, 0,
			&body[0], len(body), hp, hl, -1,
		),
	)
}

func SetTickPeriodMilliSeconds(millSec uint32) error {
	return types.StatusToError(rawhostcall.ProxySetTickPeriodMilliseconds(millSec))
}

func DispatchHttpCall(upstream string,
	headers types.Headers, body string, trailers types.Trailers,
	timeoutMillisecond uint32, callBack HttpCalloutCallBack) (calloutID uint32, err error) {
	shs := SerializeMap(headers)
	hp := &shs[0]
	hl := len(shs)

	sts := SerializeMap(trailers)
	tp := &sts[0]
	tl := len(sts)

	u := stringBytePtr(upstream)
	switch st := rawhostcall.ProxyHttpCall(u, len(upstream),
		hp, hl, stringBytePtr(body), len(body), tp, tl, timeoutMillisecond, &calloutID); st {
	case types.StatusOK:
		currentState.registerHttpCallOut(calloutID, callBack)
		return calloutID, nil
	default:
		return 0, types.StatusToError(st)
	}
}

func GetHttpCallResponseHeaders() (types.Headers, error) {
	ret, st := getMap(types.MapTypeHttpCallResponseHeaders)
	return ret, types.StatusToError(st)
}

func GetHttpCallResponseBody(start, maxSize int) ([]byte, error) {
	ret, st := getBuffer(types.BufferTypeHttpCallResponseBody, start, maxSize)
	return ret, types.StatusToError(st)
}

func GetHttpCallResponseTrailers() (types.Trailers, error) {
	ret, st := getMap(types.MapTypeHttpCallResponseTrailers)
	return ret, types.StatusToError(st)
}

func GetDownStreamData(start, maxSize int) ([]byte, error) {
	ret, st := getBuffer(types.BufferTypeDownstreamData, start, maxSize)
	return ret, types.StatusToError(st)
}

func SetDownStreamData(body []byte) error {
	var bufferData *byte
	if len(body) != 0 {
		bufferData = &body[0]
	}
	st := rawhostcall.ProxySetBufferBytes(types.BufferTypeDownstreamData, 0, len(body), bufferData, len(body))
	return types.StatusToError(st)
}

func GetUpstreamData(start, maxSize int) ([]byte, error) {
	ret, st := getBuffer(types.BufferTypeUpstreamData, start, maxSize)
	return ret, types.StatusToError(st)
}

func SetUpstreamData(body []byte) error {
	var bufferData *byte
	if len(body) != 0 {
		bufferData = &body[0]
	}
	st := rawhostcall.ProxySetBufferBytes(types.BufferTypeUpstreamData, 0, len(body), bufferData, len(body))
	return types.StatusToError(st)
}

func ContinueDownStream() error {
	return types.StatusToError(rawhostcall.ProxyContinueStream(types.StreamTypeDownstream))
}

func ContinueUpstream() error {
	return types.StatusToError(rawhostcall.ProxyContinueStream(types.StreamTypeUpstream))
}

func CloseDownStream() error {
	return types.StatusToError(rawhostcall.ProxyCloseStream(types.StreamTypeDownstream))
}

func CloseUpstream() error {
	return types.StatusToError(rawhostcall.ProxyCloseStream(types.StreamTypeUpstream))
}

func GetHttpRequestHeaders() (types.Headers, error) {
	ret, st := getMap(types.MapTypeHttpRequestHeaders)
	return ret, types.StatusToError(st)
}

func SetHttpRequestHeaders(headers types.Headers) error {
	return types.StatusToError(setMap(types.MapTypeHttpRequestHeaders, headers))
}

func GetHttpRequestHeader(key string) (string, error) {
	ret, st := getMapValue(types.MapTypeHttpRequestHeaders, key)
	return ret, types.StatusToError(st)
}

func RemoveHttpRequestHeader(key string) error {
	return types.StatusToError(removeMapValue(types.MapTypeHttpRequestHeaders, key))
}

func SetHttpRequestHeader(key, value string) error {
	return types.StatusToError(setMapValue(types.MapTypeHttpRequestHeaders, key, value))
}

func AddHttpRequestHeader(key, value string) error {
	return types.StatusToError(addMapValue(types.MapTypeHttpRequestHeaders, key, value))
}

func GetHttpRequestBody(start, maxSize int) ([]byte, error) {
	ret, st := getBuffer(types.BufferTypeHttpRequestBody, start, maxSize)
	return ret, types.StatusToError(st)
}

func SetHttpRequestBody(body []byte) error {
	var bufferData *byte
	if len(body) != 0 {
		bufferData = &body[0]
	}
	st := rawhostcall.ProxySetBufferBytes(types.BufferTypeHttpRequestBody, 0, len(body), bufferData, len(body))
	return types.StatusToError(st)
}

func GetHttpRequestTrailers() (types.Trailers, error) {
	ret, st := getMap(types.MapTypeHttpRequestTrailers)
	return ret, types.StatusToError(st)
}

func SetHttpRequestTrailers(trailers types.Trailers) error {
	return types.StatusToError(setMap(types.MapTypeHttpRequestTrailers, trailers))
}

func GetHttpRequestTrailer(key string) (string, error) {
	ret, st := getMapValue(types.MapTypeHttpRequestTrailers, key)
	return ret, types.StatusToError(st)
}

func RemoveHttpRequestTrailer(key string) error {
	return types.StatusToError(removeMapValue(types.MapTypeHttpRequestTrailers, key))
}

func SetHttpRequestTrailer(key, value string) error {
	return types.StatusToError(setMapValue(types.MapTypeHttpRequestTrailers, key, value))
}

func AddHttpRequestTrailer(key, value string) error {
	return types.StatusToError(addMapValue(types.MapTypeHttpRequestTrailers, key, value))
}

func ResumeHttpRequest() error {
	return types.StatusToError(rawhostcall.ProxyContinueStream(types.StreamTypeRequest))
}

func GetHttpResponseHeaders() (types.Headers, error) {
	ret, st := getMap(types.MapTypeHttpResponseHeaders)
	return ret, types.StatusToError(st)
}

func SetHttpResponseHeaders(headers types.Headers) error {
	return types.StatusToError(setMap(types.MapTypeHttpResponseHeaders, headers))
}

func GetHttpResponseHeader(key string) (string, error) {
	ret, st := getMapValue(types.MapTypeHttpResponseHeaders, key)
	return ret, types.StatusToError(st)
}

func RemoveHttpResponseHeader(key string) error {
	return types.StatusToError(removeMapValue(types.MapTypeHttpResponseHeaders, key))
}

func SetHttpResponseHeader(key, value string) error {
	return types.StatusToError(setMapValue(types.MapTypeHttpResponseHeaders, key, value))
}

func AddHttpResponseHeader(key, value string) error {
	return types.StatusToError(addMapValue(types.MapTypeHttpResponseHeaders, key, value))
}

func GetHttpResponseBody(start, maxSize int) ([]byte, error) {
	ret, st := getBuffer(types.BufferTypeHttpResponseBody, start, maxSize)
	return ret, types.StatusToError(st)
}

func SetHttpResponseBody(body []byte) error {
	var bufferData *byte
	if len(body) != 0 {
		bufferData = &body[0]
	}
	st := rawhostcall.ProxySetBufferBytes(types.BufferTypeHttpResponseBody, 0, len(body), bufferData, len(body))
	return types.StatusToError(st)
}

func GetHttpResponseTrailers() (types.Trailers, error) {
	ret, st := getMap(types.MapTypeHttpResponseTrailers)
	return ret, types.StatusToError(st)
}

func SetHttpResponseTrailers(trailers types.Trailers) error {
	return types.StatusToError(setMap(types.MapTypeHttpResponseTrailers, trailers))
}

func GetHttpResponseTrailer(key string) (string, error) {
	ret, st := getMapValue(types.MapTypeHttpResponseTrailers, key)
	return ret, types.StatusToError(st)
}

func RemoveHttpResponseTrailer(key string) error {
	return types.StatusToError(removeMapValue(types.MapTypeHttpResponseTrailers, key))
}

func SetHttpResponseTrailer(key, value string) error {
	return types.StatusToError(setMapValue(types.MapTypeHttpResponseTrailers, key, value))
}

func AddHttpResponseTrailer(key, value string) error {
	return types.StatusToError(addMapValue(types.MapTypeHttpResponseTrailers, key, value))
}

func ResumeHttpResponse() error {
	return types.StatusToError(rawhostcall.ProxyContinueStream(types.StreamTypeResponse))
}

func RegisterSharedQueue(name string) (uint32, error) {
	var queueID uint32
	ptr := stringBytePtr(name)
	st := rawhostcall.ProxyRegisterSharedQueue(ptr, len(name), &queueID)
	return queueID, types.StatusToError(st)
}

// TODO: not sure if the ABI is correct
func ResolveSharedQueue(vmID, queueName string) (uint32, error) {
	var ret uint32
	st := rawhostcall.ProxyResolveSharedQueue(stringBytePtr(vmID),
		len(vmID), stringBytePtr(queueName), len(queueName), &ret)
	return ret, types.StatusToError(st)
}

func DequeueSharedQueue(queueID uint32) ([]byte, error) {
	var raw *byte
	var size int
	st := rawhostcall.ProxyDequeueSharedQueue(queueID, &raw, &size)
	if st != types.StatusOK {
		return nil, types.StatusToError(st)
	}
	return RawBytePtrToByteSlice(raw, size), nil
}

func EnqueueSharedQueue(queueID uint32, data []byte) error {
	return types.StatusToError(rawhostcall.ProxyEnqueueSharedQueue(queueID, &data[0], len(data)))
}

func GetSharedData(key string) (value []byte, cas uint32, err error) {
	var raw *byte
	var size int

	st := rawhostcall.ProxyGetSharedData(stringBytePtr(key), len(key), &raw, &size, &cas)
	if st != types.StatusOK {
		return nil, 0, types.StatusToError(st)
	}
	return RawBytePtrToByteSlice(raw, size), cas, nil
}

func SetSharedData(key string, data []byte, cas uint32) error {
	st := rawhostcall.ProxySetSharedData(stringBytePtr(key),
		len(key), &data[0], len(data), cas)
	return types.StatusToError(st)
}

func GetProperty(path []string) ([]byte, error) {
	var ret *byte
	var retSize int
	raw := SerializePropertyPath(path)

	err := types.StatusToError(rawhostcall.ProxyGetProperty(&raw[0], len(raw), &ret, &retSize))
	if err != nil {
		return nil, err
	}

	return RawBytePtrToByteSlice(ret, retSize), nil

}

func SetProperty(path string, data []byte) error {
	return types.StatusToError(rawhostcall.ProxySetProperty(
		stringBytePtr(path), len(path), &data[0], len(data),
	))
}

func setMap(mapType types.MapType, headers [][2]string) types.Status {
	shs := SerializeMap(headers)
	hp := &shs[0]
	hl := len(shs)
	return rawhostcall.ProxySetHeaderMapPairs(mapType, hp, hl)
}

func getMapValue(mapType types.MapType, key string) (string, types.Status) {
	var rvs int
	var raw *byte
	if st := rawhostcall.ProxyGetHeaderMapValue(mapType, stringBytePtr(key), len(key), &raw, &rvs); st != types.StatusOK {
		return "", st
	}

	ret := RawBytePtrToString(raw, rvs)
	return ret, types.StatusOK
}

func removeMapValue(mapType types.MapType, key string) types.Status {
	return rawhostcall.ProxyRemoveHeaderMapValue(mapType, stringBytePtr(key), len(key))
}

func setMapValue(mapType types.MapType, key, value string) types.Status {
	return rawhostcall.ProxyReplaceHeaderMapValue(mapType, stringBytePtr(key), len(key), stringBytePtr(value), len(value))
}

func addMapValue(mapType types.MapType, key, value string) types.Status {
	return rawhostcall.ProxyAddHeaderMapValue(mapType, stringBytePtr(key), len(key), stringBytePtr(value), len(value))
}

func getMap(mapType types.MapType) ([][2]string, types.Status) {
	var rvs int
	var raw *byte

	st := rawhostcall.ProxyGetHeaderMapPairs(mapType, &raw, &rvs)
	if st != types.StatusOK {
		return nil, st
	}

	bs := RawBytePtrToByteSlice(raw, rvs)
	return DeserializeMap(bs), types.StatusOK
}

func getBuffer(bufType types.BufferType, start, maxSize int) ([]byte, types.Status) {
	var retData *byte
	var retSize int
	switch st := rawhostcall.ProxyGetBufferBytes(bufType, start, maxSize, &retData, &retSize); st {
	case types.StatusOK:
		// is this correct handling...?
		if retData == nil {
			return nil, types.StatusNotFound
		}
		return RawBytePtrToByteSlice(retData, retSize), st
	default:
		return nil, st
	}
}
