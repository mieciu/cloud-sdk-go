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

// SearchAllocatorsReader is a Reader for the SearchAllocators structure.
type SearchAllocatorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchAllocatorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchAllocatorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchAllocatorsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSearchAllocatorsOK creates a SearchAllocatorsOK with default headers values
func NewSearchAllocatorsOK() *SearchAllocatorsOK {
	return &SearchAllocatorsOK{}
}

/*SearchAllocatorsOK handles this case with default header values.

An overview of allocators that matched the given search query
*/
type SearchAllocatorsOK struct {
	Payload *models.AllocatorOverview
}

func (o *SearchAllocatorsOK) Error() string {
	return fmt.Sprintf("[POST /platform/infrastructure/allocators/_search][%d] searchAllocatorsOK  %+v", 200, o.Payload)
}

func (o *SearchAllocatorsOK) GetPayload() *models.AllocatorOverview {
	return o.Payload
}

func (o *SearchAllocatorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AllocatorOverview)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchAllocatorsBadRequest creates a SearchAllocatorsBadRequest with default headers values
func NewSearchAllocatorsBadRequest() *SearchAllocatorsBadRequest {
	return &SearchAllocatorsBadRequest{}
}

/*SearchAllocatorsBadRequest handles this case with default header values.

The search request failed
*/
type SearchAllocatorsBadRequest struct {
	Payload *models.BasicFailedReply
}

func (o *SearchAllocatorsBadRequest) Error() string {
	return fmt.Sprintf("[POST /platform/infrastructure/allocators/_search][%d] searchAllocatorsBadRequest  %+v", 400, o.Payload)
}

func (o *SearchAllocatorsBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SearchAllocatorsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
