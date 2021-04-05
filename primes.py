# Needs to be converted to golang

def power(x, y, p) :
    res = 1     # Initialize result
 
    # Update x if it is more
    # than or equal to p
    x = x % p
     
    if (x == 0) :
        return 0
 
    while (y > 0) :
        print("y:", y)
        # If y is odd, multiply
        # x with result
        if ((y & 1) == 1) :
            res = (res * x) % p
        print("res: ",res)
        # y must be even now
        y = y >> 1      # y = y/2
        print("y after: ", y)
        x = (x * x) % p
        print("x after: ", x)
         
    return res
