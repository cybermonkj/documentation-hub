// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
APIVersion Get the version of the API
*/
func (a *Client) APIVersion(params *APIVersionParams) (*APIVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPIVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "APIVersion",
		Method:             "GET",
		PathPattern:        "/api-version",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &APIVersionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*APIVersionOK), nil

}

/*
API Get the Api description
*/
func (a *Client) API(params *APIParams) (*APIOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPIParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Api",
		Method:             "GET",
		PathPattern:        "/api",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &APIReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*APIOK), nil

}

/*
CompileContract Compile a sophia contract from source and return byte code
*/
func (a *Client) CompileContract(params *CompileContractParams) (*CompileContractOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCompileContractParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CompileContract",
		Method:             "POST",
		PathPattern:        "/compile",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CompileContractReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CompileContractOK), nil

}

/*
DecodeCallResult Decode the result of contract call
*/
func (a *Client) DecodeCallResult(params *DecodeCallResultParams) (*DecodeCallResultOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDecodeCallResultParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DecodeCallResult",
		Method:             "POST",
		PathPattern:        "/decode-call-result",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DecodeCallResultReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DecodeCallResultOK), nil

}

/*
DecodeCalldataBytecode Identify function name and arguments in Calldata for a compiled contract
*/
func (a *Client) DecodeCalldataBytecode(params *DecodeCalldataBytecodeParams) (*DecodeCalldataBytecodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDecodeCalldataBytecodeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DecodeCalldataBytecode",
		Method:             "POST",
		PathPattern:        "/decode-calldata/bytecode",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DecodeCalldataBytecodeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DecodeCalldataBytecodeOK), nil

}

/*
DecodeCalldataSource Identify function name and arguments in Calldata for a (partial) contract
*/
func (a *Client) DecodeCalldataSource(params *DecodeCalldataSourceParams) (*DecodeCalldataSourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDecodeCalldataSourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DecodeCalldataSource",
		Method:             "POST",
		PathPattern:        "/decode-calldata/source",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DecodeCalldataSourceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DecodeCalldataSourceOK), nil

}

/*
DecodeData Decode data as retuned by a contract call. - Legacy decoding
*/
func (a *Client) DecodeData(params *DecodeDataParams) (*DecodeDataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDecodeDataParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DecodeData",
		Method:             "POST",
		PathPattern:        "/decode-data",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DecodeDataReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DecodeDataOK), nil

}

/*
EncodeCalldata Encode Sophia function call according to sophia ABI.
*/
func (a *Client) EncodeCalldata(params *EncodeCalldataParams) (*EncodeCalldataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEncodeCalldataParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EncodeCalldata",
		Method:             "POST",
		PathPattern:        "/encode-calldata",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &EncodeCalldataReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EncodeCalldataOK), nil

}

/*
GenerateACI Generate an Aeternity Contract Interface (ACI) for contract
*/
func (a *Client) GenerateACI(params *GenerateACIParams) (*GenerateACIOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGenerateACIParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GenerateACI",
		Method:             "POST",
		PathPattern:        "/aci",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GenerateACIReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GenerateACIOK), nil

}

/*
Version Get the version of the underlying Sophia compiler version
*/
func (a *Client) Version(params *VersionParams) (*VersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Version",
		Method:             "GET",
		PathPattern:        "/version",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &VersionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*VersionOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
