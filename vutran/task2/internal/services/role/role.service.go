package services

import (
	"context"
	"task2/internal/exceptions"
	repositories "task2/internal/repositories/role"
)

type RoleService interface {
	CreateRole(ctx context.Context, name string) exceptions.HttpError
	AssignRole(ctx context.Context, userID string, roleName string) exceptions.HttpError
	AddPermissionToRole(ctx context.Context, roleName string, permissionName string) exceptions.HttpError
	CheckPermission(ctx context.Context, userID string, permission string) (bool, exceptions.HttpError)
	GetUserRoles(ctx context.Context, userID string) ([]string, exceptions.HttpError)
	CheckRoleOfUser(ctx context.Context, userID, role string) (bool, exceptions.HttpError)
}

type roleService struct {
	roleRepo repositories.RoleRepository
}

func NewRoleService(roleRepo repositories.RoleRepository) RoleService {
	return &roleService{roleRepo}
}

func (s *roleService) CreateRole(ctx context.Context, name string) exceptions.HttpError {
	return s.roleRepo.InsertRole(ctx, name)
}

func (s *roleService) AssignRole(ctx context.Context, userID string, roleName string) exceptions.HttpError {
	roleID, err := s.roleRepo.FindRoleIDByName(ctx, roleName)
	if err.Code != 0 {
		return err
	}
	return s.roleRepo.AssignUserRole(ctx, userID, roleID)
}

func (s *roleService) AddPermissionToRole(ctx context.Context, roleName, permissionName string) exceptions.HttpError {
	roleID, err := s.roleRepo.FindRoleIDByName(ctx, roleName)
	if err.Code != 0 {
		return err
	}

	permID, err := s.roleRepo.FindPermissionIDByName(ctx, permissionName)
	if err.Code != 0 {
		return err
	}

	return s.roleRepo.AssignRolePermission(ctx, roleID, permID)
}

func (s *roleService) CheckPermission(ctx context.Context, userID string, permission string) (bool, exceptions.HttpError) {
	return s.roleRepo.HasUserPermission(ctx, userID, permission)
}

func (s *roleService) GetUserRoles(ctx context.Context, userID string) ([]string, exceptions.HttpError) {
	return s.roleRepo.GetUserRoleNames(ctx, userID)
}

func (s *roleService) CheckRoleOfUser(ctx context.Context, userID, role string) (bool, exceptions.HttpError) {
	return s.roleRepo.CheckRoleOfUser(ctx, userID, role)
}
