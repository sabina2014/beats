// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/TerminateClientVpnConnectionsRequest
type TerminateClientVpnConnectionsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Client VPN endpoint to which the client is connected.
	//
	// ClientVpnEndpointId is a required field
	ClientVpnEndpointId *string `type:"string" required:"true"`

	// The ID of the client connection to be terminated.
	ConnectionId *string `type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The name of the user who initiated the connection. Use this option to terminate
	// all active connections for the specified user. This option can only be used
	// if the user has established up to five connections.
	Username *string `type:"string"`
}

// String returns the string representation
func (s TerminateClientVpnConnectionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TerminateClientVpnConnectionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TerminateClientVpnConnectionsInput"}

	if s.ClientVpnEndpointId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientVpnEndpointId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/TerminateClientVpnConnectionsResult
type TerminateClientVpnConnectionsOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the Client VPN endpoint.
	ClientVpnEndpointId *string `locationName:"clientVpnEndpointId" type:"string"`

	// The current state of the client connections.
	ConnectionStatuses []TerminateConnectionStatus `locationName:"connectionStatuses" locationNameList:"item" type:"list"`

	// The user who established the terminated client connections.
	Username *string `locationName:"username" type:"string"`
}

// String returns the string representation
func (s TerminateClientVpnConnectionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opTerminateClientVpnConnections = "TerminateClientVpnConnections"

// TerminateClientVpnConnectionsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Terminates active Client VPN endpoint connections. This action can be used
// to terminate a specific client connection, or up to five connections established
// by a specific user.
//
//    // Example sending a request using TerminateClientVpnConnectionsRequest.
//    req := client.TerminateClientVpnConnectionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/TerminateClientVpnConnections
func (c *Client) TerminateClientVpnConnectionsRequest(input *TerminateClientVpnConnectionsInput) TerminateClientVpnConnectionsRequest {
	op := &aws.Operation{
		Name:       opTerminateClientVpnConnections,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &TerminateClientVpnConnectionsInput{}
	}

	req := c.newRequest(op, input, &TerminateClientVpnConnectionsOutput{})
	return TerminateClientVpnConnectionsRequest{Request: req, Input: input, Copy: c.TerminateClientVpnConnectionsRequest}
}

// TerminateClientVpnConnectionsRequest is the request type for the
// TerminateClientVpnConnections API operation.
type TerminateClientVpnConnectionsRequest struct {
	*aws.Request
	Input *TerminateClientVpnConnectionsInput
	Copy  func(*TerminateClientVpnConnectionsInput) TerminateClientVpnConnectionsRequest
}

// Send marshals and sends the TerminateClientVpnConnections API request.
func (r TerminateClientVpnConnectionsRequest) Send(ctx context.Context) (*TerminateClientVpnConnectionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TerminateClientVpnConnectionsResponse{
		TerminateClientVpnConnectionsOutput: r.Request.Data.(*TerminateClientVpnConnectionsOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TerminateClientVpnConnectionsResponse is the response type for the
// TerminateClientVpnConnections API operation.
type TerminateClientVpnConnectionsResponse struct {
	*TerminateClientVpnConnectionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TerminateClientVpnConnections request.
func (r *TerminateClientVpnConnectionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
