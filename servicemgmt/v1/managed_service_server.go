/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/servicemgmt/v1

import (
	"context"
	"net/http"

	"github.com/golang/glog"
	"github.com/openshift-online/ocm-sdk-go/errors"
)

// ManagedServiceServer represents the interface the manages the 'managed_service' resource.
type ManagedServiceServer interface {

	// Delete handles a request for the 'delete' method.
	//
	// Deletes the Managed Service
	Delete(ctx context.Context, request *ManagedServiceDeleteServerRequest, response *ManagedServiceDeleteServerResponse) error

	// Get handles a request for the 'get' method.
	//
	// Gets information on the Managed Service
	Get(ctx context.Context, request *ManagedServiceGetServerRequest, response *ManagedServiceGetServerResponse) error
}

// ManagedServiceDeleteServerRequest is the request for the 'delete' method.
type ManagedServiceDeleteServerRequest struct {
}

// ManagedServiceDeleteServerResponse is the response for the 'delete' method.
type ManagedServiceDeleteServerResponse struct {
	status int
	err    *errors.Error
}

// Status sets the status code.
func (r *ManagedServiceDeleteServerResponse) Status(value int) *ManagedServiceDeleteServerResponse {
	r.status = value
	return r
}

// ManagedServiceGetServerRequest is the request for the 'get' method.
type ManagedServiceGetServerRequest struct {
}

// ManagedServiceGetServerResponse is the response for the 'get' method.
type ManagedServiceGetServerResponse struct {
	status int
	err    *errors.Error
	body   *ManagedService
}

// Body sets the value of the 'body' parameter.
//
//
func (r *ManagedServiceGetServerResponse) Body(value *ManagedService) *ManagedServiceGetServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *ManagedServiceGetServerResponse) Status(value int) *ManagedServiceGetServerResponse {
	r.status = value
	return r
}

// dispatchManagedService navigates the servers tree rooted at the given server
// till it finds one that matches the given set of path segments, and then invokes
// the corresponding server.
func dispatchManagedService(w http.ResponseWriter, r *http.Request, server ManagedServiceServer, segments []string) {
	if len(segments) == 0 {
		switch r.Method {
		case "DELETE":
			adaptManagedServiceDeleteRequest(w, r, server)
			return
		case "GET":
			adaptManagedServiceGetRequest(w, r, server)
			return
		default:
			errors.SendMethodNotAllowed(w, r)
			return
		}
	}
	switch segments[0] {
	default:
		errors.SendNotFound(w, r)
		return
	}
}

// adaptManagedServiceDeleteRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptManagedServiceDeleteRequest(w http.ResponseWriter, r *http.Request, server ManagedServiceServer) {
	request := &ManagedServiceDeleteServerRequest{}
	err := readManagedServiceDeleteRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &ManagedServiceDeleteServerResponse{}
	response.status = 204
	err = server.Delete(r.Context(), request, response)
	if err != nil {
		glog.Errorf(
			"Can't process request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	err = writeManagedServiceDeleteResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}

// adaptManagedServiceGetRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptManagedServiceGetRequest(w http.ResponseWriter, r *http.Request, server ManagedServiceServer) {
	request := &ManagedServiceGetServerRequest{}
	err := readManagedServiceGetRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &ManagedServiceGetServerResponse{}
	response.status = 200
	err = server.Get(r.Context(), request, response)
	if err != nil {
		glog.Errorf(
			"Can't process request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	err = writeManagedServiceGetResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}
