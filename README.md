# SD Gophers Challange for September (North edition)

## Ulam's spirals


[Stanisław Marcin Ulam][ulam] was a mathematician in the early 20th century; who while board at a lecture came started drawing spirals of numbers which later came to bare his name. 
For this challange we will write a program that will generate the spirals.

A simple example of an ulam spiral: This is a 3x3 squire

```
   7 8 9
   6 1 2
   5 4 3
```

Here is another example: 4x4

```
  13 14 15 16
  12 03 04 05
  11 02 01 06
  10 09 08 07
```

You can find some greate video's on why Ulam's sprials are cool and how they relate to [primes](https://www.youtube.com/watch?v=iFuR97YcSLM) [here.](https://youtu.be/3K-12i0jclM?t=4m15s)

### Challenge 1

For this challenge create a program such that:

```
ulam -d 3
```

should generate a 3x3 sprial above:

```
   7 8 9
   6 1 2
   5 4 3
```
Or 

```
ulam -d 4
```

should generate a 4x4 sprial above:

```
  13 14 15 16
  12 03 04 05
  11 02 01 06
  10 09 08 07
```

### Challenge 2

Using a library like termbox-go highlight the prime numbers in a different colors.

### Challenge 3

Add another paramater that allow the person to give the starting number, if that number is not provided default to 1.

```
ulam -s 41 -d 3
```

should generate a 3x3 sprial above:

```
   47 48 49
   46 41 42
   45 44 43
```

Here is a video of someone implementing this using the [J programming language.](https://youtu.be/dBC5vnwf6Zw)

[ulam]: https://en.wikipedia.org/wiki/Stanislaw_Ulam "Wikipedia page on Ulam"


