import yaml
import os
import logging

logging.basicConfig()
logger = logging.getLogger()
logger.setLevel(logging.INFO)


def load_config(file_path):
    with open(file_path, "r") as f:
        return yaml.full_load(f)


def copy_file(source_file_path, target_file_path):
    try:
        with open(source_file_path, "r") as source_file:
            source_content = source_file.read()

        with open(target_file_path, "w") as target_file:
            target_file.write(source_content)
        logger.info(f"copied content from '{source_file_path}' to '{target_file_path}'")

        source_file.close()
        target_file.close()
    except FileNotFoundError:
        logger.warning(f"file '{source_file_path}' not found")
        raise FileNotFoundError
    except PermissionError:
        logger.warning(f"permission denied for file '{source_file_path}'")
        raise PermissionError
    except IOError as e:
        logger.warning(f"{e} for file '{source_file_path}'")
        raise IOError
    except Exception as e:
        logger.warning(f"{e} occurred while copying file '{source_file_path}'")
        raise Exception


def main():
    config = load_config("config.yaml")
    root = config.get("gitea-repo-path")
    files = config.get("files", {})

    try:
        for file_name, target_file_path in files.items():
            source_file_path = os.path.join("files", file_name)
            full_target_file_path = os.path.join(root, target_file_path)

            if os.path.isfile(full_target_file_path):
                copy_file(source_file_path, full_target_file_path)
            else:
                logger.warn(
                    f"target file '{full_target_file_path}' does not exist skipping"
                )
        logger.info("all files synced successfully")
    except:
        logger.warning("failed to sync all files")


if __name__ == "__main__":
    main()
