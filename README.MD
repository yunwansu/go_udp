go lang UDP 예제입니다.

go lang UDP Example

I was installed logstash 6.1.3 using APT on ubuntu 16.04

follow this url(logstash 6.x getting start)
https://www.elastic.co/guide/en/logstash/6.x/installing-logstash.html

and installed too elasticsearch 6.1.3

this here url
https://www.elastic.co/guide/en/elasticsearch/reference/current/_installation.html

and then make pipline.conf file

        input {
                udp {
                        port => "4885"
                }
        }
        output {
                elasticsearch {
                	hosts => ["http://127.0.0.1:9200"]
       		}
        }


run this command on ubuntu 16.04

        sudo bin/logstash -f /etc/logstash/first-pipeline.conf --config.reload.automatic

	cd elasticsearch-6.1.3
 	./elasticsearch

result image 
every second send from go program with udp to logstash

![alt text](https://github.com/yunwansu/go_udp/blob/master/elasticsearch_result.png)
