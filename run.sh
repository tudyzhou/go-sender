#!/bin/bash
rm -f ./main
echo "Builing main.go to main"
go build main.go
rc=$?
if [[ $rc != 0 ]]; then
	#echo "Build fail"
	exit $rc
else
	echo "Running *main"
	./main
fi
