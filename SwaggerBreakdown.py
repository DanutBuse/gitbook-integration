import yaml  # Import PyYAML
import json
import os       

def updateSwaggers(file, directory):
    # Load OpenAPI YAML
    file_path = directory + "/" + path;
    with open(file_path, "r") as f:
        openapi = yaml.safe_load(f)  # Convert YAML to Python dictionary
    
    swaggers_content = ""
    
    for path in openapi["paths"]:
        for method in openapi["paths"][path]:
            swaggers_content += f"{{% openapi src=\"{file_path}\" path=\"{path}\" method=\"{method}\" expanded=\"true\" %}}\n"
            swaggers_content += f"{{% endopenapi %}}\n\n"
    
    # Write the updated README.md
    with open("swaggers.md", "w") as f:
        f.write(swaggers_content)
    
    print("{file} processed successfully!")
    
def get_child_dirs_and_files(parent_dir="."):
    """
    Traverses all child directories of the given parent directory.
    Returns a dictionary where keys are directory names and values are lists of files inside them.
    """
    result = {}

    for root, dirs, files in os.walk(parent_dir):
        # Skip the root directory itself
        if root == parent_dir:
            continue

        # Get relative path name of the child directory
        relative_dir = os.path.relpath(root, parent_dir)
        result[relative_dir] = files

    return result

# Process all openapi files
parent_directory = "."  # Change this to any directory path
directories_and_files = get_child_dirs_and_files(parent_directory)

for dir_name, files in directories_and_files.items():
    print(f"Directory: {dir_name}")
    for file in files:
        print(f"  - {file}")
        if file == "openapi.yaml":
            updateSwaggers(file, dir_name)
