# Task Tracker CLI

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A simple command-line interface for managing your tasks, built as a solution to the [Task Tracker](https://roadmap.sh/projects/task-tracker) project on [roadmap.sh](https://roadmap.sh).

## Features

- **Add tasks:** Easily add new tasks with descriptions.
- **List tasks:** View all your pending tasks.
- **Mark tasks as complete:** Keep track of your progress.
- **Delete tasks:** Remove tasks you no longer need.
- **Data persistence:** Stores tasks in a JSON file for future reference.

## How to Run

```bash
git clone https://github.com/radityajay/task-tracker-cli.git
cd task-tracker-cli
```
Run the following command to build and run the project:

```bash
go build -o task-tracker

# To add a task
./task-tracker add "Buy groceries"

# To update a task
./task-tracker update 1 "Buy groceries and cook dinner"

# To delete a task
./task-tracker delete 1

# To mark a task as in progress/done/todo
./task-tracker mark-in-progress 1
./task-tracker mark-done 1
./task-tracker mark-todo 1

# To list all tasks
./task-tracker list
./task-tracker list done
./task-tracker list todo
./task-tracker list in-progress
```