go lang UDP 예제입니다.

go lang UDP Example

i'm installed logstash 6.1.3 using APT

follow this url(logstash 6.x getting start)
https://www.elastic.co/guide/en/logstash/6.x/installing-logstash.html

and then make pipline.conf file

<code>
input {
#       beats {
#               port => "5044"
#       }

        udp {
                port => "4885"
        }
}

# The filter part of this file is commented out to indicate that it is
# optional.
# filter {
#
# }

output {
        stdout { codec => rubydebug }
}

</code>


use this command on ubuntu 16.04

sudo bin/logstash -f /etc/logstash/first-pipeline.conf --config.reload.automatic
