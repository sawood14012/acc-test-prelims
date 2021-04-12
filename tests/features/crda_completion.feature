Feature: Validate CRDA completion command by os
    Scenario: system is darwin or linux
        Given system is running 
        When I run crda completion bash command
        Then it should run and validate completion command

    Scenario: system is windows
        Given system is running 
        When I run crda completion powershell command
        Then it should run and validate completion command