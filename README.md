# BargeDB

A lightweight SDK for distributed go applications to form RAFT consensus based distributed persistent key value store. 
Each instance of your application GETS and SETS data through the exposed methods of BargeDB instance and BargeDB internally handles the consistency across all instances via gRPC

![image](./barge.png)

- BargeDB will support [USE ](https://github.com/utkarsh-pro/use) as one of the storage options