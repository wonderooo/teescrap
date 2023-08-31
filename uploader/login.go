package uploader

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/joho/godotenv"
)

const (
	XPATH_LOGIN_A_TAG        = `/html/body/header/div/nav/div[3]/div/div[1]/div[1]/a`
	XPATH_LOGIN_USER_INP_TAG = `/html/body/div[5]/div[2]/div/div[2]/div[3]/form/div[1]/input`
	XPATH_LOGIN_PASS_INP_TAG = `/html/body/div[5]/div[2]/div/div[2]/div[3]/form/div[2]/input`
	XPATH_LOGIN_SUBM_INP_TAG = `/html/body/div[5]/div[2]/div/div[2]/div[3]/form/input[3]`
)

type LoginAction chromedp.Action

func login(user string, password string) LoginAction {

	return chromedp.ActionFunc(func(ctx context.Context) error {
		log.Printf("Received credentials %s:%s", user, password)

		var err error
		err = toModal().Do(ctx)
		if err != nil {
			log.Fatalf("Could not go to login modal: %s", err)
		}

		err = fillLoginForm(user, password).Do(ctx)
		if err != nil {
			log.Fatalf("Could not fill login modal: %s", err)
		}

		err = submitForm().Do(ctx)
		if err != nil {
			log.Fatalf("Could not submit login modal: %s", err)
		}

		return err
	})

}

func toModal() chromedp.Action {
	return chromedp.Click(XPATH_LOGIN_A_TAG, chromedp.NodeVisible)
}

func fillLoginForm(user string, password string) chromedp.Action {
	return chromedp.Tasks{
		chromedp.SendKeys(XPATH_LOGIN_USER_INP_TAG, user, chromedp.NodeVisible),
		chromedp.SendKeys(XPATH_LOGIN_PASS_INP_TAG, password, chromedp.NodeVisible),
	}
}

func submitForm() chromedp.Action {
	dur, _ := time.ParseDuration("1s")
	return chromedp.Tasks{
		chromedp.Sleep(dur),
		chromedp.Click(XPATH_LOGIN_SUBM_INP_TAG, chromedp.NodeEnabled),
	}
}

func credentials(prod bool) (string, string) {
	err := godotenv.Load(".env")
	if err != nil {
		ferr := fmt.Sprintf("Could not load .env file: %s", err)
		panic(ferr)
	}

	if prod == true {
		username := os.Getenv("TEE_USERNAME_PROD")
		password := os.Getenv("TEE_PASSWORD_PROD")

		return username, password
	} else {
		username := os.Getenv("TEE_USERNAME_TEST")
		password := os.Getenv("TEE_PASSWORD_TEST")

		return username, password
	}
}
