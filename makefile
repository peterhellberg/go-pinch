#
# Old school makefile
#
# $ export GOPATH=`pwd`
# $ go get -u github.com/peterhellberg/go-pinch
# $ make
#

all: pinch-test test1-1 test1-2 test2-1 test2-2 done

GOSRCS = 		\
	main.go		\
	pinch/pinch.go

TEST1_URL  = http://peterhellberg.github.com/pinch/test.zip
TEST1_FILE = data.json

TEST2_URL  = http://assets.c7.se/data/pinch/example2.zip
TEST2_FILE = user_details.js

pinch-test: $(GOSRCS)
	@echo =[COMPILE]======================================================
	go build -o=pinch-test

test1-1:
	@echo =[TEST1-1]======================================================
	./pinch-test $(TEST1_URL)

test1-2:
	@echo =[TEST1-2]======================================================
	./pinch-test $(TEST1_URL) $(TEST1_FILE)

test2-1:
	@echo =[TEST2-1]======================================================
	./pinch-test $(TEST2_URL)

test2-2:
	@echo =[TEST2-2]======================================================
	./pinch-test $(TEST2_URL) $(TEST2_FILE)

done:
	@echo =[DONE]=========================================================
	rm pinch-test
