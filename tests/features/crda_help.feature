Feature: Run Crda help
    Scenario: Check for crda help command
        Given system is running 
        When I run crda help command
        Then it should run and validate help command