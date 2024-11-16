#!/bin/bash

# Load environment variables from .env file
if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
fi

# Advent of Code session token from .env file
YEAR=2024
BASE_URL="https://adventofcode.com"

# Loop through each day's directory (assuming directories are named day01, day02, etc.)
for day_dir in 2024/day*/; do
    # Extract the day number from the directory name (e.g., day01 -> 01)
    day=$(basename "$day_dir" | grep -oE '[0-9]+')

    # Define the input file path
    input_file="${day_dir}inputs/complete.txt"

    # Check if the input file already has contents
    if [ ! -s "$input_file" ]; then
        echo "Fetching input for Day $day..."

        # Fetch the input from Advent of Code
        response=$(curl -s -H "Cookie: session=${SESSION_TOKEN}" "${BASE_URL}/${YEAR}/day/${day}/input")

        # Check if the response contains an error message
        if [[ "$response" == *"Please don't repeatedly request this endpoint"* ]]; then
            echo "Error: Rate limit reached or unauthorized request. Please try again later."
            exit 1
        fi

        # Create the inputs directory if it doesn't exist
        mkdir -p "${day_dir}inputs"

        # Write the fetched input to the complete.txt file
        echo "$response" > "$input_file"
        echo "Input for Day $day saved to ${input_file}."
    else
        echo "Input for Day $day already exists, skipping..."
    fi
done

echo "All inputs processed."