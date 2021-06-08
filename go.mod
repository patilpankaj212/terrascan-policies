module github.com/patilpankaj212/terrascan-policies

go 1.16

replace (
	k8s.io/api => k8s.io/api v0.19.0
	k8s.io/apimachinery => k8s.io/apimachinery v0.19.0
	k8s.io/client-go => k8s.io/client-go v0.19.0
)

require github.com/accurics/terrascan v1.6.0
