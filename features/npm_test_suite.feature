Feature: Test CRDA analyse Npm
    Scenario: Check crda analyse command for npm with vulns
        Given system is running and authenticating with valid snyk token
        When I test analyse command for npm with vulns
        Then I should Get Absolute Path
        And I Should Copy Manifest
        And I should Install Dependencies to run Stack analyses
        And I should Validate analyse for npm ecosystem
        And I should Cleanup

    Scenario: Check crda analyse command for npm with vulns with json
        Given system is running and authenticating with valid snyk token
        When I test analyse command for npm with vulns with json
        Then I should Get Absolute Path
        And I Should Copy Manifest
        And I should Install Dependencies to run Stack analyses
        And I should Validate analyse for npm ecosystem with vulns
        And I should Cleanup

    Scenario: Check crda analyse command for npm with vulns with verbose
        Given system is running and authenticating with valid snyk token
        When I test analyse command for npm with vulns with verbose
        Then I should Get Absolute Path
        And I Should Copy Manifest
        And I should Install Dependencies to run Stack analyses
        And I should Validate analyse for npm ecosystem with vulns json and verbose
        And I should Cleanup

    Scenario: Check crda analyse command for npm with vulns with debug
        Given system is running and authenticating with valid snyk token
        When I test analyse command for npm with vulns with debug
        Then I should Get Absolute Path
        And I Should Copy Manifest
        And I should Install Dependencies to run Stack analyses
        And I should Validate analyse for npm ecosystem with vulns with debug
        And I should Cleanup

    Scenario: Check crda analyse command for npm with vulns with all flags set true
        Given system is running and authenticating with valid snyk token
        When I test analyse command for npm with vulns with all flags set true
        Then I should Get Absolute Path
        And I Should Copy Manifest
        And I should Install Dependencies to run Stack analyses
        And I should Validate analyse for npm ecosystem with vulns with all flags set true
        And I should Cleanup

    Scenario: Check crda analyse command for npm without vulns with json
        Given system is running and authenticating with valid snyk token
        When I test analyse command for npm without vulns with json
        Then I should Get Absolute Path
        And I Should Copy Manifest
        And I should Install Dependencies to run Stack analyses
        And I should Validate analyse for npm ecosystem without vulns with json
        And I should Cleanup

    

    