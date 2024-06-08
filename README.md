# CVM Go

CVM algorithm implemented in Go based on paper [The CVM Algorithm for Estimating Distinct Elements in Streams](https://cs.stanford.edu/~knuth/papers/cvm-note.pdf).

## Build

Build application using bash script in ```scripts/``` directory with:

```bash
scripts/build.sh
```

This will create ```cmd/cvm-go/cvm-go``` executable binary file.

## Run

Run application by running compiled binary with:

```bash
cmd/cvm-go/cvm-go
```

## Testing

Running specific or all tests is easily available through VS Code editor in `Testing` tab.

To run all unit tests with coverage profiling, use script:

```bash
scripts/run_coverage.sh
```

Script above will try to automatically open coverage HTML report in your default browser.
