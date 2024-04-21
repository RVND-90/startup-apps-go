package constants

const (
	SUCCESS_MSG = "Success"
	SUCCESS_CODE = "00"
	SUCCESS_STATUS = 200

	ERROR_USER_NOT_FOUND_MSG = "User not found"
	ERROR_USER_NOT_FOUND_CODE = "01"
	ERROR_USER_NOT_FOUND_STATUS = 404

	ERROR_USER_ALREADY_EXISTS_MSG = "User already exists"
	ERROR_USER_ALREADY_EXISTS_CODE = "02"
	ERROR_USER_ALREADY_EXISTS_STATUS = 409

	ERROR_PARSE_MSG = "Parse error"
	ERROR_PARSE_CODE = "03"
	ERROR_PARSE_STATUS = 400

	ERROR_VALIDATION_MSG = "Validation error"
	ERROR_VALIDATION_CODE = "04"
	ERROR_VALIDATION_STATUS = 400

	ERROR_UNKNOWN_MSG = "Unknown error"
	ERROR_UNKNOWN_CODE = "50"
	ERROR_UNKNOWN_STATUS = 500

)

type RC struct {
	Message string `json:"message"`
	Code string `json:"code"`
	HttpResponseCode int `json:"-"`
}

 func GetRC(code string) *RC{
	switch code {
	case SUCCESS_CODE:
		return &RC{
			Message: SUCCESS_MSG,
			Code: SUCCESS_CODE,
			HttpResponseCode: SUCCESS_STATUS,
		}
	case ERROR_USER_NOT_FOUND_CODE:
		return &RC{
			Message: ERROR_USER_NOT_FOUND_MSG,
			Code: ERROR_USER_NOT_FOUND_CODE,
			HttpResponseCode: ERROR_USER_NOT_FOUND_STATUS,
		}
	case ERROR_USER_ALREADY_EXISTS_CODE:
		return &RC{
			Message: ERROR_USER_ALREADY_EXISTS_MSG,
			Code: ERROR_USER_ALREADY_EXISTS_CODE,
			HttpResponseCode: ERROR_USER_ALREADY_EXISTS_STATUS,
		}
	case ERROR_PARSE_CODE:
		return &RC{
			Message: ERROR_PARSE_MSG,
			Code: ERROR_PARSE_CODE,
			HttpResponseCode: ERROR_PARSE_STATUS,
		}
	case ERROR_VALIDATION_CODE:
		return &RC{
			Message: ERROR_VALIDATION_MSG,
			Code: ERROR_VALIDATION_CODE,
			HttpResponseCode: ERROR_VALIDATION_STATUS,	
		}
	}

	return nil
	
}