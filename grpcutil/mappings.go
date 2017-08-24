package grpcutil

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

// CodeFromHTTPErrorStatus given an http status code 4xx or 5xx returns a gRPC code
func CodeFromHTTPStatus(httpStatus int) codes.Code {
	switch httpStatus {

	// 1xx, 2xx, and 3xx map all to codes.OK
	case http.StatusContinue:
	case http.StatusSwitchingProtocols:
	case http.StatusProcessing:
	case http.StatusOK:
	case http.StatusCreated:
	case http.StatusAccepted:
	case http.StatusNonAuthoritativeInfo:
	case http.StatusNoContent:
	case http.StatusResetContent:
	case http.StatusPartialContent:
	case http.StatusMultiStatus:
	case http.StatusAlreadyReported:
	case http.StatusIMUsed:
	case http.StatusMultipleChoices:
	case http.StatusMovedPermanently:
	case http.StatusFound:
	case http.StatusSeeOther:
	case http.StatusNotModified:
	case http.StatusUseProxy:
	case http.StatusTemporaryRedirect:
	case http.StatusPermanentRedirect:
		return codes.OK

	case http.StatusBadRequest:
	case http.StatusMethodNotAllowed:
	case http.StatusNotAcceptable:
	case http.StatusLengthRequired:
	case http.StatusRequestEntityTooLarge:
	case http.StatusRequestURITooLong:
	case http.StatusRequestHeaderFieldsTooLarge:
	case http.StatusUnsupportedMediaType:
	case http.StatusRequestedRangeNotSatisfiable:
	case http.StatusExpectationFailed:
		return codes.InvalidArgument

	case http.StatusRequestTimeout:
		return codes.DeadlineExceeded

	case http.StatusNotFound:
	case http.StatusGone:
		return codes.NotFound

	case http.StatusConflict:
		return codes.AlreadyExists
	case http.StatusProxyAuthRequired:
		return codes.Unauthenticated

	case http.StatusForbidden:
		return codes.PermissionDenied

	case http.StatusUnauthorized:
		return codes.Unauthenticated

	case http.StatusPreconditionFailed:
		return codes.FailedPrecondition

	case http.StatusNotImplemented:
		return codes.Unimplemented

	case http.StatusInternalServerError:
	case http.StatusHTTPVersionNotSupported:
	case http.StatusVariantAlsoNegotiates:
	case http.StatusInsufficientStorage:
	case http.StatusLoopDetected:
		return codes.Internal

	case http.StatusBadGateway:
	case http.StatusServiceUnavailable:
	case http.StatusGatewayTimeout:
		return codes.Unavailable

	}
	// For the following httpStatusCode we do not have a good
	// match with gRPC codes so we return the generic 'Unknown' code:
	// StatusPaymentRequired               = 402 // RFC 7231, 6.5.2
	// StatusTeapot                        = 418 // RFC 7168, 2.3.3
	// StatusUnprocessableEntity           = 422 // RFC 4918, 11.2
	// StatusLocked                        = 423 // RFC 4918, 11.3
	// StatusFailedDependency              = 424 // RFC 4918, 11.4
	// StatusUpgradeRequired               = 426 // RFC 7231, 6.5.15
	// StatusPreconditionRequired          = 428 // RFC 6585, 3
	// StatusTooManyRequests               = 429 // RFC 6585, 4
	// StatusUnavailableForLegalReasons    = 451 // RFC 7725, 3
	// StatusNotExtended                   = 510 // RFC 2774, 7
	// StatusNetworkAuthenticationRequired = 511 // RFC 6585, 6

	return codes.Unknown
}
