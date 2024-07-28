# gitea-rebrander

Script that syncs changes to gitea repository automatically.

## Usage

### To Add New Files

1- Clone the repo:

```bash
git clone git@github.com:freeflowuniverse/gitea-rebrander.git
```

2- Configure `config.yaml` file

- Add absolute path of gitea repo after cloning it.
- Add the files that should be changed with their paths(where they are located in gitea repo)

### Example

```yaml
gitea-repo-path: /home/eyad/repos/gitea #absolute path to the gitea repo

files:
  # example of files will be updated
  'home.tmpl': templates/home.tmpl
  'footer_content.tmpl': templates/base/footer_content.tmpl
```

3- Add the changed or updated files to ./files/ directory

```bash
.
├── cmd
│   └── transporter.go
├── config.yaml
├── files
│   ├── footer_content.tmpl
│   └── home.tmpl
```

### To Sync Changes

1- Clone the repo:

```bash
git clone git@github.com:freeflowuniverse/gitea-rebrander.git
```

2- Configure `config.yaml` file

- Add absolute path of gitea repo after cloning it.

### Example

```yaml
gitea-repo-path: /home/eyad/repos/gitea #absolute path to the gitea repo
```

3- Run

```bash
make run
```

4- To apply new logo and icon before building gitea and running it, cd to gitea:

```bash
cd <path-to-gitea-repo>
```

5- Run

```bash
make generate-images
```
