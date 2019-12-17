package requests

// AuthQueryStruct struct
type AuthQueryStruct struct {
	ClientID     string `schema:"client_id"`
	RedirectURI  string `schema:"redirect_uri"`
	ResponseType string `schema:"response_type"`
	Scope        string `schema:"scope"`
	State        string `schema:"state"`
}
