# Country Flag Downloader

> This utility scrapes [countryflags.com](https://www.countryflags.com/) for the flag images it uses

This utility will download country flags in one or many different shapes. The output can either be in plain PNG files or in Base-64 encoding. You can also generate JSON lists of Base-64 encoded country flag images fairly easily with this utility.

## Building

> You must have [Go](https://golang.org/) installed in order to build this utility.

Run the following command to build the binary.

```sh
$ go get -u "github.com/nfnt/resize"
$ go build gen-flags.go
```

## Basic Execution

The following command will download all images for all countries of all shapes and place them in the `flags` directory. The directory will be automatically created if it does not exist.

```sh
./gen-flags -download -output-dir=./flags/
```

## Filters

You can filter which countries and shapes are downloaded using the following command line parameters

```sh
-filter-countries=country1,country2
-filter-shapes=flag-800,round-250
```

To get a list of valid shapes and countries run the following

```sh
./gen-flags -list-shapes    # for list of valid shapes
./gen-flags -list-countries # for list of valid countries
```

## Changing Image Scale

> Note: When scaling higher than 100 images may start to become blurry

You can change the scale percentage of output images (even when using the base-64 output types) using the following flag
```
-scale=xx # replace xx with a number greater than 0
```

## Output Types

There are several output types you can select. Decide which one you want using the `-output-type` command line argument.

- `-output-type=png`
- `-output-type=b64`
- `-output-type=b64-iso3166-numeric-json-file`
- `-output-type=b64-iso3166-alpha2-json-file`
- `-output-type=b64-iso3166-alpha3-json-file`

When using the `iso3166` output types, the result will be a json file storing the selected flags and countries defined in the following format

```json
{
    "ISO_CODE": {
        "shape1": "...base64string...",
        "shape2": "...base64string..."
    }
}
```
