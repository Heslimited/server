package handlers

import (
	"context"
	"project/internal/models"
	"project/internal/taskService" // Убедитесь, что путь к пакету правильный
	"project/internal/web/tasks"
)

type TaskHandler struct {
	Service *taskService.TaskService
}

// NewHandler создает структуру Handler
func NewTaskHandler(service *taskService.TaskService) *TaskHandler {
	return &TaskHandler{
		Service: service,
	}
}

// Реализация методов интерфейса StrictServerInterface
func (h *TaskHandler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Text,
			IsDone: &tsk.IsDone,
		}

		// Добавляем UserId в ответ только если он существует
		if tsk.UserId != nil {
			task.UserId = tsk.UserId
		}

		response = append(response, task)
	}
	return response, nil
}

// PostTasks implements tasks.StrictServerInterface.
func (h *TaskHandler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body

	taskToCreate := models.Task{
		Text:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}

	// Записать UserId, если он существует в запросе
	if taskRequest.UserId != nil {
		taskToCreate.UserId = taskRequest.UserId
	}

	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Text,
		IsDone: &createdTask.IsDone,
	}

	// Добавляем UserId в ответ, только если он существует
	if createdTask.UserId != nil {
		response.UserId = createdTask.UserId
	}

	return response, nil
}

func (h *TaskHandler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	taskID := uint(request.Id)
	taskRequest := request.Body

	taskToUpdate := models.Task{
		Text:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}

	updatedTask, err := h.Service.UpdateTaskByID(taskID, taskToUpdate)
	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Text,
		IsDone: &updatedTask.IsDone,
	}
	return response, nil
}

func (h *TaskHandler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	taskID := uint(request.Id)
	err := h.Service.DeleteTaskByID(taskID)
	if err != nil {
		return nil, err
	}

	response := tasks.DeleteTasksId200JSONResponse{
		Id: &taskID,
	}
	return response, nil
}
