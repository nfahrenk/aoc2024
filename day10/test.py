def readFile(filename):
    grid = []
    points = []
    with open(filename, 'r') as file:
        row = 0
        for line in file:
            col = 0
            entry = []
            for elem in line.strip():
                if elem == "0":
                    points.append((row, col))
                entry.append(int(elem))
                col += 1
            grid.append(entry)
            row += 1
    return grid, points

def checkBounds(grid, r, c):
    return r >= 0 and r < len(grid) and c >= 0 and c < len(grid[0])

def score(grid, start, ignoreVisited=False):
    r, c = start
    path = [(r, c, 0)]
    count = 0
    visited = set()
    while path:
        point = path.pop()
        r, c, level = point
        if level == 9:
            if ignoreVisited or (r, c) not in visited:
                count += 1
                visited.add((r, c))
            continue
        for vec in [(0, 1), (1, 0), (-1, 0), (0, -1)]:
            i, j = vec
            if checkBounds(grid, r + i, c + j) and grid[r+i][c+j] == level + 1:
                path.append((r + i, c + j, level + 1))
    return count

def part1(filename, ignoreVisited=False):
    grid, points = readFile(filename)
    return sum([score(grid, p, ignoreVisited) for p in points])

def part2(filename):
    return part1(filename, True)

if __name__ == "__main__":
    print("solution", part2("day10/input2.txt"))
