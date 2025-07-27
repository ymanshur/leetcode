class Solution:
    def solveNQueens(self, n: int) -> List[List[str]]:
        solutions = []
        board = [["."] * n for _ in range(n)]

        # states management
        column = [0] * n
        diagonal = [0] * (2 * n)
        anti_diagonal = [0] * (2 * n)

        def place(row: int):
            if row == n:
                copy = ["".join(row_state) for row_state in board]
                solutions.append(copy)
                return
            
            for col in range(n):
                if column[col] + diagonal[row + col] + anti_diagonal[n - row + col] > 0:
                    continue
                
                board[row][col] = "Q"
                column[col] = diagonal[row + col] = anti_diagonal[n - row + col] = 1

                place(row + 1)
                
                board[row][col] = "."
                column[col] = diagonal[row + col] = anti_diagonal[n - row + col] = 0

        place(0)
        return solutions
        