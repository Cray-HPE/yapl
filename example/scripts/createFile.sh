touch /tmp/demo.txt
sleep $[ ( $RANDOM % 3 )  + 1 ]s
echo "/tmp/demo.txt is created"