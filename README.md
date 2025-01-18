Found a puzzle on the notorious [Futility Closet](https://www.futilitycloset.com/) web site.

[Here's](https://www.futilitycloset.com/2025/01/15/gun-control-2/) Futility Closet's puzzle.
Mr or Ms Closet decided to call the puzzle "Gun Control".
I don't know if they chose the title to skirt controversy or something.

## Problem Statement

Marksman A hits a certain small target 75 percent of the time.
Marksman B hits it 25 percent of the time.
The two of them aim at that target and fire simultaneously.
One bullet hits it.
What’s the probability that it came from A?

---

I decided to solve this via a simulation.

## Simulation

1. Chose a random numer between 0 and 100 inclusive.
If that random number evalutes to less than 75, Marksman A hit the target.
2. Chose a random numer between 0 and 100 inclusive.
If that random number evalutes to less than 25, Marksman b hit the target.
3. If A or B (but not both!) hit the target, count who hit the target,
A or B
4. Repeat a few times, display percentages of who hit the target.

[Code repo](https://github.com/bediger4000/guncontrol)
