## Monk and Inversions

Monk's best friend Micro, who happen to be an awesome programmer, got him an integer matrix M of size  for his birthday. Monk is taking coding classes from Micro. They have just completed array inversions and Monk was successful in writing a program to count the number of inversions in an array. Now, Micro has asked Monk to find out the number of inversion in the matrix M. Number of inversions, in a matrix is defined as the number of unordered pairs of cells  such that .
Monk is facing a little trouble with this task and since you did not got him any birthday gift, you need to help him with this task.

### Video
Video approach to solve this question: https://www.youtube.com/watch?v=ys5ZTg0yIWY&list=PL1YS4hYJip07A-YteNUR8qTeA_wHQarDX&index=43

### Input:
First line consists of a single integer T denoting the number of test cases.
First line of each test case consists of one integer denoting N. Following N lines consists of N space separated integers denoting the matrix M.

### Output:
Print the answer to each test case in a new line.

### Constraints:
1 <= T <= 100
1 <= N <= 20
1 <= M[i][j] <= 1000


### Sample Input
2
3
1 2 3
4 5 6
7 8 9
2
4 3
1 4

### Sample Output
0
2

### Explanation
In first test case there is no pair of cells , ,  having , so the answer is 0.
In second test case  and , so the answer is 2.

### Limitations
Time Limit: 1.0 sec(s) for each input file
Memory Limit: 256 MB
Source Limit: 1024 KB

### Result
```
Input      Result   Time (sec) Memory (KiB) Score Your Output Correct Output
Input #1   Accepted 0.008781   2            10  
Input #2   Accepted 0.009757   2            10  
Input #3   Accepted 0.010075   2            10  
Input #4   Accepted 0.009148   2            10  
Input #5   Accepted 0.162259   104328       10  
Input #6   Accepted 0.163268   104328       10  
Input #7   Accepted 0.163371   104072       10  
Input #8   Accepted 0.16292    104328       10  
Input #9   Accepted 0.163021   104328       10  
Input #10  Accepted 0.15407    104072       10
```
