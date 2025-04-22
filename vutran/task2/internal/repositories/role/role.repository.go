package repositories

import (
	"context"
	"fmt"
	entities "task2/internal/entities/role"
	"task2/internal/exceptions"

	"github.com/jmoiron/sqlx"
)

type RoleRepository interface {
	InsertRole(ctx context.Context, name string) exceptions.HttpError
	FindRoleIDByName(ctx context.Context, name string) (string, exceptions.HttpError)
	AssignUserRole(ctx context.Context, userID string, roleID string) exceptions.HttpError
	FindPermissionIDByName(ctx context.Context, name string) (string, exceptions.HttpError)
	AssignRolePermission(ctx context.Context, roleID string, permissionID string) exceptions.HttpError
	HasUserPermission(ctx context.Context, userID string, permission string) (bool, exceptions.HttpError)
	GetUserRoleNames(ctx context.Context, userID string) ([]string, exceptions.HttpError)
	CheckRoleOfUser(ctx context.Context, userId, role string) (bool, exceptions.HttpError)
	GetByName(ctx context.Context, name string) (*entities.Role, exceptions.HttpError)
}

type roleRepo struct {
	db *sqlx.DB
}

func NewRoleRepository(db *sqlx.DB) RoleRepository {
	return &roleRepo{db}
}

func (r *roleRepo) InsertRole(ctx context.Context, name string) exceptions.HttpError {
	_, err := r.db.ExecContext(ctx, `INSERT INTO roles(name) VALUES($1) ON CONFLICT DO NOTHING`, name)

	if err != nil {
		return exceptions.NewInternal("")
	}

	return exceptions.HttpError{}
}

func (r *roleRepo) FindRoleIDByName(ctx context.Context, name string) (string, exceptions.HttpError) {
	var id string
	err := r.db.GetContext(ctx, &id, `SELECT id FROM roles WHERE name = $1`, name)

	if err != nil {
		return "", exceptions.NewNotFound(fmt.Sprintf("role with name %s not found", name))
	}

	return id, exceptions.HttpError{}
}

func (r *roleRepo) AssignUserRole(ctx context.Context, userID string, roleID string) exceptions.HttpError {
	_, err := r.db.ExecContext(ctx, `
	  INSERT INTO user_roles(user_id, role_id)
	  VALUES ($1, $2)
	  ON CONFLICT DO NOTHING`, userID, roleID)

	if err != nil {
		fmt.Println(err)
		return exceptions.NewInternal("")
	}
	return exceptions.HttpError{}
}

func (r *roleRepo) FindPermissionIDByName(ctx context.Context, name string) (string, exceptions.HttpError) {
	var id string
	err := r.db.GetContext(ctx, &id, `SELECT id FROM permissions WHERE name = $1`, name)

	if err != nil {
		return "", exceptions.NewNotFound(fmt.Sprintf("permission with name %s not found", name))
	}

	return id, exceptions.HttpError{}
}

func (r *roleRepo) AssignRolePermission(ctx context.Context, roleID, permissionID string) exceptions.HttpError {
	_, err := r.db.ExecContext(ctx, `
	  INSERT INTO role_permissions(role_id, permission_id)
	  VALUES ($1, $2)
	  ON CONFLICT DO NOTHING`, roleID, permissionID)

	if err != nil {
		return exceptions.NewInternal("")
	}

	return exceptions.HttpError{}
}

func (r *roleRepo) HasUserPermission(ctx context.Context, userID string, permission string) (bool, exceptions.HttpError) {
	var count int
	err := r.db.GetContext(ctx, &count, `
	  SELECT COUNT(*) FROM user_roles ur
	  JOIN role_permissions rp ON ur.role_id = rp.role_id
	  JOIN permissions p ON rp.permission_id = p.id
	  WHERE ur.user_id = $1 AND p.name = $2
	`, userID, permission)

	if err != nil {
		return false, exceptions.NewNotFound("")
	}

	return count > 0, exceptions.HttpError{}
}

func (r *roleRepo) GetUserRoleNames(ctx context.Context, userID string) ([]string, exceptions.HttpError) {
	var roles []string
	err := r.db.SelectContext(ctx, &roles, `
	  SELECT r.name FROM user_roles ur
	  JOIN roles r ON ur.role_id = r.id
	  WHERE ur.user_id = $1`, userID)

	if err != nil {
		return nil, exceptions.NewNotFound("")
	}

	return roles, exceptions.HttpError{}
}

func (r *roleRepo) CheckRoleOfUser(ctx context.Context, userId, role string) (bool, exceptions.HttpError) {
	var count int
	err := r.db.GetContext(ctx, &count, `
		SELECT COUNT(*) FROM user_roles ur
		JOIN roles r ON ur.role_id = r.id
		WHERE ur.user_id = $1 AND r.name = $2
	`, userId, role)

	if err != nil {
		return false, exceptions.NewInternal(err.Error())
	}

	return count > 0, exceptions.HttpError{}
}

func (r *roleRepo) GetByName(ctx context.Context, name string) (*entities.Role, exceptions.HttpError) {
	var role entities.Role
	err := r.db.GetContext(ctx, &role, "SELECT * FROM roles WHERE name=$1", name)

	if err != nil {
		return &role, exceptions.NewNotFound(fmt.Sprintf("Not found role with name = %s", name))
	}
	return &role, exceptions.HttpError{}
}
