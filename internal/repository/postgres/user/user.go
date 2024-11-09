package user

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/RoseRocket/xerrs"
	"github.com/jmoiron/sqlx"

	cons "github.com/aaalik/anton-users/internal/constant"
	"github.com/aaalik/anton-users/internal/model"
	"github.com/aaalik/anton-users/internal/service"
)

func (ur *UserRepository) CreateUser(ctx context.Context, tx *sqlx.Tx, user *model.User) error {
	query := fmt.Sprintf(`
		INSERT INTO
		%s
		(
			id, username, password, name,
			dob, gender, created_at, updated_at
		)
		VALUES
		(
			:id, :username, :password, :name,
			:dob, :gender, :created_at, :updated_at
		)
	`, cons.SqlUserTable)

	params := map[string]interface{}{
		"id":         user.Id,
		"username":   user.Username,
		"password":   user.Password,
		"name":       user.Name,
		"dob":        user.Dob,
		"gender":     user.Gender,
		"created_at": time.Now().Unix(),
		"updated_at": time.Now().Unix(),
	}

	stmt, err := tx.PrepareNamedContext(ctx, query)
	if err != nil {
		err = xerrs.Mask(err, cons.ErrorSQLUpdateUser)
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, params)
	if err != nil {
		err = xerrs.Mask(err, cons.ErrorSQLUpdateUser)
		return err
	}

	return nil
}

func (ur *UserRepository) UpdateUser(ctx context.Context, tx *sqlx.Tx, user *model.User) error {
	query := fmt.Sprintf(`
		UPDATE
		%s
		SET
		name = :name,
		dob = :dob,
		gender = :gender,
		updated_at = :updated_at
		WHERE id = :id
	`, cons.SqlUserTable)

	params := map[string]interface{}{
		"id":         user.Id,
		"name":       user.Name,
		"dob":        user.Dob,
		"gender":     user.Gender,
		"updated_at": time.Now().Unix(),
	}

	stmt, err := tx.PrepareNamedContext(ctx, query)
	if err != nil {
		err = xerrs.Mask(err, cons.ErrorSQLUpdateUser)
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, params)
	if err != nil {
		err = xerrs.Mask(err, cons.ErrorSQLUpdateUser)
		return err
	}

	return nil
}

func (ur *UserRepository) DeleteUser(ctx context.Context, tx *sqlx.Tx, id string) error {
	query := fmt.Sprintf(`
		UPDATE
		%s
		SET
		deleted_at = :deleted_at
		WHERE id = :id
	`, cons.SqlUserTable)

	params := map[string]interface{}{
		"id":         id,
		"deleted_at": time.Now().Unix(),
	}

	stmt, err := tx.PrepareNamedContext(ctx, query)
	if err != nil {
		err = xerrs.Mask(err, cons.ErrorSQLUpdateUser)
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, params)
	if err != nil {
		err = xerrs.Mask(err, cons.ErrorSQLUpdateUser)
		return err
	}

	return nil
}

func (ur *UserRepository) DetailUser(ctx context.Context, id string) (*model.User, error) {
	user := model.User{}
	query := fmt.Sprintf(`
		SELECT
		id, username, name, dob,
		gender, created_at, updated_at
		FROM %s
		WHERE id = :id
	`, cons.SqlUserTable)

	params := map[string]interface{}{
		"id": id,
	}

	stmt, err := ur.dbr.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRowxContext(ctx, params).StructScan(&user)
	if err == sql.ErrNoRows {
		return nil, xerrs.Extend(cons.ErrorDataNotFound)
	} else if err != nil {
		return nil, xerrs.Mask(err, cons.ErrorSQLRead)
	}

	return &user, nil
}

func (ur *UserRepository) ListUser(ctx context.Context, request *service.RequestListUser) ([]*model.User, error) {
	users := []*model.User{}
	selectQuery := fmt.Sprintf(`
		SELECT
		id, username, name, dob,
		gender, created_at, updated_at
		FROM %s
	`, cons.SqlUserTable)

	whereQuery, args := ur.buildUserFilterQuery(request)
	orderQuery := ur.buildSortQuery(request.Queries)

	query := strings.Join([]string{selectQuery, whereQuery, orderQuery}, " ")
	query = ur.dbr.Rebind(query)

	stmt, err := ur.dbr.PreparexContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryxContext(ctx, args...)
	if err == sql.ErrNoRows {
		return nil, xerrs.Mask(err, cons.ErrorSQLRead)
	}
	defer rows.Close()
	for rows.Next() {
		user := model.User{}
		err = rows.StructScan(&user)
		if err != nil {
			return nil, xerrs.Mask(err, cons.ErrorSQLRead)
		}
		users = append(users, &user)
	}

	return users, nil
}

func (ur *UserRepository) CountUsers(ctx context.Context, request *service.RequestListUser) (int32, error) {
	count := int32(0)
	selectQuery := fmt.Sprintf(`
		SELECT
		COUNT(u.id)
		FROM %s AS u 
	`, cons.SqlUserTable)

	whereQuery, args := ur.buildUserFilterQuery(request)

	query := strings.Join([]string{selectQuery, whereQuery}, " ")
	query = ur.dbr.Rebind(query)

	stmt, err := ur.dbr.PreparexContext(ctx, query)
	if err != nil {
		return 0, xerrs.Mask(err, cons.ErrorSQLRead)
	}
	defer stmt.Close()

	err = stmt.QueryRowxContext(ctx, args...).Scan(&count)

	if err != nil {
		return 0, xerrs.Mask(err, cons.ErrorSQLRead)
	}

	return count, nil
}

func (ur *UserRepository) buildUserFilterQuery(req *service.RequestListUser) (string, []interface{}) {
	var whereQuery string
	var whereQueries []string
	var whereArgs []interface{}

	filter := req.Includes

	if ql, arguments, err := sqlx.In("id IN (?)", filter.Ids); err == nil {
		whereQueries = append(whereQueries, ql)
		whereArgs = append(whereArgs, arguments...)
	}

	if req.Queries.Keyword != "" {
		whereQueries = append(whereQueries, "name ~* ?")
		whereArgs = append(whereArgs, req.Queries.Keyword)
	}

	if ql, arguments, err := sqlx.In("dob IN (?)", filter.Dobs); err == nil {
		whereQueries = append(whereQueries, ql)
		whereArgs = append(whereArgs, arguments...)
	}

	if filter.CreatedAt.GTE != 0 {
		whereQueries = append(whereQueries, "created_at >= ?")
		whereArgs = append(whereArgs, filter.CreatedAt.GTE)
	}

	if filter.CreatedAt.LTE != 0 {
		whereQueries = append(whereQueries, "created_at <= ?")
		whereArgs = append(whereArgs, filter.CreatedAt.LTE)
	}

	if filter.DeletedAt.GTE != 0 {
		whereQueries = append(whereQueries, "deleted_at >= ?")
		whereArgs = append(whereArgs, filter.DeletedAt.GTE)
	}

	if filter.DeletedAt.LTE != 0 {
		whereQueries = append(whereQueries, "deleted_at <= ?")
		whereArgs = append(whereArgs, filter.DeletedAt.LTE)
	}

	withDeleted := false
	for _, isDeleted := range req.Queries.WithDeleted {
		if isDeleted {
			withDeleted = true
		}
	}

	if !withDeleted {
		whereQueries = append(whereQueries, "deleted_at = 0")
	}

	if len(whereQueries) > 0 {
		whereQuery = "WHERE " + strings.Join(whereQueries, " AND ")
	}

	return whereQuery, whereArgs
}

func (ur *UserRepository) buildSortQuery(queries *service.Queries) string {
	var query string
	var limit int32
	var offset int32

	orderBy := "ORDER BY created_at DESC"
	if queries.Sort != nil {
		if queries.Sort.Order == service.ORDER_ASC {
			orderBy = fmt.Sprintf("ORDER BY %s ASC", queries.Sort.Field)
		}
		if queries.Sort.Order == service.ORDER_DESC {
			orderBy = fmt.Sprintf("ORDER BY %s DESC", queries.Sort.Field)
		}
	}

	query += orderBy

	limit = queries.Rows
	if limit == 0 {
		limit = 10
	}

	offset = 0
	if queries.Page > 0 {
		offset = (queries.Page - 1) * limit
	}

	if queries.Page != 0 && queries.Rows != 0 {
		query += fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset)
	}

	return query
}

func (ur *UserRepository) FetchUserLogin(ctx context.Context, username string) (*model.User, error) {
	user := model.User{}
	query := fmt.Sprintf(`
		SELECT
		id, username, password
		FROM %s
		WHERE username = :username
		AND deleted_at = 0
	`, cons.SqlUserTable)

	params := map[string]interface{}{
		"username": username,
	}

	stmt, err := ur.dbr.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRowxContext(ctx, params).StructScan(&user)
	if err == sql.ErrNoRows {
		return nil, xerrs.Extend(cons.ErrorDataNotFound)
	} else if err != nil {
		return nil, xerrs.Mask(err, cons.ErrorSQLRead)
	}

	return &user, nil
}
