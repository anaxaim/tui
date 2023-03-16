package utils

import (
	"net/http"
	"strings"
)

const (
	GetOperation    = "get"
	ListOperation   = "list"
	CreateOperation = "create"
	UpdateOperation = "update"
	PatchOperation  = "patch"
	DeleteOperation = "delete"
)

type RequestInfoResolver interface {
	NewRequestInfo(req *http.Request) (*RequestInfo, error)
}

type RequestInfo struct {
	Path       string
	Verb       string
	APIPrefix  string
	APIGroup   string
	APIVersion string
	Name       string
	Parts      []string
}

type RequestInfoFactory struct {
	APIPrefixes String
}

func (r *RequestInfoFactory) NewRequestInfo(req *http.Request) (*RequestInfo, error) {
	requestInfo := RequestInfo{
		Path: req.URL.Path,
		Verb: strings.ToLower(req.Method),
	}

	currentParts := splitPath(req.URL.Path)
	if len(currentParts) < 3 {
		return &requestInfo, nil
	}

	if !r.APIPrefixes.Has(currentParts[0]) {
		// return a non-resource request
		return &requestInfo, nil
	}
	requestInfo.APIPrefix = currentParts[0]
	currentParts = currentParts[1:]

	requestInfo.APIVersion = currentParts[0]
	currentParts = currentParts[1:]

	switch req.Method {
	case http.MethodPost:
		requestInfo.Verb = CreateOperation
	case http.MethodGet, http.MethodHead:
		requestInfo.Verb = GetOperation
	case http.MethodPut:
		requestInfo.Verb = UpdateOperation
	case http.MethodPatch:
		requestInfo.Verb = PatchOperation
	case http.MethodDelete:
		requestInfo.Verb = DeleteOperation
	default:
		requestInfo.Verb = ""
	}

	requestInfo.Parts = currentParts
	requestInfo.Name = requestInfo.Parts[0]

	if len(requestInfo.Name) == 0 && requestInfo.Verb == GetOperation {
		requestInfo.Verb = ListOperation
	}

	return &requestInfo, nil
}

func splitPath(path string) []string {
	path = strings.Trim(path, "/")
	if path == "" {
		return []string{}
	}
	return strings.Split(path, "/")
}
