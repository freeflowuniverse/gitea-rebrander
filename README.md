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
├── main.py
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

3- Install requirements

```bash
pip install -r requirements.txt
```

4- Run

```bash
python main.py
```

5- To apply new logo and icon before building gitea and running it, cd to gitea:

```bash
cd <path-to-gitea-repo>
```

6- Run

```bash
make generate-images
```

---

# Form manual rebranding at [Gitea source repository](https://github.com/go-gitea/gitea?tab=readme-ov-file)

- ### Changing the logo or Favicon:

To build a custom logo and/or favicon clone the Gitea source repository, replace `assets/logo.svg` and/or `assets/favicon.svg` and run

```bash
make generate-images
```

- ### Changing App Name :
  To change the app name(from Gitea to any name) at the page title and all website go to `/custom/conf/app.ini` and change it to ex:ThreeFold

![image](https://github.com/user-attachments/assets/e22f6595-404e-4c24-acdf-dca5cc5ac26f)

- ### Update home page :

  any change at home can be done here `/templates/home.tmpl` you can add or remove content.

- ### Update Footer :
  Any change to the Footer content can be done here `/templates/base/footer_content.tmpl` you can add or remove content.
