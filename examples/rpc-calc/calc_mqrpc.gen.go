// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: calc.proto

/*
	Package main is a generated protocol buffer package.

	It is generated from these files:
		calc.proto

	It has these top-level messages:
		Request
		Answer
*/
package main

import fmt "fmt"
import errors "errors"
import rpc "github.com/003random/amqp/rpc"
import proto "github.com/gogo/protobuf/proto"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Server API
type CalcServer interface {
	Eval(arg *Request) (*Answer, error)
}

// Run server API with this call
func RunServer(srv rpc.Server, handler CalcServer) {
	srv.Serve(func(funcID int32, arg []byte) ([]byte, error) {
		switch funcID {
		case Functions_Eval:
			return _Handle_Eval(handler, arg)
		default:
			return nil, errors.New(fmt.Sprintf("unknown function with code: %d", funcID))
		}
	})
}

// Client API
type CalcClient interface {
	Close()
	Eval(arg *Request) (*Answer, error)
}

func NewCalcClient(cc rpc.Client) CalcClient {
	return &calcClient{cc}
}

type calcClient struct {
	cc rpc.Client
}

// Functions enum
const (
	Functions_Eval int32 = 0
)

// Server API handlers
func _Handle_Eval(handler interface{}, arg []byte) ([]byte, error) {
	var req Request
	err := req.Unmarshal(arg)
	if err != nil {
		return nil, err
	}
	resp, err := handler.(CalcServer).Eval(&req)
	if err != nil {
		return nil, err
	}
	return resp.Marshal()
}

// Client API handlers
func (this *calcClient) Close() {
	this.cc.Close()
}
func (this *calcClient) Eval(arg *Request) (*Answer, error) {
	request, err := arg.Marshal()
	if err != nil {
		return nil, err
	}
	respData, err := this.cc.RemoteCall(rpc.Request{FuncID: Functions_Eval, Body: request})
	if err != nil {
		return nil, err
	}
	var resp Answer
	err = resp.Unmarshal(respData)
	return &resp, err
}
