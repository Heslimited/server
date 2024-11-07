package handlers

import (
	"context"
	"project/internal/taskService" // Убедитесь, что путь к пакету правильный
	"project/internal/web/tasks"
)

type Handler struct {
	Service *taskService.TaskService
}

// NewHandler создает структуру Handler
func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

// Реализация методов интерфейса StrictServerInterface

// GetTasks implements tasks.StrictServerInterface.
func (h *Handler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	// Получение всех задач из сервиса
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	// Создаем переменную response типа GetTasks200JSONResponse
	response := tasks.GetTasks200JSONResponse{}

	// Заполняем слайс response всеми задачами из БД
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Text,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	// Возвращаем response и nil (без ошибки)
	return response, nil
}

// PostTasks implements tasks.StrictServerInterface.
func (h *Handler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	// Распаковываем тело запроса
	taskRequest := request.Body
	// Обращаемся к сервису и создаем задачу
	taskToCreate := taskService.Task{
		Text:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру response
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Text,
		IsDone: &createdTask.IsDone,
	}
	// Возвращаем response
	return response, nil
}

// PostTasks implements tasks.StrictServerInterface.
func (h *Handler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	// Получаем ID задачи из параметров пути
	taskID := uint(request.Id) // Преобразуем тип int в uint. Переменная taskID имеет тип int, а метод UpdateTaskByID ожидает параметр типа uint
	// Распаковываем тело запроса
	taskRequest := request.Body
	// Обращаемся к сервису и создаем задачу
	taskToUpdate := taskService.Task{
		Text:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}

	// Обновляем задачу через сервис
	updatedTask, err := h.Service.UpdateTaskByID(taskID, taskToUpdate)
	if err != nil {
		return nil, err
	}

	// Создаем структуру response(ответа)
	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Text,
		IsDone: &updatedTask.IsDone,
	}
	// Возвращаем response
	return response, nil
}

func (h *Handler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	// Получаем ID задачи из параметров пути
	taskID := uint(request.Id) // Преобразуем тип int в uint

	// Удаляем задачу через сервис
	err := h.Service.DeleteTaskByID(taskID)
	if err != nil {
		return nil, err
	}

	// Создаем ответ с успешным статусом (без поля Message)
	response := tasks.DeleteTasksId200JSONResponse{
		Id: &taskID,
	}

	// Возвращаем response
	return response, nil
}

// func (h *Handler) PatchTaskHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPatch {
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	vars := mux.Vars(r)
// 	id := vars["id"]
// 	var task taskService.Task

// 	err := json.NewDecoder(r.Body).Decode(&task)
// 	if err != nil {
// 		http.Error(w, "Bad request", http.StatusBadRequest)
// 		return
// 	}

// 	taskID, err := strconv.ParseUint(id, 10, 64)
// 	if err != nil {
// 		http.Error(w, "Invalid ID", http.StatusBadRequest)
// 		return
// 	}

// 	updatedTask, err := h.Service.UpdateTaskByID(uint(taskID), task)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(updatedTask)
// }

// func (h *Handler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	taskID, err := strconv.ParseUint(id, 10, 64)
// 	if err != nil {
// 		http.Error(w, "Invalid ID", http.StatusBadRequest)
// 		return
// 	}

// 	err = h.Service.DeleteTaskByID(uint(taskID))
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, "Task deleted successfully")
// }
