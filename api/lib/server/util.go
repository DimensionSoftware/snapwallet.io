package server

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"log"

	"github.com/khoerling/flux/api/lib/db"
	"github.com/khoerling/flux/api/lib/db/models/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// ProfileField represents PII data which is used during the create account process
type EmailTemplateVars struct {
	OTPCode       string `json:"otpCode"`
	TransactionID string `json:"transactionID"`
	BusinessDays  int    `json:"businessDays"`
	Status        string `json:"status"`
}

// GetUserIDFromIncomingContext gets the user id from the incoming context
func GetUserIDFromIncomingContext(ctx context.Context) user.ID {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	vals := md.Get("user-id")

	if len(vals) == 0 {
		return ""
	}

	return user.ID(vals[0])
}

// LookupUserFromIncomingContext attempts to lookup the user from the incoming context
func LookupUserFromIncomingContext(ctx context.Context, db db.Db) (*user.User, error) {
	userID := GetUserIDFromIncomingContext(ctx)

	u, err := db.GetUserByID(ctx, nil, user.ID(userID))
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Unknown, "An unknown error ocurred; please try again.")
	}

	return u, nil
}

// RequireUserFromIncomingContext requires a user from the incoming context and will a throw an error if the lookup fails
func RequireUserFromIncomingContext(ctx context.Context, db db.Db) (*user.User, error) {
	u, err := LookupUserFromIncomingContext(ctx, db)
	if err != nil {
		return nil, err
	}

	if u == nil {
		return nil, status.Errorf(codes.Unauthenticated, genMsgUnauthenticatedGeneric())
	}

	return u, nil
}

/**
Generates an HTML email template. Template name should map to the
"define" block in the main HTML template (usually at the top of the file.)
*/
func genEmailTemplate(templateName string, templateVars EmailTemplateVars) (string, error) {
	errMsg := "Unable to process your email. Please contact support@snapwallet.io"

	// TODO: read into memory once
	t, err := template.ParseGlob("lib/server/templates/*")

	fmt.Println(t)
	if err != nil {
		fmt.Println(err)
		return "", status.Error(codes.Internal, errMsg)
	}

	var body bytes.Buffer
	t.ExecuteTemplate(&body, templateName, templateVars)

	return body.String(), nil
}
