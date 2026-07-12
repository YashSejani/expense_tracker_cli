# Expense Tracker CLI

A lightweight, performant Command-Line Interface (CLI) application built in Go for managing and tracking daily expenses. It uses local JSON file storage to persist records across sessions without external database dependencies.

---

## Features

- **Add Expenses:** Log individual expenses with auto-incrementing unique IDs and automatic timestamps.
- **List Expenses:** View a structured, tabular chronological display of all tracked expenses.
- **Delete Expenses:** Safely remove specific records using their unique ID with built-in validation.
- **Aggregated Summaries:** Calculate financial statistics across all inputs, with optional filtering by specific months or years.
- **Self-Contained File Storage:** Reads and writes directly to a local structured JSON database schema.

---

## Getting Started

### Prerequisites

- [Go](https://go.dev/doc/install) (version **1.21** or higher recommended due to `slices` optimization usage)

### Installation & Compilation

Clone this repository and compile the binary executable using the following terminal commands:

```bash
# Compile the project into a single binary
go build -o expense-tracker .

# (Optional) Move the binary to your local system path to run it globally
# mv expense-tracker /usr/local/bin/
```

---

## Usage Guide & Command Layout

Run the application using the subcommands detailed below. If a parameter is missed or incorrectly typed, the engine intercepts the execution and yields a global help context overview.

### 1. Log an Expense (`add`)

Appends a new expense entry. Both flags are required.

```bash
./expense-tracker add --description "Lunch at cafe" --amount 20

# Output:
Expense added successfully (ID: 1)
```

---

### 2. View All Entries (`list`)

Outputs a clear, tab-formatted visual dashboard index of your transaction history.

```bash
./expense-tracker list

# Output:
ID    Date            Description     Amount
-----------------------------------------------
1     2026-07-13      Lunch at cafe   $20
```

---

### 3. Compute Financial Totals (`summary`)

Retrieves calculated expense summaries. You can check the complete lifetime total, filter down to a specific month, or specify target calendar periods.

#### Global lifetime aggregation

```bash
./expense-tracker summary

# Output:
Total expenses: $20
```

#### Filtered down to a specific month (1–12)

```bash
./expense-tracker summary --month 7

# Output:
Total expenses for July: $20
```

#### Filtered by month and specific year

```bash
./expense-tracker summary --month 12 --year 2026
```

---

### 4. Remove an Entry (`delete`)

Safely purges data records from your local storage array map matching the designated identification criteria flag.

```bash
./expense-tracker delete --id 1

# Output:
Expense deleted successfully (ID: 1)
```

---

### 5. Access Interactive Help Menu (`help`)

Displays explicit syntactic instructions, optional parameters constraints, and input patterns.

```bash
./expense-tracker help
```

---

## Project Structure & Architecture

The application separates concerns cleanly across four modular implementation domains:

### `main.go`

Application entry point. Handles routing logic, sub-command CLI flag definitions (`flag.NewFlagSet`), and argument data validation.

### `commands.go`

Houses core logical business processes (Add, List, Delete, Summary) attached as pointer method receivers.

### `models.go`

Defines the structured application objects schema data layouts (`Expense`, `StorageSchema`, `ExpenseManager`).

### `storage.go`

Manages file input/output workflows (`json.Unmarshal`, `json.MarshalIndent`) to maintain structural integrity under `expense.json`.


---

## Acknowledgements

This project was built as a solution to the **Expense Tracker** project from roadmap.sh.

🔗 **Project Specification:**  
https://roadmap.sh/projects/expense-tracker
