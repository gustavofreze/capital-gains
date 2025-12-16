# Capital gains

[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

* [Overview](#overview)
    - [Use cases](#use_cases)
* [Installation](#installation)
    - [Repository](#repository)
    - [Configuration](#configuration)
    - [Tests](#tests)
    - [Review](#review)
    - [Reports](#reports)

<!--suppress HtmlDeprecatedAttribute -->

<div id="overview"></div> 

## Overview

Command-line (CLI) application that calculates the tax owed on profits or losses from stock market transactions.

The application processes sequences of buy and sell operations, maintaining an in-memory portfolio with the current
share quantity, weighted-average unit cost, and accumulated losses.

<div id='use_cases'></div> 

### Use cases

- [Register buy](docs/USE_CASES.md#register-buy)
- [Register sell](docs/USE_CASES.md#register-sell)
- [Calculate capital gain](docs/USE_CASES.md#calculate-capital-gain)

<div id='installation'></div> 

## Installation

<div id='repository'></div> 

### Repository

To clone the repository using the command line, run:

```bash
git clone https://github.com/gustavofreze/capital-gains.git
```

<div id='configuration'></div> 

### Configuration

To install project dependencies locally, run:

```bash
make configure
```

### Execution

Run the CLI reading operations from `stdin` and writing the JSON result to `stdout`:

```bash
make calculate < use_case.txt
```

You can also paste input directly:

```bash
echo '[{"operation":"buy","unit-cost":10.00,"quantity":100},{"operation":"sell","unit-cost":12.00,"quantity":50}]' | make calculate
```

or

```bash
make calculate << 'EOF'
[{"operation":"buy","unit-cost":10.00,"quantity":100},{"operation":"sell","unit-cost":12.00,"quantity":50}]
EOF
```

<div id='tests'></div> 

### Tests

Run all tests with coverage:

```bash
make test 
```

<div id='review'></div> 

### Review

Run static code analysis:

```bash
make review 
```

<div id='reports'></div> 

### Reports

Open static analysis reports (e.g., coverage, lints) in the browser:

```bash
make show-reports 
```

> You can check other available commands by running `make help`.
