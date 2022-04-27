**THIS IS EXAMPLE OF VIPER GOLANG**

This repository is demonstrating how to use environment variables and yaml config together.

Viper uses the following precedence order. Each item takes precedence over the item below it:

explicit call to Set
flag
**<font color="red">env</font>**
**<font color="red">config</font>**
key/value store
default

In main.go, we can set environment variables on machine and using it with the higher precedence over yaml config file. If viper cannot map the environment variables it will look for yaml config file instead.

**Be careful of data loss in case both env and yaml config file is not set.**
