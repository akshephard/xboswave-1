#!/bin/bash
DRIVER_NAME=$1
echo 'driver name is' $DRIVER_NAME
wv mke -e 50y --nopassphrase -o $ENTITY
echo "Made entity"
wv inspect $ENTITY
echo "Did inspection"
wv rtgrant --subject $ENTITY -e 3y --attester $WAVE_DEFAULT_ENTITY --indirections 0 wavemq:publish,subscribe@solarplus/$DRIVER_NAME/*
echo "Granted permissions"
wv name --public --attester $ENTITY $NAMESPACE_HASH $NAMESPACE
echo "linked namespace to hash"
wv rtprove --subject $ENTITY -o driver_proof.pem wavemq:publish,subscribe@solarplus/$DRIVER_NAME/*
echo "Did proof"
wv verify driverproof.pem
echo "Did verification"
