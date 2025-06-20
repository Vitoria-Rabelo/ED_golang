func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    inicio := image[sr][sc]
    if inicio == color {
        return image
    }
    dfs(image, sr, sc, inicio, color)
    return image
}

func dfs(image [][]int, r, c, inicio, color int) {
    if r < 0 || r >= len(image) || c < 0 || c >= len(image[0]) || image[r][c] != inicio {
        return
    }
    image[r][c] = color
    dfs(image, r+1, c, inicio, color)
    dfs(image, r-1, c, inicio, color)
    dfs(image, r, c+1, inicio, color)
    dfs(image, r, c-1, inicio, color)
}