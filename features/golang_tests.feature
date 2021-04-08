Feature: Test CRDA analyse Go
    Scenario: Check crda analyse go command with vulns
        Given system is running and authenticated with valid snyk token
        When I run crda analyse command for Go with vulns
        Then I should be able to get the absolute path
        And I should Copy Manifest
        And I should be able to copy the main file
        And I should able to run go mod tidy
        And I Should be able to run crda analyse without error
        And I should perform cleanup

    Scenario: Check crda analyse command for Go with vulns json
        Given system is running and authenticated with valid snyk token
        When I run crda analyse command for Go with vulns json
        Then I should be able to get the absolute path
        And I should Copy Manifest
        And I should be able to copy the main file
        And I should able to run go mod tidy
        And I Should be able to run crda analyse without error
        And I should perform cleanup

    Scenario: Check crda analyse command for Go with vulns and verbose
        Given system is running and authenticated with valid snyk token
        When I run crda analyse command for Go with vulns json
        Then I should be able to get the absolute path
        And I should Copy Manifest
        And I should be able to copy the main file
        And I should able to run go mod tidy
        And I Should be able to run crda analyse without error
        And I should perform cleanup

    Scenario: Check crda analyse command for Go with vulns and debug
        Given system is running and authenticated with valid snyk token
        When I run crda analyse command for Go with vulns json
        Then I should be able to get the absolute path
        And I should Copy Manifest
        And I should be able to copy the main file
        And I should able to run go mod tidy
        And I Should be able to run crda analyse without error
        And I should perform cleanup

    Scenario: Check crda analyse command for Go with vulns and all flags true
        Given system is running and authenticated with valid snyk token
        When I run crda analyse command for Go with vulns json
        Then I should be able to get the absolute path
        And I should Copy Manifest
        And I should be able to copy the main file
        And I should able to run go mod tidy
        And I Should be able to run crda analyse without error
        And I should perform cleanup

    Scenario: Check crda analyse command for Go without vulns
        Given system is running and authenticated with valid snyk token
        When I run crda analyse command for Go with vulns json
        Then I should be able to get the absolute path
        And I should Copy Manifest
        And I should be able to copy the main file
        And I should able to run go mod tidy
        And I Should be able to run crda analyse without error
        And I should perform cleanup

    Scenario: Check crda analyse command for Go without vulns json
        Given system is running and authenticated with valid snyk token
        When I run crda analyse command for Go without vulns json
        Then I should be able to get the absolute path
        And I should Copy Manifest
        And I should be able to copy the main file
        And I should able to run go mod tidy
        And I Should be able to run crda analyse without error
        And I should perform cleanup

    

    