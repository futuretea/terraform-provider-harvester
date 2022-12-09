---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harvester_vlanconfig Resource - terraform-provider-harvester"
subcategory: ""
description: |-
  
---

# harvester_vlanconfig (Resource)



## Example Usage

```terraform
resource "harvester_vlanconfig" "cluster-vlan-node1" {
  name = "cluster-vlan-node1"

  cluster_network_name = harvester_clusternetwork.cluster-vlan.name

  uplink {
    nics = [
      "eth5",
      "eth6"
    ]

    bond_mode = "active-backup"
    mtu       = 1500
  }

  node_selector = {
    "kubernetes.io/hostname" : "node1"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **cluster_network_name** (String) mgmt is a built-in cluster network and does not support creating/updating network configs.
- **name** (String) A unique name
- **uplink** (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--uplink))

### Optional

- **description** (String) Any text you want that better describes this resource
- **id** (String) The ID of this resource.
- **node_selector** (Map of String) refer to https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodeselector
- **tags** (Map of String)

### Read-Only

- **matched_nodes** (List of String)
- **message** (String)
- **state** (String)

<a id="nestedblock--uplink"></a>
### Nested Schema for `uplink`

Required:

- **nics** (List of String)

Optional:

- **bond_miimon** (Number) refer to https://www.kernel.org/doc/Documentation/networking/bonding.txt
- **bond_mode** (String)
- **mtu** (Number)

## Import

Import is supported using the following syntax:

```shell
terraform import harvester_vlanconfig.foo <Name>
```