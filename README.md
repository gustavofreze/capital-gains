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

- [Buying](docs/USE_CASES.md#buying)
- [Selling](docs/USE_CASES.md#selling)

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
