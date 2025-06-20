 func solve(board [][]byte)  {
    nl := len(board)
    nc := len(board[0])
    if nl == 0 || nc == 0 {
        return
    }


    for l := 0; l < nl; l++{
        if board[l][0] == 'O' {
            dfs(board, l, 0, nl, nc)
        }
        if board[l][nc - 1] == 'O'{
            dfs(board, l, nc - 1, nl, nc)
        }
    }

     for c := 0; c < nc; c++{
        if board[0][c] == 'O' {
            dfs(board, 0, c, nl, nc)
        }
        if board[nl - 1][c] == 'O'{
            dfs(board, nl -1 , c, nl, nc)
        }
    }

    for l := 0; l < nl; l++{
        for c := 0; c < nc; c++{
            if board[l][c] == 'O' {
                board[l][c] = 'X'
            }else if board[l][c] == 'S'{
                board[l][c] = 'O'
            }
        }
    }   
 }

 func dfs(board [][]byte,l , c, nl, nc int){
    if l < 0 || c < 0 || l >= nl || c >= nc || board[l][c] != 'O'{
         return
    }
    board[l][c] = 'S'
    dfs(board, l + 1,c,  nl, nc)
	dfs(board, l, c + 1, nl, nc )
	dfs(board, l - 1, c, nl, nc)
	dfs(board, l, c- 1, nl, nc)
 }