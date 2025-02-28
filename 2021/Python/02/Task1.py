if __name__ == "__main__":
    input = open("C:/Dateien/Programieren/Advent of Code/2021/Python/02/input.txt", "rt").read()
    commands = input.split("\n")
    aim = 0
    depth = 0
    position = 0
    for command in commands:
        command = command.split(' ')
        if command[0] == 'forward':
            position += int(command[1])
            depth += int(command[1]) * aim
        elif command[0] == 'up':
            aim -= int(command[1])
        elif command[0] == 'down':
            aim += int(command[1])

print(depth * position)