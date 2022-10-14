---
name: Gitea Comment
authors: markopolo123
description: Plugin to add comments to a Gitea Pull Request
tags: [Gitea, comment]
containerImage: mcs94/gitea-comment
containerImageUrl: https://hub.docker.com/r/mcs94/gitea-comment
url: https://github.com/markopolo123/gitea-comment-plugin
---

A Woodpecker plugin to post comments onto a Gitea Pull Request.

## Usage

Example pipeline:

```yaml
pipeline:
  comment:
    image: mcs94/gitea-comment
    settings:
      gitea_address: https://gitea.url.goes.here
      gitea_token:
        from_secret: gitea_token
      comment: >
        ✅ Build ${CI_BUILD_EVENT} of `${CI_REPO_NAME}` has status `${CI_BUILD_STATUS}`.

        📝 Commit by ${CI_COMMIT_AUTHOR} on `${CI_COMMIT_BRANCH}`:

        `${CI_COMMIT_MESSAGE}`

        🌐 ${CI_BUILD_LINK}
    when:
      event: [pull_request]
```

## Settings

|Name|Description|Default|
|---|---|---|
| `gitea_address` |URL for your gitea instance| none|
| `gitea_token` |Gitea API token| none|
| `comment` |comment to add to Pull Request|none|
