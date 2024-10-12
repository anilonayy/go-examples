# Expense Tracker

### build
```bash
go build -o expense-tracker
```

### make it workable from console
``` bash
sudo mv expense-tracker /usr/local/bin/ 
```

### help
```bash
expense-tracker -h
```

### add expense
```bash
expense-tracker add  -desc "Lunch" -price 10
# Expense added successfully (ID:1728770425)
```

### add expense
```bash
expense-tracker add  -desc "Dinner" -price 20
# Expense added successfully (ID:1728770426)
```

### list expenses
```bash
expense-tracker list
# ID              Description     Price   Date
# 1728770409      Lunch           20.00   2024-10-13T01:00:08+03:00
# 1728770421      Lunch           20.00   2024-10-13T01:00:19+03:00
# 1728770425      Lunch           20.00   2024-10-13T01:00:22+03:00
```

### delete expense
```bash
expense-tracker delete -id 1728770425
# Expense deleted successfully
```

### expense summary
```bash
expense-tracker summary
# Total expenses(3): $60
```

### expense summary with month filter
```bash
expense-tracker summary --month 8
# Total expenses(2): $40
```
