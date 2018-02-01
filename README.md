* `echo "GET http://localhost:8080/notification?message=RoyHello" | vegeta attack -rate=10 -duration=5s | tee results.bin | vegeta report`
* `echo "GET http://localhost:8080/notification?message=AwesomeHexter" | vegeta attack -rate=10 -duration=5s | tee results.bin | vegeta report`

* `brew update && brew install vegeta` or `go get -u github.com/tsenart/vegeta`