def read_input(filename):
    with open(filename, 'r') as file:
        return [list(line.strip()) for line in file]

# Directions: up, right, down, left
directions = [
    (-1, 0),  # up
    (0, 1),   # right
    (1, 0),   # down
    (0, -1)   # left
]

def find_guard_position(input_map):
    for i, row in enumerate(input_map):
        for j, cell in enumerate(row):
            if cell == '^':
                input_map[i][j] = '.'  # Replace guard's initial position with empty space
                return (i, j)
    return None

def simulate_guard_movement(input_map):
    guard_pos = find_guard_position(input_map)
    guard_dir = 0  # Initially facing up

    visited = set()
    visited.add(guard_pos)

    while True:
        next_pos = (guard_pos[0] + directions[guard_dir][0], guard_pos[1] + directions[guard_dir][1])

        # Check if the guard has left the mapped area
        if (
            next_pos[0] < 0 or next_pos[0] >= len(input_map) or
            next_pos[1] < 0 or next_pos[1] >= len(input_map[0])
        ):
            break

        if input_map[next_pos[0]][next_pos[1]] == '#':
            # Turn right 90 degrees
            guard_dir = (guard_dir + 1) % 4
        else:
            # Move forward
            guard_pos = next_pos
            visited.add(guard_pos)

    return len(visited)

if __name__ == "__main__":
    input_map = read_input('d6Input.txt')
    distinct_positions_visited = simulate_guard_movement(input_map)
    print(f"Distinct positions visited: {distinct_positions_visited}")