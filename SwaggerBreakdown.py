import yaml  # Import PyYAML
import json

# Load OpenAPI YAML
with open("openapi.yaml", "r") as f:
    openapi = yaml.safe_load(f)  # Convert YAML to Python dictionary

readme_content = "# API Documentation\n\n"

for path in openapi["paths"]:
    for method in openapi["paths"][path]:
        readme_content += f"{{% openapi src=\"./openapi.yaml\" path=\"{path}\" method=\"{method}\" expanded=\"true\" %}}\n\n"

# Write the updated README.md
with open("README.md", "w") as f:
    f.write(readme_content)

print("README.md updated successfully!")
