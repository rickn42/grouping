# grouping

## result

#### case alleviate amount = 1.0
```
repeat grouping...
person(0) score 92
[0 3 3 1 5 4 4 5 2 0 4 1 1 3 4 4 3 2 2 2 4 3 1 2 3 1 2 3 1 1 1 1 3 6 0 3 1 5 2 2 1 1 3 3 2 1 3 3 2 3]
person(1) score 90
[3 0 2 1 4 2 3 2 3 1 4 4 1 1 4 2 5 1 3 1 1 2 2 2 1 2 3 1 0 4 3 3 2 1 4 1 3 3 3 4 2 4 4 4 3 3 0 0 2 6]
...

grouping with person 50, member 5, generation 30, ideal met count 2.4
statistic score: min 56.1, max 170.1, avg 96.3
group boring score: max 3.6, avg 0.1
```

```
repeat grouping...
person(0) score 142
[0 3 3 4 5 4 6 4 5 7 5 8 7 8 1 5 6 7 3 7 5 6 5 3 5 5 4 6 6 1 7 5 6 4 3 2 5 8 3 3 3 3 5 6 6 4 5 6 5 7]
person(1) score 156
[3 0 5 1 4 6 5 7 6 8 4 5 4 3 4 2 4 3 3 5 5 4 9 7 4 5 6 6 7 2 5 4 8 6 6 6 7 7 5 4 3 6 5 4 9 3 4 4 2 5]
...

grouping with person 50, member 5, generation 60, ideal met count 4.9
statistic score: min 84.5, max 178.5, avg 127.2
group boring score: max 28.0, avg 1.6
```

```
repeat grouping...
person(0) score 165
[0 6 10 5 8 5 7 5 6 7 6 6 11 7 6 8 11 6 9 6 7 8 9 6 8 11 5 6 10 6 10 8 4 9 9 7 6 8 10 9 8 9 6 7 7 6 8 8 3 7]
person(1) score 131
[6 0 7 8 5 10 10 7 10 10 6 6 5 6 7 8 7 9 7 8 6 4 7 5 8 5 10 10 7 6 7 8 7 8 7 8 9 8 9 5 5 9 9 8 8 9 4 7 7 8]
...

grouping with person 50, member 5, generation 90, ideal met count 7.3
statistic score: min 99.1, max 243.1, avg 160.8
group boring score: max 152.1, avg 9.6
```

... not good. need improvement.

#### case alleviate amount = 0.0
```
...
person(48) score 27
[41 41 41 41 40 41 42 39 41 41 41 42 41 41 41 41 40 41 41 40 41 41 41 40 41 42 39 42 42 40 41 41 40 41 42 40 41 41 40 41 40 40 40 40 41 42 40 41 0 42]
person(49) score 27
[42 41 41 40 40 42 41 41 41 40 41 40 42 40 41 39 41 41 40 42 41 41 41 41 41 41 41 39 40 41 40 41 41 40 42 41 41 41 41 41 40 41 41 42 42 40 40 40 42 0]
grouping with person 50, member 5, generation 500, ideal met count 40.8
statistic score: min 15.3, max 41.3, avg 25.3
group boring score: max 27205331.1, avg 8922814.2
```

```
person(48) score 26
[5 4 5 4 4 5 4 4 4 4 5 5 4 3 4 3 3 5 4 4 4 4 4 4 3 5 4 3 3 4 5 5 4 5 3 5 2 5 4 4 4 5 4 3 4 4 4 5 0 4]
person(49) score 22
[3 5 4 4 4 3 5 3 4 4 5 4 4 4 4 5 5 4 4 4 3 4 4 4 4 4 4 5 4 5 3 5 4 4 4 4 6 3 5 4 4 3 5 3 4 4 4 4 4 0]
grouping with person 50, member 5, generation 50, ideal met count 4.1
statistic score: min 13.7, max 37.7, avg 21.8
group boring score: max 2465.6, avg 678.3
```

-_-??? better uniform distribution...

but, how about boring(continuous meeting) score?