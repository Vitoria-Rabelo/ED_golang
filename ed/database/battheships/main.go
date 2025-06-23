func countBattleships(board [][]byte) int {
    m := len(board)
    n := len(board[0])
    if m == 0{
        return 0
    }
    cont := 0

    for i := 0; i < m; i++{
        for j := 0; j < n; j++{
            if board[i][j] == 'X'{
                cont++
                dfs(board, i, j, m, n)
            }
        }
    }

    return cont
}

func dfs(board [][]byte, i, j, m, n int){
    if i >= m || j >= n || j < 0 || i < 0 || board[i][j] != 'X'{
        return
    }
    board[i][j] = '.'
    dfs(board, i, j- 1, m, n)
    dfs(board, i, j + 1, m, n)
    dfs(board, i - 1, j, m, n)
    dfs(board, i + 1, j, m, n)
}