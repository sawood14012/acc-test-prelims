Feature: Test CRDA auth command
    Scenario Outline: crda auth is running without snyk token
        Given system is running
        When I run crda auth without snyk token
        Then it should throw error
