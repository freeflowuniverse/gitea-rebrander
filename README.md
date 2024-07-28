# gitea-rebrander

Script that syncs changes to gitea repository automatically.

## Usage

1- Clone the repo:

```bash
git@github.com:freeflowuniverse/gitea-rebrander.git
```

2- Configure config.yaml file

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

4- Run the script

```bash
make run
```
