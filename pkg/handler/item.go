package handler

import (
	"eduProject"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary      Создать задачу
// @Description  Добавляет новую задачу в список задач
// @Tags         items
// @Accept       json
// @Produce      json
// @Param        id     path      int                 true  "ID списка задач"
// @Param        input  body      eduProject.TodoItem true  "Данные задачи"
// @Success      201    {object}  map[string]int      "ID созданной задачи"
// @Failure      400    {object}  errorResponse       "Неправильный запрос"
// @Failure      500    {object}  errorResponse       "Ошибка сервера"
// @Router       /api/lists/{id}/items [post]
func (h *Handler) createItem(c *gin.Context) {
	userId, err := ConvertUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}
	var input eduProject.TodoItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.TodoItem.Create(userId, listId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary      Получить все задачи
// @Description  Возвращает список всех задач в указанном списке
// @Tags         items
// @Accept       json
// @Produce      json
// @Param        id   path      int                    true  "ID списка задач"
// @Success      200  {array}   eduProject.TodoItem    "Список задач"
// @Failure      400  {object}  errorResponse          "Неправильный запрос"
// @Failure      500  {object}  errorResponse          "Ошибка сервера"
// @Router       /api/lists/{id}/items [get]
func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := ConvertUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id parameter")
		return
	}

	items, err := h.services.TodoItem.GetAll(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

// @Summary      Получить задачу по ID
// @Description  Возвращает одну задачу по её ID
// @Tags         items
// @Accept       json
// @Produce      json
// @Param        id   path      int                   true  "ID задачи"
// @Success      200  {object}  eduProject.TodoItem   "Задача"
// @Failure      400  {object}  errorResponse         "Неправильный запрос"
// @Failure      500  {object}  errorResponse         "Ошибка сервера"
// @Router       /api/items/{id} [get]
func (h *Handler) getItemById(c *gin.Context) {
	userId, err := ConvertUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id parameter")
		return
	}

	item, err := h.services.TodoItem.GetItemById(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

// @Summary      Обновить задачу
// @Description  Обновляет данные задачи по её ID
// @Tags         items
// @Accept       json
// @Produce      json
// @Param        id     path      int                        true  "ID задачи"
// @Param        input  body      eduProject.UpdateItemInput true  "Обновленные данные"
// @Success      200    {object}  statusResponse             "Статус обновления"
// @Failure      400    {object}  errorResponse              "Неправильный запрос"
// @Failure      500    {object}  errorResponse              "Ошибка сервера"
// @Router       /api/items/{id} [put]
func (h *Handler) updateItem(c *gin.Context) {
	userId, err := ConvertUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}
	var input eduProject.UpdateItemInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.UpdateItem(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"success update item"})
}

// @Summary      Удалить задачу
// @Description  Удаляет задачу по её ID
// @Tags         items
// @Accept       json
// @Produce      json
// @Param        id   path      int              true  "ID задачи"
// @Success      200  {object}  statusResponse   "Статус удаления"
// @Failure      400  {object}  errorResponse    "Неправильный запрос"
// @Failure      500  {object}  errorResponse    "Ошибка сервера"
// @Router       /api/items/{id} [delete]
func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := ConvertUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id parameter")
		return
	}

	err = h.services.TodoItem.DeleteItem(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"item deleted"})
}
