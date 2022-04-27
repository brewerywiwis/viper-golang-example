# THIS IS EXAMPLE OF VIPER GOLANG

<p>This repository is demonstrating how to use environment variables and yaml config together.</p>

<p>From Viper documentation, Viper uses the following precedence order. Each item takes precedence over the item below it:</p>

<ol>
<li>explicit call to Set</li>
<li>flag</li>
<li><font color="red">env</font></li>
<li><font color="red">config</font></li>
<li>key/value store</li>
<li>default</li>
</ol>

<p>In main.go, we can set environment variables on a machine and using it with the higher precedence over yaml config file. If viper cannot map the environment variables it will look for yaml config file instead.</p>

**Be careful of data loss in case both env and yaml config file is not set.**
