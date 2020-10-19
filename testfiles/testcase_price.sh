#!/bin/bash

# $2 should be reserved for test case input if necessary

func_result="`/bin/bash .oraifiles/"$1"`"
# $3 is the expected output
expected_output=$(echo "$3" |base64 --decode)
diff=$(echo "scale=2; $func_result - $expected_output" |bc) # |bc allows adding two 
# allow compare integer with float
if [[ $(echo "$diff < 10000" | bc) ]]
then 
  echo $func_result
else
  echo null
fi