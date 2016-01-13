Feature: Fail to process programs
    Scenario: Wrong address error
        Given a file named "pl0.asm" with:
        """
        LOAD A, 1
        STORE A, #(1000)
        STORE A, #(1001)
        END
        """
        When I run `pl0dashvm pl0.asm`
        Then the stderr should contain exactly:
        """
        process failed at 3: wrong address 1001
        """

    Scenario: Pop error
        Given a file named "pl0.asm" with:
        """
        LOAD A, 1
        PUSH A
        POP A
        POP A
        END
        """
        When I run `pl0dashvm pl0.asm`
        Then the stderr should contain exactly:
        """
        process failed at 4: can not pop from stack
        """

    Scenario: Value execute error
        Given a file named "pl0.asm" with:
        """
        LOAD A, 1
        STORE A, #(3)
        PUSH A
        END
        """
        When I run `pl0dashvm pl0.asm`
        Then the stderr should contain exactly:
        """
        process failed at 3: can not execute value
        """

    Scenario: Load command error
        Given a file named "pl0.asm" with:
        """
        LOAD A, #(2)
        PUSH A
        END
        """
        When I run `pl0dashvm pl0.asm`
        Then the stderr should contain exactly:
        """
        process failed at 1: expects a value
        """
