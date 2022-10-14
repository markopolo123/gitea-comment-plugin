# gitea-comment-plugin

A Woodpecker plugin to post comments onto a Gitea Pull Request.

Note this currently only works on `pull request` events.

## Usage/Examples

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

Produces something with looks like the screenshot below on pull requests:

![comments](img/comments.png)

## Authors

* [@markopolo123](https://www.github.com/markopolo123)

## Running just the container

```bash
docker run \
-e PLUGIN_COMMENT="test comment" \
-e PLUGIN_GITEA_TOKEN="tokenhere" \
-e PLUGIN_GITEA_ADDRESS="https://gitea.url.here" \
-e CI_REPO_OWNER="repoowwer" \
-e CI_REPO_NAME="yourrepo" \
-e CI_COMMIT_PULL_REQUEST=8 \
test-gitea
```
