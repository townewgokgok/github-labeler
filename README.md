github-labeler
==============

CLI that sets GitHub labels exactly as written in YAML file

## Abilities added in this fork version

- Can be used for Github Enterprise by `-base-url` option.
- API token can also be specified by `-token` option.
- Target repositories can also be specified by `-repo` option.
- All labels defined in `labels` will be applied if `.repos[].labels` is omitted.
- Label definitions can also be retrieved from the source repository specified by the `-copy-repo` option like:
  
  ```bash
  github-labeler -copy-repo foo/bar -repo foo/baz
  ```
  If the `-repo` is not specified, the retrieved definition is simply displayed.

## Usage

```console
$ go build ./cmd/github-labeler
$ ./github-labeler
```

1. Write `labels.yaml`.
2. Do `read -s TOKEN` and input your API token of Github (or Github Enterprise).
3. Use like this:
   
   ```bash
   github-labeler -base-url 'https://github-enterprise.example.com/api/v3/' \
     -token "$TOKEN" \
     -manifest labels.yaml \
     -repo foo/bar
   ```
   `-dry-run` option is also available.

## What this app does

- Create a label (e.g. when no label described in YAML)
- Edit a label (e.g. when its color was changed)
- Delete a label (e.g. when the label not described in YAML exists on GitHub)

## YAML example

```yaml
labels:
  - name: kind/proactive
    description: Categorizes issue or PR as related to proactive tasks.
    color: 9CCC65
  - name: kind/reactive
    description: Categorizes issue or PR as related to reactive tasks.
    color: FFA000

repos:
  - name: org/repo
    labels:
      - kind/proactive
      - kind/reactive
```

## Author

- @b4b4r07 (the original author)
- @townewgokgok (the author of this fork version)

## License

MIT
