Feature: Test for invalid command
    Scenario: Check for invalid command error
        Given system is running 
        When I run invalid crda command
        Then it should throw invalid command error