// Package model contains structs for Google API authorization and Google+ event pictures.
package model

// authorizationResponse represents the response of a Google API Access Token request.
type AuthorizationResponse struct {
    Access_token string `json:"access_token"`
    Token_type   string `json:"token_type"`
    Expires_in   int    `json:"expires_in"`
}

// picture reprents a Google+ picture with its URL for the task payload.
type Picture struct {
    Url string `json:"url"`
}