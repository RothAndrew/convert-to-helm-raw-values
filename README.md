# convert-to-helm-raw-values

## What

convert-to-helm-raw-values converts a standard Kubernetes YAML document into one that can be used by the Helm chart [incubator/raw](https://github.com/helm/charts/tree/master/incubator/raw)

## Why

Sometimes, particularly when dealing with Kubernetes Operators, the instructions say to just apply a kubernetes yaml file. [Here's an example from the docs website for the Mattermost Operator](https://docs.mattermost.com/install/install-kubernetes.html):

```sh
kubectl create ns mattermost-operator
kubectl apply -n mattermost-operator -f https://raw.githubusercontent.com/mattermost/mattermost-operator/master/docs/mattermost-operator/mattermost-operator.yaml
```

This is fine, but what if your Kubernetes development and deployment workflow centers around deploying and managing Helm charts and/or [Helmfiles](https://github.com/roboll/helmfile)? The wonderful little chart `incubator/raw` takes in arbitrary K8s yaml and applies it using the standard Helm format. However, the format for a values file for incubator/raw is different from that of standard Kubernetes YAML.

Kubernetes YAML:

```yaml
a: Easy!
b:
  c: 2
  d: [3, 4]
---
a: Easy!
b:
  c: 2
  d: [3, 4]
```

Needs to be converted to a format that incubator/raw accepts:

```yaml
resources:
  - a: Easy!
    b:
      c: 2
      d:
        - 3
        - 4
  - a: Easy!
    b:
      c: 2
      d:
        - 3
        - 4
```

## How

Using the same Mattermost Operator example as above, you can now do this:

```sh
wget https://raw.githubusercontent.com/mattermost/mattermost-operator/master/docs/mattermost-operator/mattermost-operator.yaml
convert-to-helm-raw-values -i mattermost-operator.yaml -o mattermost-operator-values.yaml
helm install --name mattermost-operator incubator/raw -f mattermost-operator-values.yaml
```

## Usage

```sh
$ convert-to-helm-raw-values --help
Converts K8s YAML to a version that helm incubator/raw can use

Usage:
  convert-to-helm-raw-values [flags]

Flags:
  -h, --help             help for convert-to-helm-raw-values
  -i, --infile string    Input file - Must be compliant with K8s YAML
  -o, --outfile string   Output file - Will be formatted such that it can be used as values.yaml for helm chart incubator/raw. Will always overwrite if the file already exists.
```

## Contributor Guide

### Releasing new versions

To release a new version, just push a new tag. Automation will cover the rest. Make sure the tag conforms to SemVer. [goreleaser](https://goreleaser.com/) is used in the pipeline to create a GitHub Release.
