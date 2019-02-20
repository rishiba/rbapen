docker create network elk_network

sudo  docker run -d --name elasticsearch --net elk_network -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:6.5.4

sudo docker run -d --name kibana --net elk_network -p 5601:5601 kibana:6.5.4

