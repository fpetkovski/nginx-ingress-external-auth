docker_build('electroma/ingress-demo-echosvc-amd64:0.1', 'echosvc')
k8s_yaml('deploy/echo-service.yaml')

docker_build('electroma/ingress-demo-authsvc-amd64:0.1', 'authsvc')
k8s_yaml('deploy/auth-service.yaml')
