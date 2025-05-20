func largestRectangleArea(heights []int) int {
    res := 0
    var stack []int
    
    heights = append(heights, 0)
    for i, h := range heights {
        for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
            height := heights[stack[len(stack)-1]]
            stack = stack[:len(stack)-1]
            width := i
            if len(stack) > 0 {
                width = i - stack[len(stack)-1] - 1
            }
            res = max(res, height * width)
        }
        stack = append(stack, i)
    }
    return res
}
