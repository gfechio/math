# Bifurcation Diagram in Go

This project generates a bifurcation diagram for the logistic map, a mathematical equation that exhibits chaotic behavior. The bifurcation diagram visually represents the different states the logistic map can be in based on the value of a parameter 'r'.

## Problem definition

- Z(n+1) = Zn^2 + C
- X(n+1) = RXn(1-Xn)
- Feigenbaum Constant ( 4.669... )

## Prerequisites

You need to have Go installed on your machine. You can download it from the [official Go website](https://golang.org/dl/).

## Installation


Install the required Go package `gonum/plot`:

```
go get -u gonum.org/v1/plot
```

## Running the Program

To run the program, use the following command:

```
go run main.go
```

This will generate a PNG image file named "bifurcation_diagram.png" in the project directory, containing the bifurcation diagram.

## Understanding the Bifurcation Diagram

The bifurcation diagram shows the possible stable states (fixed points or periodic orbits) of the logistic map for different values of the parameter 'r'. As 'r' increases, the system undergoes bifurcations, where a stable state splits into two. This is a characteristic behavior of chaotic systems.

## License

This project is open source and available under the [MIT License](LICENSE).


