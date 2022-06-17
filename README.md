# VERSIONER
CLI application that reads and/or edits version from specific types of source code files:
- Maven pom.xml (version)
- Node package.json (version)
- OAS3 specification.yaml (version)
- Helm Chart.yaml (version, appVersion)

## GitHub Action

### Usage

#### Read Version

```yaml

- uses: raitonbl.com/versioner
  with:
    command: get
    runtime: helm
    object : version
    file   : specification.yaml
```

#### Edit Version

```yaml

- uses: raitonbl.com/versioner
  with:
    command: set
    runtime: helm
    object : version
    file   : specification.yaml
```

#### Set Stamped Version

```yaml

- uses: raitonbl.com/versioner
  with:
    command: set-stamped-version
    runtime: helm
    object : version
    file   : specification.yaml
```