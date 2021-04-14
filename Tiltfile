# Load external restart proces
load('ext://restart_process', 'docker_build_with_restart')

# Only allow the follow context
allow_k8s_contexts('k3d-dts')

# File service deployment and live development
# k8s_yaml(helm('deployments/file-service', name='file-service'))

# docker_build_with_restart(
#   'k3d-dts-registry:5000/file-service', 
#   './services/file-service',
#   entrypoint='/start_app',
#   target='dev',
#   dockerfile='./services/file-service/Dockerfile',
#   live_update=[
#     sync('./services/file-service', '/usr/src/app'),
#   ],
# )

# Workflow manager service deployment and live development
k8s_yaml(helm('deployments/workflow-manager-service', name='workflow-manager-service'))

docker_build_with_restart(
  'k3d-dts-registry:5000/workflow-manager-service', 
  './services/workflow-manager-service',
  entrypoint='/start_app',
  target='dev',
  dockerfile='./services/workflow-manager-service/Dockerfile',
  live_update=[
    sync('./services/workflow-manager-service', '/usr/src/app'),
  ],
)
