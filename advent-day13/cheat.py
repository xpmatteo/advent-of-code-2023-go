import numpy as np

def mirrorpos(arr, axis=0, diff=0):
    m = np.array(arr, dtype=int)
    if axis == 1:
        m = m.T
    for i in range(m.shape[0] - 1):
        upper_flipped = np.flip(m[: i + 1], axis=0)
        lower = m[i + 1 :]
        rows = min(upper_flipped.shape[0], lower.shape[0])
        if np.count_nonzero(upper_flipped[:rows] - lower[:rows]) == diff:
            return i + 1
    return 0

with open("day13.txt", "r") as file:
    data = file.read().split("\n\n")
for i in range(2):
    total = 0
    for puzzle in data:
        arr = []
        for line in puzzle.splitlines():
            arr.append([*line.strip().replace(".", "0").replace("#", "1")])
        h = 100 * mirrorpos(arr, axis=0, diff=i)
        v = mirrorpos(arr, axis=1, diff=i)
        print(puzzle, h, v)
        total += h + v
    print(total)


