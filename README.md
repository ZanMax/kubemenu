# KubeMenu

kubemenu is a simple tool for connecting to Kubernetes cluster written on GO.

### How to use:

- Create folder inside ~/.kube folder ( like dev_kube )
- Move config file from root ~/.kube to dev_kube
- Repeat for all your cluster
- Run ./kubemenu and enjoy

#### Example
	kube
	├── dev
	│   └── config
	├── other
	│   └── config
	├── prod
	│   └── config
	├── stage
	│   └── config

### Compilation:
```bash
go build kubemenu.go
```
Other OS:
```bash
env GOOS=linux GOARCH=amd64 go build kubemenu.go
```
### How to install:
```bash
mv kubemenu /usr/local/bin
chmod +x /usr/local/bin/kubemenu
kubemenu
```