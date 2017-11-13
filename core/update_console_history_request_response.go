// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
	"net/http"
)

// UpdateConsoleHistoryRequest wrapper for the UpdateConsoleHistory operation
type UpdateConsoleHistoryRequest struct {

	// The OCID of the console history.
	InstanceConsoleHistoryID *string `mandatory:"true" contributesTo:"path" name:"instanceConsoleHistoryId"`

	// Update instance fields
	UpdateConsoleHistoryDetails `contributesTo:"body"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	// parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	// will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`
}

func (request UpdateConsoleHistoryRequest) String() string {
	return common.PointerString(request)
}

// UpdateConsoleHistoryResponse wrapper for the UpdateConsoleHistory operation
type UpdateConsoleHistoryResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ConsoleHistory instance
	ConsoleHistory `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestID *string `presentIn:"header" name:"opc-request-id"`
}

func (response UpdateConsoleHistoryResponse) String() string {
	return common.PointerString(response)
}