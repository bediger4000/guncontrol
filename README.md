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

<!-- 90% of the time, why do you ask? -->

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

## Build and run

```
$ cd $GOPATH/src
$ git clone https://github.com/bediger4000/guncontrol.git
$ cd guncontrol
$ go build gc1.go
$ ./gc1
```

## Usage

```
  -A int
        Proability in percent that Marksman A hits the target (default 75)
  -B int
        Proability in percent that Marksman B hits the target (default 25)
  -N int
        Number of rounds of shooting to simulate (default 10000)
```

I put 10,000 iterations as a default, but 100,000 iterations works better,
and is computationally quite quick.
