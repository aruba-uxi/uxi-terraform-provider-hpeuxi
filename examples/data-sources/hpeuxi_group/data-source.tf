# Retrieve data for level 1 group
data "hpeuxi_group" "level_1" {
  filter = {
    id = "<level_1_group_id>"
  }
}
