func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }

    lines := make([][]byte, numRows)
    i, step := 0, 1
    for _, c := range s {
        lines[i] = append(lines[i], byte(c))
        i += step
        if i == 0 || i == numRows-1 {
            step = -step;
        }
    }
    return string(bytes.Join(lines, nil))
}
