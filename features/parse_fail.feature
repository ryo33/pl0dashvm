Feature: Fail to parse programs
    Scenario: Fail
        Given a file named "pl0.asm" with:
        """
        LOAD A, 1 
        LOAD B 2
        LOAD C, #(A)
        PRINT A, B
        PRINT C
        PRINTLN
        PPRINT B
        PRINTLN A
        END
        """
        When I run `pl0dashvm pl0.asm`
        Then the stderr should contain exactly:
        """
        parse failed at 1,10:	expects newline
        parse failed at 2,8:	expects ,
        parse failed at 3,11:	expects number
        parse failed at 4,8:	expects newline
        parse failed at 7,2:	wrong command
        parse failed at 8,8:	expects newline
        """
