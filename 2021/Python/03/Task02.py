input = open('./input.txt', 'r').read()

def getMostCommon(liste):
    rotated = list(zip(*(liste[::-1])))
    return [max(lis, key=lis.count) for lis in rotated]

def getLeastCommon(liste):
    rotated = list(zip(*(liste[::-1])))
    return [min(lis, key=lis.count) for lis in rotated]

input = [list(line) for line in input.split()]


ogr = input
iterator = 0
print(ogr)
while len(ogr) > 1:
    ogr = [line for line in ogr if line[iterator] == getMostCommon(ogr)[iterator]]
    iterator += 1
    print(ogr)

csr = input
iterator = 0
while len(csr) > 1:
    csr = [line for line in csr if line[iterator] == getLeastCommon(csr)[iterator]]
    iterator += 1
    print(csr)

print(int(''.join(ogr[0]), 2) * int(''.join(csr[0]), 2))