# fratbots

## Level creation

* Draw the level in Charaster.
* Save as plain text.
* Clean from non-printable characters: tr -cd '[:print:]' < level.txt > clean.txt
* Create header in the first line of map file (clean.txt) in the following format: map|80x40
(according to width and height of the map respectively).
