package models

import "encoding/json"

// APIResponse is the generic envelope for all API responses.
type APIResponse struct {
	Data  json.RawMessage `json:"data"`
	Error *string         `json:"error"`
}
