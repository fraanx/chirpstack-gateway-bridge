#!/bin/bash

oldstring=github.com/brocaar/chirpstack-api
newstring=github.com/fraanx/chirpstack-v3-api

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
wrongold="oldstring=github.com/fraanx/chirpstack-v3-api"
correctold="oldstring=github.com/brocaar/chirpstack-api"

sed -i "s#$wrongold#$correctold#g" $0

fi

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
wrongold="oldstring=github.com/fraanx/chirpstack-gateway-bridge"
correctold="oldstring=github.com/brocaar/chirpstack-gateway-bridge"

sed -i "s#$wrongold#$correctold#g" $0

fi

