# Foxyfy
![build](https://github.com/danjhl/foxyfy/actions/workflows/build.yml/badge.svg)

## Commands

### Listing bookmarks in bookmark directory
`foxyfy ls -b music -db ../path/to/places.sqlite`

Bookmarks:
```
music
  |_ song1 (https://youtube.com?v=abc1)
  |_ subdir
      |_ song2 (https://youtube.com?v=abc2)
      |_ song3 (https://youtube.com?v=abc3)
```

Console Output:
```
song1 - https://youtube.com?v=abc1
---- subdir ----
song2 - https://youtube.com?v=abc2 
song3 - https://youtube.com?v=abc3
```