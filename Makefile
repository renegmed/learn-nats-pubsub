init-project:
	go mod init github.com/renegmed/learn-nats-pubsub
.PHONY: init-project

up:
	docker-compose up --build -d
.PHONY: up

#----- publisher 1 ----------- #
log-pub1:
	docker logs -f publisher-1
.PHONY: log-pub1

pub-info-pub1:
	curl -X POST  http://localhost:8282/pub-product-info/2 
	curl -X POST  http://localhost:8282/pub-product-info/4 
	curl -X POST  http://localhost:8282/pub-product-info/6
.PHONY: pub-info-pub1


#----- subscriber 1 ----------- #
log-sub1:
	docker logs -f subscriber-1
.PHONY: log-sub1

list-sub1:
	curl -X GET  http://localhost:8484/products 
.PHONY: list-sub1


#----- subscriber 2 ----------- #
log-sub2:
	docker logs -f subscriber-2
.PHONY: log-sub2

list-sub2:
	curl -X GET http://localhost:8686/products 
.PHONY: list-sub2


#----- NATS server ----------- #
info-server:
	watch -n 1 curl http://localhost:8222/varz --silent
	# message throutput
	# watch -n 1 'curl http://localhost:8222/varz --silent | grep "\(msgs\|byte\)"'
.PHONY: info-server

# statistics and metadata about the clients currently connected to the server.
info-conn:
	watch -n 1 curl http://localhost:8222/connz --silent
.PHONY: info-conn

# cumulative stats about internal state of the sublist data structur that server maintains.
info-sublist:
	watch -n 1 curl http://localhost:8222/subsz --silent
.PHONY: info-sublist
