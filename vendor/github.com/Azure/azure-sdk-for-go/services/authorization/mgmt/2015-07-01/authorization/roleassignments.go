package authorization

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// RoleAssignmentsClient is the role based access control provides you a way to apply granular level policy
// administration down to individual resources or resource groups. These operations enable you to manage role
// definitions and role assignments. A role definition describes the set of actions that can be performed on resources.
// A role assignment grants access to Azure Active Directory users.
type RoleAssignmentsClient struct {
	ManagementClient
}

// NewRoleAssignmentsClient creates an instance of the RoleAssignmentsClient client.
func NewRoleAssignmentsClient(subscriptionID string) RoleAssignmentsClient {
	return NewRoleAssignmentsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRoleAssignmentsClientWithBaseURI creates an instance of the RoleAssignmentsClient client.
func NewRoleAssignmentsClientWithBaseURI(baseURI string, subscriptionID string) RoleAssignmentsClient {
	return RoleAssignmentsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates a role assignment.
//
// scope is the scope of the role assignment to create. The scope can be any REST resource instance. For example, use
// '/subscriptions/{subscription-id}/' for a subscription,
// '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}' for a resource group, and
// '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}/providers/{resource-provider}/{resource-type}/{resource-name}'
// for a resource. roleAssignmentName is the name of the role assignment to create. It can be any valid GUID.
// parameters is parameters for the role assignment.
func (client RoleAssignmentsClient) Create(scope string, roleAssignmentName string, parameters RoleAssignmentCreateParameters) (result RoleAssignment, err error) {
	req, err := client.CreatePreparer(scope, roleAssignmentName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client RoleAssignmentsClient) CreatePreparer(scope string, roleAssignmentName string, parameters RoleAssignmentCreateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"roleAssignmentName": autorest.Encode("path", roleAssignmentName),
		"scope":              scope,
	}

	const APIVersion = "2015-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client RoleAssignmentsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client RoleAssignmentsClient) CreateResponder(resp *http.Response) (result RoleAssignment, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateByID creates a role assignment by ID.
//
// roleAssignmentID is the fully qualified ID of the role assignment, including the scope, resource name and resource
// type. Use the format, /{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}. Example:
// /subscriptions/{subId}/resourcegroups/{rgname}//providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}.
// parameters is parameters for the role assignment.
func (client RoleAssignmentsClient) CreateByID(roleAssignmentID string, parameters RoleAssignmentCreateParameters) (result RoleAssignment, err error) {
	req, err := client.CreateByIDPreparer(roleAssignmentID, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "CreateByID", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateByIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "CreateByID", resp, "Failure sending request")
		return
	}

	result, err = client.CreateByIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "CreateByID", resp, "Failure responding to request")
	}

	return
}

// CreateByIDPreparer prepares the CreateByID request.
func (client RoleAssignmentsClient) CreateByIDPreparer(roleAssignmentID string, parameters RoleAssignmentCreateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"roleAssignmentId": roleAssignmentID,
	}

	const APIVersion = "2015-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{roleAssignmentId}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateByIDSender sends the CreateByID request. The method will close the
// http.Response Body if it receives an error.
func (client RoleAssignmentsClient) CreateByIDSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateByIDResponder handles the response to the CreateByID request. The method always
// closes the http.Response Body.
func (client RoleAssignmentsClient) CreateByIDResponder(resp *http.Response) (result RoleAssignment, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a role assignment.
//
// scope is the scope of the role assignment to delete. roleAssignmentName is the name of the role assignment to
// delete.
func (client RoleAssignmentsClient) Delete(scope string, roleAssignmentName string) (result RoleAssignment, err error) {
	req, err := client.DeletePreparer(scope, roleAssignmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client RoleAssignmentsClient) DeletePreparer(scope string, roleAssignmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"roleAssignmentName": autorest.Encode("path", roleAssignmentName),
		"scope":              scope,
	}

	const APIVersion = "2015-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client RoleAssignmentsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client RoleAssignmentsClient) DeleteResponder(resp *http.Response) (result RoleAssignment, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteByID deletes a role assignment.
//
// roleAssignmentID is the fully qualified ID of the role assignment, including the scope, resource name and resource
// type. Use the format, /{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}. Example:
// /subscriptions/{subId}/resourcegroups/{rgname}//providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}.
func (client RoleAssignmentsClient) DeleteByID(roleAssignmentID string) (result RoleAssignment, err error) {
	req, err := client.DeleteByIDPreparer(roleAssignmentID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "DeleteByID", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteByIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "DeleteByID", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteByIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "DeleteByID", resp, "Failure responding to request")
	}

	return
}

// DeleteByIDPreparer prepares the DeleteByID request.
func (client RoleAssignmentsClient) DeleteByIDPreparer(roleAssignmentID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"roleAssignmentId": roleAssignmentID,
	}

	const APIVersion = "2015-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{roleAssignmentId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteByIDSender sends the DeleteByID request. The method will close the
// http.Response Body if it receives an error.
func (client RoleAssignmentsClient) DeleteByIDSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteByIDResponder handles the response to the DeleteByID request. The method always
// closes the http.Response Body.
func (client RoleAssignmentsClient) DeleteByIDResponder(resp *http.Response) (result RoleAssignment, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get the specified role assignment.
//
// scope is the scope of the role assignment. roleAssignmentName is the name of the role assignment to get.
func (client RoleAssignmentsClient) Get(scope string, roleAssignmentName string) (result RoleAssignment, err error) {
	req, err := client.GetPreparer(scope, roleAssignmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RoleAssignmentsClient) GetPreparer(scope string, roleAssignmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"roleAssignmentName": autorest.Encode("path", roleAssignmentName),
		"scope":              scope,
	}

	const APIVersion = "2015-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RoleAssignmentsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RoleAssignmentsClient) GetResponder(resp *http.Response) (result RoleAssignment, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetByID gets a role assignment by ID.
//
// roleAssignmentID is the fully qualified ID of the role assignment, including the scope, resource name and resource
// type. Use the format, /{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}. Example:
// /subscriptions/{subId}/resourcegroups/{rgname}//providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}.
func (client RoleAssignmentsClient) GetByID(roleAssignmentID string) (result RoleAssignment, err error) {
	req, err := client.GetByIDPreparer(roleAssignmentID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "GetByID", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "GetByID", resp, "Failure sending request")
		return
	}

	result, err = client.GetByIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "GetByID", resp, "Failure responding to request")
	}

	return
}

// GetByIDPreparer prepares the GetByID request.
func (client RoleAssignmentsClient) GetByIDPreparer(roleAssignmentID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"roleAssignmentId": roleAssignmentID,
	}

	const APIVersion = "2015-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{roleAssignmentId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetByIDSender sends the GetByID request. The method will close the
// http.Response Body if it receives an error.
func (client RoleAssignmentsClient) GetByIDSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetByIDResponder handles the response to the GetByID request. The method always
// closes the http.Response Body.
func (client RoleAssignmentsClient) GetByIDResponder(resp *http.Response) (result RoleAssignment, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all role assignments for the subscription.
//
// filter is the filter to apply on the operation. Use $filter=atScope() to return all role assignments at or above the
// scope. Use $filter=principalId eq {id} to return all role assignments at, above or below the scope for the specified
// principal.
func (client RoleAssignmentsClient) List(filter string) (result RoleAssignmentListResult, err error) {
	req, err := client.ListPreparer(filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client RoleAssignmentsClient) ListPreparer(filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/roleAssignments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client RoleAssignmentsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RoleAssignmentsClient) ListResponder(resp *http.Response) (result RoleAssignmentListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client RoleAssignmentsClient) ListNextResults(lastResults RoleAssignmentListResult) (result RoleAssignmentListResult, err error) {
	req, err := lastResults.RoleAssignmentListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListComplete gets all elements from the list without paging.
func (client RoleAssignmentsClient) ListComplete(filter string, cancel <-chan struct{}) (<-chan RoleAssignment, <-chan error) {
	resultChan := make(chan RoleAssignment)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.List(filter)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}

// ListForResource gets role assignments for a resource.
//
// resourceGroupName is the name of the resource group. resourceProviderNamespace is the namespace of the resource
// provider. parentResourcePath is the parent resource identity. resourceType is the resource type of the resource.
// resourceName is the name of the resource to get role assignments for. filter is the filter to apply on the
// operation. Use $filter=atScope() to return all role assignments at or above the scope. Use $filter=principalId eq
// {id} to return all role assignments at, above or below the scope for the specified principal.
func (client RoleAssignmentsClient) ListForResource(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, filter string) (result RoleAssignmentListResult, err error) {
	req, err := client.ListForResourcePreparer(resourceGroupName, resourceProviderNamespace, parentResourcePath, resourceType, resourceName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "ListForResource", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListForResourceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "ListForResource", resp, "Failure sending request")
		return
	}

	result, err = client.ListForResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "ListForResource", resp, "Failure responding to request")
	}

	return
}

// ListForResourcePreparer prepares the ListForResource request.
func (client RoleAssignmentsClient) ListForResourcePreparer(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"parentResourcePath":        parentResourcePath,
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"resourceName":              autorest.Encode("path", resourceName),
		"resourceProviderNamespace": autorest.Encode("path", resourceProviderNamespace),
		"resourceType":              resourceType,
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}/providers/Microsoft.Authorization/roleAssignments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListForResourceSender sends the ListForResource request. The method will close the
// http.Response Body if it receives an error.
func (client RoleAssignmentsClient) ListForResourceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListForResourceResponder handles the response to the ListForResource request. The method always
// closes the http.Response Body.
func (client RoleAssignmentsClient) ListForResourceResponder(resp *http.Response) (result RoleAssignmentListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListForResourceNextResults retrieves the next set of results, if any.
func (client RoleAssignmentsClient) ListForResourceNextResults(lastResults RoleAssignmentListResult) (result RoleAssignmentListResult, err error) {
	req, err := lastResults.RoleAssignmentListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "ListForResource", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListForResourceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "ListForResource", resp, "Failure sending next results request")
	}

	result, err = client.ListForResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "ListForResource", resp, "Failure responding to next results request")
	}

	return
}

// ListForResourceComplete gets all elements from the list without paging.
func (client RoleAssignmentsClient) ListForResourceComplete(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, filter string, cancel <-chan struct{}) (<-chan RoleAssignment, <-chan error) {
	resultChan := make(chan RoleAssignment)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListForResource(resourceGroupName, resourceProviderNamespace, parentResourcePath, resourceType, resourceName, filter)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListForResourceNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}

// ListForResourceGroup gets role assignments for a resource group.
//
// resourceGroupName is the name of the resource group. filter is the filter to apply on the operation. Use
// $filter=atScope() to return all role assignments at or above the scope. Use $filter=principalId eq {id} to return
// all role assignments at, above or below the scope for the specified principal.
func (client RoleAssignmentsClient) ListForResourceGroup(resourceGroupName string, filter string) (result RoleAssignmentListResult, err error) {
	req, err := client.ListForResourceGroupPreparer(resourceGroupName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "ListForResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListForResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "ListForResourceGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ListForResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "ListForResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListForResourceGroupPreparer prepares the ListForResourceGroup request.
func (client RoleAssignmentsClient) ListForResourceGroupPreparer(resourceGroupName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Authorization/roleAssignments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListForResourceGroupSender sends the ListForResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client RoleAssignmentsClient) ListForResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListForResourceGroupResponder handles the response to the ListForResourceGroup request. The method always
// closes the http.Response Body.
func (client RoleAssignmentsClient) ListForResourceGroupResponder(resp *http.Response) (result RoleAssignmentListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListForResourceGroupNextResults retrieves the next set of results, if any.
func (client RoleAssignmentsClient) ListForResourceGroupNextResults(lastResults RoleAssignmentListResult) (result RoleAssignmentListResult, err error) {
	req, err := lastResults.RoleAssignmentListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "ListForResourceGroup", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListForResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "ListForResourceGroup", resp, "Failure sending next results request")
	}

	result, err = client.ListForResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "ListForResourceGroup", resp, "Failure responding to next results request")
	}

	return
}

// ListForResourceGroupComplete gets all elements from the list without paging.
func (client RoleAssignmentsClient) ListForResourceGroupComplete(resourceGroupName string, filter string, cancel <-chan struct{}) (<-chan RoleAssignment, <-chan error) {
	resultChan := make(chan RoleAssignment)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListForResourceGroup(resourceGroupName, filter)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListForResourceGroupNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}

// ListForScope gets role assignments for a scope.
//
// scope is the scope of the role assignments. filter is the filter to apply on the operation. Use $filter=atScope() to
// return all role assignments at or above the scope. Use $filter=principalId eq {id} to return all role assignments
// at, above or below the scope for the specified principal.
func (client RoleAssignmentsClient) ListForScope(scope string, filter string) (result RoleAssignmentListResult, err error) {
	req, err := client.ListForScopePreparer(scope, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "ListForScope", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListForScopeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "ListForScope", resp, "Failure sending request")
		return
	}

	result, err = client.ListForScopeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "ListForScope", resp, "Failure responding to request")
	}

	return
}

// ListForScopePreparer prepares the ListForScope request.
func (client RoleAssignmentsClient) ListForScopePreparer(scope string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scope": scope,
	}

	const APIVersion = "2015-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Authorization/roleAssignments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListForScopeSender sends the ListForScope request. The method will close the
// http.Response Body if it receives an error.
func (client RoleAssignmentsClient) ListForScopeSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListForScopeResponder handles the response to the ListForScope request. The method always
// closes the http.Response Body.
func (client RoleAssignmentsClient) ListForScopeResponder(resp *http.Response) (result RoleAssignmentListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListForScopeNextResults retrieves the next set of results, if any.
func (client RoleAssignmentsClient) ListForScopeNextResults(lastResults RoleAssignmentListResult) (result RoleAssignmentListResult, err error) {
	req, err := lastResults.RoleAssignmentListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "ListForScope", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListForScopeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "ListForScope", resp, "Failure sending next results request")
	}

	result, err = client.ListForScopeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentsClient", "ListForScope", resp, "Failure responding to next results request")
	}

	return
}

// ListForScopeComplete gets all elements from the list without paging.
func (client RoleAssignmentsClient) ListForScopeComplete(scope string, filter string, cancel <-chan struct{}) (<-chan RoleAssignment, <-chan error) {
	resultChan := make(chan RoleAssignment)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListForScope(scope, filter)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListForScopeNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}