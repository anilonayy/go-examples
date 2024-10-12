# task-cli app

### build 
```bash
go build -o task-cli
```

### make it workable from console
``` bash
sudo mv task-cli /usr/local/bin/ 
```


### add task
```bash
# task-cli add <title>
task-cli add new2
```

### list task
```bash
task-cli list
```

### list task with filter
```bash
# task-cli list (todo,in-progress,done)
task-cli list todo
```

### update task title
```bash
# task-cli update <task-id> <title>
task-cli update 1 "New Title"
```

### update task status
```bash
# task-cli delete <task-id>
task-cli mark-todo 1
task-cli mark-in-progress 2
task-cli mark-done 3
```

### delete task
```bash
# task-cli delete <task-id>
task-cli delete 1
```