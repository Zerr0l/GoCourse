package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"hw3/proto"
	"log"
	"time"
)

type Command struct {
	Port    int
	Host    string
	Cmd     string
	Name    string
	NewName string
	Amount  int64
}

func (c *Command) Do() error {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", c.Host, c.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("new grpc client failed: %v", err)
	}

	defer func() {
		_ = conn.Close()
	}()

	switch c.Cmd {
	case "create":
		return c.create(conn)
	case "get":
		return c.get(conn)
	case "delete":
		return c.delete(conn)
	case "change_name":
		return c.changeName(conn)
	case "change_amount":
		return c.changeAmount(conn)
	default:
		return fmt.Errorf("unknown command: %s", c.Cmd)
	}
}

func (cmd *Command) create(conn *grpc.ClientConn) error {
	client := proto.NewBankAccountManagerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := client.CreateAccount(ctx, &proto.CreateAccountRequest{Name: cmd.Name, Amount: cmd.Amount})
	if err != nil {
		return err
	}

	return nil
}

func (cmd *Command) get(conn *grpc.ClientConn) error {
	client := proto.NewBankAccountManagerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	respond, err := client.GetAccount(ctx, &proto.GetAccountRequest{Name: cmd.Name})
	if err != nil {
		return err
	}

	fmt.Printf("Name: %s, Amount: %d\n", respond.Name, respond.Amount)

	return nil
}

func (cmd *Command) delete(conn *grpc.ClientConn) error {
	client := proto.NewBankAccountManagerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := client.DeleteAccount(ctx, &proto.DeleteAccountRequest{Name: cmd.Name})
	if err != nil {
		return err
	}

	return nil
}

func (cmd *Command) changeName(conn *grpc.ClientConn) error {
	client := proto.NewBankAccountManagerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := client.ChangeAccountName(ctx, &proto.ChangeAccountNameRequest{Name: cmd.Name, NewName: cmd.NewName})
	if err != nil {
		return err
	}

	return nil
}

func (cmd *Command) changeAmount(conn *grpc.ClientConn) error {
	client := proto.NewBankAccountManagerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := client.ChangeAccountAmount(ctx, &proto.ChangeAccountAmountRequest{Name: cmd.Name, Amount: cmd.Amount})
	if err != nil {
		return err
	}

	return nil
}

func main() {
	portVal := flag.Int("port", 8080, "server port")
	hostVal := flag.String("host", "0.0.0.0", "server host")
	cmdVal := flag.String("cmd", "", "command to execute")
	nameVal := flag.String("name", "", "name of account")
	newNameVal := flag.String("new_name", "", "new name of account")
	amountVal := flag.Int64("amount", 0, "amount of account")

	flag.Parse()

	cmd := Command{
		Port:    *portVal,
		Host:    *hostVal,
		Cmd:     *cmdVal,
		Name:    *nameVal,
		NewName: *newNameVal,
		Amount:  *amountVal,
	}

	if err := cmd.Do(); err != nil {
		log.Fatalf("command failed: %v", err)
	}
}
