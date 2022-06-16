# ProgLog

Go example commit log service.

A commit log is a sequence of records where each record has its own unique identifier. Records and only be appended at the end of a commit log. Records are immutable once written.

Reading a commit log happens from left to right. There's no querying so offsets are used to specify the start and end points of the read.

E.g. 
Kafka
Individual records are identified using <topic, partition, offset>.

Topic: movies
============= 0      1      2      3      4      ...
Partion 0     video  video  video
Partion 1     video  video  
Partion 2     video  video  video  video


Commit logs provide a source of truth for what occurred in a system and in what order.