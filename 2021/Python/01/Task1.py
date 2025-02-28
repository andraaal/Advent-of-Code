if(__name__ == "__main__"):
    input = open('C:/Dateien/Programieren/Advent of Code/2021/Python/01/input.txt', 'rt').read()
    depths = input.split('\n')
    output = 0

    lastDepth = int(depths[0])

    for depth in depths:
        depth = int(depth)
        if depth > lastDepth:
            output += 1
        lastDepth = depth

    print(output)

