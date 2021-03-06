package logic

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
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-08-01-preview/logic instead.
// IntegrationAccountCertificatesClient is the REST API for Azure Logic Apps.
type IntegrationAccountCertificatesClient struct {
	BaseClient
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-08-01-preview/logic instead.
// NewIntegrationAccountCertificatesClient creates an instance of the IntegrationAccountCertificatesClient client.
func NewIntegrationAccountCertificatesClient(subscriptionID string) IntegrationAccountCertificatesClient {
	return NewIntegrationAccountCertificatesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-08-01-preview/logic instead.
// NewIntegrationAccountCertificatesClientWithBaseURI creates an instance of the IntegrationAccountCertificatesClient
// client.
func NewIntegrationAccountCertificatesClientWithBaseURI(baseURI string, subscriptionID string) IntegrationAccountCertificatesClient {
	return IntegrationAccountCertificatesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-08-01-preview/logic instead.
// CreateOrUpdate creates or updates an integration account certificate.
//
// resourceGroupName is the resource group name. integrationAccountName is the integration account name.
// certificateName is the integration account certificate name. certificate is the integration account certificate.
func (client IntegrationAccountCertificatesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, integrationAccountName string, certificateName string, certificate IntegrationAccountCertificate) (result IntegrationAccountCertificate, err error) {
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, integrationAccountName, certificateName, certificate)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountCertificatesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountCertificatesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountCertificatesClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-08-01-preview/logic instead.
// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client IntegrationAccountCertificatesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, integrationAccountName string, certificateName string, certificate IntegrationAccountCertificate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"certificateName":        autorest.Encode("path", certificateName),
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/certificates/{certificateName}", pathParameters),
		autorest.WithJSON(certificate),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-08-01-preview/logic instead.
// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationAccountCertificatesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-08-01-preview/logic instead.
// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client IntegrationAccountCertificatesClient) CreateOrUpdateResponder(resp *http.Response) (result IntegrationAccountCertificate, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-08-01-preview/logic instead.
// Delete deletes an integration account certificate.
//
// resourceGroupName is the resource group name. integrationAccountName is the integration account name.
// certificateName is the integration account certificate name.
func (client IntegrationAccountCertificatesClient) Delete(ctx context.Context, resourceGroupName string, integrationAccountName string, certificateName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(ctx, resourceGroupName, integrationAccountName, certificateName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountCertificatesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountCertificatesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountCertificatesClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-08-01-preview/logic instead.
// DeletePreparer prepares the Delete request.
func (client IntegrationAccountCertificatesClient) DeletePreparer(ctx context.Context, resourceGroupName string, integrationAccountName string, certificateName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"certificateName":        autorest.Encode("path", certificateName),
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/certificates/{certificateName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-08-01-preview/logic instead.
// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationAccountCertificatesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-08-01-preview/logic instead.
// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client IntegrationAccountCertificatesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-08-01-preview/logic instead.
// Get gets an integration account certificate.
//
// resourceGroupName is the resource group name. integrationAccountName is the integration account name.
// certificateName is the integration account certificate name.
func (client IntegrationAccountCertificatesClient) Get(ctx context.Context, resourceGroupName string, integrationAccountName string, certificateName string) (result IntegrationAccountCertificate, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, integrationAccountName, certificateName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountCertificatesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountCertificatesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountCertificatesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-08-01-preview/logic instead.
// GetPreparer prepares the Get request.
func (client IntegrationAccountCertificatesClient) GetPreparer(ctx context.Context, resourceGroupName string, integrationAccountName string, certificateName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"certificateName":        autorest.Encode("path", certificateName),
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/certificates/{certificateName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-08-01-preview/logic instead.
// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationAccountCertificatesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-08-01-preview/logic instead.
// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client IntegrationAccountCertificatesClient) GetResponder(resp *http.Response) (result IntegrationAccountCertificate, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-08-01-preview/logic instead.
// List gets a list of integration account certificates.
//
// resourceGroupName is the resource group name. integrationAccountName is the integration account name. top is the
// number of items to be included in the result.
func (client IntegrationAccountCertificatesClient) List(ctx context.Context, resourceGroupName string, integrationAccountName string, top *int32) (result IntegrationAccountCertificateListResultPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, integrationAccountName, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountCertificatesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.iaclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountCertificatesClient", "List", resp, "Failure sending request")
		return
	}

	result.iaclr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountCertificatesClient", "List", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-08-01-preview/logic instead.
// ListPreparer prepares the List request.
func (client IntegrationAccountCertificatesClient) ListPreparer(ctx context.Context, resourceGroupName string, integrationAccountName string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/certificates", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-08-01-preview/logic instead.
// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationAccountCertificatesClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-08-01-preview/logic instead.
// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client IntegrationAccountCertificatesClient) ListResponder(resp *http.Response) (result IntegrationAccountCertificateListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client IntegrationAccountCertificatesClient) listNextResults(lastResults IntegrationAccountCertificateListResult) (result IntegrationAccountCertificateListResult, err error) {
	req, err := lastResults.integrationAccountCertificateListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.IntegrationAccountCertificatesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic.IntegrationAccountCertificatesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountCertificatesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-08-01-preview/logic instead.
// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client IntegrationAccountCertificatesClient) ListComplete(ctx context.Context, resourceGroupName string, integrationAccountName string, top *int32) (result IntegrationAccountCertificateListResultIterator, err error) {
	result.page, err = client.List(ctx, resourceGroupName, integrationAccountName, top)
	return
}
