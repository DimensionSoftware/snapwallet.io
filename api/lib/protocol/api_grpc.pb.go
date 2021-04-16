// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protocol

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FluxClient is the client API for Flux service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FluxClient interface {
	// Get viewer data
	//
	// Provides user (viewer) data associated with the access token
	ViewerData(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ViewerDataResponse, error)
	// Get viewer profile data
	//
	// Provides user (viewer) data associated with the access token
	ViewerProfileData(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ProfileDataInfo, error)
	// Change users email (viewer based on jwt)
	//
	// requires an otp code and the desired email address change
	ChangeViewerEmail(ctx context.Context, in *ChangeViewerEmailRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Change users phone (viewer based on jwt)
	//
	// requires an otp code and the desired phone change
	ChangeViewerPhone(ctx context.Context, in *ChangeViewerPhoneRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
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
	// Exchange a refresh token for new token material; refresh tokens can only be used once
	// If refresh tokens are used more than once RTR dictates that any access tokens which were created by it should be immediately revoked
	// this is because this indicates an attack (something is wrong)
	TokenExchange(ctx context.Context, in *TokenExchangeRequest, opts ...grpc.CallOption) (*TokenExchangeResponse, error)
	// Post chosen bank info from plaid in order to create a new ACH pyment method in wyre
	//
	// requires a plaid processor token which in turn requires a plaid widget interaction where the user selects the account id
	PlaidConnectBankAccounts(ctx context.Context, in *PlaidConnectBankAccountsRequest, opts ...grpc.CallOption) (*PlaidConnectBankAccountsResponse, error)
	// https://plaid.com/docs/link/link-token-migration-guide/
	PlaidCreateLinkToken(ctx context.Context, in *PlaidCreateLinkTokenRequest, opts ...grpc.CallOption) (*PlaidCreateLinkTokenResponse, error)
	// SaveProfileData saves profile data items for the user
	//
	// ...
	SaveProfileData(ctx context.Context, in *SaveProfileDataRequest, opts ...grpc.CallOption) (*ProfileDataInfo, error)
	WyreWebhook(ctx context.Context, in *WyreWebhookRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	WyreGetPaymentMethods(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*WyrePaymentMethods, error)
	WyreCreateTransfer(ctx context.Context, in *WyreCreateTransferRequest, opts ...grpc.CallOption) (*WyreTransfer, error)
	WyreConfirmTransfer(ctx context.Context, in *WyreConfirmTransferRequest, opts ...grpc.CallOption) (*WyreTransfer, error)
	WyreGetTransfer(ctx context.Context, in *WyreGetTransferRequest, opts ...grpc.CallOption) (*WyreTransfer, error)
	WyreGetTransfers(ctx context.Context, in *WyreGetTransfersRequest, opts ...grpc.CallOption) (*WyreTransfers, error)
	WidgetGetShortUrl(ctx context.Context, in *SnapWidgetConfig, opts ...grpc.CallOption) (*WidgetGetShortUrlResponse, error)
	// UploadFile uploads a file and returns a file id
	//
	// ...
	UploadFile(ctx context.Context, in *UploadFileRequest, opts ...grpc.CallOption) (*UploadFileResponse, error)
	// GetImage returns an image with optionally specified resize proportions
	//
	// The image is reference via a file ID; the blob data will be returned as well as the mimetype and size.
	//
	// If the file is not of an image mime type, you will get an InvalidArguments error
	GetImage(ctx context.Context, in *GetImageRequest, opts ...grpc.CallOption) (*GetImageResponse, error)
	Goto(ctx context.Context, in *GotoRequest, opts ...grpc.CallOption) (*GotoResponse, error)
}

type fluxClient struct {
	cc grpc.ClientConnInterface
}

func NewFluxClient(cc grpc.ClientConnInterface) FluxClient {
	return &fluxClient{cc}
}

func (c *fluxClient) ViewerData(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ViewerDataResponse, error) {
	out := new(ViewerDataResponse)
	err := c.cc.Invoke(ctx, "/Flux/ViewerData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fluxClient) ViewerProfileData(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ProfileDataInfo, error) {
	out := new(ProfileDataInfo)
	err := c.cc.Invoke(ctx, "/Flux/ViewerProfileData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fluxClient) ChangeViewerEmail(ctx context.Context, in *ChangeViewerEmailRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Flux/ChangeViewerEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fluxClient) ChangeViewerPhone(ctx context.Context, in *ChangeViewerPhoneRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Flux/ChangeViewerPhone", in, out, opts...)
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

func (c *fluxClient) TokenExchange(ctx context.Context, in *TokenExchangeRequest, opts ...grpc.CallOption) (*TokenExchangeResponse, error) {
	out := new(TokenExchangeResponse)
	err := c.cc.Invoke(ctx, "/Flux/TokenExchange", in, out, opts...)
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

func (c *fluxClient) SaveProfileData(ctx context.Context, in *SaveProfileDataRequest, opts ...grpc.CallOption) (*ProfileDataInfo, error) {
	out := new(ProfileDataInfo)
	err := c.cc.Invoke(ctx, "/Flux/SaveProfileData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fluxClient) WyreWebhook(ctx context.Context, in *WyreWebhookRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Flux/WyreWebhook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fluxClient) WyreGetPaymentMethods(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*WyrePaymentMethods, error) {
	out := new(WyrePaymentMethods)
	err := c.cc.Invoke(ctx, "/Flux/WyreGetPaymentMethods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fluxClient) WyreCreateTransfer(ctx context.Context, in *WyreCreateTransferRequest, opts ...grpc.CallOption) (*WyreTransfer, error) {
	out := new(WyreTransfer)
	err := c.cc.Invoke(ctx, "/Flux/WyreCreateTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fluxClient) WyreConfirmTransfer(ctx context.Context, in *WyreConfirmTransferRequest, opts ...grpc.CallOption) (*WyreTransfer, error) {
	out := new(WyreTransfer)
	err := c.cc.Invoke(ctx, "/Flux/WyreConfirmTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fluxClient) WyreGetTransfer(ctx context.Context, in *WyreGetTransferRequest, opts ...grpc.CallOption) (*WyreTransfer, error) {
	out := new(WyreTransfer)
	err := c.cc.Invoke(ctx, "/Flux/WyreGetTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fluxClient) WyreGetTransfers(ctx context.Context, in *WyreGetTransfersRequest, opts ...grpc.CallOption) (*WyreTransfers, error) {
	out := new(WyreTransfers)
	err := c.cc.Invoke(ctx, "/Flux/WyreGetTransfers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fluxClient) WidgetGetShortUrl(ctx context.Context, in *SnapWidgetConfig, opts ...grpc.CallOption) (*WidgetGetShortUrlResponse, error) {
	out := new(WidgetGetShortUrlResponse)
	err := c.cc.Invoke(ctx, "/Flux/WidgetGetShortUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fluxClient) UploadFile(ctx context.Context, in *UploadFileRequest, opts ...grpc.CallOption) (*UploadFileResponse, error) {
	out := new(UploadFileResponse)
	err := c.cc.Invoke(ctx, "/Flux/UploadFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fluxClient) GetImage(ctx context.Context, in *GetImageRequest, opts ...grpc.CallOption) (*GetImageResponse, error) {
	out := new(GetImageResponse)
	err := c.cc.Invoke(ctx, "/Flux/GetImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fluxClient) Goto(ctx context.Context, in *GotoRequest, opts ...grpc.CallOption) (*GotoResponse, error) {
	out := new(GotoResponse)
	err := c.cc.Invoke(ctx, "/Flux/Goto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FluxServer is the server API for Flux service.
// All implementations must embed UnimplementedFluxServer
// for forward compatibility
type FluxServer interface {
	// Get viewer data
	//
	// Provides user (viewer) data associated with the access token
	ViewerData(context.Context, *emptypb.Empty) (*ViewerDataResponse, error)
	// Get viewer profile data
	//
	// Provides user (viewer) data associated with the access token
	ViewerProfileData(context.Context, *emptypb.Empty) (*ProfileDataInfo, error)
	// Change users email (viewer based on jwt)
	//
	// requires an otp code and the desired email address change
	ChangeViewerEmail(context.Context, *ChangeViewerEmailRequest) (*emptypb.Empty, error)
	// Change users phone (viewer based on jwt)
	//
	// requires an otp code and the desired phone change
	ChangeViewerPhone(context.Context, *ChangeViewerPhoneRequest) (*emptypb.Empty, error)
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
	// Exchange a refresh token for new token material; refresh tokens can only be used once
	// If refresh tokens are used more than once RTR dictates that any access tokens which were created by it should be immediately revoked
	// this is because this indicates an attack (something is wrong)
	TokenExchange(context.Context, *TokenExchangeRequest) (*TokenExchangeResponse, error)
	// Post chosen bank info from plaid in order to create a new ACH pyment method in wyre
	//
	// requires a plaid processor token which in turn requires a plaid widget interaction where the user selects the account id
	PlaidConnectBankAccounts(context.Context, *PlaidConnectBankAccountsRequest) (*PlaidConnectBankAccountsResponse, error)
	// https://plaid.com/docs/link/link-token-migration-guide/
	PlaidCreateLinkToken(context.Context, *PlaidCreateLinkTokenRequest) (*PlaidCreateLinkTokenResponse, error)
	// SaveProfileData saves profile data items for the user
	//
	// ...
	SaveProfileData(context.Context, *SaveProfileDataRequest) (*ProfileDataInfo, error)
	WyreWebhook(context.Context, *WyreWebhookRequest) (*emptypb.Empty, error)
	WyreGetPaymentMethods(context.Context, *emptypb.Empty) (*WyrePaymentMethods, error)
	WyreCreateTransfer(context.Context, *WyreCreateTransferRequest) (*WyreTransfer, error)
	WyreConfirmTransfer(context.Context, *WyreConfirmTransferRequest) (*WyreTransfer, error)
	WyreGetTransfer(context.Context, *WyreGetTransferRequest) (*WyreTransfer, error)
	WyreGetTransfers(context.Context, *WyreGetTransfersRequest) (*WyreTransfers, error)
	WidgetGetShortUrl(context.Context, *SnapWidgetConfig) (*WidgetGetShortUrlResponse, error)
	// UploadFile uploads a file and returns a file id
	//
	// ...
	UploadFile(context.Context, *UploadFileRequest) (*UploadFileResponse, error)
	// GetImage returns an image with optionally specified resize proportions
	//
	// The image is reference via a file ID; the blob data will be returned as well as the mimetype and size.
	//
	// If the file is not of an image mime type, you will get an InvalidArguments error
	GetImage(context.Context, *GetImageRequest) (*GetImageResponse, error)
	Goto(context.Context, *GotoRequest) (*GotoResponse, error)
	mustEmbedUnimplementedFluxServer()
}

// UnimplementedFluxServer must be embedded to have forward compatible implementations.
type UnimplementedFluxServer struct {
}

func (UnimplementedFluxServer) ViewerData(context.Context, *emptypb.Empty) (*ViewerDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewerData not implemented")
}
func (UnimplementedFluxServer) ViewerProfileData(context.Context, *emptypb.Empty) (*ProfileDataInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewerProfileData not implemented")
}
func (UnimplementedFluxServer) ChangeViewerEmail(context.Context, *ChangeViewerEmailRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeViewerEmail not implemented")
}
func (UnimplementedFluxServer) ChangeViewerPhone(context.Context, *ChangeViewerPhoneRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeViewerPhone not implemented")
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
func (UnimplementedFluxServer) TokenExchange(context.Context, *TokenExchangeRequest) (*TokenExchangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TokenExchange not implemented")
}
func (UnimplementedFluxServer) PlaidConnectBankAccounts(context.Context, *PlaidConnectBankAccountsRequest) (*PlaidConnectBankAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaidConnectBankAccounts not implemented")
}
func (UnimplementedFluxServer) PlaidCreateLinkToken(context.Context, *PlaidCreateLinkTokenRequest) (*PlaidCreateLinkTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaidCreateLinkToken not implemented")
}
func (UnimplementedFluxServer) SaveProfileData(context.Context, *SaveProfileDataRequest) (*ProfileDataInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveProfileData not implemented")
}
func (UnimplementedFluxServer) WyreWebhook(context.Context, *WyreWebhookRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WyreWebhook not implemented")
}
func (UnimplementedFluxServer) WyreGetPaymentMethods(context.Context, *emptypb.Empty) (*WyrePaymentMethods, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WyreGetPaymentMethods not implemented")
}
func (UnimplementedFluxServer) WyreCreateTransfer(context.Context, *WyreCreateTransferRequest) (*WyreTransfer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WyreCreateTransfer not implemented")
}
func (UnimplementedFluxServer) WyreConfirmTransfer(context.Context, *WyreConfirmTransferRequest) (*WyreTransfer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WyreConfirmTransfer not implemented")
}
func (UnimplementedFluxServer) WyreGetTransfer(context.Context, *WyreGetTransferRequest) (*WyreTransfer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WyreGetTransfer not implemented")
}
func (UnimplementedFluxServer) WyreGetTransfers(context.Context, *WyreGetTransfersRequest) (*WyreTransfers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WyreGetTransfers not implemented")
}
func (UnimplementedFluxServer) WidgetGetShortUrl(context.Context, *SnapWidgetConfig) (*WidgetGetShortUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WidgetGetShortUrl not implemented")
}
func (UnimplementedFluxServer) UploadFile(context.Context, *UploadFileRequest) (*UploadFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedFluxServer) GetImage(context.Context, *GetImageRequest) (*GetImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImage not implemented")
}
func (UnimplementedFluxServer) Goto(context.Context, *GotoRequest) (*GotoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Goto not implemented")
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

func _Flux_ViewerData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxServer).ViewerData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Flux/ViewerData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxServer).ViewerData(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Flux_ViewerProfileData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxServer).ViewerProfileData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Flux/ViewerProfileData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxServer).ViewerProfileData(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Flux_ChangeViewerEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeViewerEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxServer).ChangeViewerEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Flux/ChangeViewerEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxServer).ChangeViewerEmail(ctx, req.(*ChangeViewerEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Flux_ChangeViewerPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeViewerPhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxServer).ChangeViewerPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Flux/ChangeViewerPhone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxServer).ChangeViewerPhone(ctx, req.(*ChangeViewerPhoneRequest))
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

func _Flux_TokenExchange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenExchangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxServer).TokenExchange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Flux/TokenExchange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxServer).TokenExchange(ctx, req.(*TokenExchangeRequest))
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

func _Flux_SaveProfileData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveProfileDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxServer).SaveProfileData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Flux/SaveProfileData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxServer).SaveProfileData(ctx, req.(*SaveProfileDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Flux_WyreWebhook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WyreWebhookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxServer).WyreWebhook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Flux/WyreWebhook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxServer).WyreWebhook(ctx, req.(*WyreWebhookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Flux_WyreGetPaymentMethods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxServer).WyreGetPaymentMethods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Flux/WyreGetPaymentMethods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxServer).WyreGetPaymentMethods(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Flux_WyreCreateTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WyreCreateTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxServer).WyreCreateTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Flux/WyreCreateTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxServer).WyreCreateTransfer(ctx, req.(*WyreCreateTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Flux_WyreConfirmTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WyreConfirmTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxServer).WyreConfirmTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Flux/WyreConfirmTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxServer).WyreConfirmTransfer(ctx, req.(*WyreConfirmTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Flux_WyreGetTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WyreGetTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxServer).WyreGetTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Flux/WyreGetTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxServer).WyreGetTransfer(ctx, req.(*WyreGetTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Flux_WyreGetTransfers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WyreGetTransfersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxServer).WyreGetTransfers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Flux/WyreGetTransfers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxServer).WyreGetTransfers(ctx, req.(*WyreGetTransfersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Flux_WidgetGetShortUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SnapWidgetConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxServer).WidgetGetShortUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Flux/WidgetGetShortUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxServer).WidgetGetShortUrl(ctx, req.(*SnapWidgetConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _Flux_UploadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxServer).UploadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Flux/UploadFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxServer).UploadFile(ctx, req.(*UploadFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Flux_GetImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxServer).GetImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Flux/GetImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxServer).GetImage(ctx, req.(*GetImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Flux_Goto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxServer).Goto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Flux/Goto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxServer).Goto(ctx, req.(*GotoRequest))
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
			MethodName: "ViewerData",
			Handler:    _Flux_ViewerData_Handler,
		},
		{
			MethodName: "ViewerProfileData",
			Handler:    _Flux_ViewerProfileData_Handler,
		},
		{
			MethodName: "ChangeViewerEmail",
			Handler:    _Flux_ChangeViewerEmail_Handler,
		},
		{
			MethodName: "ChangeViewerPhone",
			Handler:    _Flux_ChangeViewerPhone_Handler,
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
			MethodName: "TokenExchange",
			Handler:    _Flux_TokenExchange_Handler,
		},
		{
			MethodName: "PlaidConnectBankAccounts",
			Handler:    _Flux_PlaidConnectBankAccounts_Handler,
		},
		{
			MethodName: "PlaidCreateLinkToken",
			Handler:    _Flux_PlaidCreateLinkToken_Handler,
		},
		{
			MethodName: "SaveProfileData",
			Handler:    _Flux_SaveProfileData_Handler,
		},
		{
			MethodName: "WyreWebhook",
			Handler:    _Flux_WyreWebhook_Handler,
		},
		{
			MethodName: "WyreGetPaymentMethods",
			Handler:    _Flux_WyreGetPaymentMethods_Handler,
		},
		{
			MethodName: "WyreCreateTransfer",
			Handler:    _Flux_WyreCreateTransfer_Handler,
		},
		{
			MethodName: "WyreConfirmTransfer",
			Handler:    _Flux_WyreConfirmTransfer_Handler,
		},
		{
			MethodName: "WyreGetTransfer",
			Handler:    _Flux_WyreGetTransfer_Handler,
		},
		{
			MethodName: "WyreGetTransfers",
			Handler:    _Flux_WyreGetTransfers_Handler,
		},
		{
			MethodName: "WidgetGetShortUrl",
			Handler:    _Flux_WidgetGetShortUrl_Handler,
		},
		{
			MethodName: "UploadFile",
			Handler:    _Flux_UploadFile_Handler,
		},
		{
			MethodName: "GetImage",
			Handler:    _Flux_GetImage_Handler,
		},
		{
			MethodName: "Goto",
			Handler:    _Flux_Goto_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
