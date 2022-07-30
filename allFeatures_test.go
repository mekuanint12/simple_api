package main

import (
	"github.com/cucumber/godog"
)

func errorMessageDisplayedWithInvalidData() error {
	// config, err := util.LoadConfig(".")
	// if err != nil {
	// 	log.Fatal("Cant load config file: ", err)
	// }
	// con, err := sql.Open(config.DBDriver, config.DBSource)
	// if err != nil {
	// 	log.Fatal("Can not connect to the database:", err)
	// }
	// create_account := db.NewUser(con)

	// server := api.NewServer(create_account)
	return godog.ErrPending

}

func errorMessageDisplayedWithWrongPassword() error {
	return godog.ErrPending
}

func userClickButton(arg1 string) error {
	return godog.ErrPending
}

func userClicksButton(arg1 string) error {
	return godog.ErrPending
}

func userEnterInTheField(arg1, arg2 string) error {
	return godog.ErrPending
}

func userEntersTheFollowingDetails(arg1 *godog.Table) error {
	return godog.ErrPending
}

func userEntersWith(arg1, arg2 string) error {
	return godog.ErrPending
}

func userEntersWrongCharacters() error {
	return godog.ErrPending
}

func userFollowsThePage(arg1 string) error {
	return godog.ErrPending
}

func userIsOnTheHomepage() error {
	return godog.ErrPending
}

func userPressTheButton(arg1 string) error {
	return godog.ErrPending
}

func userReturnsBackOnLoginPage() error {
	return godog.ErrPending
}

func userReturnsBackOnRegistrationPage() error {
	return godog.ErrPending
}

func userShouldSee(arg1 string) error {
	return godog.ErrPending
}

func userShouldSeeMessage(arg1 string) error {
	return godog.ErrPending
}

func userShouldSeeThePage(arg1 string) error {
	return godog.ErrPending
}

func userVisitsPage(arg1 string) error {
	return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^error message displayed with invalid data$`, errorMessageDisplayedWithInvalidData)
	ctx.Step(`^error message displayed with wrong password$`, errorMessageDisplayedWithWrongPassword)
	ctx.Step(`^User click "([^"]*)" button$`, userClickButton)
	ctx.Step(`^user click "([^"]*)" button$`, userClickButton)
	ctx.Step(`^user clicks "([^"]*)" button$`, userClicksButton)
	ctx.Step(`^User enter "([^"]*)" in the "([^"]*)" field$`, userEnterInTheField)
	ctx.Step(`^User enters the following details$`, userEntersTheFollowingDetails)
	ctx.Step(`^user enters "([^"]*)" with "([^"]*)"$`, userEntersWith)
	ctx.Step(`^user enters wrong characters$`, userEntersWrongCharacters)
	ctx.Step(`^User follows the "([^"]*)" page$`, userFollowsThePage)
	ctx.Step(`^User is on the homepage$`, userIsOnTheHomepage)
	ctx.Step(`^User press the "([^"]*)" button$`, userPressTheButton)
	ctx.Step(`^user returns back on login page$`, userReturnsBackOnLoginPage)
	ctx.Step(`^user returns back on registration page$`, userReturnsBackOnRegistrationPage)
	ctx.Step(`^user should see "([^"]*)"$`, userShouldSee)
	ctx.Step(`^User should see "([^"]*)" message$`, userShouldSeeMessage)
	ctx.Step(`^User should see the "([^"]*)" page$`, userShouldSeeThePage)
	ctx.Step(`^User visits "([^"]*)" page$`, userVisitsPage)
}
