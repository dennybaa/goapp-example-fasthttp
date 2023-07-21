# go fasthttp demo application chart

This chart deploys a basic demo go application [docker/dennybaa/goapp-example-fasthttp](https://hub.docker.com/repository/docker/dennybaa/goapp-example-fasthttp)

## TL;DR

```console
git clone https://github.com/dennybaa/goapp-example-fasthttp
cd goapp-example-fasthttp

helm install app chart/

pod=$(kubectl get pods -l app.kubernetes.io/name=demo -o jsonpath='{.items[].metadata.name}')
kubectl exec -t $pod -- wget -qO - http://app-demo:8080/hello
```

## Source Code

* <https://github.com/dennybaa/goapp-example-fasthttp/tree/main/chart>

## Parameters

### App chart parameters

| Name                   | Description                                                                                                            | Value             |
| ---------------------- | ---------------------------------------------------------------------------------------------------------------------- | ----------------- |
| `app.name`             | Specifies the application name                                                                                         | `demo`            |
| `app.workload.enabled` | Specifies whether the default workload resource is generated (Deployment/StatefulSet etc)                              | `true`            |
| `app.workload.type`    | Specifies type of the main workload resource                                                                           | `deployment`      |
| `app.components`       | Specifies list of components to enable used in direct mode (it respectively expects .Values.[component] to be present) | `[]`              |
| `selector.matchLabels` | Specifies additional selector labels for the workload resources and services                                           | `{}`              |
| `reuse`                | Enables reuse/merge of the upper-level component values (applicable for containers/initContainers)                     | `false`           |
| `containers`           | Specify a map of additional pod containers                                                                             | `{}`              |
| `initContainers`       | Specifies initContainers **(use, values map for order and data)**                                                      | `{}`              |
| `env.ENVIRONMENT`      | Specifies environment  (production/development)                                                                        | `production`      |
| `env.BACKEND`          | Specifies backend mode (logfile/sqlite)                                                                                | `sqlite`          |
| `env.FILEPATH`         | Data file store path                                                                                                   | `/data/sqlite.db` |
| `env.PORT`             | Specifies port to listen on                                                                                            | `""`              |
| `env.LOGLEVEL`         | Specifies the log level for application logger                                                                         | `""`              |
| `envFrom`              | Configures of envFrom to include into the main container                                                               | `[]`              |
| `volumes`              | Specify volumes for the main pod                                                                                       | `{}`              |
| `volumeMounts`         | Specify volumeMounts for the main container                                                                            | `{}`              |
| `configMaps`           | Creates application сonfigMaps (note the name is prefixed with the app name)                                           | `{}`              |
| `secrets`              | Creates application secrets (note the name is prefixed with the app name)                                              | `{}`              |
| `templateChecksums`    | Specifies list of template files to add as an annotation checksum into the pod.                                        | `[]`              |


### Global parameters

| Name                      | Description                                     | Value |
| ------------------------- | ----------------------------------------------- | ----- |
| `global.imageRegistry`    | Global Docker image registry                    | `""`  |
| `global.imagePullSecrets` | Global Docker registry secret names as an array | `[]`  |
| `global.storageClass`     | Global StorageClass for Persistent Volume(s)    | `""`  |


### Common parameters

| Name                     | Description                                                                             | Value           |
| ------------------------ | --------------------------------------------------------------------------------------- | --------------- |
| `kubeVersion`            | Override Kubernetes version                                                             | `""`            |
| `nameOverride`           | String to partially override common.names.name                                          | `""`            |
| `fullnameOverride`       | String to fully override common.names.fullname                                          | `""`            |
| `namespaceOverride`      | String to fully override common.names.namespace                                         | `""`            |
| `commonLabels`           | Labels to add to all deployed objects                                                   | `{}`            |
| `commonAnnotations`      | Annotations to add to all deployed objects                                              | `{}`            |
| `clusterDomain`          | Kubernetes cluster domain name                                                          | `cluster.local` |
| `extraDeploy`            | Array of extra objects to deploy with the release                                       | `[]`            |
| `diagnosticMode.enabled` | Enable diagnostic mode (all probes will be disabled and the command will be overridden) | `false`         |
| `diagnosticMode.command` | Command to override all containers in the deployment                                    | `["sleep"]`     |
| `diagnosticMode.args`    | Args to override all containers in the deployment                                       | `["infinity"]`  |


### Main pod Parameters

| Name                                              | Description                                                                                                                                                       | Value                             |
| ------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------- |
| `image.registry`                                  | image registry                                                                                                                                                    | `""`                              |
| `image.repository`                                | image repository                                                                                                                                                  | `dennybaa/goapp-example-fasthttp` |
| `image.tag`                                       | image tag (immutable tags are recommended)                                                                                                                        | `0.1.1`                           |
| `image.digest`                                    | image digest in the way sha256:aa.... Please note this parameter, if set, will override the tag image tag (immutable tags are recommended)                        | `""`                              |
| `image.pullPolicy`                                | image pull policy                                                                                                                                                 | `IfNotPresent`                    |
| `image.pullSecrets`                               | image pull secrets                                                                                                                                                | `[]`                              |
| `image.debug`                                     | Enable image debug mode                                                                                                                                           | `false`                           |
| `replicaCount`                                    | Number of replicas to deploy                                                                                                                                      | `1`                               |
| `containerPorts.http`                             | Specifies app container port                                                                                                                                      | `8080`                            |
| `livenessProbe.enabled`                           | Enable livenessProbe on containers                                                                                                                                | `false`                           |
| `livenessProbe.initialDelaySeconds`               | Initial delay seconds for livenessProbe                                                                                                                           | `0`                               |
| `livenessProbe.periodSeconds`                     | Period seconds for livenessProbe                                                                                                                                  | `10`                              |
| `livenessProbe.timeoutSeconds`                    | Timeout seconds for livenessProbe                                                                                                                                 | `1`                               |
| `livenessProbe.failureThreshold`                  | Failure threshold for livenessProbe                                                                                                                               | `3`                               |
| `livenessProbe.successThreshold`                  | Success threshold for livenessProbe                                                                                                                               | `1`                               |
| `livenessProbe.exec`                              | Specifies exec action (must set before enabling)                                                                                                                  |                                   |
| `livenessProbe.grpc`                              | Specifies GRPC action (must set before enabling)                                                                                                                  |                                   |
| `livenessProbe.httpGet`                           | Specifies HTTPGet action (must set before enabling)                                                                                                               |                                   |
| `livenessProbe.tcpSocket`                         | Specifies TCPSocket action (must set before enabling)                                                                                                             |                                   |
| `readinessProbe.enabled`                          | Enable readinessProbe on containers                                                                                                                               | `true`                            |
| `readinessProbe.initialDelaySeconds`              | Initial delay seconds for readinessProbe                                                                                                                          | `0`                               |
| `readinessProbe.periodSeconds`                    | Period seconds for readinessProbe                                                                                                                                 | `10`                              |
| `readinessProbe.timeoutSeconds`                   | Timeout seconds for readinessProbe                                                                                                                                | `1`                               |
| `readinessProbe.failureThreshold`                 | Failure threshold for readinessProbe                                                                                                                              | `3`                               |
| `readinessProbe.successThreshold`                 | Success threshold for readinessProbe                                                                                                                              | `1`                               |
| `readinessProbe.exec`                             | Specifies exec action (one of the actions must be set before enabling)                                                                                            |                                   |
| `readinessProbe.grpc`                             | Specifies GRPC action (one of the actions must be set before enabling)                                                                                            |                                   |
| `readinessProbe.httpGet`                          | Specifies HTTPGet action (one of the actions must be set before enabling)                                                                                         |                                   |
| `readinessProbe.httpGet.path`                     | Specifies path of readiness endpoint                                                                                                                              | `/hello`                          |
| `readinessProbe.httpGet.port`                     | Specifies the container port                                                                                                                                      | `http`                            |
| `readinessProbe.tcpSocket`                        | Specifies TCPSocket action (one of the actions must be set before enabling)                                                                                       |                                   |
| `startupProbe.enabled`                            | Enable startupProbe on containers                                                                                                                                 | `false`                           |
| `startupProbe.initialDelaySeconds`                | Initial delay seconds for startupProbe                                                                                                                            | `0`                               |
| `startupProbe.periodSeconds`                      | Period seconds for startupProbe                                                                                                                                   | `10`                              |
| `startupProbe.timeoutSeconds`                     | Timeout seconds for startupProbe                                                                                                                                  | `1`                               |
| `startupProbe.failureThreshold`                   | Failure threshold for startupProbe                                                                                                                                | `3`                               |
| `startupProbe.successThreshold`                   | Success threshold for startupProbe                                                                                                                                | `1`                               |
| `startupProbe.exec`                               | Specifies exec action (one of the actions must be set before enabling)                                                                                            |                                   |
| `startupProbe.grpc`                               | Specifies GRPC action (one of the actions must be set before enabling)                                                                                            |                                   |
| `startupProbe.httpGet`                            | Specifies HTTPGet action (one of the actions must be set before enabling)                                                                                         |                                   |
| `startupProbe.tcpSocket`                          | Specifies TCPSocket action (one of the actions must be set before enabling)                                                                                       |                                   |
| `customLivenessProbe`                             | Custom livenessProbe that overrides the default one                                                                                                               | `{}`                              |
| `customReadinessProbe`                            | Custom readinessProbe that overrides the default one                                                                                                              | `{}`                              |
| `customStartupProbe`                              | Custom startupProbe that overrides the default one                                                                                                                | `{}`                              |
| `resources.requests`                              | The requested resources for the containers                                                                                                                        | `{}`                              |
| `resources.limits.cpu`                            |                                                                                                                                                                   | `100m`                            |
| `resources.limits.memory`                         |                                                                                                                                                                   | `128Mi`                           |
| `podSecurityContext.enabled`                      | Enabled pods' Security Context                                                                                                                                    | `true`                            |
| `podSecurityContext.fsGroup`                      | Set pod's Security Context fsGroup                                                                                                                                | `1001`                            |
| `containerSecurityContext.enabled`                | Enabled containers' Security Context                                                                                                                              | `true`                            |
| `containerSecurityContext.runAsUser`              | Set containers' Security Context runAsUser                                                                                                                        | `1001`                            |
| `containerSecurityContext.runAsNonRoot`           | Set containers' Security Context runAsNonRoot                                                                                                                     | `true`                            |
| `containerSecurityContext.readOnlyRootFilesystem` | Set containers' Security Context runAsNonRoot                                                                                                                     | `false`                           |
| `command`                                         | Override default container command (useful when using custom images)                                                                                              | `[]`                              |
| `args`                                            | Override default container args (useful when using custom images)                                                                                                 | `[]`                              |
| `hostAliases`                                     | pods host aliases                                                                                                                                                 | `[]`                              |
| `podLabels`                                       | Extra labels for pods                                                                                                                                             | `{}`                              |
| `podAnnotations`                                  | Annotations for pods                                                                                                                                              | `{}`                              |
| `podAffinityPreset`                               | Pod affinity preset. Ignored if `affinity` is set. Allowed values: `soft` or `hard`                                                                               | `""`                              |
| `podAntiAffinityPreset`                           | Pod anti-affinity preset. Ignored if `affinity` is set. Allowed values: `soft` or `hard`                                                                          | `soft`                            |
| `pdb.create`                                      | Enable/disable a Pod Disruption Budget creation                                                                                                                   | `false`                           |
| `pdb.minAvailable`                                | Minimum number/percentage of pods that should remain scheduled                                                                                                    | `1`                               |
| `pdb.maxUnavailable`                              | Maximum number/percentage of pods that may be made unavailable                                                                                                    | `""`                              |
| `autoscaling.enabled`                             | Enable autoscaling for %%MAIN_OBJECT_BLOCK%%                                                                                                                      | `false`                           |
| `autoscaling.minReplicas`                         | Minimum number of %%MAIN_OBJECT_BLOCK%% replicas                                                                                                                  | `""`                              |
| `autoscaling.maxReplicas`                         | Maximum number of %%MAIN_OBJECT_BLOCK%% replicas                                                                                                                  | `""`                              |
| `autoscaling.targetCPU`                           | Target CPU utilization percentage                                                                                                                                 | `""`                              |
| `autoscaling.targetMemory`                        | Target Memory utilization percentage                                                                                                                              | `""`                              |
| `autoscaling.behavior`                            | Specifies separate scale-up and scale-down behaviors                                                                                                              | `{}`                              |
| `nodeAffinityPreset.type`                         | Node affinity preset type. Ignored if `affinity` is set. Allowed values: `soft` or `hard`                                                                         | `""`                              |
| `nodeAffinityPreset.key`                          | Node label key to match. Ignored if `affinity` is set                                                                                                             | `""`                              |
| `nodeAffinityPreset.values`                       | Node label values to match. Ignored if `affinity` is set                                                                                                          | `[]`                              |
| `affinity`                                        | Affinity for pods assignment                                                                                                                                      | `{}`                              |
| `nodeSelector`                                    | Node labels for pods assignment                                                                                                                                   | `{}`                              |
| `tolerations`                                     | Tolerations for pods assignment                                                                                                                                   | `[]`                              |
| `updateStrategy.type`                             | statefulset strategy type                                                                                                                                         | `RollingUpdate`                   |
| `dnsPolicy`                                       | Set DNS policy for the pod. Defaults to "ClusterFirst".                                                                                                           | `nil`                             |
| `hostNetwork`                                     | Host networking requested for this pod. Use the host's network namespace. If this option is set, the ports that will be used must be specified. Default to false. | `nil`                             |
| `automountServiceAccountToken`                    | Automount service account token for the pod. Defaults to "true"                                                                                                   | `nil`                             |
| `podManagementPolicy`                             | Statefulset Pod management policy, it needs to be Parallel to be able to complete the cluster join                                                                | `OrderedReady`                    |
| `priorityClassName`                               | pods' priorityClassName                                                                                                                                           | `""`                              |
| `topologySpreadConstraints`                       | Topology Spread Constraints for pod assignment spread across your cluster among failure-domains. Evaluated as a template                                          | `[]`                              |
| `schedulerName`                                   | Name of the k8s scheduler (other than default) for pods                                                                                                           | `""`                              |
| `terminationGracePeriodSeconds`                   | Seconds Redmine pod needs to terminate gracefully                                                                                                                 | `""`                              |
| `lifecycleHooks`                                  | for the container(s) to automate configuration before or after startup                                                                                            | `{}`                              |
| `extraEnvVars`                                    | Array with extra environment variables to add to nodes                                                                                                            | `[]`                              |
| `extraEnvVarsCM`                                  | Name of existing ConfigMap containing extra env vars for nodes                                                                                                    | `""`                              |
| `extraEnvVarsSecret`                              | Name of existing Secret containing extra env vars for nodes                                                                                                       | `""`                              |
| `extraVolumes`                                    | Optionally specify extra list of additional volumes for the pod(s)                                                                                                | `[]`                              |
| `extraVolumeMounts`                               | Optionally specify extra list of additional volumeMounts for the container(s)                                                                                     | `[]`                              |
| `sidecars`                                        | Add additional sidecar containers to the pod(s)                                                                                                                   | `[]`                              |
| `extraInitContainers`                             | Add additional init containers to the pod(s) (go after .initContainers)                                                                                           | `[]`                              |


### Traffic Exposure Parameters

| Name                               | Description                                                                                                                      | Value                    |
| ---------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | ------------------------ |
| `service.type`                     | service type                                                                                                                     | `ClusterIP`              |
| `service.ports.http`               | Specify service the http service port                                                                                            | `8080`                   |
| `service.clusterIP`                | service Cluster IP                                                                                                               | `""`                     |
| `service.loadBalancerIP`           | service Load Balancer IP                                                                                                         | `""`                     |
| `service.loadBalancerSourceRanges` | service Load Balancer sources                                                                                                    | `[]`                     |
| `service.externalTrafficPolicy`    | service external traffic policy                                                                                                  | `Cluster`                |
| `service.annotations`              | Additional custom annotations for service                                                                                        | `{}`                     |
| `service.extraPorts`               | Extra ports to expose in service (normally used with the `sidecars` value)                                                       | `[]`                     |
| `service.sessionAffinity`          | Control where client requests go, to the same pod or round-robin                                                                 | `None`                   |
| `service.sessionAffinityConfig`    | Additional settings for the sessionAffinity                                                                                      | `{}`                     |
| `ingress.enabled`                  | Enable ingress record generation                                                                                                 | `false`                  |
| `ingress.namespace`                | Specify custom namespace for the ingress (has priority both over the release namespace and namespaceOverride)                    | `""`                     |
| `ingress.customName`               | Specify custom name for the Ingress                                                                                              | `""`                     |
| `ingress.serviceName`              | Specify service ingress points too (uses the main service by default)                                                            | `""`                     |
| `ingress.servicePort`              | Specifies the service port (must be provided)                                                                                    | `nil`                    |
| `ingress.pathType`                 | Ingress path type                                                                                                                | `ImplementationSpecific` |
| `ingress.apiVersion`               | Force Ingress API version (automatically detected if not set)                                                                    | `""`                     |
| `ingress.hostname`                 | Default host for the ingress record                                                                                              | `app.local`              |
| `ingress.ingressClassName`         | IngressClass that will be be used to implement the Ingress (Kubernetes 1.18+)                                                    | `""`                     |
| `ingress.path`                     | Default path for the ingress record                                                                                              | `/`                      |
| `ingress.annotations`              | Additional annotations for the Ingress resource. To enable certificate autogeneration, place here your cert-manager annotations. | `{}`                     |
| `ingress.tls`                      | Enable TLS configuration for the host defined at `ingress.hostname` parameter                                                    | `false`                  |
| `ingress.selfSigned`               | Create a TLS secret for this ingress record using self-signed certificates generated by Helm                                     | `false`                  |
| `ingress.selfSignedDays`           | Validity of the self-signed certificates generated by Helm                                                                       | `365`                    |
| `ingress.extraHosts`               | An array with additional hostname(s) to be covered with the ingress record                                                       | `[]`                     |
| `ingress.extraPaths`               | An array with additional arbitrary paths that may need to be added to the ingress under the main host                            | `[]`                     |
| `ingress.extraTls`                 | TLS configuration for additional hostname(s) to be covered with this ingress record                                              | `[]`                     |
| `ingress.secrets`                  | Custom TLS certificates as secrets                                                                                               | `[]`                     |
| `ingress.extraRules`               | Additional rules to be covered with this ingress record                                                                          | `[]`                     |


### Persistence Parameters

| Name                        | Description                                                                                             | Value               |
| --------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------- |
| `persistence.enabled`       | Enable persistence using Persistent Volume Claims                                                       | `true`              |
| `persistence.emptyDir`      | Enable emptyDir persistence instead of a PVC                                                            | `false`             |
| `persistence.mountName`     | Persistent volume name                                                                                  | `data`              |
| `persistence.mountPath`     | Path to mount the volume at                                                                             | `/data`             |
| `persistence.subPath`       | The subdirectory of the volume to mount to, useful in dev environments and one PV for multiple services | `""`                |
| `persistence.storageClass`  | Storage class of backing PVC                                                                            | `""`                |
| `persistence.annotations`   | Persistent Volume Claim annotations                                                                     | `{}`                |
| `persistence.accessModes`   | Persistent Volume Access Modes                                                                          | `["ReadWriteOnce"]` |
| `persistence.size`          | Size of data volume                                                                                     | `1Gi`               |
| `persistence.existingClaim` | The name of an existing PVC to use for persistence                                                      | `""`                |
| `persistence.selector`      | Selector to match an existing Persistent Volume for WordPress data PVC                                  | `{}`                |
| `persistence.dataSource`    | Custom PVC data source                                                                                  | `{}`                |


### Init Container Parameters

| Name                                                   | Description                                              | Value                   |
| ------------------------------------------------------ | -------------------------------------------------------- | ----------------------- |
| `volumePermissions.command`                            | Command to execute in volumePermissions container.       | `[]`                    |
| `volumePermissions.image.registry`                     | Bitnami Shell image registry                             | `docker.io`             |
| `volumePermissions.image.repository`                   | Bitnami Shell image repository                           | `bitnami/bitnami-shell` |
| `volumePermissions.image.tag`                          | Bitnami Shell image tag (immutable tags are recommended) | `11-debian-11`          |
| `volumePermissions.image.pullPolicy`                   | Bitnami Shell image pull policy                          | `Always`                |
| `volumePermissions.image.pullSecrets`                  | Bitnami Shell image pull secrets                         | `[]`                    |
| `volumePermissions.resources.limits`                   | The resources limits for the init container              | `{}`                    |
| `volumePermissions.resources.requests`                 | The requested resources for the init container           | `{}`                    |
| `volumePermissions.containerSecurityContext.runAsUser` | Set init container's Security Context runAsUser          | `0`                     |


### Other Parameters

| Name                                          | Description                                                                                                              | Value   |
| --------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------ | ------- |
| `minReadySeconds`                             | minimum seconds for pod to become ready. 0 is the default k8s value                                                      | `0`     |
| `rbac.create`                                 | Specifies whether RBAC resources should be created                                                                       | `false` |
| `rbac.rules`                                  | Custom RBAC rules to set                                                                                                 | `[]`    |
| `serviceAccount.create`                       | Specifies whether a ServiceAccount should be created                                                                     | `true`  |
| `serviceAccount.name`                         | The name of the ServiceAccount to use.                                                                                   | `""`    |
| `serviceAccount.annotations`                  | Additional Service Account annotations (evaluated as a template)                                                         | `{}`    |
| `serviceAccount.automountServiceAccountToken` | Automount service account token for the server service account                                                           | `true`  |
| `metrics.enabled`                             | Enable the export of Prometheus metrics                                                                                  | `false` |
| `metrics.serviceMonitor.enabled`              | if `true`, creates a Prometheus Operator ServiceMonitor (also requires `metrics.enabled` to be `true`)                   | `false` |
| `metrics.serviceMonitor.path`                 | Specifies HTTP path to scrape for metrics. If empty, Prometheus uses the default value (e.g. /metrics).                  | `""`    |
| `metrics.serviceMonitor.port`                 | Name of the service port this endpoint refers to                                                                         | `nil`   |
| `metrics.serviceMonitor.targetPort`           | Name or number of the target port of the Pod behind the Service, the port must be specified with container port property | `nil`   |
| `metrics.serviceMonitor.namespace`            | Namespace in which Prometheus is running                                                                                 | `""`    |
| `metrics.serviceMonitor.annotations`          | Additional custom annotations for the ServiceMonitor                                                                     | `{}`    |
| `metrics.serviceMonitor.labels`               | Extra labels for the ServiceMonitor                                                                                      | `{}`    |
| `metrics.serviceMonitor.jobLabel`             | The name of the label on the target service to use as the job name in Prometheus                                         | `""`    |
| `metrics.serviceMonitor.honorLabels`          | honorLabels chooses the metric's labels on collisions with target labels                                                 | `false` |
| `metrics.serviceMonitor.interval`             | Interval at which metrics should be scraped.                                                                             | `""`    |
| `metrics.serviceMonitor.scrapeTimeout`        | Timeout after which the scrape is ended                                                                                  | `""`    |
| `metrics.serviceMonitor.metricRelabelings`    | Specify additional relabeling of metrics                                                                                 | `[]`    |
| `metrics.serviceMonitor.relabelings`          | Specify general relabeling                                                                                               | `[]`    |
| `metrics.serviceMonitor.selector`             | Prometheus instance selector labels                                                                                      | `{}`    |


