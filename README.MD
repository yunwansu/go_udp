go lang UDP 예제입니다.

go lang UDP Example

i'm installed logstash 6.1.3 using APT on ubuntu 16.04

follow this url(logstash 6.x getting start)
https://www.elastic.co/guide/en/logstash/6.x/installing-logstash.html

and then make pipline.conf file

        input {
                udp {
                        port => "4885"
                }
        }
        output {
                stdout { codec => rubydebug }
        }


run this command on ubuntu 16.04

        sudo bin/logstash -f /etc/logstash/first-pipeline.conf --config.reload.automatic


result image 
every second send from go program with udp to logstash

![alt text](https://github.com/yunwansu/go_udp/blob/master/result.png)
