# action-docker-manifest

Github Action to create a docker manifest from several docker images

## Inputs

### `images`

**Required** List of images to include in the manifest.

### `manifest`

**Required** Name of the manifest to create.

### `push`

**Optional** Push the manifest to the registry. Default `true`.

## Example usage

Potential usage:

```yaml
uses: actions/action-docker-manifest@v1
with:
    images: |
        myorg/myimage-linux-arm64:1.0.0
        myorg/myimage-linux-amd64:1.0.0
    manifest: myorg/myimage:1.0.0
    push: true
```
