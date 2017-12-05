//-----------------------------------------------------------------------------
// Copyright (c) 2017 Oracle and/or its affiliates.  All rights reserved.
// This program is free software: you can modify it and/or redistribute it
// under the terms of:
//
// (i)  the Universal Permissive License v 1.0 or at your option, any
//      later version (http://oss.oracle.com/licenses/upl); and/or
//
// (ii) the Apache License v 2.0. (http://www.apache.org/licenses/LICENSE-2.0)
//-----------------------------------------------------------------------------

package grpcutil

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

// CodeFromHTTPStatus given an http status code 4xx or 5xx returns a gRPC code
func CodeFromHTTPStatus(httpStatus int) codes.Code {
	switch httpStatus {

	// 1xx, 2xx, and 3xx all map to codes.OK
	case
		http.StatusContinue,
		http.StatusSwitchingProtocols,
		http.StatusProcessing,
		http.StatusOK,
		http.StatusCreated,
		http.StatusAccepted,
		http.StatusNonAuthoritativeInfo,
		http.StatusNoContent,
		http.StatusResetContent,
		http.StatusPartialContent,
		http.StatusMultiStatus,
		http.StatusAlreadyReported,
		http.StatusIMUsed,
		http.StatusMultipleChoices,
		http.StatusMovedPermanently,
		http.StatusFound,
		http.StatusSeeOther,
		http.StatusNotModified,
		http.StatusUseProxy,
		http.StatusTemporaryRedirect,
		http.StatusPermanentRedirect:
		return codes.OK

	case
		http.StatusBadRequest,
		http.StatusMethodNotAllowed,
		http.StatusNotAcceptable,
		http.StatusLengthRequired,
		http.StatusRequestEntityTooLarge,
		http.StatusRequestURITooLong,
		http.StatusRequestHeaderFieldsTooLarge,
		http.StatusUnsupportedMediaType,
		http.StatusRequestedRangeNotSatisfiable,
		http.StatusExpectationFailed:
		return codes.InvalidArgument

	case http.StatusRequestTimeout:
		return codes.DeadlineExceeded

	case
		http.StatusNotFound,
		http.StatusGone:
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

	case
		http.StatusInternalServerError,
		http.StatusHTTPVersionNotSupported,
		http.StatusVariantAlsoNegotiates,
		http.StatusInsufficientStorage,
		http.StatusLoopDetected:
		return codes.Internal

	case
		http.StatusBadGateway,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout:
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
