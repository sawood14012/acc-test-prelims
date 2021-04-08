Feature: Run Crda analyse without any file
    Given system is running
    When I run crda analyse command without path
    Then validate analyse without file throws error

    