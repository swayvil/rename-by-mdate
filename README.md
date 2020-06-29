# rename-by-mdate
## Usage
rename-by-mdate takes the root directory path in parameter
```
go build
./rename-by-mdate <Path> [Optional prefixe]
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

Images are renamed with the folder name suffixed by the images date-time:
```
» ls -l ~/Pictures/holidays
total 16464
-rw-r--r--@ 1 swayvil  staff  1816906 29 jui 16:00 holidays_2020-06-29_16.00.05.jpeg
-rw-r--r--@ 1 swayvil  staff  1803723 29 jui 16:00 holidays_2020-06-29_16.00.06.jpeg
-rw-r--r--@ 1 swayvil  staff  1811887 29 jui 16:00 holidays_2020-06-29_16.00.08.jpeg
-rw-r--r--@ 1 swayvil  staff  2488200 29 jui 16:00 holidays_2020-06-29_16.00.08_1.jpeg
```

Optionaly, we can run "rename-by-mdate" with a prefixe:
```
» ./rename-by-mdate ~/Pictures/holidays family
```

Images are renamed with the prefixe followed by the images date-time:

```
» ls -l ~/Pictures/holidays
total 16464
-rw-r--r--@ 1 swayvil  staff  1816906 29 jui 16:00 family_2020-06-29_16.00.05.jpeg
-rw-r--r--@ 1 swayvil  staff  1803723 29 jui 16:00 family_2020-06-29_16.00.06.jpeg
-rw-r--r--@ 1 swayvil  staff  1811887 29 jui 16:00 family_2020-06-29_16.00.08.jpeg
-rw-r--r--@ 1 swayvil  staff  2488200 29 jui 16:00 family_2020-06-29_16.00.08_1.jpeg
```