package main

import (
	"context"
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	connectionString := "host=0.0.0.0 port=5432 dbname=postgres password=mysecretpassword"

	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot open database")
	}

	defer func() {
		_ = db.Close()
	}()
	if err := db.Ping(); err != nil {
		return c.String(http.StatusInternalServerError, "cannot connect to database")
	}

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

	ctx := context.Background()

	_, err = db.ExecContext(ctx, "INSERT INTO accounts(name, balance) VALUES ($1, $2)", request.Name, request.Amount)
	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot insert new account into database")
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
	connectionString := "host=0.0.0.0 port=5432 dbname=postgres password=mysecretpassword"

	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot open database")
	}

	defer func() {
		_ = db.Close()
	}()
	if err := db.Ping(); err != nil {
		return c.String(http.StatusInternalServerError, "cannot connect to database")
	}

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

	ctx := context.Background()

	_, err = db.ExecContext(ctx, "DELETE FROM accounts WHERE (name = $1)", request.Name)
	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot delete account from database")
	}

	delete(h.accounts, request.Name)

	h.guard.Unlock()

	return c.NoContent(http.StatusOK)
}

// Меняет баланс
func (h *Handler) ChangeAccountAmount(c echo.Context) error {
	connectionString := "host=0.0.0.0 port=5432 dbname=postgres password=mysecretpassword"

	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot open database")
	}

	defer func() {
		_ = db.Close()
	}()
	if err := db.Ping(); err != nil {
		return c.String(http.StatusInternalServerError, "cannot connect to database")
	}

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

	ctx := context.Background()

	_, err = db.ExecContext(ctx, "UPDATE accounts SET balance = $1 WHERE name = $2", h.accounts[request.Name].Amount+request.Amount, request.Name)
	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot update account amount")
	}

	h.accounts[request.Name].Amount += request.Amount

	h.guard.Unlock()

	return c.NoContent(http.StatusOK)
}

// Меняет имя
func (h *Handler) ChangeAccountName(c echo.Context) error {
	connectionString := "host=0.0.0.0 port=5432 dbname=postgres password=mysecretpassword"

	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot open database")
	}

	defer func() {
		_ = db.Close()
	}()

	if err := db.Ping(); err != nil {
		return c.String(http.StatusInternalServerError, "cannot connect to database")
	}

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

	ctx := context.Background()

	_, err = db.ExecContext(ctx, "UPDATE accounts SET name = $1 WHERE name = $2", request.NewName, request.Name)
	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot update account name")
	}

	h.accounts[request.NewName] = &models.Account{
		Name:   request.NewName,
		Amount: account.Amount,
	}

	h.guard.Unlock()

	return c.NoContent(http.StatusOK)
}

// Написать клиент консольный, который делает запросы

func main() {
	connectionString := "host=0.0.0.0 port=5432 dbname=postgres password=mysecretpassword"

	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = db.Close()
	}()

	if err := db.Ping(); err != nil {
		panic(err)
	}

	ctx := context.Background()

	rows, err := db.QueryContext(ctx, "SELECT name, balance FROM accounts")
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = rows.Close()
	}()

	accountsHandler := New()

	for rows.Next() {
		var account models.Account

		if err := rows.Scan(&account.Name, &account.Amount); err != nil {
			panic(err)
		}

		accountsHandler.accounts[account.Name] = &account
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/account", accountsHandler.GetAccount)
	e.POST("/account/create", accountsHandler.CreateAccount)
	e.POST("/account/delete", accountsHandler.DeleteAccount)
	e.POST("/account/change_name", accountsHandler.ChangeAccountName)
	e.POST("/account/change_amount", accountsHandler.ChangeAccountAmount)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
