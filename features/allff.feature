@alluserflow
Feature: User Registration and Login

  Background: 
    Given User is on the homepage

  @registration
  Scenario: Create New User account
    When User follows the "Register" page
    And User click "create an account" button
    And User enters the following details
      | Username   |   meku1       |
      | First Name | Meku         |
      | Last Name  | Legese       |
      | Password   | 123456       |
      |Email       | meku@meku.com|
    And User click "register" button
    Then User should see "Welcome" message

  @registration
  Scenario: User does not follow form validations
    When user enters wrong characters
    Then error message displayed with invalid data
    And user returns back on registration page

  @login
  Scenario: Login Functionality
    When User visits "login" page
    And User enter "Meku" in the "user name" field
    And User enter "123456" in the "password" field
    And User press the "login" button
    Then User should see the "welcome" page

  @login
  Scenario: Verification of Login Function
    When User visits "login" page
    And user enters "email address" with "meku@gmail.com"
    And user enters "password" with "123456"
    And user click "log in" button
    Then user should see "My Account"

  @login
  Scenario: Unsuccessful login
    When User visits "login" page
    And User enter "Meku" in the "user name" field
    And User enter "121212" in the "password" field
    And user clicks "login" button
    Then error message displayed with wrong password
    And user returns back on login page
