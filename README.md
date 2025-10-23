# SLSA Generator Demo
Demonstration of `slsa-github-generator`.

- [slsa-github-generator](https://github.com/slsa-framework/slsa-github-generator)
- [Generation of SLSA3+ provenance for container images](https://github.com/slsa-framework/slsa-github-generator/blob/main/internal/builders/container/README.md)


[Getting Started](https://github.com/slsa-framework/slsa-github-generator/blob/main/internal/builders/container/README.md#getting-started)を参考にGitHub ActionsのWorkflowを作成した。
dockerのActionsが古いバージョンなので、最新のバージョンに更新した。
(そうしないと`set-output`のdeprecatedのエラーが出る)

- [成果物の構成証明と再利用可能なワークフローを使用して SLSA v1 ビルド レベル 3 を実現する](https://docs.github.com/ja/actions/how-tos/secure-your-work/use-artifact-attestations/increase-security-rating#building-with-a-reusable-workflow)
  - ビルドジョブをReusable Workflowとして実行するよう指示があるため従った

## Attestation
タグの一覧表示
Attestationは`<DIGEST>.att`という名前でコンテナレジストリに保存されている。

```bash
$ crane ls ghcr.io/mochizuki875/container-build-with-slsa-generator
main
sha256-38ec019b69846ff9b6e33a9d051d95beb89fc82a2e2f14b4d4377ef9b1534377.att 
```
```bash
$ crane digest ghcr.io/mochizuki875/container-build-with-slsa-generator:main
sha256:38ec019b69846ff9b6e33a9d051d95beb89fc82a2e2f14b4d4377ef9b1534377
```

Attestationの参照
`subject`にイメージのダイジェスト値が含まれている。
```bash
cosign download attestation ghcr.io/mochizuki875/container-build-with-slsa-generator:main | jq '.payload |= (@base64d | fromjson)'
```
<details>

<summary>Build Attestation</summary>

```json
{
  "payloadType": "application/vnd.in-toto+json",
  "payload": {
    "_type": "https://in-toto.io/Statement/v0.1",
    "predicateType": "https://slsa.dev/provenance/v0.2",
    "subject": [
      {
        "name": "ghcr.io/mochizuki875/container-build-with-slsa-generator",
        "digest": {
          "sha256": "38ec019b69846ff9b6e33a9d051d95beb89fc82a2e2f14b4d4377ef9b1534377"
        }
      }
    ],
    "predicate": {
      "builder": {
        "id": "https://github.com/slsa-framework/slsa-github-generator/.github/workflows/generator_container_slsa3.yml@refs/tags/v2.1.0"
      },
      "buildType": "https://github.com/slsa-framework/slsa-github-generator/container@v1",
      "invocation": {
        "configSource": {
          "uri": "git+https://github.com/mochizuki875/container-build-with-slsa-generator@refs/heads/main",
          "digest": {
            "sha1": "031382406ea1ccfb1f59946f11a95f7c3b68c050"
          },
          "entryPoint": ".github/workflows/image-build.yaml"
        },
        "parameters": {
          "vars": {}
        },
        "environment": {
          "github_actor": "mochizuki875",
          "github_actor_id": "37737691",
          "github_base_ref": "",
          "github_event_name": "push",
          "github_event_payload": {
            "after": "031382406ea1ccfb1f59946f11a95f7c3b68c050",
            "base_ref": null,
            "before": "10a268d38feacdd30ee153690e0032efaff69c2b",
            "commits": [
              {
                "author": {
                  "email": "mzk875@gmail.com",
                  "name": "mochizuki875",
                  "username": "mochizuki875"
                },
                "committer": {
                  "email": "mzk875@gmail.com",
                  "name": "mochizuki875",
                  "username": "mochizuki875"
                },
                "distinct": true,
                "id": "031382406ea1ccfb1f59946f11a95f7c3b68c050",
                "message": "first commit",
                "timestamp": "2025-10-23T17:54:57Z",
                "tree_id": "7aef30644b34fb5f2785418a3f252f268444716e",
                "url": "https://github.com/mochizuki875/container-build-with-slsa-generator/commit/031382406ea1ccfb1f59946f11a95f7c3b68c050"
              }
            ],
            "compare": "https://github.com/mochizuki875/container-build-with-slsa-generator/compare/10a268d38fea...031382406ea1",
            "created": false,
            "deleted": false,
            "forced": false,
            "head_commit": {
              "author": {
                "email": "mzk875@gmail.com",
                "name": "mochizuki875",
                "username": "mochizuki875"
              },
              "committer": {
                "email": "mzk875@gmail.com",
                "name": "mochizuki875",
                "username": "mochizuki875"
              },
              "distinct": true,
              "id": "031382406ea1ccfb1f59946f11a95f7c3b68c050",
              "message": "first commit",
              "timestamp": "2025-10-23T17:54:57Z",
              "tree_id": "7aef30644b34fb5f2785418a3f252f268444716e",
              "url": "https://github.com/mochizuki875/container-build-with-slsa-generator/commit/031382406ea1ccfb1f59946f11a95f7c3b68c050"
            },
            "pusher": {
              "email": "37737691+mochizuki875@users.noreply.github.com",
              "name": "mochizuki875"
            },
            "ref": "refs/heads/main",
            "repository": {
              "allow_forking": true,
              "archive_url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator/{archive_format}{/ref}",
              "archived": false,
              "assignees_url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator/assignees{/user}",
              "blobs_url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator/git/blobs{/sha}",
              "branches_url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator/branches{/branch}",
              "clone_url": "https://github.com/mochizuki875/container-build-with-slsa-generator.git",
              "collaborators_url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator/collaborators{/collaborator}",
              "comments_url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator/comments{/number}",
              "commits_url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator/commits{/sha}",
              "compare_url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator/compare/{base}...{head}",
              "contents_url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator/contents/{+path}",
              "contributors_url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator/contributors",
              "created_at": 1761232676,
              "default_branch": "main",
              "deployments_url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator/deployments",
              "description": null,
              "disabled": false,
              "downloads_url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator/downloads",
              "events_url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator/events",
              "fork": false,
              "forks": 0,
              "forks_count": 0,
              "forks_url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator/forks",
              "full_name": "mochizuki875/container-build-with-slsa-generator",
              "git_commits_url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator/git/commits{/sha}",
              "git_refs_url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator/git/refs{/sha}",
              "git_tags_url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator/git/tags{/sha}",
              "git_url": "git://github.com/mochizuki875/container-build-with-slsa-generator.git",
              "has_discussions": false,
              "has_downloads": true,
              "has_issues": true,
              "has_pages": false,
              "has_projects": true,
              "has_wiki": true,
              "homepage": null,
              "hooks_url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator/hooks",
              "html_url": "https://github.com/mochizuki875/container-build-with-slsa-generator",
              "id": 1081989864,
              "is_template": false,
              "issue_comment_url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator/issues/comments{/number}",
              "issue_events_url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator/issues/events{/number}",
              "issues_url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator/issues{/number}",
              "keys_url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator/keys{/key_id}",
              "labels_url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator/labels{/name}",
              "language": "Go",
              "languages_url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator/languages",
              "license": null,
              "master_branch": "main",
              "merges_url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator/merges",
              "milestones_url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator/milestones{/number}",
              "mirror_url": null,
              "name": "container-build-with-slsa-generator",
              "node_id": "R_kgDOQH3a6A",
              "notifications_url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator/notifications{?since,all,participating}",
              "open_issues": 0,
              "open_issues_count": 0,
              "owner": {
                "avatar_url": "https://avatars.githubusercontent.com/u/37737691?v=4",
                "email": "37737691+mochizuki875@users.noreply.github.com",
                "events_url": "https://api.github.com/users/mochizuki875/events{/privacy}",
                "followers_url": "https://api.github.com/users/mochizuki875/followers",
                "following_url": "https://api.github.com/users/mochizuki875/following{/other_user}",
                "gists_url": "https://api.github.com/users/mochizuki875/gists{/gist_id}",
                "gravatar_id": "",
                "html_url": "https://github.com/mochizuki875",
                "id": 37737691,
                "login": "mochizuki875",
                "name": "mochizuki875",
                "node_id": "MDQ6VXNlcjM3NzM3Njkx",
                "organizations_url": "https://api.github.com/users/mochizuki875/orgs",
                "received_events_url": "https://api.github.com/users/mochizuki875/received_events",
                "repos_url": "https://api.github.com/users/mochizuki875/repos",
                "site_admin": false,
                "starred_url": "https://api.github.com/users/mochizuki875/starred{/owner}{/repo}",
                "subscriptions_url": "https://api.github.com/users/mochizuki875/subscriptions",
                "type": "User",
                "url": "https://api.github.com/users/mochizuki875",
                "user_view_type": "public"
              },
              "private": false,
              "pulls_url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator/pulls{/number}",
              "pushed_at": 1761242099,
              "releases_url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator/releases{/id}",
              "size": 15,
              "ssh_url": "git@github.com:mochizuki875/container-build-with-slsa-generator.git",
              "stargazers": 0,
              "stargazers_count": 0,
              "stargazers_url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator/stargazers",
              "statuses_url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator/statuses/{sha}",
              "subscribers_url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator/subscribers",
              "subscription_url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator/subscription",
              "svn_url": "https://github.com/mochizuki875/container-build-with-slsa-generator",
              "tags_url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator/tags",
              "teams_url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator/teams",
              "topics": [],
              "trees_url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator/git/trees{/sha}",
              "updated_at": "2025-10-23T17:52:25Z",
              "url": "https://api.github.com/repos/mochizuki875/container-build-with-slsa-generator",
              "visibility": "public",
              "watchers": 0,
              "watchers_count": 0,
              "web_commit_signoff_required": false
            },
            "sender": {
              "avatar_url": "https://avatars.githubusercontent.com/u/37737691?v=4",
              "events_url": "https://api.github.com/users/mochizuki875/events{/privacy}",
              "followers_url": "https://api.github.com/users/mochizuki875/followers",
              "following_url": "https://api.github.com/users/mochizuki875/following{/other_user}",
              "gists_url": "https://api.github.com/users/mochizuki875/gists{/gist_id}",
              "gravatar_id": "",
              "html_url": "https://github.com/mochizuki875",
              "id": 37737691,
              "login": "mochizuki875",
              "node_id": "MDQ6VXNlcjM3NzM3Njkx",
              "organizations_url": "https://api.github.com/users/mochizuki875/orgs",
              "received_events_url": "https://api.github.com/users/mochizuki875/received_events",
              "repos_url": "https://api.github.com/users/mochizuki875/repos",
              "site_admin": false,
              "starred_url": "https://api.github.com/users/mochizuki875/starred{/owner}{/repo}",
              "subscriptions_url": "https://api.github.com/users/mochizuki875/subscriptions",
              "type": "User",
              "url": "https://api.github.com/users/mochizuki875",
              "user_view_type": "public"
            }
          },
          "github_head_ref": "",
          "github_ref": "refs/heads/main",
          "github_ref_type": "branch",
          "github_repository_id": "1081989864",
          "github_repository_owner": "mochizuki875",
          "github_repository_owner_id": "37737691",
          "github_run_attempt": "1",
          "github_run_id": "18757361050",
          "github_run_number": "21",
          "github_sha1": "031382406ea1ccfb1f59946f11a95f7c3b68c050"
        }
      },
      "metadata": {
        "buildInvocationID": "18757361050-1",
        "completeness": {
          "parameters": true,
          "environment": false,
          "materials": false
        },
        "reproducible": false
      },
      "materials": [
        {
          "uri": "git+https://github.com/mochizuki875/container-build-with-slsa-generator@refs/heads/main",
          "digest": {
            "sha1": "031382406ea1ccfb1f59946f11a95f7c3b68c050"
          }
        }
      ]
    }
  },
  "signatures": [
    {
      "keyid": "",
      "sig": "MEQCIDNExxYg1nHzmyxjELZxdSMUeXXaMVYtDpHERy8xxNoDAiBl4Hh4qzTvlhe03V8Li+/SzUmKHVe/6fAtV/kNxnakRQ=="
    }
  ]
}
```

</details>


## SLSA Verification

イメージのProvenanceを検証する時は、イメージのダイジェスト値を指定する必要がある。
```bash
IMAGE=ghcr.io/mochizuki875/container-build-with-slsa-generator:main
IMAGE="${IMAGE}@"$(crane digest "${IMAGE}")
```


Attestationに含まれる署名を検証し、期待するBuilderによって作成されたものであることを検証する。
OPを付けることで、ブランチやリポジトリタグとかも検証に含めることができる。

```bash
slsa-verifier verify-image "$IMAGE" \
    --source-uri github.com/mochizuki875/container-build-with-slsa-generator
```
```bash
Verified build using builder "https://github.com/slsa-framework/slsa-github-generator/.github/workflows/generator_container_slsa3.yml@refs/tags/v2.1.0" at commit 031382406ea1ccfb1f59946f11a95f7c3b68c050
PASSED: SLSA verification passed
```

## 透過ログの確認
GitHub Actionsの`Create and sign provenance`ステップに`tlog entry created with index: xxxxxxx`の形でindexが出力されるので、そのindexを使ってrekorで透過ログを確認できる。

- [rekor](https://search.sigstore.dev/)






