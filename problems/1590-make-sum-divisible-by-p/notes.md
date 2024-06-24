
### Search

d0: [3,1,4,2]-10

d1: [3,1,4]-7 [3,1,2]-6 [3,4,2]-9 [1,4,2]-7

d2: [3,1] [3,4] [1,4] ~~[3,1]~~ [3,2] [1,2] [4,2] ~~[3,2]~~ ~~[3,4]~~ ~~[4,2]~~ ~~[1,2]~~ ~~[1,4]~~

d3: [3] [1] [4] [2]

### Prefix Sum

#### Example 1

nums = [3,1,4,2], p = 6, remainder = 4

prefixSums = [3,4,8,10], remainders = [3,4,2,4]


#### Example 2

nums = [6,3,5,2], p = 9, remainder = 7

prefixSums = [6,9,14,16], remainders = [6,0,5,7]

prefixSums = [6,0,11,18], remainders = [6,0,2,0]