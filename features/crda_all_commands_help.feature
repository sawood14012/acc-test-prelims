Feature: Validate there is a help page for all commands
    Scenario: check analyse command has help page
        Given system is running 
        When I run "crda analyse --help" command
        Then it should run and validate "crda analyse --help" command

    Scenario: check auth command has help page
        Given system is running 
        When I run "crda auth --help" command
        Then it should run and validate "crda auth --help" command

    Scenario: check completion command has help page
        Given system is running 
        When I run "crda completion --help" command
        Then it should run and validate "crda completion --help" command

    Scenario: check version command has help page
        Given system is running 
        When I run "crda version --help" command
        Then it should run and validate "crda version --help" command

    Scenario: check help command has help page
        Given system is running 
        When I run "crda help bash" command
        Then it should run and validate "crda hrlp bash" command