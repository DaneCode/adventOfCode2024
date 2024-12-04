def read_word_search(file_path):
    with open(file_path, 'r') as file:
        return [line.strip() for line in file.readlines()]

def count_word_occurrences(grid, word):
    rows, cols = len(grid), len(grid[0])
    word_len = len(word)
    count = 0

    def check_direction(x, y, dx, dy):
        for i in range(word_len):
            if not (0 <= x + i * dx < rows and 0 <= y + i * dy < cols):
                return False
            if grid[x + i * dx][y + i * dy] != word[i]:
                return False
        return True

    for i in range(rows):
        for j in range(cols):
            if grid[i][j] == word[0]:
                directions = [(1, 0), (0, 1), (1, 1), (1, -1), (-1, 0), (0, -1), (-1, -1), (-1, 1)]
                for dx, dy in directions:
                    if check_direction(i, j, dx, dy):
                        count += 1

    return count

if __name__ == "__main__":
    file_path = 'day4Input.txt'
    word_search = read_word_search(file_path)
    word = "XMAS"
    occurrences = count_word_occurrences(word_search, word)
    print(f"The word '{word}' appears {occurrences} times in the word search.")