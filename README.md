# gobazel
Trying out bazel with go microservices.

Clone repo
Run `dep ensure` to create vendor folder

Use builder image to build targets

Bazel requires a build file in each package. When adding dependencies to the vendor
folder using `dep ensure -add pkg` newly added packages will not have the build file
by default. Run `bazel run //:gazelle` to automatically add build files to vendor
pachages.


To make using the system way easier it's important to map the minikube ip address to any names used. Get the minikube ip address by issuing `minikube ip` and then update hosts file

/etc/hosts host entry
`minikubeIP      kubernetes registry.minikube prometheus.minikube grafana.minikube gobazel.minikube traefik.minikube`

example:
`192.168.64.19	kubernetes registry.minikube prometheus.minikube grafana.minikube gobazel.minikube traefik.minikube`

## Minikube
Install minikube check the current releases on https://github.com/kubernetes/minikube/releases/ install is like:
`curl -Lo minikube https://storage.googleapis.com/minikube/releases/v1.0.0/minikube-darwin-amd64 && chmod +x minikube && sudo cp minikube /usr/local/bin/ && rm minikube`

The project depends on minikube running. 
`minikube start --vm-driver=hyperkit` on mac with hyperkit hypervisor
To stop minikube run
`minikube stop`
To start fresh delete the current cluster using 
`minikube delete` then issue start command



# Notes
Open files error
For OS X Sierra (10.12.X) you need to:

1. Create a file at /Library/LaunchDaemons/limit.maxfiles.plist and paste the following in (feel free to change the two numbers (which are the soft and hard limits, respectively):

<?xml version="1.0" encoding="UTF-8"?>  
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN"  
        "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">  
  <dict>
    <key>Label</key>
    <string>limit.maxfiles</string>
    <key>ProgramArguments</key>
    <array>
      <string>launchctl</string>
      <string>limit</string>
      <string>maxfiles</string>
      <string>64000</string>
      <string>524288</string>
    </array>
    <key>RunAtLoad</key>
    <true/>
    <key>ServiceIPC</key>
    <false/>
  </dict>
</plist> 
2. Change the owner of your new file:

sudo chown root:wheel /Library/LaunchDaemons/limit.maxfiles.plist
3. Load these new settings:

sudo launchctl load -w /Library/LaunchDaemons/limit.maxfiles.plist
4. Finally, check that the limits are correct:

launchctl limit maxfiles


Where docker rules expects the binary
```
/app/service1/service1-image.runfiles/__main__/service1/service1
```

what is should 
```
/app/service1/service1.runfiles/__main__/service1/linux_amd64_pure_stripped/service1
```

why time sync issue happens
https://github.com/kubernetes/minikube/issues/1378

how to work arond
`make sync`

Testing with vegeta

`vegeta attack -rate=100 -duration=10s -targets=testdata/targets.txt | vegeta report`


## Hosts file
127.0.0.1	localhost
255.255.255.255	broadcasthost
::1             localhost 
192.168.64.19	kubernetes tracing.minikube prometheus.minikube grafana.minikube gobazel.minikube traefik.minikube
