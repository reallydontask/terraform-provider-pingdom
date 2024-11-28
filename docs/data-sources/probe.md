---
layout: "pingdom"
page_title: "pingdom_probe"
sidebar_current: "docs-pingdom-data-probe"
description: |-
  Provides a data for a probe
---

# pingdom\_probe

Provides a way to fetch all probe servers.

## Example Usage

```hcl
data "pingdom_probe" "all" {
}
```

## Attributes Reference

* `active` - Whether the probe is active.
* `country` - The probe's country.
* `countryiso` - The probe's iso country code.
* `city` - The probe's city.
* `hostname` - The probe's hostname.
* `ip` - The probe's ipv4 address.
* `ipv6` - The probe's ipv6 address.
* `name` - The probe's name.
* `region` - The probe's region.
