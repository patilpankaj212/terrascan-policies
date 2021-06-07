provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "resource_group_details" {
  name     = "example-resources"
  location = "West US 2"
}

resource "azurerm_managed_disk" "checkDataDisksEncrypted" {
  name                 = "acctestmd"
  location             = "West US 2"
  resource_group_name  = azurerm_resource_group.resource_group_details.name
  storage_account_type = "Standard_LRS"
  create_option        = "Empty"
  disk_size_gb         = "1"

  tags = {
    environment = "staging"
  }

  encryption_settings {
    enabled = false
  }
}
