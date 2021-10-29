n = 265149
i = 1
while i*i < n:
    i += 2
pivots = [i*i - k*(i-1) for k in range(4)]
for p in pivots:
    dist = abs(p - n)
    if dist <= (i-1)//2:
        print(i-1-dist)
        break