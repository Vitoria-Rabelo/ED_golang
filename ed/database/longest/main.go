type Pos struct {
    i, j int
}

func getNeib(p Pos) []Pos {
    return []Pos{
        {p.i - 1, p.j}, 
        {p.i + 1, p.j},
        {p.i, p.j - 1}, 
        {p.i, p.j + 1}, 
    }
}

func inside(p Pos, matrix [][]int) bool {
    return p.i >= 0 && p.j >= 0 && p.i < len(matrix) && p.j < len(matrix[0])
}

func longestIncreasingPath(matrix [][]int) int {
    if len(matrix) == 0 {
        return 0
    }

    m, n := len(matrix), len(matrix[0])
    maxLength := 1

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            currentLength := dfs(matrix, Pos{i, j})
            if currentLength > maxLength {
                maxLength = currentLength
            }
        }
    }

    return maxLength
}

func dfs(matrix [][]int, p Pos) int {
    maxDepth := 1

    for _, neighbor := range getNeib(p) {
        if inside(neighbor, matrix) && matrix[neighbor.i][neighbor.j] > matrix[p.i][p.j] {
            currentDepth := 1 + dfs(matrix, neighbor)
            if currentDepth > maxDepth {
                maxDepth = currentDepth
            }
        }
    }

    return maxDepth
}