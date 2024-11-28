# Create level 1 group attached to the root node
resource "hpeuxi_group" "level_1" {
  name            = "Parent Group"
}

# Create level 2 group attached to level 1 group
resource "hpeuxi_group" "level_2" {
  name            = "Child Group"
  parent_group_id = hpeuxi_group.level_1.id
}
