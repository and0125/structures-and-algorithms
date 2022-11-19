# Big O Notation

## Overview

Big O notation is the language and metric used to describe the efficiency of an algorithm. This helps you judge the performance of an algorithm.

## Time Complexity

Big O notation represents time complexity; the asymptotic runtime of an algorithm.

As a general rule, constant time will eventually run faster than linear time at a certain time depending on the size of the data transformation the algorithm is performing.

Some common runtimes are:

- O(1): constant time
- O(N): linear time
- O(log(n)): logarithmic time
- O(n \* log(n)): n log n time
- O(N^2): squared time (?)
- O(2^N): exponential time (?)

also there can be multiple variables in the time complexity formula of a specific algorithm. The example given was that the height and width of a wall both effect the amoutn of time needed to paint the wall, and could be described by a time complexity formula of O(w\*h).

## Different Bounds of an Algorithm

- Big O (O): used to describe the upper bound of the time complexity. Technically, you can say list any function here, but its most useful when you find the expression for the lowest upper bound of the time complexity.
- Big Omega: this is the lower bound of the time complexity.
- Big Theta: this indicates that the upper bound and lower bound of a function are the same.

In the workplace, Big O notation is used colloquially to describe big theta notation. That is Big O notation is meant to determine the tightest description of the runtime complexity of an algorithm.

### Describing Scenarios

You can use these terms to describe the best case, worst case, and expected time complexity of an algorithm.

Often, the best case and the worst case are the same in practice.

In translation, the best worst and expected cases describe the Big O or Big Theta time for particular inputs or scenarios.

## Space Complexity

We also care about the amount of memory or space an algorithm requires. We describe this space complexity with the same big O notation.

Some common space complexities are:

- O(1): constant space
- O(N): linear space
- O(log(n)): logarithmic space
- O(n \* log(n)): n log n space
- O(N^2): squared space (?)
- O(2^N): exponential space (?)

Note that just because a function may make N stack calls (i.e. output N times, or adjust a variable n times to work) it doesn't mean it takes O(n) space.

## Asymptotic Behavior

It's also very possible for O(N) code to run faster than O(1) code for specif inputs. Big O notation is meant to describe the rate of increase in either space or time as the inputs to the algorithm get larger.

For this reason, because constants are small relative to the increase, constant time or constant space terms are often dropped from the final big O expression of an algorithm.

The same reasoning applies to constant multipliers, so O(2\*N) is equivalent to O(N) because the constant doesn't change the rate of increase with respect to the input.

### Dropping Non-Dominant Terms

This is very similar to asymptotic limits of math functions. if there are exponential factors in a time complexity, they increase faster than linear terms in that function, and as the size of input grows, those exponential factors will take over.

For this reason, O(N^2 + N) is equivalent to O(N^2) because the exponential term increases faster.

Remember the whole notation is meant to capture the overall rate of increase, and the fastest increasing term in the algorithm is the dominant term.

Examples:

- O(N^2 + N) is equivalent to O(N^2) (polynomial terms increase faster than linear terms)
- O(N + log(N)) is equivalent to O(N) (linear increases faster than logarithmic)
- O(5\*2^N + 1000\*N^100) is equivalent to O(2^N) (exponential terms increase faster than polynomial terms)

Note that these reductions work when the variable used in each term of the algorithm is the same input, therefore we can compare them. Much like a math problem with multiple variables, unless the two variables can be compared in an equation, the terms with the different variables must be maintained in the final time or space complexity of the algorithm:

- so O(2^M + N) cannot be reduced to a simpler complexity expression without some knowledge of how M and N are related. If they aren't, then this is the final time or space complexity.

Order of non-constant runtimes (from best to worst):

- O(log(n))
- O(n)
- O( n\* log(n))
- O(n^2); note that 2 could be any constant
- O(2^n); note that 2 could be any constant
- O(n!)

## Analyzing Multi-Part Algorithms

for functions within an algorithm that are run sequentially and not consecutively, the complexity of the two functions are added.

For functions that are nested and run one after the other (nested for loops for example), the complexity is multiplied.

shorthand:

- if algorithm does one thing then another, the runtimes are additive
- if every time the algorithm does one thing it has to do another, the runtimes are multiplicative.

## Amortized Time

This is a way to capture the fact that sometimes the most complex case of an algorithm will not happen very often.

Amortized time is a way to state that a functions worst case will only happen every once in awhile.

So for a function that will usually run in constant time, O(1), and only in the worst case runs in O(N) time, we say the algorithm worst case takes O(N) time, and but has O(1) amortized time.

## Logarithmic Runtimes

Logarithmic run times occur when the number of elements in the problem space gets reduced by half at every step in the algorithm. Binary search is an example of an O(log_2(n)) algorithm.

When considering big O notation, the base for the logarithm is droppped because logarithms of different bases only differ by a constant factor.

for any number n, if we have log_2(n), we can convert this to log_10(n) (or just log(n)) with:

- log_10(n) = log_2(n) / log_2(10)

So this constant term, (log_2(10)), is the only difference, and we ignore constant terms in big O notation.

Therefore, we can omit the base of any logarithmic time/space complexity, and write these commonly as O(log(n)).

## Recursive Runtimes

These can be tricky, but often you can analyze these functions by looking at the number of function calls made, and turning this into a tree diagram, then assessing the number of responses or nodes or outputs at each step of the recursive call.

Once you have this in a diagram, look for the pattern between consecutive layers of the recursive call stack.

This relationship will define the runtime complexity.

In the tree diagram, you can roughly think of the time complexity as O(b^n), where b is the number of branches, and n is the depth of the tree.

The example given was reduced to O(2^N -1) time complexity, where the -1 excludes the base case.

Note that the base of the exponential term matters in assessing Time complexity, because the difference in the base results in an exponential increase in the input, not a constant term difference.

The space complexity is different, and is O(N), because the space complexity only denotes the amount of memory allotted to each function call. In this recursive example, the function is only called four times, even though, at the lowest level, 8 nodes are produced. Remember that difference for space assessments, as it could seem like the space complexity is the same, but it can be different as in this case.

## Pointers in working through the Analysis of algorithms

- label the variables individually and concretely. Make sure you represent each variable individually, so you can express the runtime as a function of those variables, and can assess whether the variables are independent or can be reduced to expressions of one another.

- constant time terms are constant time, even if there are 100k iterations of a for loop. 100k is still a constant.

- Remember the order of dominant terms above.

- Don't think that the mathematical expresssion used in a function factors into its runtime just by assumption. Really assess the number of loops the algorithm does, not the complexity of math expressions within it. Each math expression is still a constant time function.
