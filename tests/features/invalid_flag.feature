Feature: Test for invalid flag throws error
    Scenario: Check for invalid flag error
        Given system is running and authenticated with valid snyk token 
        When I run crda analyse command with invalid flag
        Then it should throw invalid flag error