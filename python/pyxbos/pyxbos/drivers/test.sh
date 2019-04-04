#!/bin/bash
DRIVER_NAME=$1
ENTITY=$DRIVER_NAME".ent"
echo 'driver name is' $DRIVER_NAME
echo 'entity name is' $ENTITY
wv mke -e 50y --nopassphrase -o $ENTITY
echo "Made entity"
wv inspect $ENTITY
echo "Did inspection"
wv rtgrant --subject $ENTITY -e 3y --attester $WAVE_DEFAULT_ENTITY --indirections 0 wavemq:publish,subscribe@solarplus/$DRIVER_NAME/*
echo "Granted permissions"
echo $ENTITY
echo $NAMESPACE
echo $NAMESPACE_HASH
wv name --public --attester $ENTITY $NAMESPACE_HASH $NAMESPACE
echo "linked namespace to hash"
wv rtprove --subject $ENTITY -o driverproof.pem wavemq:publish,subscribe@solarplus/$DRIVER_NAME/*
echo "Did proof"
wv verify driverproof.pem
echo "Did verification"
