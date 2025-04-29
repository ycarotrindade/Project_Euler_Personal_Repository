import os
import json

ROOT = "."
output = []

for folder in os.listdir(ROOT):
    folder_path = os.path.join(ROOT, folder)
    if os.path.isdir(folder_path):
        entry = {
            "name": folder.replace(" ", "_"),
            "python": f"https://github.com/ycarotrindade/Project_Euler_Personal_Repository/tree/main/{folder}/Python",
            "golang": f"https://github.com/ycarotrindade/Project_Euler_Personal_Repository/tree/main/{folder}/Golang"
        }
        output.append(entry)

with open("puzzles.json", "w") as f:
    json.dump(output, f, indent=4)
