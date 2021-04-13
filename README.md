This is a test for setting up development tooling. 

1. What do I have to install for local development?
   - [Docker](https://www.docker.com/)
   - [Helm](https://helm.sh/docs/intro/quickstart/)
   - [Kubectl](https://kubernetes.io/docs/tasks/tools/)
   - [K3d](https://k3d.io/#installation)
   - [Go](https://golang.org/)
   - [Telepresence](https://www.telepresence.io/)

2. Copy the .env.template file to an .env file and edit it to your liking.
```
cp .env.template .env
```

3. Create a deamon.json file in /etc/docker/ that enables insecure registries at port 5000.
```
./scripts/add_registry_to_docker.sh
```

4. The cluster can be launched with the following script.
```
./scripts/start_cluster.sh
```

- Creates a new Docker image registry.
- Creates a local Kubernetes cluster.
- Add the secrets for your microservices.

To reach your cluster:
 - port 9080 (HTTP)
 - port 9443 (HTTPS)