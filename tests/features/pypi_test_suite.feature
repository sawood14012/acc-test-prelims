Feature: Test CRDA analyse Pypi
    Scenario: Check crda analyse command for pypi with vulns
        Given system is running and authenticating with valid snyk token
        When I test analyse command for pypi with vulns
        Then I Should be able to get the absolute path
        And I should Copy Manifest
        And I should able to run pip install
        And I Should be able to run analyse without error
        And I should perform cleanup

    Scenario: Check crda analyse command for pypi with vulns json
        Given system is running and authenticating with valid snyk token
        When I test analyse command for pypi with vulns json
        Then I Should be able to get the absolute path
        And I should Copy Manifest
        And I should able to run pip install
        And I Should be able to run analyse without error
        And I should perform cleanup

    Scenario: Check crda analyse command for pypi with vulns and verbose
        Given system is running and authenticating with valid snyk token
        When I test analyse command for pypi with vulns and verbose
        Then I Should be able to get the absolute path
        And I should Copy Manifest
        And I should able to run pip install
        And I Should be able to run analyse without error
        And I should perform cleanup

    Scenario: Check crda analyse command for pypi with vulns and debug
        Given system is running and authenticating with valid snyk token
        When I test analyse command for pypi with vulns and debug
        Then I Should be able to get the absolute path
        And I should Copy Manifest
        And I should able to run pip install
        And I Should be able to run analyse without error
        And I should perform cleanup

    Scenario: Check crda analyse command for pypi with vulns and all flags true
        Given system is running and authenticating with valid snyk token
        When I test analyse command for pypi with vulns and all flags true
        Then I Should be able to get the absolute path
        And I should Copy Manifest
        And I should able to run pip install
        And I Should be able to run analyse without error
        And I should perform cleanup

    Scenario: Check crda analyse command for pypi without vulns
        Given system is running and authenticating with valid snyk token
        When I test analyse command for pypi without vulns
        Then I Should be able to get the absolute path
        And I should Copy Manifest
        And I should able to run pip install
        And I Should be able to run analyse without error
        And I should perform cleanup

    Scenario: Check crda analyse command for pypi without vulns json
        Given system is running and authenticating with valid snyk token
        When I test analyse command for pypi without vulns json
        Then I Should be able to get the absolute path
        And I should Copy Manifest
        And I should able to run pip install
        And I Should be able to run analyse without error
        And I should perform cleanup