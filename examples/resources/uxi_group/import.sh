# Import level 1 group using its ID
terraform import uxi_group.level_1 <level_1_id>

# Import level 1 group using its ID with an import block
import {
  to = uxi_group.level_1
  id = "level_1_id"
}

# Import level 2 group using its ID
terraform import uxi_group.level_2 <level_2_id>

# Import level 2 group using its ID with an import block
import {
  to = uxi_group.level_2
  id = "level_2_id"
}
