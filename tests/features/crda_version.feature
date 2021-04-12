Feature: Run Crda version
    Scenario: check which crda version is installed
        Given system is running
        When I run crda version command
        Then it runs and validate CLI version