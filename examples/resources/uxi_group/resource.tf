# Create a parent group attached to the root node
resource "uxi_group" "level_1" {
  name            = "Parent Group"
}

# Create a child group attached to the parent group
resource "uxi_group" "level_2" {
  name            = "Child Group"
  parent_group_id = uxi_group.level_1.id
}
