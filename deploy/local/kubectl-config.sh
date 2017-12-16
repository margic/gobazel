#!/bin/bash

echo copying kubectl certs from $HOME/.minikube

if [ ! -d $HOME/.minikube ] 
    then  echo '$HOME/.minikube' not found, is minikube installed?
    exit 1 
fi

# copy kube ctl files across assuming minikube and kubectl installed and working
if [ -e $HOME/.minikube/client.crt ]
 then cp $HOME/.minikube/client.crt kube/client.crt
 echo copied client.crt
fi
if [ -e $HOME/.minikube/client.key ]
 then cp $HOME/.minikube/client.key kube/client.key
 echo copied client.key
fi
if [ -e $HOME/.minikube/ca.crt ]
 then cp $HOME/.minikube/ca.crt kube/ca.crt
 echo copied client.crt
fi