package errors

const (
	// 未指定错误
	errNoUnknown = -1

	// 成功
	errNoSuccess = 0

	// 参数错误
	errNoInvalidInput = 1001
	errNoMissingField = 1002
	errNoOutOfRange   = 1003

	// 认证和权限错误
	errNoUnauthorized = 2001
	errNoForbidden    = 2002
	errNoExpiredToken = 2003

	// 资源错误
	errNoNotFound               = 3001
	errNoDuplicateEntry         = 3002
	errNoInsufficientPermission = 3003

	// 内部错误
	errNoInternal      = 4001
	errNoDatabaseError = 4002
	errNoNetworkError  = 4003

	// 请求错误
	errNoBadRequest           = 5001
	errNoMethodNotAllowed     = 5002
	errNoUnsupportedMediaType = 5003

	// 文件和上传错误
	errNoFileTooLarge      = 6001
	errNoInvalidFileFormat = 6002
	errNoFileNotFound      = 6003

	// 连接和通信错误
	errNoConnectionRefused  = 7001
	errNoConnectionTimeout  = 7002
	errNoNetworkUnreachable = 7003

	// 第三方服务错误
	errNoExternalServiceUnavailable = 8001
	errNoExternalServiceTimeout     = 8002
	errNoExternalServiceError       = 8003

	// 安全和加密错误
	errNoEncryptionFailed = 9001
	errNoDecryptionFailed = 9002
	errNoInvalidSignature = 9003
)

var errorMessages = map[int]string{
	// 未指定错误
	errNoUnknown: "Unknown Error",

	// 成功
	errNoSuccess: "Success",

	// 参数错误
	errNoInvalidInput: "Invalid input",
	errNoMissingField: "Missing field",
	errNoOutOfRange:   "Out of range",

	// 认证和权限错误
	errNoUnauthorized: "Unauthorized",
	errNoForbidden:    "Forbidden",
	errNoExpiredToken: "Expired token",

	// 资源错误
	errNoNotFound:               "Not found",
	errNoDuplicateEntry:         "Duplicate entry",
	errNoInsufficientPermission: "Insufficient permission",

	// 内部错误
	errNoInternal:      "Internal error",
	errNoDatabaseError: "Database error",
	errNoNetworkError:  "Network error",

	// 请求错误
	errNoBadRequest:           "Bad request",
	errNoMethodNotAllowed:     "Method not allowed",
	errNoUnsupportedMediaType: "Unsupported media type",

	// 文件和上传错误
	errNoFileTooLarge:      "File too large",
	errNoInvalidFileFormat: "Invalid file format",
	errNoFileNotFound:      "File not found",

	// 连接和通信错误
	errNoConnectionRefused:  "Connection refused",
	errNoConnectionTimeout:  "Connection timeout",
	errNoNetworkUnreachable: "Network unreachable",

	// 第三方服务错误
	errNoExternalServiceUnavailable: "External service unavailable",
	errNoExternalServiceTimeout:     "External service timeout",
	errNoExternalServiceError:       "External service error",

	// 安全和加密错误
	errNoEncryptionFailed: "Encryption failed",
	errNoDecryptionFailed: "Decryption failed",
	errNoInvalidSignature: "Invalid signature",
}
