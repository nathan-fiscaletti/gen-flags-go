package main

import (
    "encoding/base64"
    "encoding/json"
    "flag"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "strings"
)

var s_shapes = [...]string{
    "flag-800",
    "flag-square-250",
    "round-250",
    "flag-wave-250",
    "flag-waving-250",
    "flag-button-square-250",
    "flag-button-round-250",
    "flag-3d-250",
    "flag-3d-round-250",
    "flag-heart-3d-250",
}

var s_countries = [...]string{
    "argentina", "antigua-and-barbuda", "austria", "australia", "armenia", "angola", "afghanistan", "azerbaijan", 
    "andorra", "algeria", "albania", "brunei", "burundi", "burkina-faso", "bulgaria", "brazil", "benin", "botswana", 
    "bosnia-and-herzegovina", "bolivia", "belize", "bahrain", "belarus", "belgium", "barbados", "bangladesh", "bhutan", 
    "bahamas", "cuba", "cote-d-ivoire", "czech-republic", "cyprus", "croatia", "comoros", "costa-rica", 
    "congo-democratic-republic-of-the", "congo-republic-of-the", "colombia", "china", "chile", "chad", 
    "central-african-republic", "cape-verde", "canada", "cameroon", "cambodia", "dominican-republic", "dominica", 
    "djibouti", "denmark", "ethiopia", "estonia", "eritrea", "equatorial-guinea", "el-salvador", "egypt", "ecuador", 
    "east-timor", "france", "finland", "fiji", "guyana", "guinea-bissau", "guinea", "guatemala", "grenada", "greece", 
    "ghana", "georgia", "gambia", "gabon", "germany", "hungary", "honduras", "haiti", "italy", "israel", "ireland", 
    "iraq", "iran", "indonesia", "india", "iceland", "jordan", "japan", "jamaica", "kyrgyzstan", "kazakhstan", "kosovo", 
    "kuwait", "kiribati", "kenya", "luxembourg", "lithuania", "liechtenstein", "libya", "liberia", "laos", "lesotho", 
    "lebanon", "latvia", "monaco", "myanmar", "mozambique", "morocco", "montenegro", "mongolia", "moldova", 
    "micronesia", "mexico", "mauritius", "mauritania", "marshall-islands", "malta", "mali", "maldives", "malaysia", 
    "malawi", "madagascar", "macedonia", "nigeria", "norway", "north-korea", "niger", "nicaragua", "new-zealand", 
    "niue", "netherlands", "nepal", "nauru", "namibia", "oman", "papua-new-guinea", "portugal", "poland", "philippines", 
    "peru", "paraguay", "panama", "pakistan", "palau", "qatar", "rwanda", "russia", "romania", 
    "saint-vincent-and-the-grenadines", "sao-tome-and-principe", "syria", "suriname", "switzerland", "sweden", 
    "swaziland", "sudan", "south-africa", "spain", "sri-lanka", "south-sudan", "south-korea", "somalia", 
    "solomon-islands", "slovenia", "slovakia", "singapore", "sierra-leone", "seychelles", "serbia", "senegal", 
    "saudi-arabia", "san-marino", "samoa", "saint-lucia", "saint-kitts-and-nevis", "togo", "turkmenistan", 
    "turkey", "tunisia", "tuvalu", "trinidad-and-tobago", "thailand", "tanzania", "tajikistan", "tonga", "taiwan", 
    "ukraine", "uzbekistan", "uruguay", "united-states-of-america", "united-kingdom", "united-arab-emirates", "uganda", 
    "vietnam", "venezuela", "vatican-city", "yemen", "zambia", "zimbabwe",
}

var s_iso3166_a2 [196]string = [196]string{"AR", "GW", "AT", "AU", "AM", "AO", "AF", "AZ", "AD", "DZ", "AL", "BN", "BI", "BF", "BG", "BR", "BJ", "BW", "BA", "BO", "BZ", "BH", "BY", "BE", "BB", "BD", "BT", "BS", "CU", "CI", "CZ", "CY", "HR", "KM", "CR", "CD", "CD", "CO", "CN", "CL", "TD", "CF", "CV", "CA", "CM", "KH", "DO", "DM", "DJ", "DK", "ET", "EE", "ER", "GQ", "SV", "EG", "EC", "TL", "FR", "FI", "FJ", "GY", "GW", "GN", "GT", "GD", "GR", "GH", "GE", "GM", "GA", "DE", "HU", "HN", "HT", "IT", "IL", "IE", "IQ", "IQ", "ID", "IN", "IS", "JO", "JP", "JM", "KG", "KZ", "CI", "KW", "KI", "KE", "LU", "LT", "LI", "LY", "LR", "MO", "LS", "LB", "LV", "MC", "MM", "MZ", "MA", "ME", "MN", "MD", "FM", "MX", "MU", "MR", "MH", "MT", "ML", "MV", "MY", "MW", "MG", "MK", "NG", "NO", "KR", "NE", "NI", "NZ", "NU", "NL", "NP", "NR", "NA", "OM", "PG", "PT", "PL", "PH", "PE", "PY", "PA", "PK", "PW", "QA", "RW", "RU", "RO", "VC", "ST", "SY", "SR", "CH", "SE", "BR", "SD", "ZA", "ES", "LK", "SS", "SS", "SO", "SB", "SI", "SK", "SG", "SL", "SC", "RS", "SN", "SA", "SM", "WS", "LC", "KN", "TG", "TM", "TR", "TN", "TV", "TT", "TH", "TZ", "TJ", "TO", "TW", "UA", "UZ", "UY", "US", "GB", "AE", "UG", "VN", "VE", "VU", "YE", "ZM", "ZW"}
var s_iso3166_a3 [196]string = [196]string{"ARG", "GNB", "AUT", "AUS", "ARM", "AGO", "AFG", "AZE", "AND", "DZA", "ALB", "BRN", "BDI", "BFA", "BGR", "BRA", "BEN", "BWA", "BIH", "BOL", "BLZ", "BHR", "BLR", "BEL", "BRB", "BGD", "BTN", "BHS", "CUB", "CIV", "CZE", "CYP", "HRV", "COM", "CRI", "COD", "COD", "COL", "CHN", "CHL", "TCD", "CAF", "CPV", "CAN", "CMR", "KHM", "DOM", "DMA", "DJI", "DNK", "ETH", "EST", "ERI", "GNQ", "SLV", "EGY", "ECU", "TLS", "FRA", "FIN", "FJI", "GUY", "GNB", "GIN", "GTM", "GRD", "GRC", "GHA", "GEO", "GMB", "GAB", "DEU", "HUN", "HND", "HTI", "ITA", "ISR", "IRL", "IRQ", "IRQ", "IDN", "IND", "ISL", "JOR", "JPN", "JAM", "KGZ", "KAZ", "CIV", "KWT", "KIR", "KEN", "LUX", "LTU", "LIE", "LBY", "LBR", "MAC", "LSO", "LBN", "LVA", "MCO", "MMR", "MOZ", "MAR", "MNE", "MNG", "MDA", "FSM", "MEX", "MUS", "MRT", "MHL", "MLT", "MLI", "MDV", "MYS", "MWI", "MDG", "MKD", "NGA", "NOR", "KOR", "NER", "NIC", "NZL", "NIU", "NLD", "NPL", "NRU", "NAM", "OMN", "PNG", "PRT", "POL", "PHL", "PER", "PRY", "PAN", "PAK", "PLW", "QAT", "RWA", "RUS", "ROU", "VCT", "STP", "SYR", "SUR", "CHE", "SWE", "BRA", "SDN", "ZAF", "ESP", "LKA", "SSD", "SSD", "SOM", "SLB", "SVN", "SVK", "SGP", "SLE", "SYC", "SRB", "SEN", "SAU", "SMR", "WSM", "LCA", "KNA", "TGO", "TKM", "TUR", "TUN", "TUV", "TTO", "THA", "TZA", "TJK", "TON", "TWN", "UKR", "UZB", "URY", "USA", "GBR", "ARE", "UGA", "VNM", "VEN", "VUT", "YEM", "ZMB", "ZWE"}
var s_iso3166_nu [196]string = [196]string{"032", "624", "040", "036", "051", "024", "004", "031", "020", "012", "008", "096", "108", "854", "100", "076", "204", "072", "070", "068", "084", "048", "112", "056", "052", "050", "064", "044", "192", "384", "203", "196", "191", "174", "188", "180", "180", "170", "156", "152", "148", "140", "132", "124", "120", "116", "214", "212", "262", "208", "231", "233", "232", "226", "222", "818", "218", "626", "250", "246", "242", "328", "624", "324", "320", "308", "300", "288", "268", "270", "266", "276", "348", "340", "332", "380", "376", "372", "368", "368", "360", "356", "352", "400", "392", "388", "417", "398", "384", "414", "296", "404", "442", "440", "438", "434", "430", "446", "426", "422", "428", "492", "104", "508", "504", "499", "496", "498", "583", "484", "480", "478", "584", "470", "466", "462", "458", "454", "450", "807", "566", "578", "410", "562", "558", "554", "570", "528", "524", "520", "516", "512", "598", "620", "616", "608", "604", "600", "591", "586", "585", "634", "646", "643", "642", "670", "678", "760", "740", "756", "752", "076", "729", "710", "724", "144", "728", "728", "706", "090", "705", "703", "702", "694", "690", "688", "686", "682", "674", "882", "662", "659", "768", "795", "792", "788", "798", "780", "764", "834", "762", "776", "158", "804", "860", "858", "840", "826", "784", "800", "704", "862", "548", "887", "894", "716"}

const URL_FORMAT = "https://cdn.countryflags.com/thumbs/%v/%v.png"

type DownloadTarget struct {
    ISO3166 ISO3166
    Country string
    Shape string
    URL string
}

type ISO3166 struct {
    A2 string
    A3 string
    NU string
}

func list_contains(list *[]string, val string) bool {
    for _,v := range *list {
        if v == val {
            return true
        }
    }

    return false
}

func get_image_list(countries *[]string, shapes *[]string) []DownloadTarget {
    results := []DownloadTarget{}
    for idx,country := range s_countries {
        if countries == nil || list_contains(countries, country) {
            for _,shape := range s_shapes {
                if shapes == nil || list_contains(shapes, shape) {
                    download_url := fmt.Sprintf(URL_FORMAT, country, shape)
                    results = append(results, DownloadTarget{
                        Country: country,
                        Shape: shape,
                        URL: download_url,
                        ISO3166: ISO3166{
                            A2: s_iso3166_a2[idx],
                            A3: s_iso3166_a3[idx],
                            NU: s_iso3166_nu[idx],
                        },
                    })
                }
            }
        }
    }

    return results
}

func download(src string) ([]byte, error) {
    resp, err := http.Get(src)
    if err != nil {
        return []byte{}, err
    }

    defer resp.Body.Close()
    return ioutil.ReadAll(resp.Body)
}

func download_all(countries *[]string, shapes *[]string, output_dir string, output_type string) {
    images := get_image_list(countries, shapes)

    if output_type == "png" {
        _ = os.Mkdir(output_dir, 0700)
        for _,image := range images {
            out_path := output_dir + "/" + image.Country + "-" + image.Shape + "." + output_type
            fmt.Printf("Downloading %s ... ", out_path)
            data, err := download(image.URL)
            if err != nil {
                fmt.Printf("[%s: %v]\n", "FAILED", err)
                continue;
            }

            err = ioutil.WriteFile(out_path, data, 0700)
            if err != nil {
                fmt.Printf("[%s: %v]\n", "FAILED", err)
                continue;
            }

            fmt.Printf("%s\n", "[Done]")
        }
    } else if output_type == "b64" {
        _ = os.Mkdir(output_dir, 0700)
        for _,image := range images {
            out_path := output_dir + "/" + image.Country + "-" + image.Shape + "." + output_type
            fmt.Printf("downloading %s ... ", out_path)
            data, err := download(image.URL)
            if err != nil {
                fmt.Printf("[%s: %v]\n", "FAILED", err)
                continue;
            }

            err = ioutil.WriteFile(out_path, []byte(base64.StdEncoding.EncodeToString(data)), 0700)
            if err != nil {
                fmt.Printf("[%s: %v]\n", "FAILED", err)
                continue;
            }

            fmt.Printf("%s\n", "[Done]")
        }
    } else if output_type == "b64-iso3166-numeric-json-file" {
        _ = os.Mkdir(output_dir, 0700)
        jsonResult := map[string]map[string]string{}
        for _,image := range images {
            fmt.Printf("downloading %s(%s)-%s ... ", image.Country, image.ISO3166.NU, image.Shape)
            data, err := download(image.URL)
            if err != nil {
                fmt.Printf("[%s: %v]\n", "FAILED", err)
                continue;
            }

            if jsonResult[image.ISO3166.NU] == nil {
                jsonResult[image.ISO3166.NU] = map[string]string{}
            }

            jsonResult[image.ISO3166.NU][image.Shape] = base64.StdEncoding.EncodeToString(data)
            fmt.Printf("[Done]\n")
        }

        fmt.Printf("%s", "building json data ... ")
        jsonStrBytes, err := json.Marshal(jsonResult)
        if err != nil {
            fmt.Printf("[%s: %v]\n", "error", err)
            os.Exit(1)
        }
        fmt.Printf("%s\n", "[Done]")

        out_path := output_dir + "/flags_iso3166_numeric.json"

        fmt.Printf("%s", "writing json file ... ")
        err = ioutil.WriteFile(out_path, jsonStrBytes, 0700)
        if err != nil {
            fmt.Printf("[%s: %v]\n", "error", err)
            os.Exit(1)
        }
        fmt.Printf("%s\n", "[Done]")

        fmt.Printf("%s: %s\n", "output iso3166-numeric json file to", out_path)
    } else if output_type == "b64-iso3166-alpha2-json-file" {
        _ = os.Mkdir(output_dir, 0700)
        jsonResult := map[string]map[string]string{}
        for _,image := range images {
            fmt.Printf("downloading %s(%s)-%s ... ", image.Country, image.ISO3166.A2, image.Shape)
            data, err := download(image.URL)
            if err != nil {
                fmt.Printf("[%s: %v]\n", "FAILED", err)
                continue;
            }

            if jsonResult[image.ISO3166.A2] == nil {
                jsonResult[image.ISO3166.A2] = map[string]string{}
            }

            jsonResult[image.ISO3166.A2][image.Shape] = base64.StdEncoding.EncodeToString(data)
            fmt.Printf("[Done]\n")
        }

        fmt.Printf("%s", "building json data ... ")
        jsonStrBytes, err := json.Marshal(jsonResult)
        if err != nil {
            fmt.Printf("[%s: %v]\n", "error", err)
            os.Exit(1)
        }
        fmt.Printf("%s\n", "[Done]")

        out_path := output_dir + "/flags_iso3166_a2.json"

        fmt.Printf("%s", "writing json file ... ")
        err = ioutil.WriteFile(out_path, jsonStrBytes, 0700)
        if err != nil {
            fmt.Printf("[%s: %v]\n", "error", err)
            os.Exit(1)
        }
        fmt.Printf("%s\n", "[Done]")

        fmt.Printf("%s: %s\n", "output iso3166-alpha-2 json file to", out_path)
    } else if output_type == "b64-iso3166-alpha3-json-file" {
        _ = os.Mkdir(output_dir, 0700)
        jsonResult := map[string]map[string]string{}
        for _,image := range images {
            fmt.Printf("downloading %s(%s)-%s ... ", image.Country, image.ISO3166.A3, image.Shape)

            data, err := download(image.URL)
            if err != nil {
                fmt.Printf("[%s: %v]\n", "FAILED", err)
                continue;
            }

            if jsonResult[image.ISO3166.A3] == nil {
                jsonResult[image.ISO3166.A3] = map[string]string{}
            }

            jsonResult[image.ISO3166.A3][image.Shape] = base64.StdEncoding.EncodeToString(data)
            fmt.Printf("[Done]\n")
        }

        fmt.Printf("%s", "building json data ... ")
        jsonStrBytes, err := json.Marshal(jsonResult)
        if err != nil {
            fmt.Printf("[%s: %v]\n", "error", err)
            os.Exit(1)
        }
        fmt.Printf("%s\n", "[Done]")

        out_path := output_dir + "/flags_iso3166_a3.json"

        fmt.Printf("%s", "writing json file ... ")
        err = ioutil.WriteFile(out_path, jsonStrBytes, 0700)
        if err != nil {
            fmt.Printf("[%s: %v]\n", "error", err)
            os.Exit(1)
        }
        fmt.Printf("%s\n", "[Done]")

        fmt.Printf("%s: %s\n", "output iso3166-alpha-3 json file to", out_path)
    } else {
        fmt.Printf("error: %s: %s\n", "invalid output type", output_type)
        os.Exit(1)
    }
}

func list_from_csv(csv string) *[]string {
    if csv == "" {
        return nil
    }

    result := strings.Split(csv, ",")
    return &result
}

func main() {
    listCountries := flag.Bool("list-countries", false, "list available countries")
    listShapes := flag.Bool("list-shapes", false, "list available shapes")
    doDownload := flag.Bool("download", false, "download all images within the supplied filter")
    output := flag.String("output-dir", "./", "the directory in which to store the output images")
    output_type := flag.String("output-type", "png", "the output type, valid values are 'png', 'b64', 'b64-iso3166-numeric-json-file', 'b64-iso3166-alpha2-json-file', 'b64-iso3166-alpha3-json-file'")
    filterCountries_str := flag.String("filter-countries", "", "the list of countries to filter, comma separated")
    filterShapes_str := flag.String("filter-shapes", "", "the list of shapes to filter, comma separated")

    flag.Parse()

    filter_countries := list_from_csv(*filterCountries_str)
    filter_shapes := list_from_csv(*filterShapes_str)

    if *listCountries {
        fmt.Printf("%s\n\n", "Available Countries:")
        for _,country := range s_countries {
            fmt.Printf("    - %s\n", country)
        }
        os.Exit(0)
    }

    if *listShapes {
        fmt.Printf("%s\n\n", "Available Shapes:")
        for _,shape := range s_shapes {
            fmt.Printf("    - %s\n", shape)
        }
        os.Exit(0)
    }

    if *doDownload {
        download_all(filter_countries, filter_shapes, *output, *output_type)
    }
}