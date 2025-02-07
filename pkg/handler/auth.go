package handler

import (
	"eduProject"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary      Регистрация
// @Description  Создает нового пользователя в системе
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body      eduProject.User   true  "Данные пользователя"
// @Success      201    {object}  map[string]int    "ID созданного пользователя"
// @Failure      400    {object}  errorResponse     "Неправильный запрос"
// @Failure      500    {object}  errorResponse     "Ошибка сервера"
// @Router       /auth/sign-up [post]
func (h *Handler) SignUp(c *gin.Context) {
	var input eduProject.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary      Авторизация
// @Description  Авторизует пользователя и возвращает токен
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body      signInInput       true  "Данные для входа"
// @Success      200    {object}  map[string]string "Токен авторизации"
// @Failure      400    {object}  errorResponse     "Неправильный запрос"
// @Failure      500    {object}  errorResponse     "Ошибка сервера"
// @Router       /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
