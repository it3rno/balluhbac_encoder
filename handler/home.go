package handler

type HomeHandler struct {}

func (h *HomeHandler) HandleHome(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}
