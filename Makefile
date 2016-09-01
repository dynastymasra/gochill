SHELL := bash

VERSION ?= latest
TEST_PREFIX ?= test-
PROJECT_NAME = gochill

.PHONY: test

test_version.txt:
	# put version tag for test image into a test_version.txt
	# e.g. test-ab12cd34
	echo -n "$(TEST_PREFIX)" > test_version.txt
	cat /dev/urandom | LC_CTYPE=C tr -dc 'a-f0-9' | fold -w 8 | head -n 1 \
		>> test_version.txt

test: unit-test

unit-test: test_version.txt
	go test -v -cover
