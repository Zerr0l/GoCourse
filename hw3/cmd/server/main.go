package main

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"hw3/accounts/models"
	"hw3/proto"
	"log"
	"net"
	"sync"
)

func NewBankAccountManagerServer() *Server {
	return &Server{
		accounts: make(map[string]*models.Account),
		guard:    &sync.RWMutex{},
	}
}

type Server struct {
	proto.UnimplementedBankAccountManagerServer
	accounts map[string]*models.Account
	guard    *sync.RWMutex
}

func (s *Server) CreateAccount(ctx context.Context, request *proto.CreateAccountRequest) (*proto.CreateAccountResponse, error) {
	if len(request.Name) == 0 {
		return nil, errors.New("empty name")
	}

	s.guard.Lock()

	if _, ok := s.accounts[request.Name]; ok {
		s.guard.Unlock()

		return nil, errors.New("account already exists")
	}

	s.accounts[request.Name] = &models.Account{
		Name:   request.Name,
		Amount: request.Amount,
	}

	s.guard.Unlock()

	response := proto.CreateAccountResponse{Result: "account created"}
	return &response, nil
}

func (s *Server) GetAccount(ctx context.Context, request *proto.GetAccountRequest) (*proto.GetAccountResponse, error) {
	name := request.Name // {"name": "alice"}
	if len(name) == 0 {
		return nil, errors.New("empty name")
	}

	s.guard.RLock()
	account, ok := s.accounts[name]
	s.guard.RUnlock()

	if !ok {
		return nil, errors.New("account not found")
	}

	response := proto.GetAccountResponse{Name: account.Name, Amount: account.Amount}

	return &response, nil
}

// Удаляет аккаунт
func (s *Server) DeleteAccount(ctx context.Context, request *proto.DeleteAccountRequest) (*proto.DeleteAccountResponse, error) {
	if len(request.Name) == 0 {
		return nil, errors.New("empty name")
	}

	s.guard.Lock()

	if _, ok := s.accounts[request.Name]; !ok {
		s.guard.Unlock()

		return nil, errors.New("account does not exist")
	}

	delete(s.accounts, request.Name)

	s.guard.Unlock()

	response := proto.DeleteAccountResponse{Result: "account deleted"}

	return &response, nil
}

// Меняет баланс
func (s *Server) ChangeAccountAmount(ctx context.Context, request *proto.ChangeAccountAmountRequest) (*proto.ChangeAccountAmountResponse, error) {
	if len(request.Name) == 0 {
		return nil, errors.New("empty name")
	}

	s.guard.Lock()

	if _, ok := s.accounts[request.Name]; !ok {
		s.guard.Unlock()

		return nil, errors.New("account does not exist")
	}

	s.accounts[request.Name].Amount += request.Amount

	s.guard.Unlock()

	response := proto.ChangeAccountAmountResponse{Result: "account amount changed"}
	return &response, nil
}

// Меняет имя
func (s *Server) ChangeAccountName(ctx context.Context, request *proto.ChangeAccountNameRequest) (*proto.ChangeAccountNameResponse, error) {
	if len(request.Name) == 0 {
		return nil, errors.New("empty name")
	}
	if len(request.NewName) == 0 {
		return nil, errors.New("empty new name")
	}

	s.guard.Lock()

	account, ok := s.accounts[request.Name]
	if !ok {
		return nil, errors.New("account not found")
	}

	delete(s.accounts, request.Name)

	if _, ok := s.accounts[request.NewName]; ok {
		s.accounts[request.Name] = account
		s.guard.Unlock()

		return nil, errors.New("new name already exists")
	}

	s.accounts[request.NewName] = &models.Account{
		Name:   request.NewName,
		Amount: account.Amount,
	}

	s.guard.Unlock()

	response := proto.ChangeAccountNameResponse{Result: "account name changed"}
	return &response, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterBankAccountManagerServer(s, NewBankAccountManagerServer())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
