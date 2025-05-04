#!/bin/bash

declare -A files=(
    ["2d-2e20.sym.egr"]="https://userweb.cs.txstate.edu/~burtscher/research/ECLgraph/2d-2e20.sym.egr"
    ["USA-road-d.NY.egr"]="https://userweb.cs.txstate.edu/~burtscher/research/ECLgraph/USA-road-d.NY.egr"
    ["internet.egr"]="https://userweb.cs.txstate.edu/~burtscher/research/ECLgraph/internet.egr"
    ["rmat16.sym.egr"]="https://userweb.cs.txstate.edu/~burtscher/research/ECLgraph/rmat16.sym.egr"
    ["citationCiteseer.egr"]="https://userweb.cs.txstate.edu/~burtscher/research/ECLgraph/citationCiteseer.egr"
    ["coPapersDBLP.egr"]="https://userweb.cs.txstate.edu/~burtscher/research/ECLgraph/coPapersDBLP.egr"
)

download_dir="./experiment-graphs/"

mkdir -p "$download_dir"

for file_name in "${!files[@]}"; do
    url="${files[$file_name]}"
    file_path="$download_dir/$file_name"

    if [ -f "$file_path" ]; then
        echo "File $file_name already exists. Skipping download."
    else
        echo "Downloading $file_name from $url to $download_dir..."
        wget -q -O "$file_path" "$url"

        if [ $? -eq 0 ]; then
            echo "Successfully downloaded $file_name"
        else
            echo "ERROR: Failed to download $file_name"
        fi
    fi

done

echo "Download complete!"
