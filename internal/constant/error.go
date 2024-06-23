package constant

import (
	"github.com/RoseRocket/xerrs"
)

var (
	// 404 Errors
	ErrorDataNotFound = xerrs.New("data not found")

	// 422 Errors
	ErrorInvalidUser = xerrs.New("User invalid")

	// 500 Errors
	ErrorSQLCreateTransaction   = xerrs.New("failed to write data")
	ErrorSQLCommitTransaction   = xerrs.New("failed to commit data")
	ErrorSQLRollbackTransaction = xerrs.New("failed to rollback data")
	ErrorSQLRead                = xerrs.New("failed to retrieve data")
	ErrFileProcessing           = xerrs.New("failed to process file")
	ErrInvalidContentType       = xerrs.New("invalid content type")

	ErrorSQLCreateUser = xerrs.New("failed to create user")
	ErrorSQLUpdateUser = xerrs.New("failed to update user")
	ErrorSQLDeleteUser = xerrs.New("failed to delete user")
)

func init() {
	xerrs.SetData(ErrorDataNotFound, "status_code", 404)

	xerrs.SetData(ErrorInvalidUser, "status_code", 422)
}
