input = open('./input.txt', 'r').read()

input = list(zip(*([list(line) for line in input.split()][::-1])))
outputHighest = ''.join([max(lis, key=lis.count) for lis in input])
outputLowest = ''.join([min(lis, key=lis.count) for lis in input])

print(int(outputHighest, 2) * int(outputLowest, 2))
