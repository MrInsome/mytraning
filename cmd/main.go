package main

import (
	"github.com/labstack/echo/v4"
	"net/http"

	"github.com/labstack/echo/v4/middleware"

	"mytraning/internal"
	"mytraning/pkg"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/accounts", createAccount)
	e.GET("/accounts/:id", getAccount)
	e.DELETE("/accounts/:id", deleteAccount)

	e.Logger.Fatal(e.Start(":8080"))
}

func createAccount(c echo.Context) error {
	account := new(pkg.Account)
	if err := c.Bind(account); err != nil {
		return err
	}

	db := internal.NewDatabase()
	db.AddAccount(internal.Account{
		AccessToken:  account.AccessToken,
		RefreshToken: account.RefreshToken,
		ExpiresIn:    account.ExpiresIn,
		AccountID:    account.AccountID,
		Integration: internal.Integration{
			SecretKey:        account.Integration.SecretKey,
			ClientID:         account.Integration.ClientID,
			RedirectURL:      account.Integration.RedirectURL,
			AuthorizationURL: account.Integration.AuthorizationURL,
		},
	})

	return c.JSON(http.StatusCreated, account)
}

func getAccount(c echo.Context) error {
	accountID := c.Param("id")

	db := internal.NewDatabase()
	account, ok := db.GetAccount(accountID)
	if !ok {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, account)
}

func deleteAccount(c echo.Context) error {
	accountID := c.Param("id")

	db := internal.NewDatabase()
	db.RemoveAccount(accountID)

	return c.NoContent(http.StatusNoContent)
}
