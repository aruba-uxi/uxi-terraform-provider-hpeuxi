# Retrieve data for level 1 group
data "uxi_group" "level_1" {
  filter = {
    id = "<level_1_id>"
  }
}

# Retrieve data for level 2 group
data "uxi_group" "level_2" {
  filter = {
    id = "<level_2_id>"
  }
}
