package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/userhubdev/go-sdk/adminapi"
	"github.com/userhubdev/go-sdk/userapi"
)

func run() error {
	ctx := context.Background()

	adminKey := os.Getenv("USERHUB_ADMIN_KEY")
	if adminKey == "" {
		return fmt.Errorf("USERHUB_ADMIN_KEY required")
	}

	userKey := os.Getenv("USERHUB_USER_KEY")
	if userKey == "" {
		return fmt.Errorf("USERHUB_USER_KEY required")
	}

	adminApi, err := adminapi.New(adminKey)
	if err != nil {
		return err
	}

	res, err := adminApi.Users().List(ctx, &adminapi.UserListInput{
		PageSize: 5,
	})
	if err != nil {
		return err
	}

	for _, user := range res.Users {
		if user.Disabled {
			continue
		}

		apiSession, err := adminApi.Users().CreateApiSession(ctx, user.Id, nil)
		if err != nil {
			return err
		}

		userApi, err := userapi.New(userKey, apiSession.AccessToken)
		if err != nil {
			return err
		}

		session, err := userApi.Session().Get(ctx, nil)
		if err != nil {
			return err
		}

		fmt.Println("User:")
		fmt.Println(" - ID:", session.User.Id)
		fmt.Println(" - Name:", session.User.DisplayName)
		fmt.Println(" - Memberships:", len(session.Memberships))

		break
	}

	_, err = adminApi.Users().Get(ctx, "fail", nil)
	if err != nil {
		fmt.Println()
		fmt.Println(err.Error())

		var e *adminapi.Error
		if errors.As(err, &e) {
			fmt.Println()
			fmt.Println("UserHub error:")
			fmt.Println(" - ApiCode:", e.ApiCode())
			fmt.Println(" - Message:", e.Message())
			fmt.Println(" - Reason:", e.Reason())
			fmt.Println(" - Param:", e.Param())
			fmt.Println(" - Call:", e.Call())
		}
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
