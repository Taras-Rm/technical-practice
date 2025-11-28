# Task Tracker CLI

Task Tracker is a simple command line interface (CLI) application used to track and manage your tasks.

With this tool, you can:

- Add new tasks
- Update and delete existing tasks
- Mark tasks as **todo**, **in-progress**, or **done**
- List tasks by different statuses

---

## Features

The Task Tracker CLI supports the following actions:

- **Add a task**
- **Update a task**
- **Delete a task**
- **Mark a task as in progress**
- **Mark a task as done**
- **List all tasks**
- **List tasks filtered by status**:
  - `todo`
  - `in-progress`
  - `done`

---

## Example Usage

Below is an example of how the CLI might be used from the terminal:

```bash
# Adding a new task
./task-tracker add "Buy groceries"

# Updating a task
./task-tracker update 1 "Buy groceries and cook dinner"

# Deleting a task
./task-tracker delete 1

# Marking a task as in progress
./task-tracker mark-in-progress 1

# Marking a task as done
./task-tracker mark-done 1

# Listing all tasks
./task-tracker list

# Listing tasks by status
./task-tracker list done
./task-tracker list todo
./task-tracker list in-progress
