@alluserflow
Feature: User Login

  Background: 
    Given User is on "/login" page

  @login
  Scenario: Login Functionality
    When User enters the following details
      | Username   |   meku1      |
      | Password   | 123456       |
    And User sends a "login" requiest
    Then User should see the "welcome" page

  @login
  Scenario: Unsuccessful login
    When User enters the following details
      | Username   |   meku122      |
      | Password   |  123456        |
    And User sends "/login" GET requiest
    Then Error message displayed with the with invalid data
