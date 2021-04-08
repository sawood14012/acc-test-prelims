Feature: Run Crda analyse without any file
    Scenario: Check for crda analyse without file error
        Given system is running
        When I run crda analyse command without path
        Then validate analyse without file throws error