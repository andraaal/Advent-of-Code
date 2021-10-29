def spirals(n):
    locs = {} # Will map a tuple (i,j) to a value
    locs[(0,0)] = 1

    def get_loc(i,j):
        return locs[(i,j)] if (i,j) in locs else 0

    def turn_ccw(dx, dy):
        return [-dy, dx]

    [dx, dy] = [1, 0] # point right
    [x, y] = [0, 0] # start at origin
    travel_distance = 1
    traveled_so_far = 0
    increase_travel_distance = False #
    while(locs[(x,y)] < n):
        x += dx
        y += dy

        locs[(x,y)] = (get_loc(x+1,y+1) + get_loc(x,y+1) + get_loc(x-1,y+1) + 
                      get_loc(x+1,y) + get_loc(x-1,y) + 
                      get_loc(x+1,y-1) + get_loc(x,y-1) + get_loc(x-1,y-1))

        traveled_so_far += 1
        if (traveled_so_far == travel_distance):
            traveled_so_far = 0
            [dx, dy] = turn_ccw(dx, dy)
            if (increase_travel_distance):
                travel_distance += 1
                increase_travel_distance = False
            else:
                increase_travel_distance = True

    return locs[x,y]
print("The solution is: " + str(spirals(265149)))