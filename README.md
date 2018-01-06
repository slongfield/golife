# golife
[![Build Status](https://travis-ci.org/slongfield/golife.svg?branch=master)](https://travis-ci.org/slongfield/golife)
[Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life)

A slight modification on the traditional game of life--when a cell dies, it
leaves behind a record, which decays over the course of the next few iterations.

This comes with two rendering engines, one which emits a simple ASCII rendering,
(cmd/gol-text), and another which uses termbox to emit a colorful rendering at
the full size of the terminal.

Requires Go 1.7+ for testing.Run
