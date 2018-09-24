# grouping

## result

```
--- grouping by random shuffle ---
...
person(48) score 200
meet count [2 2 8 2 2 2 7 3 3 2 5 4 4 2 3 8 2 5 2 4 2 3 2 8 2 4 5 1 4 0 7 4 5 4 3 5 1 1 3 7 6 3 4 3 6 2 5 0 0 4]
person(49) score 179
meet count [5 8 4 2 4 3 4 1 4 4 5 5 4 0 4 5 2 4 2 2 2 0 3 3 4 3 2 3 4 2 2 3 4 3 4 6 4 0 4 6 5 11 5 4 5 4 5 1 4 0]

grouping with person 50, member of per group 3, total turn 90, ideal meet turn 24.5, ideal meet count 3.7

* statistics score
min 107.8, max 232.9, avg 174.4

* group boring cost
max 79233.0, avg 10431.8

* meet turn interval cost
min 29957.8 max 42321.2 avg 35262
```

```
--- grouping by regression ---
person(0) score 14
meet count [0 3 3 4 3 4 4 3 4 4 3 4 4 4 4 3 4 3 4 4 3 3 4 4 5 3 3 4 4 4 4 4 3 4 3 4 4 4 3 4 4 3 3 3 4 4 3 3 3 4]
...
person(49) score 14
meet count [4 4 4 4 3 3 3 4 4 3 4 4 3 3 3 5 4 4 4 4 3 3 4 4 4 4 4 3 3 3 3 3 4 3 4 3 3 4 4 3 4 4 4 3 4 3 4 4 3 0]

grouping with person 50, member of per group 3, total turn 90, ideal meet turn 24.5, ideal meet count 3.7

* statistics score
min 11.5, max 20.2, avg 13.9

* group boring cost
max 20408.5, avg 5060.5

* meet turn interval cost
min 16151.8 max 25681.8 avg 20300
```

#### much better grouping now! ^^