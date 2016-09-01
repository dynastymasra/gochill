SHELL := bash

VERSION ?= latest
TEST_PREFIX ?= test-
PROJECT_NAME = gochill

.PHONY: build

build:
	echo "building..."

test_version.txt:
	# put version tag for test image into a test_version.txt
	# e.g. test-ab12cd34
	echo -n "$(TEST_PREFIX)" > test_version.txt
	cat /dev/urandom | LC_CTYPE=C tr -dc 'a-f0-9' | fold -w 8 | head -n 1 \
		>> test_version.txt

clean: clean-txt clean-out clean-xml

test: unit-test

unit-test: test_version.txt
		$(MAKE) -e VERSION=$$(cat test_version.txt) build
		go get -u github.com/jstemmer/go-junit-report
		go get github.com/axw/gocov/gocov
		go get github.com/AlekSi/gocov-xml
		./test-unit.sh

clean-txt:
		rm *.txt

clean-out:
		rm *.out

clean-xml:
		rm *.xml
