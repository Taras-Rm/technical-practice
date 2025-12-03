# Expense Tracker CLI

Expense Tracker CLI is a simple command line application that helps you manage your personal finances.  
With this tool, you can easily add, update, delete, and view your expenses, as well as generate summaries for better budgeting.

---

## 🚀 Features

The Expense Tracker CLI supports the following core functionality:

### **Core Features**
- Add a new expense with a description, category and amount  
- Update an existing expense  
- Delete an expense by ID  
- View all expenses  
- View a total summary of all expenses  
- View expenses summary for a specific month of the current year  

---

## 📌 Requirements

Your CLI application must:

- Run from the terminal  
- Accept user commands and flags as arguments  
- Store expense data in a simple text file (JSON, CSV, or similar)  
- Persist data between runs

---

## 🧪 Example Usage

### **Add expenses**
```bash
$ go run main.go add --description "Lunch" --amount 20 --category food
# Expense added successfully (ID: 1)

$ go run main.go update --id a0302ff0-9ec2-421e-bba3-2cde5f9010d4 --description "Lunch" --amount 20 --category food
# Expense updated successfully (ID: a0302ff0-9ec2-421e-bba3-2cde5f9010d4)

$ go run main.go delete --id a0302ff0-9ec2-421e-bba3-2cde5f9010d4
# Expense deleted successfully (ID: a0302ff0-9ec2-421e-bba3-2cde5f9010d4)

$ go run main.go summary
# Total expenses: $3233234

$ go run main.go summary --month 3
# Total expenses: $12


$ go run main.go list
