* [Buying](#buying)
* [Selling](#selling)

<div id='buying'></div>

## Buying

###### It is the operation of acquiring shares and updating the portfolio state.

### Request

Each `buy` operation in the input list must contain the following fields:

| Field       |  Type   | Description                                  | Constraints                             | Required |
|:------------|:-------:|:---------------------------------------------|:----------------------------------------|:--------:|
| `operation` | String  | Type of the operation.                       | Must be exactly `"buy"`.                |   Yes    |
| `unit-cost` | Decimal | Unit price paid per share.                   | Positive decimal value (e.g., `10.00`). |   Yes    |
| `quantity`  | Integer | Number of shares purchased in the operation. | Positive integer (e.g., `1000`).        |   Yes    |

```json
{
  "operation": "buy",
  "unit-cost": 10.00,
  "quantity": 1000
}
```

### Response

> Note: `buy` operations do not have an explicit response payload.  
> They only update the internal portfolio state (quantity and weighted-average unit cost) and do not generate any tax.

---

<div id='selling'></div>

## Selling

###### It is the operation of disposing shares, realizing profit or loss, and possibly generating tax.

### Request

Each `sell` operation in the input list must contain the following fields:

| Field       |  Type   | Description                                 | Constraints                                          | Required |
|:------------|:-------:|:--------------------------------------------|:-----------------------------------------------------|:--------:|
| `operation` | String  | Type of the operation.                      | Must be exactly `"sell"`.                            |   Yes    |
| `unit-cost` | Decimal | Unit price received per share (sale price). | Positive decimal value (e.g., `15.00`).              |   Yes    |
| `quantity`  | Integer | Number of shares sold in the operation.     | Positive integer, not greater than shares available. |   Yes    |

```json
{
  "operation": "sell",
  "unit-cost": 15.00,
  "quantity": 500
}
```

### Response

For each `sell` operation, the output list contains an object with a single field:

| Field |  Type   | Description                              | Constraints                                                                                                            | Required |
|:------|:-------:|:-----------------------------------------|:-----------------------------------------------------------------------------------------------------------------------|:--------:|
| `tax` | Decimal | Tax amount calculated for the operation. | Decimal value greater than or equal to `0.00`. Calculated at 20% over taxable profit, after loss offset and threshold. |   Yes    |

```json
{
  "tax": 1000.00
}
```
