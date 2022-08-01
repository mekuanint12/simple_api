@alluserflow
Feature: User Registration

  Background: 
    Given User is on "/register" page

  @registration
  Scenario: Creating New User Account
    When User enters the following details
      | Username   |   meku1      |
      | First Name | Meku         |
      | Last Name  | Legese       |
      | Password   | 123456       |
      |Email       | meku@meku.com|
    And User sends "register" POST requiest
    Then User should see "Welcome" message

  @registration
  Scenario: User Does NOT Follow the form validations
    When User enters wrong characters
      | Username   |   meku1      |
      | First Name | Meku         |
      | Last Name  | Legese       |
      | Password   | 123456       |
      | Email      | meku.com     |
    Then Error message displayed with invalid data