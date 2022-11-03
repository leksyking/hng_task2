### HNG Stage-2 Task

- Create an (POST) api endoint that takes the following sample json:
```
{ 
    “operation_type”: Enum <addition | subtraction | multiplication>,
    “x”: Integer,
    “y”: Integer,
}
```
- And return 
```
{ 
    “slackUsername”: String,
    "operation_type" : Enum.value,
    “result”: Integer 
}
```
