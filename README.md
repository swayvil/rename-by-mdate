# rename-by-mdate
## Usage
```
go build
./rename-by-mdate <source directory path> <target directory path> [optional suffix]
```
## Example
"~/Pictures/holidays" folder contains images with default names:

```
» ls -l ~/Pictures/holidays
total 16464
-rw-r--r--@ 1 swayvil  staff  1816906 29 jui 16:00 IMG_2305.jpeg
-rw-r--r--@ 1 swayvil  staff  1803723 29 jui 16:00 IMG_2306.jpeg
-rw-r--r--@ 1 swayvil  staff  1811887 29 jui 16:00 IMG_2307.jpeg
-rw-r--r--@ 1 swayvil  staff  2488200 29 jui 16:00 IMG_2308.jpeg
```

Run "rename-by-mdate" with the folder path in parameter:
```
» ./rename-by-mdate ~/Pictures/holidays
```

Images are renamed with the date-time of the file, suffixed by the folder name:
```
» ls -l ~/Pictures/holidays
total 16464
-rw-r--r--@ 1 swayvil  staff  1816906 29 jui 16:00 2020-06-29_16.00.05_holidays.jpeg
-rw-r--r--@ 1 swayvil  staff  1803723 29 jui 16:00 2020-06-29_16.00.06_holidays.jpeg
-rw-r--r--@ 1 swayvil  staff  1811887 29 jui 16:00 2020-06-29_16.00.08_holidays.jpeg
-rw-r--r--@ 1 swayvil  staff  2488200 29 jui 16:00 2020-06-29_16.00.08_holidays_1.jpeg
```

Optionaly, "rename-by-mdate" can be run with a suffix:
```
» ./rename-by-mdate ~/Pictures/holidays ~/Pictures/timeline family
```

Images are renamed with the date-time of the file, followed by the suffix:

```
» ls -l ~/Pictures/timeline
total 16464
-rw-r--r--@ 1 swayvil  staff  1816906 29 jui 16:00 2020-06-29_16.00.05_family.jpeg
-rw-r--r--@ 1 swayvil  staff  1803723 29 jui 16:00 2020-06-29_16.00.06_family.jpeg
-rw-r--r--@ 1 swayvil  staff  1811887 29 jui 16:00 2020-06-29_16.00.08_family.jpeg
-rw-r--r--@ 1 swayvil  staff  2488200 29 jui 16:00 2020-06-29_16.00.08_family_1.jpeg
```

## Useful Linux commands
### Copy files to a server and keep the timestamps
```
rsync -a -P -e "ssh -t -p <PORT>" /local/path <USER>@<HOST>:/target/path
```

### Find all files which are not images or videos
```
find . -type f -not \( -iname \*.jpg -o -iname \*.png -o -iname \*.jpeg -o -iname \*.bmp -o -iname \*.mov -o -iname \*.heic -o -iname \*.mp4 -o -iname \*.avi -o -iname \*.gif -o -iname \*.mpg -o -iname \*.tif -o -iname \*.wmv -o -iname \*.wav \) > ../verif.txt
```

### Find files and run commands
```
find . -type f \( -name "*.heic" -o -name "*.jpeg" \) -exec cp -rp "{}" /path  ';'
find . -type f \( -iname \*.dcm \) -exec mv '{}' /path ';'
find . -type f \( -iname \*.jpe \) -exec rm -f '{}' ';'
```