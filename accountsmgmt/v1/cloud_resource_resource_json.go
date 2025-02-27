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

package v1 // github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1

import (
	"io"
	"net/http"
)

func readCloudResourceDeleteRequest(request *CloudResourceDeleteServerRequest, r *http.Request) error {
	return nil
}
func writeCloudResourceDeleteRequest(request *CloudResourceDeleteRequest, writer io.Writer) error {
	return nil
}
func readCloudResourceDeleteResponse(response *CloudResourceDeleteResponse, reader io.Reader) error {
	return nil
}
func writeCloudResourceDeleteResponse(response *CloudResourceDeleteServerResponse, w http.ResponseWriter) error {
	return nil
}
func readCloudResourceGetRequest(request *CloudResourceGetServerRequest, r *http.Request) error {
	return nil
}
func writeCloudResourceGetRequest(request *CloudResourceGetRequest, writer io.Writer) error {
	return nil
}
func readCloudResourceGetResponse(response *CloudResourceGetResponse, reader io.Reader) error {
	var err error
	response.body, err = UnmarshalCloudResource(reader)
	return err
}
func writeCloudResourceGetResponse(response *CloudResourceGetServerResponse, w http.ResponseWriter) error {
	return MarshalCloudResource(response.body, w)
}
func readCloudResourceUpdateRequest(request *CloudResourceUpdateServerRequest, r *http.Request) error {
	var err error
	request.body, err = UnmarshalCloudResource(r.Body)
	return err
}
func writeCloudResourceUpdateRequest(request *CloudResourceUpdateRequest, writer io.Writer) error {
	return MarshalCloudResource(request.body, writer)
}
func readCloudResourceUpdateResponse(response *CloudResourceUpdateResponse, reader io.Reader) error {
	var err error
	response.body, err = UnmarshalCloudResource(reader)
	return err
}
func writeCloudResourceUpdateResponse(response *CloudResourceUpdateServerResponse, w http.ResponseWriter) error {
	return MarshalCloudResource(response.body, w)
}
