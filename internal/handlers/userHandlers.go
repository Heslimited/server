package handlers

import (
	"context"
	"project/internal/models"
	"project/internal/userService" // Убедитесь, что путь к пакету правильный
	"project/internal/web/users"
)

type UserHandler struct {
	Service *userService.UserService
}

// NewHandler создает структуру Handler
func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

// Реализация методов интерфейса StrictServerInterface

func (h *UserHandler) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}
	for _, usr := range allUsers {
		user := users.User{
			Id:       &usr.ID,
			Email:    &usr.Email,
			Password: &usr.Password,
		}
		response = append(response, user)
	}
	return response, nil
}

func (h *UserHandler) GetTasksByUserID(ctx context.Context, request users.GetTasksByUserIDRequestObject) (users.GetTasksByUserIDResponseObject, error) {
	userID := uint(request.Id)
	allTasks, err := h.Service.GetTasksForUser(userID)

	if err != nil {
		return nil, err
	}

	response := users.GetTasksByUserID200JSONResponse{}
	for _, tsk := range allTasks {
		task := users.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Text,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}
	return response, nil
}

func (h *UserHandler) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body
	userToCreate := models.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}
	createdUser, err := h.Service.CreateUser(userToCreate)
	if err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}
	return response, nil
}

func (h *UserHandler) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	userID := uint(request.Id)
	userRequest := request.Body

	userToUpdate := models.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}

	updatedUser, err := h.Service.UpdateUserByID(userID, userToUpdate)
	if err != nil {
		return nil, err
	}

	response := users.PatchUsersId200JSONResponse{
		Id:       &updatedUser.ID,
		Email:    &updatedUser.Email,
		Password: &updatedUser.Password,
	}
	return response, nil
}

func (h *UserHandler) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	userID := uint(request.Id)
	err := h.Service.DeleteUserByID(userID)
	if err != nil {
		return nil, err
	}

	response := users.DeleteUsersId200JSONResponse{
		Id: &userID,
	}
	return response, nil
}
