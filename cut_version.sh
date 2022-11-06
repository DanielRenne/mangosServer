set -e
currentDir=$(pwd)
echo "New version number (with out the v)"
read version
echo "Whats your title/description for the release?"
read title
gh release create "v$version" --title "$title" --generate-notes
