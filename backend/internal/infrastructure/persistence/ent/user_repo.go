package entrepo

import (
	"cafe-pos/internal/db/models"
	"cafe-pos/internal/db/models/user"
	"cafe-pos/internal/domain/errors"
	"cafe-pos/internal/domain/repository"
	"context"

	"github.com/google/uuid"
)

// struct layer
type userRepository struct {
	client *models.Client
}

func NewUserRepository(client *models.Client) repository.UserRepository {
	return &userRepository{
		client: client,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, domainUser *repository.User) error {
	modelUser := toModelUser(domainUser)
	_, err := r.client.User.Create().
		SetID(modelUser.ID).
		SetCode(modelUser.Code).
		SetFirstName(modelUser.FirstName).
		SetLastName(modelUser.LastName).
		SetUsername(modelUser.Username).
		SetEmail(modelUser.Email).
		SetPassword(modelUser.Password).
		SetRoleID(modelUser.RoleID).
		SetStatus(modelUser.Status).
		Save(ctx)
	
	if err != nil {
		return convertToDomainError(err, "User")
	}
	
	return nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*repository.User, error) {
	modelUser, err := r.client.User.Query().Where(user.ID(id)).Only(ctx)
	if err != nil {
		return nil, convertToDomainError(err, "User")
	}
	return toDomainUser(modelUser), nil
}

func (r *userRepository) GetUser(ctx context.Context, filter *repository.UserFilter) (*repository.User, error) {
	modelFilter := toModelFilter(filter)
	query := applyUserFilter(r.client.User.Query(), modelFilter)
	modelUser, err := query.Only(ctx)
	if err != nil {
		// Return nil for not found, don't convert to error (used for checking existence)
		if models.IsNotFound(err) {
			return nil, nil
		}
		return nil, convertToDomainError(err, "User")
	}
	return toDomainUser(modelUser), nil
}

func (r *userRepository) GetUsers(ctx context.Context, filter *repository.UserFilter) ([]*repository.User, error) {
	modelFilter := toModelFilter(filter)
	query := applyUserFilter(r.client.User.Query(), modelFilter)
	modelUsers, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	return toDomainUsers(modelUsers), nil
}

func (r *userRepository) UpdateUser(ctx context.Context, domainUser *repository.User) error {
	modelUser := toModelUser(domainUser)
	updater := r.client.User.UpdateOneID(modelUser.ID).
		SetCode(modelUser.Code).
		SetFirstName(modelUser.FirstName).
		SetLastName(modelUser.LastName).
		SetUsername(modelUser.Username).
		SetEmail(modelUser.Email).
		SetPassword(modelUser.Password).
		SetRoleID(modelUser.RoleID).
		SetStatus(modelUser.Status)

	_, err := updater.Save(ctx)
	return err
}

func (r *userRepository) UpdateUserWithFields(ctx context.Context, domainUser *repository.User, fields []string, omitFields []string) error {
	modelUser := toModelUser(domainUser)
	updater := r.client.User.UpdateOneID(modelUser.ID)
	updater = applyUserUpdates(updater, modelUser, fields, omitFields)
	_, err := updater.Save(ctx)
	return err
}

func (r *userRepository) UpdateUsers(ctx context.Context, domainUsers []*repository.User) error {
	for _, u := range domainUsers {
		if err := r.UpdateUser(ctx, u); err != nil {
			return err
		}
	}
	return nil
}

func (r *userRepository) UpdateUsersWithFields(ctx context.Context, domainUsers []*repository.User, fields []string, omitFields []string) error {
	for _, u := range domainUsers {
		if err := r.UpdateUserWithFields(ctx, u, fields, omitFields); err != nil {
			return err
		}
	}
	return nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	err := r.client.User.DeleteOneID(id).Exec(ctx)
	return err
}

// toDomainUser converts infrastructure model to domain model
func toDomainUser(modelUser *models.User) *repository.User {
	if modelUser == nil {
		return nil
	}
	return &repository.User{
		ID:        modelUser.ID,
		Code:      modelUser.Code,
		FirstName: modelUser.FirstName,
		LastName:  modelUser.LastName,
		Username:  modelUser.Username,
		Email:     modelUser.Email,
		Password:  modelUser.Password,
		RoleID:    modelUser.RoleID,
		Status:    modelUser.Status,
		CreatedAt: modelUser.CreatedAt,
		UpdatedAt: modelUser.UpdatedAt,
	}
}

// toDomainUsers converts slice of infrastructure models to domain models
func toDomainUsers(modelUsers []*models.User) []*repository.User {
	if modelUsers == nil {
		return nil
	}
	domainUsers := make([]*repository.User, len(modelUsers))
	for i, u := range modelUsers {
		domainUsers[i] = toDomainUser(u)
	}
	return domainUsers
}

// toModelUser converts domain model to infrastructure model
func toModelUser(domainUser *repository.User) *models.User {
	if domainUser == nil {
		return nil
	}
	return &models.User{
		ID:        domainUser.ID,
		Code:      domainUser.Code,
		FirstName: domainUser.FirstName,
		LastName:  domainUser.LastName,
		Username:  domainUser.Username,
		Email:     domainUser.Email,
		Password:  domainUser.Password,
		RoleID:    domainUser.RoleID,
		Status:    domainUser.Status,
		CreatedAt: domainUser.CreatedAt,
		UpdatedAt: domainUser.UpdatedAt,
	}
}

// toModelFilter converts domain filter to infrastructure filter
func toModelFilter(domainFilter *repository.UserFilter) *models.UserFilter {
	if domainFilter == nil {
		return nil
	}
	return &models.UserFilter{
		ID:       domainFilter.ID,
		Code:     domainFilter.Code,
		Username: domainFilter.Username,
		Email:    domainFilter.Email,
		Status:   domainFilter.Status,
		RoleID:   domainFilter.RoleID,
	}
}

// applyUserFilter applies filter conditions to the query
func applyUserFilter(query *models.UserQuery, filter *models.UserFilter) *models.UserQuery {
	if filter == nil {
		return query
	}

	if filter.ID != nil {
		query = query.Where(user.ID(*filter.ID))
	}
	if filter.Code != nil {
		query = query.Where(user.Code(*filter.Code))
	}
	if filter.Username != nil {
		query = query.Where(user.Username(*filter.Username))
	}
	if filter.Email != nil {
		query = query.Where(user.Email(*filter.Email))
	}
	if filter.Status != nil {
		query = query.Where(user.Status(*filter.Status))
	}
	if filter.RoleID != nil {
		query = query.Where(user.RoleID(*filter.RoleID))
	}

	return query
}

// applyUserUpdates applies field updates to the updater based on fields and omitFields
func applyUserUpdates(updater *models.UserUpdateOne, u *models.User, fields []string, omitFields []string) *models.UserUpdateOne {
	omitMap := make(map[string]bool, len(omitFields))
	for _, field := range omitFields {
		omitMap[field] = true
	}

	allowedFields := make(map[string]bool, len(fields))
	for _, field := range fields {
		allowedFields[field] = true
	}

	shouldUpdate := func(fieldName string) bool {
		return (len(fields) == 0 || allowedFields[fieldName]) && !omitMap[fieldName]
	}

	if shouldUpdate("code") {
		updater = updater.SetCode(u.Code)
	}
	if shouldUpdate("first_name") {
		updater = updater.SetFirstName(u.FirstName)
	}
	if shouldUpdate("last_name") {
		updater = updater.SetLastName(u.LastName)
	}
	if shouldUpdate("username") {
		updater = updater.SetUsername(u.Username)
	}
	if shouldUpdate("email") {
		updater = updater.SetEmail(u.Email)
	}
	if shouldUpdate("password") {
		updater = updater.SetPassword(u.Password)
	}
	if shouldUpdate("role_id") {
		updater = updater.SetRoleID(u.RoleID)
	}
	if shouldUpdate("status") {
		updater = updater.SetStatus(u.Status)
	}

	return updater
}

// convertToDomainError converts infrastructure errors to domain errors
func convertToDomainError(err error, resource string) error {
	if err == nil {
		return nil
	}

	// Handle ent errors
	if models.IsNotFound(err) {
		return errors.NewNotFound(resource, "").WithError(err)
	}

	if models.IsConstraintError(err) {
		// Try to extract field information from constraint error
		return errors.NewConflict(resource, "", "").WithError(err)
	}

	if models.IsValidationError(err) {
		return errors.NewValidationFailed(err.Error()).WithError(err)
	}

	// Default to internal error for unknown errors
	return errors.NewInternal("database operation failed").WithError(err)
}
