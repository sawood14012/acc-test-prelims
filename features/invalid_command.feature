Feature: Test for invalid command
    Given system is running 
    When I run invalid crda command
    Then it should throw invalid command error