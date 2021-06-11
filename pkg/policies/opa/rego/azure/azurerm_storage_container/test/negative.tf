provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "rg_details" {
  name     = "example-resources"
  location = "West Europe"
}

resource "azurerm_storage_account" "storage_details" {
  name                     = "examplestoraccount"
  resource_group_name      = azurerm_resource_group.rg_details.name
  location                 = azurerm_resource_group.rg_details.location
  account_tier             = "Standard"
  account_replication_type = "LRS"

  tags = {
    environment = "staging"
  }
}

resource "azurerm_storage_container" "checkStorageContainerAccess" {
  name                  = "vhds"
  storage_account_name  = azurerm_storage_account.storage_details.name
  container_access_type = "private"
}
