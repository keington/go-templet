package code

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/11/20 21:51
 * @file: code_const's.go
 * @description: code const's
 */
var (
	Nil                       = statusCode{-1, ""}                             // No error code specified.
	Success                   = statusCode{0, "Success"}                       // It is Success.
	InternalError             = statusCode{50, "Internal Error"}               // An error occurred internally.
	ValidationFailed          = statusCode{51, "Validation Failed"}            // Detail validation failed.
	DbOperationError          = statusCode{52, "Database Operation Error"}     // Database operation error.
	InvalidParameter          = statusCode{53, "Invalid Parameter"}            // The given parameter for current operation is invalid.
	MissingParameter          = statusCode{54, "Missing Parameter"}            // Parameter for current operation is missing.
	InvalidOperation          = statusCode{55, "Invalid Operation"}            // The function cannot be used like this.
	InvalidConfiguration      = statusCode{56, "Invalid Configuration"}        // The configuration is invalid for current operation.
	MissingConfiguration      = statusCode{57, "Missing Configuration"}        // The configuration is missing for current operation.
	NotImplemented            = statusCode{58, "Not Implemented"}              // The operation is not implemented yet.
	NotSupported              = statusCode{59, "Not Supported"}                // The operation is not supported yet.
	OperationFailed           = statusCode{60, "Operation Failed"}             // I tried, but I cannot give you what you want.
	NotAuthorized             = statusCode{61, "Not Authorized"}               // Not Authorized.
	SecurityReason            = statusCode{62, "Security Reason"}              // Security Reason.
	ServerBusy                = statusCode{63, "Server Is Busy"}               // Server is busy, please try again later.
	Unknown                   = statusCode{64, "Unknown Error"}                // Unknown error.
	NotFound                  = statusCode{65, "Not Found"}                    // Resource does not exist.
	InvalidRequest            = statusCode{66, "Invalid Request"}              // Invalid request.
	NecessaryPackageNotImport = statusCode{67, "Necessary Package Not Import"} // It needs necessary package import.
	InternalPanic             = statusCode{68, "Internal Panic"}               // A panic occurred internally.
	BusinessValidationFailed  = statusCode{300, "Business Validation Failed"}  // Business validation failed.
)
