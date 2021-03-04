// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protocol

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FluxClient is the client API for Flux service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FluxClient interface {
	// Get user data
	//
	// Provides user data associated with the access token
	UserData(ctx context.Context, in *UserDataRequest, opts ...grpc.CallOption) (*UserDataResponse, error)
	// Get pricing data
	//
	// Provides pricing data for all markets with rate maps
	PricingData(ctx context.Context, in *PricingDataRequest, opts ...grpc.CallOption) (*PricingDataResponse, error)
	// Post email or phone in exchange for a one time passcode
	//
	// Will cause your email or phone to receive a one time passcode.
	// This can be used in the verify step to obtain a token for login
	OneTimePasscode(ctx context.Context, in *OneTimePasscodeRequest, opts ...grpc.CallOption) (*OneTimePasscodeResponse, error)
	// Post one time passcode in exchange for an access token
	//
	// The passcode received in either email or phone text message should be provided here in order to obtain on access token
	OneTimePasscodeVerify(ctx context.Context, in *OneTimePasscodeVerifyRequest, opts ...grpc.CallOption) (*OneTimePasscodeVerifyResponse, error)
	// Post chosen bank info from plaid in order to create a new ACH pyment method in wyre
	//
	// requires a plaid processor token which in turn requires a plaid widget interaction where the user selects the account id
	PlaidConnectBankAccounts(ctx context.Context, in *PlaidConnectBankAccountsRequest, opts ...grpc.CallOption) (*PlaidConnectBankAccountsResponse, error)
	// PlaidCreateLinkToken implements this flow: https://plaid.com/docs/link/link-token-migration-guide/
	PlaidCreateLinkToken(ctx context.Context, in *PlaidCreateLinkTokenRequest, opts ...grpc.CallOption) (*PlaidCreateLinkTokenResponse, error)
}

type fluxClient struct {
	cc grpc.ClientConnInterface
}

func NewFluxClient(cc grpc.ClientConnInterface) FluxClient {
	return &fluxClient{cc}
}

func (c *fluxClient) UserData(ctx context.Context, in *UserDataRequest, opts ...grpc.CallOption) (*UserDataResponse, error) {
	out := new(UserDataResponse)
	err := c.cc.Invoke(ctx, "/Flux/UserData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fluxClient) PricingData(ctx context.Context, in *PricingDataRequest, opts ...grpc.CallOption) (*PricingDataResponse, error) {
	out := new(PricingDataResponse)
	err := c.cc.Invoke(ctx, "/Flux/PricingData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fluxClient) OneTimePasscode(ctx context.Context, in *OneTimePasscodeRequest, opts ...grpc.CallOption) (*OneTimePasscodeResponse, error) {
	out := new(OneTimePasscodeResponse)
	err := c.cc.Invoke(ctx, "/Flux/OneTimePasscode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fluxClient) OneTimePasscodeVerify(ctx context.Context, in *OneTimePasscodeVerifyRequest, opts ...grpc.CallOption) (*OneTimePasscodeVerifyResponse, error) {
	out := new(OneTimePasscodeVerifyResponse)
	err := c.cc.Invoke(ctx, "/Flux/OneTimePasscodeVerify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fluxClient) PlaidConnectBankAccounts(ctx context.Context, in *PlaidConnectBankAccountsRequest, opts ...grpc.CallOption) (*PlaidConnectBankAccountsResponse, error) {
	out := new(PlaidConnectBankAccountsResponse)
	err := c.cc.Invoke(ctx, "/Flux/PlaidConnectBankAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fluxClient) PlaidCreateLinkToken(ctx context.Context, in *PlaidCreateLinkTokenRequest, opts ...grpc.CallOption) (*PlaidCreateLinkTokenResponse, error) {
	out := new(PlaidCreateLinkTokenResponse)
	err := c.cc.Invoke(ctx, "/Flux/PlaidCreateLinkToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FluxServer is the server API for Flux service.
// All implementations must embed UnimplementedFluxServer
// for forward compatibility
type FluxServer interface {
	// Get user data
	//
	// Provides user data associated with the access token
	UserData(context.Context, *UserDataRequest) (*UserDataResponse, error)
	// Get pricing data
	//
	// Provides pricing data for all markets with rate maps
	PricingData(context.Context, *PricingDataRequest) (*PricingDataResponse, error)
	// Post email or phone in exchange for a one time passcode
	//
	// Will cause your email or phone to receive a one time passcode.
	// This can be used in the verify step to obtain a token for login
	OneTimePasscode(context.Context, *OneTimePasscodeRequest) (*OneTimePasscodeResponse, error)
	// Post one time passcode in exchange for an access token
	//
	// The passcode received in either email or phone text message should be provided here in order to obtain on access token
	OneTimePasscodeVerify(context.Context, *OneTimePasscodeVerifyRequest) (*OneTimePasscodeVerifyResponse, error)
	// Post chosen bank info from plaid in order to create a new ACH pyment method in wyre
	//
	// requires a plaid processor token which in turn requires a plaid widget interaction where the user selects the account id
	PlaidConnectBankAccounts(context.Context, *PlaidConnectBankAccountsRequest) (*PlaidConnectBankAccountsResponse, error)
	// PlaidCreateLinkToken implements this flow: https://plaid.com/docs/link/link-token-migration-guide/
	PlaidCreateLinkToken(context.Context, *PlaidCreateLinkTokenRequest) (*PlaidCreateLinkTokenResponse, error)
	mustEmbedUnimplementedFluxServer()
}

// UnimplementedFluxServer must be embedded to have forward compatible implementations.
type UnimplementedFluxServer struct {
}

func (UnimplementedFluxServer) UserData(context.Context, *UserDataRequest) (*UserDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserData not implemented")
}
func (UnimplementedFluxServer) PricingData(context.Context, *PricingDataRequest) (*PricingDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PricingData not implemented")
}
func (UnimplementedFluxServer) OneTimePasscode(context.Context, *OneTimePasscodeRequest) (*OneTimePasscodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OneTimePasscode not implemented")
}
func (UnimplementedFluxServer) OneTimePasscodeVerify(context.Context, *OneTimePasscodeVerifyRequest) (*OneTimePasscodeVerifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OneTimePasscodeVerify not implemented")
}
func (UnimplementedFluxServer) PlaidConnectBankAccounts(context.Context, *PlaidConnectBankAccountsRequest) (*PlaidConnectBankAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaidConnectBankAccounts not implemented")
}
func (UnimplementedFluxServer) PlaidCreateLinkToken(context.Context, *PlaidCreateLinkTokenRequest) (*PlaidCreateLinkTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaidCreateLinkToken not implemented")
}
func (UnimplementedFluxServer) mustEmbedUnimplementedFluxServer() {}

// UnsafeFluxServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FluxServer will
// result in compilation errors.
type UnsafeFluxServer interface {
	mustEmbedUnimplementedFluxServer()
}

func RegisterFluxServer(s grpc.ServiceRegistrar, srv FluxServer) {
	s.RegisterService(&Flux_ServiceDesc, srv)
}

func _Flux_UserData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxServer).UserData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Flux/UserData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxServer).UserData(ctx, req.(*UserDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Flux_PricingData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PricingDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxServer).PricingData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Flux/PricingData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxServer).PricingData(ctx, req.(*PricingDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Flux_OneTimePasscode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OneTimePasscodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxServer).OneTimePasscode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Flux/OneTimePasscode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxServer).OneTimePasscode(ctx, req.(*OneTimePasscodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Flux_OneTimePasscodeVerify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OneTimePasscodeVerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxServer).OneTimePasscodeVerify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Flux/OneTimePasscodeVerify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxServer).OneTimePasscodeVerify(ctx, req.(*OneTimePasscodeVerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Flux_PlaidConnectBankAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaidConnectBankAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxServer).PlaidConnectBankAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Flux/PlaidConnectBankAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxServer).PlaidConnectBankAccounts(ctx, req.(*PlaidConnectBankAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Flux_PlaidCreateLinkToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaidCreateLinkTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxServer).PlaidCreateLinkToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Flux/PlaidCreateLinkToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxServer).PlaidCreateLinkToken(ctx, req.(*PlaidCreateLinkTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Flux_ServiceDesc is the grpc.ServiceDesc for Flux service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Flux_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Flux",
	HandlerType: (*FluxServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserData",
			Handler:    _Flux_UserData_Handler,
		},
		{
			MethodName: "PricingData",
			Handler:    _Flux_PricingData_Handler,
		},
		{
			MethodName: "OneTimePasscode",
			Handler:    _Flux_OneTimePasscode_Handler,
		},
		{
			MethodName: "OneTimePasscodeVerify",
			Handler:    _Flux_OneTimePasscodeVerify_Handler,
		},
		{
			MethodName: "PlaidConnectBankAccounts",
			Handler:    _Flux_PlaidConnectBankAccounts_Handler,
		},
		{
			MethodName: "PlaidCreateLinkToken",
			Handler:    _Flux_PlaidCreateLinkToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
