# Port Proxy

Provides local mapping from local ports to remote ports

## How to use

The utility maps between local ports and remote ports using property files.
Define your local ports in the local.properties file and name them at the process.
Define your report host and port in the remote.properties file and use the same name
you used for the local port.

The utility maps from local ports to remote hosts using the names in the two files.


##### Sample local.properties file:

```
memcached.host=6601
zookeeper.host=6602
kafka.host=6603
mysql.host=6604
cassandra.cql.host=6605
cassandra.thrift.host=6606
trc001.host=6607
trc001.debug=6608
mysql001.host=6609
vertica001.host=6610
cassandra_cql001.host=6611
```

##### Sample remote.properties file:

```
memcached.host=qa-nmd00023.taboolasyndication.com:28070
zookeeper.host=qa-nmd00003.taboolasyndication.com:23565
kafka.host=qa-nmd00022.taboolasyndication.com:26653
mysql.host=qa-nmd00021.taboolasyndication.com:23472
cassandra.cql.host=qa-nmd00028.taboolasyndication.com:26852
cassandra.thrift.host=qa-nmd00028.taboolasyndication.com:21997
trc001.host=qa-nmd00013.taboolasyndication.com:21306
trc001.debug=qa-nmd00013.taboolasyndication.com:30884
mysql001.host=qa-nmd00021.taboolasyndication.com:23472
vertica001.host=qa-nmd00008.taboolasyndication.com:23696
cassandra_cql001.host=qa-nmd00028.taboolasyndication.com:21997
```

##### The files above create the following mappings:
```
localhost:6601 ==> qa-nmd00023.taboolasyndication.com:28070
localhost:6602 ==> qa-nmd00003.taboolasyndication.com:23565
localhost:6603 ==> qa-nmd00022.taboolasyndication.com:26653
localhost:6604 ==> qa-nmd00021.taboolasyndication.com:23472
localhost:6605 ==> qa-nmd00028.taboolasyndication.com:26852
localhost:6606 ==> qa-nmd00028.taboolasyndication.com:21997
localhost:6607 ==> qa-nmd00013.taboolasyndication.com:21306
localhost:6608 ==> qa-nmd00013.taboolasyndication.com:30884
localhost:6609 ==> qa-nmd00021.taboolasyndication.com:23472
localhost:6610 ==> qa-nmd00008.taboolasyndication.com:23696
localhost:6611 ==> qa-nmd00028.taboolasyndication.com:21997
```

#### How to run the code:

1. Make sure you have go 112 installed
2. Run ```git clone https://github.com/guypeled76/portproxy.git```
3. cd into 'portproxy' directory
4. Change the content of the local.properties / remote.properties files to feet your mappings
5. run the ./portproxy.sh bash script your terminal 

##### After running the bash script you will get the following output:

```
2019/07/09 12:06:03 Initializing port proxy based on local.properties to remote.properties.
2019/07/09 12:06:03 Creating 'memcached.host' proxy from localhost:6601 to qa-nmd00023.taboolasyndication.com:28070.
2019/07/09 12:06:03 Creating 'cassandra_cql001.host' proxy from localhost:6611 to qa-nmd00028.taboolasyndication.com:21997.
2019/07/09 12:06:03 Creating 'kafka.host' proxy from localhost:6603 to qa-nmd00022.taboolasyndication.com:26653.
2019/07/09 12:06:03 Creating 'mysql001.host' proxy from localhost:6609 to qa-nmd00021.taboolasyndication.com:23472.
2019/07/09 12:06:03 Creating 'cassandra.cql.host' proxy from localhost:6605 to qa-nmd00028.taboolasyndication.com:26852.
2019/07/09 12:06:03 Creating 'cassandra.thrift.host' proxy from localhost:6606 to qa-nmd00028.taboolasyndication.com:21997.
2019/07/09 12:06:03 Creating 'zookeeper.host' proxy from localhost:6602 to qa-nmd00003.taboolasyndication.com:23565.
2019/07/09 12:06:03 Creating 'mysql.host' proxy from localhost:6604 to qa-nmd00021.taboolasyndication.com:23472.
2019/07/09 12:06:03 Creating 'trc001.debug' proxy from localhost:6608 to qa-nmd00013.taboolasyndication.com:30884.
2019/07/09 12:06:03 Creating 'vertica001.host' proxy from localhost:6610 to qa-nmd00008.taboolasyndication.com:23696.
2019/07/09 12:06:03 Creating 'trc001.host' proxy from localhost:6607 to qa-nmd00013.taboolasyndication.com:21306.
```

