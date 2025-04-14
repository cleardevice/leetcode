function convert(s: string, numRows: number): string {
    if (numRows === 1) {
        return s;
    }

    const lines = new Array(numRows).fill('');
    let i = 0;
    let toDown = true;
    for (const c of s) {
        lines[i] += c;
        i += toDown ? 1 : -1;
        if (i === 0 || i === numRows-1) {
            toDown = !toDown;
        }
    }
    return lines.reduce((acc, s) => acc + s);
}
