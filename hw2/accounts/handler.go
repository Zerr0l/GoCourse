package accounts

import (
	"github.com/labstack/echo/v4"
	"hw2/accounts/dto"
	"hw2/accounts/models"
	"net/http"
	"sync"
)

func New() *Handler {
	return &Handler{
		accounts: make(map[string]*models.Account),
		guard:    &sync.RWMutex{},
	}
}

type Handler struct {
	accounts map[string]*models.Account
	guard    *sync.RWMutex
}

func (h *Handler) CreateAccount(c echo.Context) error {
	var request dto.CreateAccountRequest // {"name": "alice", "amount": 50}
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}

	h.guard.Lock()

	if _, ok := h.accounts[request.Name]; ok {
		h.guard.Unlock()

		return c.String(http.StatusForbidden, "account already exists")
	}

	h.accounts[request.Name] = &models.Account{
		Name:   request.Name,
		Amount: request.Amount,
	}

	h.guard.Unlock()

	return c.NoContent(http.StatusCreated)
}

func (h *Handler) GetAccount(c echo.Context) error {
	name := c.QueryParams().Get("name") // {"name": "alice"}
	if len(name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}

	h.guard.RLock()

	account, ok := h.accounts[name]

	h.guard.RUnlock()

	if !ok {
		return c.String(http.StatusNotFound, "account not found")
	}

	response := dto.GetAccountResponse{
		Name:   account.Name,
		Amount: account.Amount,
	}

	return c.JSON(http.StatusOK, response)
}

// Удаляет аккаунт
func (h *Handler) DeleteAccount(c echo.Context) error {
	var request dto.DeleteAccountRequest // {"name": "alice"}
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}

	h.guard.Lock()

	if _, ok := h.accounts[request.Name]; !ok {
		h.guard.Unlock()

		return c.String(http.StatusNotFound, "account does not exist")
	}

	delete(h.accounts, request.Name)

	h.guard.Unlock()

	return c.NoContent(http.StatusOK)
}

// Меняет баланс
func (h *Handler) ChangeAccountAmount(c echo.Context) error {
	var request dto.ChangeAccountAmountRequest // {"name": "alice", "amount": "10"}
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}

	h.guard.Lock()

	if _, ok := h.accounts[request.Name]; !ok {
		h.guard.Unlock()

		return c.String(http.StatusNotFound, "account does not exist")
	}

	h.accounts[request.Name].Amount += request.Amount

	h.guard.Unlock()

	return c.NoContent(http.StatusOK)
}

// Меняет имя
func (h *Handler) ChangeAccountName(c echo.Context) error {
	var request dto.ChangeAccountNameRequest // {"name": "alice", "new_name": "aline"}
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}
	if len(request.NewName) == 0 {
		return c.String(http.StatusBadRequest, "empty new name")
	}

	h.guard.Lock()

	account, ok := h.accounts[request.Name]
	if !ok {
		return c.String(http.StatusNotFound, "account not found")
	}

	delete(h.accounts, request.Name)

	if _, ok := h.accounts[request.NewName]; ok {
		h.accounts[request.Name] = account
		h.guard.Unlock()

		return c.String(http.StatusBadRequest, "new name already exists")
	}

	h.accounts[request.NewName] = &models.Account{
		Name:   request.NewName,
		Amount: account.Amount,
	}

	h.guard.Unlock()

	return c.NoContent(http.StatusOK)
}

// Написать клиент консольный, который делает запросы
