# BargeDB [In progress]

Specification: https://easy-river-387.notion.site/BargeDB-86b73c3a75f84347b43f0859b2d30f84

Consensus layer based on RAFT with pluggable storage and transport interface. Default storage is just go maps and default transport is gRPC. 
Reference: https://raft.github.io/raft.pdf
![image](./barge.png)

- BargeDB will support [USE ](https://github.com/utkarsh-pro/use) as one of the storage options
