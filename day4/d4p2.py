def read_grid(filename):
    with open(filename, 'r') as file:
        return [line.strip() for line in file]

def is_xmas_pattern(grid, row, col):
    # Check if the X-MAS pattern fits within the grid boundaries
    if row + 2 >= len(grid) or col + 2 >= len(grid[0]):
        return False
    
    # Check the X-MAS pattern
    if (grid[row][col] == 'M' and grid[row][col + 2] == 'S' and
        grid[row + 1][col + 1] == 'A' and
        grid[row + 2][col] == 'M' and grid[row + 2][col + 2] == 'S'):
        return True
    
    # Check the reversed X-MAS pattern
    if (grid[row][col] == 'S' and grid[row][col + 2] == 'M' and
        grid[row + 1][col + 1] == 'A' and
        grid[row + 2][col] == 'S' and grid[row + 2][col + 2] == 'M'):
        return True
    # Check the downward pattern
    if (grid[row][col] == 'M' and grid[row][col + 2] == 'M' and
        grid[row + 1][col + 1] == 'A' and
        grid[row + 2][col] == 'S' and grid[row + 2][col + 2] == 'S'):
        return True
    # Check the upward pattern
    if (grid[row][col] == 'S' and grid[row][col + 2] == 'S' and
        grid[row + 1][col + 1] == 'A' and
        grid[row + 2][col] == 'M' and grid[row + 2][col + 2] == 'M'):
        return True
    
    return False

def count_xmas_patterns(grid):
    count = 0
    for row in range(len(grid)):
        for col in range(len(grid[0])):
            if is_xmas_pattern(grid, row, col):
                count += 1
    return count

if __name__ == "__main__":
    grid = read_grid('day4Input.txt')
    xmas_count = count_xmas_patterns(grid)
    print(f"Number of X-MAS patterns: {xmas_count}")