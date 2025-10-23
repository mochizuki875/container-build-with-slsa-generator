# SLSA Generator Demo
Demonstration of `slsa-github-generator`.

- [slsa-github-generator](https://github.com/slsa-framework/slsa-github-generator)
- [Generation of SLSA3+ provenance for container images](https://github.com/slsa-framework/slsa-github-generator/blob/main/internal/builders/container/README.md)


[Getting Started](https://github.com/slsa-framework/slsa-github-generator/blob/main/internal/builders/container/README.md#getting-started)を参考にGitHub ActionsのWorkflowを作成した。
dockerのActionsが古いバージョンなので、最新のバージョンに更新した。
(そうしないと`set-output`のdeprecatedのエラーが出る)


## Attestation
タグの一覧表示
Attestationは`<DIGEST>.att`という名前でコンテナレジストリに保存されている。

```bash
$ crane ls ghcr.io/mochizuki875/slsa-generator-demo
main
<DIGEST>.att
```
```bash
$ crane digest ghcr.io/mochizuki875/slsa-generator-demo:main
<DIGEST>
```


Attestationの参照
`subject`にイメージのダイジェスト値が含まれている。
```bash
cosign download attestation ghcr.io/mochizuki875/slsa-generator-demo:main | jq '.payload |= (@base64d | fromjson)'
```
```json

```

Provenanceの確認
`subject`にイメージのダイジェスト値が含まれている。
```bash
cosign download attestation ghcr.io/mochizuki875/slsa-generator-demo:main | jq -r '.payload' | base64 -d | jq . 
```


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

```

※`xxxxxxxxxxxxxx`はリポジトリ(`mochizuki875/slsa-generator-demo`)のコミットID




## 透過ログの確認
GitHub Actionsの`Create and sign provenance`ステップに`tlog entry created with index: xxxxxxx`の形でindexが出力されるので、そのindexを使ってrekorで透過ログを確認できる。

- [rekor](https://search.sigstore.dev/)






