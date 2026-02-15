package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"vago/internal/application/task"
	"vago/internal/config/code"
	"vago/internal/transport/http/api/response"
	"vago/internal/transport/http/shared/template"

	"github.com/gin-gonic/gin"
)

func Tasks(service *task.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := template.TplWithMetaData(c, "Задачи пользователя")

		tasks, err := service.GetAllByUser(data[code.UserId].(int64))
		if err != nil {
			c.HTML(http.StatusInternalServerError, "error.html", gin.H{
				"Message": "Не удалось загрузить задачи",
			})
			return
		}

		data["Tasks"] = tasks
		c.HTML(http.StatusOK, "tasks.html", data)
	}
}

func TasksAPI(service *task.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		currentId, errGetUSerId := c.Get(code.UserId)
		if !errGetUSerId {
			response.Error(c, http.StatusUnauthorized, "Пользователь не аутентифицировался")
			return
		}

		tasks, _ := service.GetAllByUser(currentId.(int64))

		response.OK(c, "Задачи", tasksToDTO(tasks))
	}
}

func PostTaskAPI(service *task.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto PostTaskDTO
		if err := c.ShouldBindJSON(&dto); err != nil {
			response.Error(c, http.StatusBadRequest, "Некорректные данные")
			return
		}

		currentId, errGetUSerId := c.Get(code.UserId)
		if !errGetUSerId {
			response.Error(c, http.StatusUnauthorized, "Пользователь не аутентифицировался")
			return
		}

		err := service.PostTask(dto.Name, dto.Description, dto.Completed, currentId.(int64))
		if err != nil {
			response.Error(c, http.StatusInternalServerError, err.Error())
			return
		}

		response.OKNoData(c, "Задача создана")
	}
}

func PostTask(service *task.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.PostForm("name")
		desc := c.PostForm("description")
		completed := c.PostForm("completed")

		if name == "" {
			c.String(http.StatusBadRequest, "Название задачи обязательно")
			return
		}

		sessionUserID, ok := c.Get(code.UserId)
		if !ok {
			c.HTML(http.StatusInternalServerError, "error.html", gin.H{
				"Message": "Нет ключа в session",
				"Error":   fmt.Sprintf("Значение ключа: %v", code.UserId),
			})
		}

		err := service.PostTask(name, desc, completed == "on", sessionUserID.(int64))
		if err != nil {
			ShowError(c, "Ошибка создания задачи", err.Error())
			return
		}

		c.Redirect(http.StatusSeeOther, "/tasks")
	}
}

func DeleteTaskAPI(service *task.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		parseId, parseIdErr := strconv.ParseInt(c.Param("id"), 10, 64)
		if parseIdErr != nil {
			response.Error(c, http.StatusBadRequest, "Некорректные данные")
			return
		}

		err := service.DeleteTask(parseId)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, err.Error())
			return
		}

		response.OKNoData(c, "Задача удалена.")
	}
}

func DeleteTask(service *task.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		parseId, parseIdErr := strconv.ParseInt(id, 10, 64)
		if parseIdErr != nil {
			response.Error(c, http.StatusBadRequest, "Некорректные данные")
			return
		}

		err := service.DeleteTask(parseId)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error deleting task")
			return
		}

		c.Redirect(http.StatusSeeOther, "/tasks")
	}
}

func UpdateTaskAPI(service *task.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		parseId, parseIdErr := strconv.ParseInt(c.Param("id"), 10, 64)
		if parseIdErr != nil {
			response.Error(c, http.StatusBadRequest, "Некорректные данные")
			return
		}

		var dto UpdateTaskDTO
		if err := c.ShouldBindJSON(&dto); err != nil {
			response.Error(c, http.StatusBadRequest, "Некорректные данные")
			return
		}

		currentId, errGetUSerId := c.Get(code.UserId)
		if !errGetUSerId {
			response.Error(c, http.StatusUnauthorized, "Пользователь не аутентифицировался")
			return
		}

		errUpdate := service.UpdateCompleted(parseId, currentId.(int64), dto.Completed)
		if errUpdate != nil {
			response.Error(c, http.StatusInternalServerError, errUpdate.Error())
			return
		}

		response.OKNoData(c, "Успешное обновление задачи")

		/*userID := c.MustGet(code.UserId).(int64)


		if errUpdate != nil {
			appCtx.Log.Errorw("failed to update task", "error", errUpdate)
			c.JSON(http.StatusBadRequest, gin.H{"error": errUpdate.Error()})
			return
		}

		appCtx.Log.Debugw("Update task", "taskID", taskID, "userID", userID)
		c.Redirect(http.StatusSeeOther, "/tasks")*/
	}
}

func UpdateTask(service *task.Service) gin.HandlerFunc {
	type reqBody struct {
		Completed bool `json:"completed"`
	}

	return func(c *gin.Context) {
		taskID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}

		var body reqBody
		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
			return
		}

		userID := c.MustGet(code.UserId).(int64)

		errUpdate := service.UpdateCompleted(int64(taskID), userID, body.Completed)
		if errUpdate != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errUpdate.Error()})
			return
		}

		c.Redirect(http.StatusSeeOther, "/tasks")
	}
}
