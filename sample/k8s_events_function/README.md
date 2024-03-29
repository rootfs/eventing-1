# k8s_events_function

A simple function that receives Kubernetes events and prints them out the after decoding
from base64 encoding. Because we do **not** have an in-cluster event delivery mechanism yet, uses a
Knative route as an endpoint.

## Prerequisites

1. [Setup your development environment](../../DEVELOPMENT.md#getting-started)
2. [Start Knative](../../README.md#start-knative)
3. Decide on the DNS name that git can then call. Update knative/serving/domain-config.yaml domainSuffix.
For example I used aikas.org as my hostname, so my domain-config.yaml looks like so:

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: domain-config
  namespace: knative-serving-system
data:
  aikas.org: |
```

If you were already running the knative controllers, you will need reapply this config map so the cluster
will see the changes.

4. Install k8s events as an event source
```shell
ko apply -f pkg/sources/k8sevents/
```

## Running

You can deploy this to Knative from the root directory via:
```shell
ko apply -f sample/k8s_events_function/route.yaml
ko apply -f sample/k8s_events_function/configuration.yaml
```

Once deployed, you can inspect the created resources with `kubectl` commands:

```shell
# This will show the Route that we created:
kubectl get route -o yaml

# This will show the Configuration that we created:
kubectl get configurations -o yaml

# This will show the Revision that was created by our configuration:
kubectl get revisions -o yaml

# This will show the available EventSources that you can bind to:
kubectl get eventsources -oyaml

# This will show the available EventTypes that you can bind to:
kubectl get eventtypes -oyaml

```

To make this service accessible to github, we first need to determine its ingress address
(might have to wait a little while until 'ADDRESS' gets assigned):
```shell
$ watch kubectl get ingress
NAME                              HOSTS                                                                           ADDRESS           PORTS     AGE
k8s-events-function-ela-ingress   k8s-events-function.default.aikas.org,*.k8s-events-function.default.aikas.org   104.197.125.124   80        48s

```

Once the `ADDRESS` gets assigned to the cluster, you need to assign a DNS name for that IP address. This DNS address needs to be:
k8s-events-function.default.<domainsuffix you created> so for me, I would create a DNS entry from:
k8s-events-function.default.aikas.org pointing to 104.197.125.124
[Using GCP DNS](https://support.google.com/domains/answer/3290350)

So, you'd need to create an A record for k8s-events-function.default.aikas.org pointing to 104.197.125.124

To now bind the k8s_events_function for k8s system events with the function we created above, you need to
create a Bind object. Modify sample/k8s_events_function/bind.yaml to specify the namespace you want to
watch events for ('default' in this example):

```yaml
apiVersion: feeds.knative.dev/v1alpha1
kind: Bind
metadata:
  name: k8s-events-example
  namespace: default
spec:
  trigger:
    eventType: receiveevent
    resource: k8sevents/receiveevent
    service: k8sevents
    parameters:
      namespace: default
  action:
    routeName: k8s-events-example
```

Then create the binding so that you can see changes

```shell
 kubectl create -f sample/k8s_events_function/bind.yaml
```


This will create a receive_adapter that runs in the cluster and receives native k8s events
and pushes them to the consuming function.

```shell
$kubectl -n knative-eventing-system get pods

NAME                                                        READY     STATUS    RESTARTS   AGE
bind-controller-dddb99dfc-jzp7z                             1/1       Running   0          1d
sub-a3095905-f9c8-4f32-87ac-3c8fec9b51f9-85db55dc48-2mbm9   1/1       Running   0          1m

```

Then look at the logs for the function:

```shell
$ kubectl get pods
NAME                                                    READY     STATUS    RESTARTS   AGE
k8s-events-function-00001-deployment-68864b8c7d-rgx2w   3/3       Running   0          12m


# Replace k8s-events-function with the pod name from above:
$ kubectl logs k8s-events-function user-container
```

## Removing a binding

Remove the binding and things get cleaned up (including removing the receive adapter to k8s events)

```shell
kubectl delete binds k8s-events-example
```

## Cleaning up

To clean up the sample service:

```shell
ko delete -f sample/k8s_events_function/
```
