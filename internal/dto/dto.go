package dto

import "net/url"

type ServiceDTO interface {
	ParseRequest(url.Values)
	Validate()
}
