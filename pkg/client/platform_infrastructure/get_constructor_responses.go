// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by go-swagger; DO NOT EDIT.

package platform_infrastructure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// GetConstructorReader is a Reader for the GetConstructor structure.
type GetConstructorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConstructorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConstructorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetConstructorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetConstructorOK creates a GetConstructorOK with default headers values
func NewGetConstructorOK() *GetConstructorOK {
	return &GetConstructorOK{}
}

/*GetConstructorOK handles this case with default header values.

The information for the constructor specified by {constructor_id}.
*/
type GetConstructorOK struct {
	Payload *models.ConstructorInfo
}

func (o *GetConstructorOK) Error() string {
	return fmt.Sprintf("[GET /platform/infrastructure/constructors/{constructor_id}][%d] getConstructorOK  %+v", 200, o.Payload)
}

func (o *GetConstructorOK) GetPayload() *models.ConstructorInfo {
	return o.Payload
}

func (o *GetConstructorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConstructorInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConstructorNotFound creates a GetConstructorNotFound with default headers values
func NewGetConstructorNotFound() *GetConstructorNotFound {
	return &GetConstructorNotFound{}
}

/*GetConstructorNotFound handles this case with default header values.

The constructor specified by {constructor_id} cannot be found. (code: `constructors.constructor_not_found`)
*/
type GetConstructorNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *GetConstructorNotFound) Error() string {
	return fmt.Sprintf("[GET /platform/infrastructure/constructors/{constructor_id}][%d] getConstructorNotFound  %+v", 404, o.Payload)
}

func (o *GetConstructorNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetConstructorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
