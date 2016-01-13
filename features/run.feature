Feature: Run programs successfully
    Scenario: PRINT and PRINTLN
        Given a file named "pl0.asm" with:
        """
        LOAD A, 1
        LOAD B, 2
        LOAD C, 3
        PRINT A
        PRINT B
        PRINTLN
        PRINT C
        PRINTLN
        END
        """
        When I successfully run `pl0dashvm pl0.asm`
        Then the output should contain exactly:
        """
        12
        3
        """

    Scenario: LOAD and STORE
        Given a file named "pl0.asm" with:
        """
        LOAD A, 1
        STORE A, #(801)
        LOAD B, #(801)
        PRINT B
        PRINTLN
        LOAD B, 801
        STORE A, #(B)
        LOAD C, #(801)
        PRINT C
        PRINTLN
        END
        """
        When I successfully run `pl0dashvm pl0.asm`
        Then the output should contain exactly:
        """
        1
        1
        """

    Scenario: JMP and JPC
        Given a file named "pl0.asm" with:
        """
        LOAD A, 1
        LOAD B, 2
        JMP 6
        PRINT A
        PRINTLN
        PRINT B
        PRINTLN
        LOAD C, 1
        JPC 12
        PRINT A
        PRINTLN
        PRINT B
        PRINTLN
        LOAD C, 0
        JPC 18
        PRINT A
        PRINTLN
        PRINT B
        PRINTLN
        END
        """
        When I successfully run `pl0dashvm pl0.asm`
        Then the output should contain exactly:
        """
        2
        1
        2
        2
        """

    Scenario: PUSH and POP
        Given a file named "pl0.asm" with:
        """
        LOAD A, 1
        LOAD B, 2
        LOAD C, 3
        PUSH A
        POP A
        PUSH B
        PUSH C
        POP B
        POP C
        PRINT A
        PRINTLN
        PRINT B
        PRINTLN
        PRINT C
        PRINTLN
        END
        """
        When I successfully run `pl0dashvm pl0.asm`
        Then the output should contain exactly:
        """
        1
        3
        2
        """

    Scenario: PLUS, MINUS, MULTI, and DIV
        Given a file named "pl0.asm" with:
        """
        LOAD A, 5
        LOAD B, 2
        PLUS
        PRINT C
        PRINTLN
        MINUS
        PRINT C
        PRINTLN
        MULTI
        PRINT C
        PRINTLN
        DIV
        PRINT C
        PRINTLN
        END
        """
        When I successfully run `pl0dashvm pl0.asm`
        Then the output should contain exactly:
        """
        7
        3
        10
        2
        """

    Scenario: CMPODD
        Given a file named "pl0.asm" with:
        """
        LOAD A, 1
        LOAD C, 100
        CMPODD
        PRINT C
        PRINTLN
        LOAD A, 2
        LOAD C, 100
        CMPODD
        PRINT C
        PRINTLN
        END
        """
        When I successfully run `pl0dashvm pl0.asm`
        Then the output should contain exactly:
        """
        1
        0
        """

    Scenario: CMPEQ
        Given a file named "pl0.asm" with:
        """
        LOAD A, 1
        LOAD B, 1
        LOAD C, 100
        CMPEQ
        PRINT C
        PRINTLN
        LOAD A, 1
        LOAD B, 2
        LOAD C, 100
        CMPEQ
        PRINT C
        PRINTLN
        LOAD A, 2
        LOAD B, 1
        LOAD C, 100
        CMPEQ
        PRINT C
        PRINTLN
        END
        """
        When I successfully run `pl0dashvm pl0.asm`
        Then the output should contain exactly:
        """
        1
        0
        0
        """

    Scenario: CMPNOTEQ
        Given a file named "pl0.asm" with:
        """
        LOAD A, 1
        LOAD B, 1
        LOAD C, 100
        CMPNOTEQ
        PRINT C
        PRINTLN
        LOAD A, 1
        LOAD B, 2
        LOAD C, 100
        CMPNOTEQ
        PRINT C
        PRINTLN
        LOAD A, 2
        LOAD B, 1
        LOAD C, 100
        CMPNOTEQ
        PRINT C
        PRINTLN
        END
        """
        When I successfully run `pl0dashvm pl0.asm`
        Then the output should contain exactly:
        """
        0
        1
        1
        """

    Scenario: CMPLT
        Given a file named "pl0.asm" with:
        """
        LOAD A, 1
        LOAD B, 1
        LOAD C, 100
        CMPLT
        PRINT C
        PRINTLN
        LOAD A, 1
        LOAD B, 2
        LOAD C, 100
        CMPLT
        PRINT C
        PRINTLN
        LOAD A, 2
        LOAD B, 1
        LOAD C, 100
        CMPLT
        PRINT C
        PRINTLN
        END
        """
        When I successfully run `pl0dashvm pl0.asm`
        Then the output should contain exactly:
        """
        0
        1
        0
        """

    Scenario: CMPGT
        Given a file named "pl0.asm" with:
        """
        LOAD A, 1
        LOAD B, 1
        LOAD C, 100
        CMPGT
        PRINT C
        PRINTLN
        LOAD A, 1
        LOAD B, 2
        LOAD C, 100
        CMPGT
        PRINT C
        PRINTLN
        LOAD A, 2
        LOAD B, 1
        LOAD C, 100
        CMPGT
        PRINT C
        PRINTLN
        END
        """
        When I successfully run `pl0dashvm pl0.asm`
        Then the output should contain exactly:
        """
        0
        0
        1
        """

    Scenario: CMPLE
        Given a file named "pl0.asm" with:
        """
        LOAD A, 1
        LOAD B, 1
        LOAD C, 100
        CMPLE
        PRINT C
        PRINTLN
        LOAD A, 1
        LOAD B, 2
        LOAD C, 100
        CMPLE
        PRINT C
        PRINTLN
        LOAD A, 2
        LOAD B, 1
        LOAD C, 100
        CMPLE
        PRINT C
        PRINTLN
        END
        """
        When I successfully run `pl0dashvm pl0.asm`
        Then the output should contain exactly:
        """
        1
        1
        0
        """

    Scenario: CMPGE
        Given a file named "pl0.asm" with:
        """
        LOAD A, 1
        LOAD B, 1
        LOAD C, 100
        CMPGE
        PRINT C
        PRINTLN
        LOAD A, 1
        LOAD B, 2
        LOAD C, 100
        CMPGE
        PRINT C
        PRINTLN
        LOAD A, 2
        LOAD B, 1
        LOAD C, 100
        CMPGE
        PRINT C
        PRINTLN
        END
        """
        When I successfully run `pl0dashvm pl0.asm`
        Then the output should contain exactly:
        """
        1
        0
        1
        """
