Feature: Test CRDA auth command
    Scenario: crda auth is running without snyk token
        Given system is running
        When I run crda auth without snyk token
        Then it should throw error

    Scenario: crda auth is running with invalid snyk token
        Given system is running
        When I run crda auth with invalid snyk token
        Then it should throw error

    Scenario: crda auth is running with valid snyk token
        Given system is running
        When I run crda auth with valid snyk token
        Then it should not throw error