
provider "box" {
    base_url = "https://api.box.com/2.0/",
    access_token = "NvocCaINeboZfFLGXzzJGC5A3IURu8S5"
}

resource "box_folder" "vinayv9" {
    name = "voteForChange"
    id = "0"
}

