* [Register buy](#register_buy)
* [Register sell](#register_sell)
* [Calculate capital gain](#calculate_capital_gain)
* [FAQ](#faq)
* [Reference](#reference)

<div id='register_buy'></div>

## Register buy

###### It registers a **buy** operation, updating the portfolio state (share quantity and weighted-average unit cost).

### Explanation

Represents a single `buy` operation from the input list.  
It increases the current position quantity and recalculates the weighted-average unit cost (WAC).

> Common rules (WAC formula, rounding) are listed in the [FAQ](#faq).

### Request

| Field       |  Type   | Description                                  | Constraints                             | Required |
|:------------|:-------:|:---------------------------------------------|:----------------------------------------|:--------:|
| `operation` | String  | Type of the operation.                       | Must be exactly `"buy"`.                |   Yes    |
| `unit-cost` | Decimal | Unit price paid per share.                   | Positive decimal value (e.g., `10.00`). |   Yes    |
| `quantity`  | Integer | Number of shares purchased in the operation. | Positive integer (e.g., `1000`).        |   Yes    |

Example as it appears inside an input line:

```json
[
  {
    "operation": "buy",
    "unit-cost": 10.00,
    "quantity": 1000
  }
]
```

### Response

Register buy does not generate tax by itself.

> Note: In calculate capital gain, the corresponding output element for a buy operation is always:
>
> ```json
> {"tax":0.00}
> ```


<div id='register_sell'></div>

## Register sell

###### It registers a **sell** operation, realizing profit or loss and possibly generating tax.

### Explanation

Represents a single `sell` operation from the input list.  
It decreases the current position quantity and may produce profit or loss for the operation.

Whether tax is due depends on the exemption threshold and on the final profit after applying accumulated losses.

> Common rules (threshold, loss carryforward, tax rate, rounding) are listed in the [FAQ](#faq).

### Request

| Field       |  Type   | Description                                 | Constraints                                          | Required |
|:------------|:-------:|:--------------------------------------------|:-----------------------------------------------------|:--------:|
| `operation` | String  | Type of the operation.                      | Must be exactly `"sell"`.                            |   Yes    |
| `unit-cost` | Decimal | Unit price received per share (sale price). | Positive decimal value (e.g., `15.00`).              |   Yes    |
| `quantity`  | Integer | Number of shares sold in the operation.     | Positive integer, not greater than shares available. |   Yes    |

Example as it appears inside an input line:

```json
[
  {
    "operation": "sell",
    "unit-cost": 15.00,
    "quantity": 500
  }
]
```

### Response

Register sell does not print output by itself.  
In calculate capital gain, the corresponding output element for a sell operation is:

| Field |  Type   | Description                              | Constraints                                        | Required |
|:------|:-------:|:-----------------------------------------|:---------------------------------------------------|:--------:|
| `tax` | Decimal | Tax amount calculated for the operation. | Decimal value **greater than or equal to** `0.00`. |   Yes    |

Example element:

```json
{
  "tax": 1000.00
}
```

---

<div id='calculate_capital_gain'></div>

## Calculate capital gain

###### It reads a list of stock operations from `stdin`, runs an independent simulation per input line, and prints the tax for each operation to `stdout`.

### Explanation

- The program reads **one JSON array per line** from standard input.
- Each line is processed as an **independent simulation** (state is not shared across lines).
- For each input line, the program prints **one JSON array** with the computed `tax` for each operation.

> Common rules (tax logic, threshold, WAC, loss carryforward, rounding) are listed in the [FAQ](#faq).

See [Execution](../README.md#execution) for how to run the program using `make calculate`.

### Request

Each input line must be a JSON array where every operation contains the following fields:

| Field       |  Type   | Description                               | Constraints                             | Required |
|:------------|:-------:|:------------------------------------------|:----------------------------------------|:--------:|
| `operation` | String  | Type of the operation.                    | Must be exactly `"buy"` or `"sell"`.    |   Yes    |
| `unit-cost` | Decimal | Unit price per share (2 decimal places).  | Positive decimal value (e.g., `10.00`). |   Yes    |
| `quantity`  | Integer | Number of shares traded in the operation. | Positive integer (e.g., `1000`).        |   Yes    |

Example (single input line):

```json
[
  {
    "operation": "buy",
    "unit-cost": 10.00,
    "quantity": 10000
  },
  {
    "operation": "sell",
    "unit-cost": 20.00,
    "quantity": 5000
  },
  {
    "operation": "sell",
    "unit-cost": 5.00,
    "quantity": 5000
  }
]
```

### Response

For each input line, the program outputs a JSON array with the **same length** as the input operation list.  
Each element contains a single field:

| Field |  Type   | Description                              | Constraints                                        | Required |
|:------|:-------:|:-----------------------------------------|:---------------------------------------------------|:--------:|
| `tax` | Decimal | Tax amount calculated for the operation. | Decimal value **greater than or equal to** `0.00`. |   Yes    |

Example (output line for the input above):

```json
[
  {
    "tax": 0.00
  },
  {
    "tax": 10000.00
  },
  {
    "tax": 0.00
  }
]
```

<div id='faq'></div>

## FAQ

### What is the expected input format?

- One JSON array per line, where each element is an operation object.
- The last line of the input is an empty line.

### Is the portfolio state shared across input lines?

No. Each input line is an independent simulation. The portfolio and accumulated losses must be reset for every new line.

### How is the weighted-average unit cost (WAC) calculated?

On every buy, recalculate:

`new-weighted-average-unit-cost = ((current-share-quantity * current-weighted-average-unit-cost) + (buy-share-quantity * buy-unit-cost)) / (current-share-quantity + buy-share-quantity)`

### When is tax due?

- Tax is calculated only for **sale** operations that produce **profit** (sell price greater than WAC).
- The tax rate is **20%** over the final taxable profit (after deducting accumulated losses).

### What is the exemption threshold?

If the **total sell operation value** (`unit-cost * quantity`) is **less than or equal to 20000.00**, then **no tax is
paid** for that sell
operation.

Important:

- The exemption is based on the operation value, not on profit.
- Losses still must be accumulated and used to offset future profits, even for exempt sells.

### How does loss carryforward work?

Losses from sells must be accumulated and used to offset future profits until fully deducted.

### Are buys taxed?

No. Buy operations always have tax `0.00`.

### How are decimals handled?

- Monetary values are rounded to **2 decimal places**.
- Output `tax` must be a non-negative decimal.

### Can I assume the input is valid?

Yes. You may assume input lines follow the contract and will not break parsing.

### Can the program print other messages?

No. The program must print only the JSON outputs (no prompts, logs, or explanatory text).

<div id='reference'></div>

## Reference

[Code Challenge: Capital Gains](code_challenge_ptbr.pdf)
