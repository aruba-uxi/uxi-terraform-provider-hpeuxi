# Create a root group (no parent_group_id)
resource "uxi_group" "parent_group" {
  name            = "Parent Group"
}

# Create a child group of the root group
resource "uxi_group" "child_group" {
  name            = "Child Group"
  parent_group_id = uxi_group.parent_group.id
}
