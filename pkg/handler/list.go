package handler

import (
	"eduProject"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetAllListsResponse struct {
	Data []eduProject.TodoList `json:"data"`
}

// @Summary      Создать список задач
// @Description  Добавляет новый список задач в систему
// @Tags         lists
// @Accept       json
// @Produce      json
// @Param        input  body  eduProject.TodoList  true  "Данные списка задач"
// @Success      201  {object}  map[string]int  "ID созданного списка"
// @Failure      400  {object}  errorResponse  "Неправильный запрос"
// @Failure      500  {object}  errorResponse  "Ошибка сервера"
// @Router       /api/lists [post]
func (h *Handler) createList(c *gin.Context) {
	userId, err := ConvertUserId(c)
	if err != nil {
		return
	}

	var input eduProject.TodoList
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary      Получить все списки задач
// @Description  Возвращает все списки задач пользователя
// @Tags         lists
// @Produce      json
// @Success      200  {object}  GetAllListsResponse  "Список задач"
// @Failure      500  {object}  errorResponse  "Ошибка сервера"
// @Router       /api/lists [get]
func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := ConvertUserId(c)
	if err != nil {
		return
	}
	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, GetAllListsResponse{
		Data: lists,
	})
}

// @Summary      Получить список задач по ID
// @Description  Возвращает список задач по его ID
// @Tags         lists
// @Produce      json
// @Param        id  path  int  true  "ID списка"
// @Success      200  {object}  eduProject.TodoList  "Список задач"
// @Failure      400  {object}  errorResponse  "Неверный ID"
// @Failure      500  {object}  errorResponse  "Ошибка сервера"
// @Router       /api/lists/{id} [get]
func (h *Handler) getListById(c *gin.Context) {
	userId, err := ConvertUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}
	list, err := h.services.TodoList.GetById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}

// @Summary      Обновить список задач
// @Description  Изменяет название или описание списка задач
// @Tags         lists
// @Accept       json
// @Produce      json
// @Param        id     path   int                      true  "ID списка"
// @Param        input  body   eduProject.UpdateListInput  true  "Обновленные данные"
// @Success      200    {object}  statusResponse  "Успешное обновление"
// @Failure      400    {object}  errorResponse  "Неверный ID или данные"
// @Failure      500    {object}  errorResponse  "Ошибка сервера"
// @Router       /api/lists/{id} [put]
func (h *Handler) updateList(c *gin.Context) {
	userId, err := ConvertUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}
	var input eduProject.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"success update list"})
}

// @Summary      Удалить список задач
// @Description  Удаляет список задач по ID
// @Tags         lists
// @Param        id  path  int  true  "ID списка"
// @Success      200  {object}  statusResponse  "Успешное удаление"
// @Failure      400  {object}  errorResponse  "Неверный ID"
// @Failure      500  {object}  errorResponse  "Ошибка сервера"
// @Router       /api/lists/{id} [delete]
func (h *Handler) deleteList(c *gin.Context) {
	userId, err := ConvertUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}
	err = h.services.TodoList.Delete(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
