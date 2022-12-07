with open("inputs_day7.txt", "r") as inputs:
    day7_inputs = inputs.read()

terminal_history = day7_inputs.split("\n")
# print(terminal_history)
file_tree = {}
current_directory = "/"
layers_deep = 0
path = ""

for cmd in terminal_history:
    # command processing
    if cmd[0] == "$":
        if ".." in cmd:
            for idx, char in enumerate(path[:-1]):
                if char == "/":
                    last_slash = idx
            path = path[:last_slash+1]
            # print("path change with ..! new path:")
            # print(path)
        elif "cd" in cmd:
            change_directory = cmd[5:]
            if change_directory == "/":
                path = change_directory
            else:
                path = path + change_directory + "/"
            file_tree[path] = []
            print("path change! new path:")
            print(path)
        elif "ls" in cmd:
            pass
    else:
        file_tree[path].append(cmd.split(" "))

print(file_tree)

file_sizes = {}
all_directories = []
subdirectories = {}
# iterate through all directories in key
for key in file_tree:
    all_directories.append(key)

for folder in all_directories:
    subfolders = 0
    for idx, char in enumerate(folder):
        if char == "/":
            subfolders += 1
    subdirectories[folder] = subfolders-1

maxsubs = 0
for key in subdirectories.keys():
    if subdirectories[key] > maxsubs:
        maxsubs = subdirectories[key]

i = maxsubs
while i != -1:
    current_level_dir = []
    for key in subdirectories.keys():
        if subdirectories[key] == i:
            current_level_dir.append(key)
    folder_size = 0
    for folder in current_level_dir:
        print("scanning:")
        print(folder)
        for entry in file_tree[folder]:
            if entry[0] != "dir":
                folder_size += int(entry[0])
            if entry[0] == "dir":
                folder_size += int(entry[2])
        for idx, char in enumerate(folder[:-1]):
                if char == "/":
                    last_slash = idx
        folder_to_find = folder.split("/")[-2]
        path = folder[:last_slash+1]
        print("parent_path:")
        print(path)
        print("folder to find:")
        print(folder_to_find)
        print("inside:")
        print(file_tree[path])
        for entry in file_tree[path]:
            if entry[1] == folder_to_find:
                print("Found it!")
                print(entry)
                entry.append(folder_size)
                print(entry)
        print(file_tree[folder])
    i -= 1

finalSoln = 0
print("TO DELETE")
for key in file_tree.keys():
    for subdir in file_tree[key]:
        if subdir[0] == "dir":
            if subdir[2] < 100000:
                finalSoln += subdir[2]
                print(subdir[1])
print("============================================================")
print(file_tree)
print(finalSoln)











