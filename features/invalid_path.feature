Feature: Test for invalid path throws error
    Given system is running and authenticated with valid snyk token 
    When I run crda analyse command with invalid path
    Then it should throw invalid path error