### Simple Golang/NATS Pub and Sub application for learning ###

To run:

1. Build and run NATS server, publisher 1, subscriber 1, & subscriber 2 docker containers

        make up

2. Monitor publisher 1 logs (new terminal)

        make log-pub1

3. Monitor subscriber 1 logs (new terminal)

        make log-sub1

4. Moniter subscriber 2 logs (new terminal)

        make log-sub2

5. Publisher 1 publish product info by proving product ID. Example:

        curl -X POST  http://localhost:8282/pub-product-info/1 

        curl -X POST  http://localhost:8282/pub-product-info/2 

Note:
    Each application must run go mod vendor before building containers




