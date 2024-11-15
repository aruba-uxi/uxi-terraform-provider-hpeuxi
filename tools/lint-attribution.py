import glob
import sys
from datetime import datetime

current_year = datetime.now().year

notice = f"""
/*
Copyright {current_year} Hewlett Packard Enterprise Development LP.
*/
"""

def check_files_for_word(directory, word):
    go_files = glob.glob(f'{directory}/**/*.go', recursive=True)
    for file_path in go_files:
        with open(file_path, 'r') as file:
            content = file.read()
            if word not in content:
                print(f"Error: '{word}' was not found in file '{file_path}'")
                sys.exit(1)
    print("All files contain the word.")

def check_and_add_copyright(directory, word):
    go_files = glob.glob(f'{directory}/**/*.go', recursive=True)
    for file_path in go_files:
        with open(file_path, 'r+') as file:
            content = file.read()
            if word not in content:
                print(f"Adding copyright notice to file '{file_path}'")
                file.seek(0, 0)
                file.write(f"{notice}\n" + content)
    print("All files checked and updated if necessary.")

def main():
    if len(sys.argv) < 2:
        print("Usage: python lint-attribution.py <lint|format> <directory>")
        sys.exit(1)

    command = sys.argv[1]
    directory = sys.argv[2] if len(sys.argv) > 2 else '.'

    if command == 'lint':
        check_files_for_word(directory, 'Hewlett Packard Enterprise Development LP.')
    elif command == 'format':
        check_and_add_copyright(directory, 'Hewlett Packard Enterprise Development LP.')
    else:
        print("Invalid command. Use 'lint' or 'format'.")
        sys.exit(1)

if __name__ == "__main__":
    main()
