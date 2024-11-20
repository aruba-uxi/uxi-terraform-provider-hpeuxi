resource "uxi_group" "parent_group" {
  name            = "Parent Group"
}

resource "uxi_group" "child_group" {
  name            = "Child Group"
  parent_group_id = uxi_group.parent_group.id
}
