// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// AccessPackageSubjectRequestBuilder is request builder for AccessPackageSubject
type AccessPackageSubjectRequestBuilder struct{ BaseRequestBuilder }

// Request returns AccessPackageSubjectRequest
func (b *AccessPackageSubjectRequestBuilder) Request() *AccessPackageSubjectRequest {
	return &AccessPackageSubjectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AccessPackageSubjectRequest is request for AccessPackageSubject
type AccessPackageSubjectRequest struct{ BaseRequest }

// Get performs GET request for AccessPackageSubject
func (r *AccessPackageSubjectRequest) Get(ctx context.Context) (resObj *AccessPackageSubject, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AccessPackageSubject
func (r *AccessPackageSubjectRequest) Update(ctx context.Context, reqObj *AccessPackageSubject) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AccessPackageSubject
func (r *AccessPackageSubjectRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}