package Chit_Chat_Server

var (
	SUCCESS = &Result{
		Message: "Success",
		Code:    "1",
		Reason:  "",
	}

	UNAUTHERIZED_USER = &Result{
		Message: "Failure",
		Code:    "2",
		Reason:  "Unauthorized user",
	}

	BAD_UPDATE = &Result{
		Message: "Failure",
		Code:    "3",
		Reason:  "Unsuccessful update.",
	}

	MISSING_PARAMS = &Result{
		Message: "Failure",
		Code:    "4",
		Reason:  "Some parameters were missing or incorrect.",
	}

	BAD_EMAIL = &Result{
		Message: "Failure",
		Code:    "5",
		Reason:  "email was ether not given or with a improper format. Email must contain champlain.edu",
	}

	FAILED_TO_MODIFY = &Result{
		Message: "Failure",
		Code:    "6",
		Reason:  "Could not modify database",
	}

	BAD_METHOD = &Result{
		Message: "Failure",
		Code: "7",
		Reason: "Bad method given",
	}
)
