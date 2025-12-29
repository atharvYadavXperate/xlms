package cutomeerror

import "errors"

var (
	ContentTypeConflictMustJson error = errors.New("Content Type must be json")
	InvalidJson                 error = errors.New("Invalid Json formate")
	FieldsAreRequired           error = errors.New("All filed are required")
	UserNotApproved             error = errors.New("User not approved")
	InvalidOtp                  error = errors.New("Invalid Otp")
)
