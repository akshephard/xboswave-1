#!/bin/bash

echo "Positional Parameters"
echo '$0 = ' $0
echo '$1 = ' $1
echo '$2 = ' $2
echo '$3 = ' $3
DRIVER_NAME=$1
echo 'driver name is' $DRIVER_NAME
ENTITY=$DRIVER_NAME".ent"
echo $ENTITY

wv mke -e 50y --nopassphrase -o $ENTITY
wv inspect $ENTITY
wv rtgrant --subject $ENTITY -e 3y --attester $WAVE_DEFAULT_ENTITY --indirections 0 wavemq:publish,subscribe@solarplus/$DRIVER_NAME/*
wv name --public --attester $ENTITY $NAMESPACE_HASH $NAMESPACE
wv rtprove --subject $ENTITY -o driver_proof.pem wavemq:publish,subscribe@solarplus/$DRIVER_NAME/*
wv verify driverproof.pem
