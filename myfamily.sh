#curl -s https://learn.reboot01.com/assets/superhero/all.json | jq ' .[] | select(.id == 1) | .relatives '
curl -s https://learn.reboot01.com/assets/superhero/all.json | jq '.[] | select(.id == 1) | .relatives'