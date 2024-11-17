#!/bin/bash

# Load environment variables from .env file
if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
fi

# Function to create directory if it does not exist
create_directory() {
    if [ ! -d "$1" ]; then
        mkdir -p "$1"
        echo "Directory $1 created."
    else
        echo "Directory $1 already exists, skipping creation."
    fi
}

# Define the year and day arguments
year=$1
day=$2

# Check if year and day arguments are provided
if [ -z "$year" ] || [ -z "$day" ]; then
    echo "Usage: $0 <year> <day>"
    exit 1
fi

# Define paths
year_dir="${year}"
day_dir="${year}/Day${day}"
scaffold_dir="scaffolds"

# Create year and day directories
create_directory "$year_dir"
create_directory "$day_dir"

# Copy scaffold files (excluding input directory and files)
for file in $scaffold_dir/*; do
    if [ -f "$file" ]; then
        base_name=$(basename "$file")
        new_name=$(echo "$base_name" | sed "s/Scaffold/Day${day}/")
        cp "$file" "$day_dir/$new_name"
        echo "Copied $(basename "$file") to $day_dir as $new_name."
    fi
done

# Call the fetch-aoc-inputs.sh script to import the input for the specific day
script_dir=$(dirname "$0")
"$script_dir/fetch_aoc_inputs.sh"

# Inform user of completion
echo "Scaffold files copied and input fetched for Year $year, Day $day."
