# Create a parent group attached to the root node
resource "uxi_group" "parent_group" {
  name            = "Parent Group"
}

# Create a child group attached to the parent group
resource "uxi_group" "child_group" {
  name            = "Child Group"
  parent_group_id = uxi_group.parent_group.id
}
