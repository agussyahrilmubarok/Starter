package exception

import "net/http"

type Exception struct {
	Code    int
	Message string
	Err     map[string]string
}

func HttpErrMap(key, err string) map[string]string {
	return map[string]string{
		key: err,
	}
}

func (e *Exception) Error() string {
	return e.Message
}

// 4xx Client Errors

// 400
func NewBadRequest(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusBadRequest, msg, errors}
}

// 401
func NewUnauthorized(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusUnauthorized, msg, errors}
}

// 402
func NewPaymentRequired(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusPaymentRequired, msg, errors}
}

// 403
func NewForbidden(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusForbidden, msg, errors}
}

// 404
func NewNotFound(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusNotFound, msg, errors}
}

// 405
func NewMethodNotAllowed(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusMethodNotAllowed, msg, errors}
}

// 406
func NewNotAcceptable(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusNotAcceptable, msg, errors}
}

// 407
func NewProxyAuthRequired(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusProxyAuthRequired, msg, errors}
}

// 408
func NewRequestTimeout(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusRequestTimeout, msg, errors}
}

// 409
func NewConflict(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusConflict, msg, errors}
}

// 410
func NewGone(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusGone, msg, errors}
}

// 411
func NewLengthRequired(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusLengthRequired, msg, errors}
}

// 412
func NewPreconditionFailed(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusPreconditionFailed, msg, errors}
}

// 413
func NewRequestEntityTooLarge(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusRequestEntityTooLarge, msg, errors}
}

// 414
func NewRequestURITooLong(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusRequestURITooLong, msg, errors}
}

// 415
func NewUnsupportedMediaType(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusUnsupportedMediaType, msg, errors}
}

// 416
func NewRequestedRangeNotSatisfiable(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusRequestedRangeNotSatisfiable, msg, errors}
}

// 417
func NewExpectationFailed(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusExpectationFailed, msg, errors}
}

// 418
func NewTeapot(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusTeapot, msg, errors}
}

// 421
func NewMisdirectedRequest(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusMisdirectedRequest, msg, errors}
}

// 422
func NewUnprocessableEntity(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusUnprocessableEntity, msg, errors}
}

// 423
func NewLocked(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusLocked, msg, errors}
}

// 424
func NewFailedDependency(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusFailedDependency, msg, errors}
}

// 425
func NewTooEarly(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusTooEarly, msg, errors}
}

// 426
func NewUpgradeRequired(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusUpgradeRequired, msg, errors}
}

// 428
func NewPreconditionRequired(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusPreconditionRequired, msg, errors}
}

// 429
func NewTooManyRequests(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusTooManyRequests, msg, errors}
}

// 431
func NewRequestHeaderFieldsTooLarge(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusRequestHeaderFieldsTooLarge, msg, errors}
}

// 451
func NewUnavailableForLegalReasons(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusUnavailableForLegalReasons, msg, errors}
}

// 5xx Server Errors

// 500
func NewInternalServer(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusInternalServerError, msg, errors}
}

// 501
func NewNotImplemented(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusNotImplemented, msg, errors}
}

// 502
func NewBadGateway(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusBadGateway, msg, errors}
}

// 503
func NewServiceUnavailable(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusServiceUnavailable, msg, errors}
}

// 504
func NewGatewayTimeout(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusGatewayTimeout, msg, errors}
}

// 505
func NewHTTPVersionNotSupported(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusHTTPVersionNotSupported, msg, errors}
}

// 506
func NewVariantAlsoNegotiates(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusVariantAlsoNegotiates, msg, errors}
}

// 507
func NewInsufficientStorage(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusInsufficientStorage, msg, errors}
}

// 508
func NewLoopDetected(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusLoopDetected, msg, errors}
}

// 510
func NewNotExtended(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusNotExtended, msg, errors}
}

// 511
func NewNetworkAuthenticationRequired(msg string, errors map[string]string) *Exception {
	return &Exception{http.StatusNetworkAuthenticationRequired, msg, errors}
}
