def getDepthSum(lastDepths):
    print(lastDepths[-1] + lastDepths[-2] + lastDepths[-3])
    return lastDepths[-1] + lastDepths[-2] + lastDepths[-3]

if(__name__ == "__main__"):
    input = open('./input.txt', 'rt').read()
    depths = input.split('\n')
    output = 0

    lastDepths = [int(depths[0]), int(depths[1]), int(depths[2])]

    for depth in depths:
        depth = int(depth)
        lastSum = getDepthSum(lastDepths)
        lastDepths.append(depth)
        currentSum = getDepthSum(lastDepths)
        if currentSum > lastSum:
            output += 1

    print(output)

