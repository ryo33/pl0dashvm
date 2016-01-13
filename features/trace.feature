Feature: Run programs with trace option
    Scenario: 
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
        When I successfully run `pl0dashvm -t pl0.asm`
        Then the output should contain exactly:
        """
        OUTPUT	PC	SP	A	B	C	COMMAND
        	1	1000	0	0	0	LOAD A, 1
        	2	1000	1	0	0	LOAD B, 2
        	3	1000	1	2	0	LOAD C, 3
        	4	1000	1	2	3	PRINT A
        1	5	1000	1	2	3	PRINT B
        2	6	1000	1	2	3	PRINTLN

        	7	1000	1	2	3	PRINT C
        3	8	1000	1	2	3	PRINTLN

        	9	1000	1	2	3	END
        """
