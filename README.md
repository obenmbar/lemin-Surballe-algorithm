# Lem-in   
<a id="readme-top"></a>

[](data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNTguNzUiIGhlaWdodD0iMzUiIHZpZXdCb3g9IjAgMCAxNTguNzUgMzUiPjxyZWN0IHdpZHRoPSI4NC4yMTY2NjcxNzUyOTI5NyIgaGVpZ2h0PSIzNSIgZmlsbD0iIzRhOTBlMiIvPjxyZWN0IHg9Ijg0LjIxNjY2NzE3NTI5Mjk3IiB3aWR0aD0iNzQuNTMzMzMyODI0NzA3MDMiIGhlaWdodD0iMzUiIGZpbGw9IiNjY2Q4ZTciLz48dGV4dCB4PSI0Mi4xMDgzMzM1ODc2NDY0ODQiIHk9IjIxLjUiIGZvbnQtc2l6ZT0iMTIiIGZvbnQtZmFtaWx5PSInUm9ib3RvJywgc2Fucy1zZXJpZiIgZmlsbD0iI0ZGRkZGRiIgdGV4dC1hbmNob3I9Im1pZGRsZSIgbGV0dGVyLXNwYWNpbmc9IjIiPkdPTEFORzwvdGV4dD48dGV4dCB4PSIxMjEuNDgzMzMzNTg3NjQ2NDgiIHk9IjIxLjUiIGZvbnQtc2l6ZT0iMTIiIGZvbnQtZmFtaWx5PSInTW9udHNlcnJhdCcsIHNhbnMtc2VyaWYiIGZpbGw9IiMwZTBlMGUiIHRleHQtYW5jaG9yPSJtaWRkbGUiIGZvbnQtd2VpZ2h0PSI5MDAiIGxldHRlci1zcGFjaW5nPSIyIj4xLjIyLjM8L3RleHQ+PC9zdmc+)


<a href="https://gitpoint.co/">
    <img alt="GitPoint" title="GitPoint" src="https://cdn.creazilla.com/cliparts/24564/1524593971-clipart-xl.png" width="350">
  </a>

## Description

This project simulates an ant colony navigating through interconnected rooms — from a start point to an exit — using optimized pathfinding.

The objective is to move all ants to the end room in the minimum number of turns, while respecting movement constraints.

**Repository**: [lem-in](https://learn.zone01oujda.ma/git/obenmbar/lem-in)






## Features

- Efficient pathfinding using Dijkstra’s and Suurballe’s algorithms
- Handles multiple disjoint paths and optimizes ant distribution
- Detects and handles invalid input or malformed graphs
- Minimizes total turns for all ants
- Clean modular structure for clarity and testing

## Project Structure

```
lem-in/
├── functions/
│   ├── Check_helpers.go   # Helper functions for validating rooms and tunnels
│   ├── Check_utils.go     # Utility checkers for input and syntax rules
│   ├── Solve_helpers.go   # Helper utilities for building and linking paths
│   ├── algo.go            # Core algorithms: BFS, Dijkstra, and path building
│   ├── check.go           # Main parser and validator for input data
│   ├── moove.go           # Handles ant movement turn-by-turn
│   ├── suurballe.go       # Implements Suurballe's algorithm for path merging
│   └── README.md          # Project documentation
├── go.mod                 # Go module configuration
└── main.go                # Entry point
```

## Algorithm

The project implements:
- BFS / Dijkstra — finds shortest paths through the graph
- Suurballe’s algorithm — builds multiple edge-disjoint paths
- Ant Distribution — assigns ants across paths to minimize total turns

## Installation

```bash
git clone https://learn.zone01oujda.ma/git/obenmbar/lem-in.git
cd lem-in
```

## Usage

```bash
go run main.go test_file.txt
```

### Input Format

```
number_of_ants
##start
start_room x y
room1 x y
room2 x y
##end
end_room x y
start_room-room1
room1-room2
room2-end_room
```

### Output Format

```
[Input content]

L1-room1 L2-start_room
L1-room2 L2-room1 L3-start_room
L1-end_room L2-room2 L3-room1
...
```

Each Lx-roomy shows ant x moving to room y.

### Example : 
```bash
$ go run . level-2.txt
4
##start
start 0 0
1 1 0
2 2 0
##end
end 3 0
start-1
start-2
1-end
2-end

L1-1 L2-2
L1-end L2-end L3-1 L4-2
L3-end L4-end
```

## Rules

- Ants begin in the ##start room and must reach ##end
- One ant per room at a time (except start and end)
- Each turn, ants may move one room forward
- The goal is to minimize total turns

## Instructions

- We must use Go's standard library only
- The project must compile without errors
- Imports must be formatted correctly
- Functions, variables, constants, types must follow good practices
- Code must be well-structured and properly tested
- It is recommended to have test files with `_test` suffix for unit testing

## Allowed Packages

Only the following Go standard library packages are allowed:
- `os`
- `fmt`
- `strconv`
- `strings`
- `math`

## Learning Objectives

This project helps us learn about:
- Graph theory and pathfinding algorithms
- Optimizing multiple path traversal
- Input parsing and validation
- Data structure efficiency
- Algorithmic problem solving
---



## Authors

- **Othmane Benmbarek** - [obenmbar](https://learn.zone01oujda.ma/git/obenmbar)
- **Mohamed Nouri** - [mohnouri](https://learn.zone01oujda.ma/git/mohnouri)
- **BEMAMORY Nomenjanahary Luciano Loic** - [bnomenja](https://learn.zone01oujda.ma/git/bnomenja)


<p align="right">(<a href="#readme-top">back to top</a>)</p>

---