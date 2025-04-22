package services

// import (
// 	"context"
// 	"task2/internal/exceptions"
// 	repositories "task2/internal/repositories/role"

// 	"github.com/google/uuid"
// )

// type RoleService interface {
// 	CreateRole(ctx context.Context, name string) exceptions.HttpError
// 	AssignRole(ctx context.Context, userID uuid.UUID, roleName string) exceptions.HttpError
// 	AddPermissionToRole(ctx context.Context, roleName string, permissionName string) exceptions.HttpError
// 	CheckPermission(ctx context.Context, userID uuid.UUID, permission string) (bool, exceptions.HttpError)
// 	GetUserRoles(ctx context.Context, userID uuid.UUID) ([]string, exceptions.HttpError)
// }

// type roleService struct {
// 	roleRepo repositories.RoleRepository
// }

// func NewRoleService(roleRepo repositories.RoleRepository) RoleService {
// 	return &roleService{roleRepo}
// }

// func (s *roleService) CreateRole(ctx context.Context, name string) exceptions.HttpError {
// 	return s.roleRepo.InsertRole(ctx, name)
// }

// func (s *roleService) AssignRole(ctx context.Context, userID uuid.UUID, roleName string) exceptions.HttpError {
// 	roleID, err := s.roleRepo.FindRoleIDByName(ctx, roleName)
// 	if err.Code != 0 {
// 		return err
// 	}
// 	return s.roleRepo.AssignUserRole(ctx, userID, roleID)
// }

// func (s *roleService) AddPermissionToRole(ctx context.Context, roleName, permissionName string) exceptions.HttpError {
// 	roleID, err := s.repo.FindRoleIDByName(ctx, roleName)
// 	if err != nil {
// 		return err
// 	}

// 	permID, err := s.repo.FindPermissionIDByName(ctx, permissionName)
// 	if err != nil {
// 		return err
// 	}

// 	return s.repo.AssignRolePermission(ctx, roleID, permID)
// }

// func (s *roleService) CheckPermission(ctx context.Context, userID uuid.UUID, permission string) (bool, exceptions.HttpError) {
// 	return s.repo.HasUserPermission(ctx, userID, permission)
// }

// func (s *roleService) GetUserRoles(ctx context.Context, userID uuid.UUID) ([]string, exceptions.HttpError) {
// 	return s.repo.GetUserRoleNames(ctx, userID)
// }
