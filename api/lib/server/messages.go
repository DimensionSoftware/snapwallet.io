package server

import "github.com/khoerling/flux/api/lib/db/models/onetimepasscode"

func genMsgUnauthenticatedOTP(login onetimepasscode.LoginKind) string {
	switch login {
	case onetimepasscode.LoginKindPhone:
		return "The phone code provided was not valid. Please try again."
	default:
		return "The email code provided was not valid. Please try again."
	}
}

func genMsgUnauthenticatedGeneric() string {
	return "Could not authenticate"
}
