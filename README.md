# Task: Flea circus

A 30×30 grid of squares contains 900 fleas, initially one flea per square.
When a bell is rung, each flea jumps to an adjacent square at random (usually 4 possibilities, except for fleas on the edge of the grid or at the corners).

What is the expected number of unoccupied squares after 50 rings of the bell? Give your answer rounded to six decimal places.

Implementation plan

 - Implement single simulation
 - Run multiple simulations and calculate average
 - Run simulations in parallel
 - Optimize for better speed, less memory allocations

## Usage

Build the project first

```sh
go build -o fleas
```

Run single simulation
```sh
./fleas --times=1
```

Run multple simulations in sequence
```sh
./fleas --times=10
```

Run multple simulations in parallel
```sh
./fleas --times=10 --parallel=true
```

Run `./fleas --help` to see details:
```
Usage of ./fleas:
  -display-grid
        Determines if grid matrix should be displayed in termial (Supported with simulation in sequence only).
  -parallel
        Determines if simulation should run in parallel.
  -profile
        Determines if simulation should run with profiler enabled.
  -times int
        Number of times to run simulation. (default 1)
```

NOTE: after simulation was run `debug.log` file is created with resources usage details


### Performance optimization
According to profiler most expensive function call is function related to rand numbers (see profiler graph below)
[![profile-before.png](https://i.postimg.cc/MG6RBZ8y/profile-before.png)](https://postimg.cc/5Xr6dVy0)

This is profiler graph after randomizer optimization
[![profile-after.png](https://i.postimg.cc/hG2ywP8r/profile-after.png)](https://postimg.cc/Z9yxvmMB)

