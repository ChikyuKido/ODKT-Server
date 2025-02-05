# Documentation for helper Package
## Card Types

### Chance Card

#### Settings
- **Types**: Represented as a string, separated by commas, stored in `card_chance`.

#### Cards
Each chance card consists of:
- **Type**: Determines the card's behavior. Possible values:
  - `pay`
  - `receive`
  - `land-tax`
  - `move`
  - `house-repair`
  - `hold`
- **Text**: A string containing the card's description.
- **Payload**: Additional data encoded based on the type:
  - `pay`, `receive`, `land-tax`: 4-byte signed integer (int32)
  - `move`: String representing the destination field
  - `house-repair`: First 4 bytes are a signed integer for house value, second 4 bytes are a signed integer for hotel value
  - `hold`: String representing the action

---

### Bank Card

#### Settings
- **Types**: Represented as a string, separated by commas, stored in `card_bank`.

#### Cards
Each bank card consists of:
- **Type**: Determines the card's behavior. Possible values:
  - `pay`
  - `receive`
  - `move`
- **Text**: A string containing the card's description.
- **Payload**:
  - `pay`, `receive`: 4-byte signed integer (int32)
  - `move`: String representing the destination field

---

### Special Card

#### Settings
Stored in `card_special` as a concatenated byte array.
- **Start**: 4-byte signed integer
- **Multiplier**: 4-byte signed integer
- **Text**: String description

#### Cards
Each special card consists of:
- **Name**: Name of the special card
- **Price**: Integer value representing the price
- **Price Name**: Name associated with the price

---

### Street Card

#### Cards
Each street card consists of:
- **City**: Name of the city
- **Name**: Name of the street
- **Rent**: List of integer values, stored as a comma-separated string
- **Price**: Integer representing the base price
- **House Price**: Integer representing the cost of a house
- **Hotel Price**: Integer representing the cost of a hotel
---

### Railroad Card

#### Settings
Stored in `card_railroad` as a concatenated byte array.
- **Start**: 4-byte signed integer
- **Multiplier**: 4-byte signed integer
- **Text**: String description

#### Cards
Each railroad card consists of:
- **Name**: Name of the railroad
- **Price**: Integer value representing the price

---

### Other Fields Card

#### Cards
Each other-fields card consists of:
- **Name**: Name of the card
- **Payload** (only for `vermögensabgabe` type):
  - **Percent**: 4-byte signed integer
  - **Maximum**: 4-byte signed integer

---

## Summary
This document outlines the structure of the different card types imported into the database, including their settings, attributes, and payload structures. Each card type has its own format and specific encoding requirements for storing additional data.

