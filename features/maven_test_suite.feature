Feature: Test CRDA analyse Maven
    Scenario: Check crda analyse command for maven with no vulns
        Given system is running
        When I test analyse command for Maven with no vulns
        Then Should be able to get the absolute path
        And Copy Manifest
        And Should be able to run analyse without error
        And I should perform cleanup

    Scenario: Check crda analyse command for maven with no vulns json
        Given system is running
        When I test analyse command for Maven with no vulns json
        Then Should be able to get the absolute path
        And Copy Manifest
        And Should be able to run analyse without error
        And I should perform cleanup

    Scenario: Check crda analyse command for maven with vulns
        Given system is running
        When I test analyse command for Maven with vulns
        Then Should be able to get the absolute path
        And Copy Manifest
        And Should be able to run analyse without error
        And I should perform cleanup

    Scenario: Check crda analyse command for maven with vulns and json
        Given system is running
        When I test analyse command for Maven with vulns and json
        Then Should be able to get the absolute path
        And Copy Manifest
        And Should be able to run analyse without error
        And I should perform cleanup

    Scenario: Check crda analyse command for maven with vulns and verbose
        Given system is running
        When I test analyse command for Maven with vulns and verbose
        Then Should be able to get the absolute path
        And Copy Manifest
        And Should be able to run analyse without error
        And I should perform cleanup

    Scenario: Check crda analyse command for maven with vulns and debug
        Given system is running
        When I test analyse command for Maven with vulns and debug
        Then Should be able to get the absolute path
        And Copy Manifest
        And Should be able to run analyse without error
        And I should perform cleanup

    Scenario: Check crda analyse command for maven with vulns and all flags true
        Given system is running
        When I test analyse command for Maven with vulns and all flags true
        Then Should be able to get the absolute path
        And Copy Manifest
        And Should be able to run analyse without error
        And I should perform cleanup

    