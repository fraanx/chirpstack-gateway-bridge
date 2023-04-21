#!/bin/bash

############ LORAWAN

oldstring=github.com/brocaar/lorawan
newstring=github.com/fraanx/lorawan

read -p  " Do you want to replace repository from $oldstring to $newstring ?"
if [[ $REPLY == "y" ]]
then

targetlist=$(grep -r $oldstring|cut -d ':' -f 1)

for fname in $targetlist
do
echo file is $fname
sed -i "s#$oldstring#$newstring#g" "$fname"

done

#restore reporeplace strings
wrongold="oldstring=github.com/fraanx/lorawan"
correctold="oldstring=github.com/brocaar/lorawan"

sed -i "s#$wrongold#$correctold#g" $0

fi


##### chirpstack-gateway-bridge
oldstring=github.com/brocaar/chirpstack-gateway-bridge
newstring=github.com/fraanx/chirpstack-gateway-bridge

read -p  " Do you want to replace repository from $oldstring to $newstring ?"
if [[ $REPLY == "y" ]]
then

targetlist=$(grep -r $oldstring|cut -d ':' -f 1)

for fname in $targetlist
do
echo file is $fname
sed -i "s#$oldstring#$newstring#g" "$fname"

done

#restore reporeplace strings
wrongold="oldstring=github.com/chirpstack/chirpstack-gateway-bridge"
correctold="oldstring=github.com/brocaar/chirpstack-gateway-bridge"

sed -i "s#$wrongold#$correctold#g" $0

fi

##### chirpstack
oldstring=github.com/chirpstack/chirpstack/
newstring=github.com/fraanx/chirpstack/

read -p  " Do you want to replace repository from $oldstring to $newstring ?"
if [[ $REPLY == "y" ]]
then

targetlist=$(grep -r $oldstring|cut -d ':' -f 1)

for fname in $targetlist
do
echo file is $fname
sed -i "s#$oldstring#$newstring#g" "$fname"

done

#restore reporeplace strings
wrongold="oldstring=github.com/chirpstack/chirpstack"
correctold="oldstring=github.com/chirpstack/chirpstack"

sed -i "s#$wrongold#$correctold#g" $0

fi

